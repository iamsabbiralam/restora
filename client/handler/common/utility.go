package common

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/Masterminds/sprig"
	"github.com/gorilla/csrf"
	"golang.org/x/crypto/bcrypt"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

var (
	IgnorePath = []string{LoginPath, RegistrationPath}
)

var ValidateImgFileType = []string{
	"image/jpeg", "image/png", "image/jpg",
}

type SessionUser struct {
	ID              string
	Email           string
	Image           string
	RoleID          string
	RoleName        string
	UserName        string
	FirstName       string
	LastName        string
	DesignationName string
}

type Authenticator struct {
	BaseURL   string
	LogoutURL string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func (s *Server) GetSessionUser(r *http.Request) *SessionUser {
	sess, err := store.Get(r, SessionCookieName)
	if err != nil {
		s.Logger.Infof("Unable to get session %+v", err)
		return &SessionUser{}
	}

	userID := fmt.Sprintf("%v", sess.Values["authUserID"])
	if userID == "" {
		s.Logger.Infof("Session does not contain user id %+v", err)
		return &SessionUser{}
	}

	return &SessionUser{
		ID:        userID,
		FirstName: fmt.Sprintf("%v", sess.Values[SessionFirstName]),
		LastName:  fmt.Sprintf("%v", sess.Values[SessionLastName]),
		Email:     fmt.Sprintf("%v", sess.Values[SessionEmail]),
		UserName:  fmt.Sprintf("%v", sess.Values[SessionUserName]),
	}
}

func (s *Server) ParseTemplates() error {
	templates := template.New("client-templates").Funcs(template.FuncMap{
		"assetHash": func(n string) string {
			return path.Join("/", s.Assets.HashName(strings.TrimPrefix(path.Clean(n), "/")))
		},
		"activeStatus": func(status int32) string {
			if status == 1 {
				return "Active"
			}
			return "Inactive"
		},
		"incrementKey": func(status int) int {
			return status + 1
		},
		"formatDate": func(ts *tspb.Timestamp, layout string) string {
			if !ts.IsValid() {
				return ""
			}
			return ts.AsTime().Format(layout)
		},

		"countPaginate": func(a, b int32) int32 {
			if a > 0 {
				c := a / b
				if a%b != 0 {
					c = c + 1
				}
				return c
			}
			return 0
		},
		"noScape": func(str string) template.HTML {
			if str == "" {
				return template.HTML("<h1>Content not found</h1>")
			}
			return template.HTML(str)
		},
		"nowTime": func() string {
			return time.Now().Format("02 Jan 2006")
		},
		"permissionChecked": func(res string, act string, allPerm map[string][]string) string {
			if val, ok := allPerm[res]; ok {
				for _, v := range val {
					if v == act {
						return "checked"
					}
				}
				return ""
			}
			return ""
		},
		"permission": func(res string) bool {
			return true
		}, "urls": func(url string, params ...string) string {
			for _, v := range params {
				a := strings.Split(v, "_")
				if len(a) == 2 {
					url = strings.Replace(url, "{"+a[0]+"}", a[1], 1)
				}
			}
			return url
		},
	}).Funcs(sprig.FuncMap())

	tmpl, err := templates.ParseFS(s.Assets, "templates/*/*.html", "templates/*/*/*.html")
	if err != nil {
		return err
	}

	s.Templates = tmpl
	return nil
}

func (s *Server) GetErrorHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := s.DoTemplate(w, r, "error.html", http.StatusTemporaryRedirect); err != nil {
			s.Logger.WithError(err).Error("unable to load error template")
		}
	})
}

func IsPartialTemplate(name string) bool {
	return strings.HasSuffix(name, ".part.html")
}

func (s *Server) DoTemplate(w http.ResponseWriter, r *http.Request, name string, status int) error {
	template := s.LookupTemplate(name)
	if template == nil || IsPartialTemplate(name) {
		template, status = s.Templates.Lookup("error.html"), http.StatusNotFound
	}

	w.WriteHeader(status)
	return template.Execute(w, s.TemplateData(r))
}

func (s *Server) LookupTemplate(name string) *template.Template {
	if s.Env == "development" {
		if err := s.ParseTemplates(); err != nil {
			s.Logger.WithError(err).Error("template reload")
			return nil
		}
	}

	return s.Templates.Lookup(name)
}

type TemplateData struct {
	Env       string
	CSRFField template.HTML
	Form      TemplateForm
}

type TemplateForm struct {
	ErrorCode    string
	ErrorDetails string
}

func (s *Server) TemplateData(r *http.Request) TemplateData {
	return TemplateData{
		Env:       s.Env,
		CSRFField: csrf.TemplateField(r),
		Form: TemplateForm{
			ErrorCode:    "500",
			ErrorDetails: "Internal error",
		},
	}
}

func CacheStaticFiles(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if asset is hashed extend cache to 180 days
		e := `"4FROTHS24N"`
		w.Header().Set("Etag", e)
		w.Header().Set("Cache-Control", "max-age=15552000")
		if match := r.Header.Get("If-None-Match"); match != "" {
			if strings.Contains(match, e) {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func checkGuestPath(r *http.Request) bool {
	for _, val := range IgnorePath {
		if strings.HasPrefix(r.URL.Path, val) {
			return true
		}
	}
	return false
}

func (s *Server) GetAuthMiddleware(next http.Handler) http.Handler {
	s.Logger.WithField("method", "handler.utility.GetAuthMiddleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.Cookies.Get(r, SessionCookieName)
		if err != nil {
			log.Fatal(err)
		}

		authUserID := session.Values["authUserID"]
		if authUserID != nil {
			if checkGuestPath(r) {
				http.Redirect(w, r, DashboardPath, http.StatusSeeOther)
			}

			next.ServeHTTP(w, r)
		} else if authUserID == nil {
			if !checkGuestPath(r) {
				http.Redirect(w, r, LoginPath, http.StatusSeeOther)
			}

			next.ServeHTTP(w, r)
		}

		http.Redirect(w, r, HomePath, http.StatusSeeOther)
	})
}

func (s *Server) SessionResetLoginData(r *http.Request, w http.ResponseWriter) error {
	uID, err := middleware.GetUserID(r.Context())
	if err != nil {
		return fmt.Errorf(" Unable to get user id from hydra session %+v", err)
	}

	userInfo, err := s.User.GetUserByID(r.Context(), &user.GetUserByIDRequest{
		ID: uID,
	})
	if err != nil {
		return fmt.Errorf(" Unable to get user info %+v", err)
	}

	sess, err := s.Sess.Get(r, SessionCookieName)
	if err != nil {
		return fmt.Errorf(" Unable to get session %+v", err)
	}

	sess.Values[SessionUserID] = userInfo.GetUser().GetID()
	sess.Values[SessionEmail] = userInfo.GetUser().GetEmail()
	sess.Values[SessionUserName] = userInfo.GetUser().GetUsername()
	sess.Values[SessionProfileImage] = ""
	if err := s.Sess.Save(r, w, sess); err != nil {
		return fmt.Errorf(" Unable to save add session info %+v", err)
	}

	return nil
}

func (s *Server) StringToDate(date string) time.Time {
	layout := "2006-01-02"
	fDate, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
	}

	return fDate
}

func (s *Server) ValidateSingleFileType(r *http.Request, name string, types []string) error {
	f, _, err := r.FormFile(name)
	if err != nil {
		return err
	}

	defer f.Close()
	c, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	ft := http.DetectContentType(c)
	rtn := false
	for _, t := range types {
		if ft == t {
			rtn = true
		}
	}

	if !rtn {
		return errors.New("invalid file format")
	}

	return nil
}

func RandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, n)
	read, err := crand.Read(bytes)
	// Note that err == nil if we read len(b) bytes.
	if err != nil {
		return ""
	}

	if read != n {
		return ""
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes)
}

func (s *Server) SaveImage(file multipart.File, fileHeader *multipart.FileHeader, imagePath string) (string, error) {
	if fileHeader == nil {
		return "", nil
	}

	tt := time.Now().Local()
	image := tt.Format("20060102") + RandomString(5) + fileHeader.Filename
	dest, err := os.Create(fmt.Sprintf(imagePath+"%s", image))
	if err != nil {
		return "", err
	}
	defer dest.Close()
	if _, err := io.Copy(dest, file); err != nil {
		fmt.Println(err.Error())
	}

	return image, err
}
