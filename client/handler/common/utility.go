package common

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
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

type DynamicQueryString struct {
	SearchTerm   string
	StartDate    string
	SortBy       string
	SortByColumn string
	EndDate      string
	PageNumber   int32
	CurrentPage  int32
	Offset       int32
	OthersValue  map[string]string
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

func (s *Server) StringToDate(date string) time.Time {
	layout := "2006-01-02"
	fDate, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
	}

	return fDate
}

func GetQueryStringData(r *http.Request, keys []string, isNotDefault bool) *DynamicQueryString {
	var data DynamicQueryString
	queryParams := r.URL.Query()
	var err error
	data.SearchTerm, err = url.PathUnescape(queryParams.Get("search-term"))
	if err != nil {
		data.SearchTerm = ""
	}

	data.StartDate, err = url.PathUnescape(queryParams.Get("start"))
	if err != nil {
		data.StartDate = ""
	}

	data.EndDate, err = url.PathUnescape(queryParams.Get("end"))
	if err != nil {
		data.EndDate = ""
	}

	data.SortBy, err = url.PathUnescape(queryParams.Get("sort-by"))
	if err != nil {
		data.SortBy = ""
	}

	data.SortByColumn, err = url.PathUnescape(queryParams.Get("sort-by-column"))
	if err != nil {
		data.SortByColumn = ""
	}

	page, err := url.PathUnescape(queryParams.Get("page"))
	if err != nil || page == "" {
		page = "1"
	}

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		pageNumber = 1
	}

	data.PageNumber = int32(pageNumber)
	var offset int32 = 0
	currentPage := pageNumber
	if currentPage <= 0 {
		currentPage = 1
	} else {
		offset = LimitPerPage*int32(currentPage) - LimitPerPage
	}

	data.CurrentPage = int32(currentPage)
	data.Offset = offset

	if isNotDefault {
		data = DynamicQueryString{}
	}

	if len(keys) > 0 {
		myMap := make(map[string]string)
		for _, v := range keys {
			myMap[v], err = url.PathUnescape(queryParams.Get(v))
			if err != nil {
				myMap[v] = ""
			}
		}

		data.OthersValue = myMap
	}

	return &data
}
