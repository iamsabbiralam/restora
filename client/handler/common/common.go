package common

import (
	"errors"
	"html/template"
	"os"

	"github.com/benbjohnson/hashfs"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/iamsabbiralam/restora/proto/v1/usermgm/auth"
	"github.com/iamsabbiralam/restora/proto/v1/usermgm/user"

	"github.com/iamsabbiralam/restora/proto/v1/server/brand"
	"github.com/iamsabbiralam/restora/proto/v1/server/category"
)

type Server struct {
	Env       string
	Config    *viper.Viper
	Logger    *logrus.Entry
	Assets    *hashfs.FS
	Decoder   *schema.Decoder
	Auth      *Authenticator
	Cookies   *sessions.CookieStore
	Templates *template.Template
	Sess      *sessions.Session
	User      user.UserServiceClient
	Login     auth.LoginServiceClient
	Category  category.CategoryServiceClient
	Brand     brand.BrandServiceClient
}

type (
	PublicTemplateData struct {
		UserInfo *SessionUser
	}
	Status struct {
		ID   int32
		Name string
	}

	JsonErrorFormat struct {
		Status   bool   `Json:"status"`
		Code     int32  `Json:"code"`
		Message  string `Json:"message"`
		ErrorMas map[string]string
	}
)

var (
	// NotFound is returned when the requested resource does not exist.
	NotFound = status.Error(codes.NotFound, "not found")
	// Conflict is returned when trying to create the same resource twice.
	Conflict = status.Error(codes.AlreadyExists, "conflict")
	// UsernameExists is returned when the username already exists in storage.
	UsernameExists = errors.New("username already exists")
	// EmailExists is returned when signup email already exists in storage.
	EmailExists = errors.New("email already exists")
	// InvCodeExists is returned when invitation code already exists in storage.
	InvCodeExists = errors.New("invitation code already exists")
)

var store = sessions.NewCookieStore([]byte(os.Getenv("secret")))

const (
	SessionCookieName          = "restora-session"
	SessionCookieState         = "state"
	AuthCodeURL                = "somerandomstring"
	GenericErrMsg              = "Please contact the administrator."
	SessionCookieLoginRedirect = "loginRedirect"
	SessionCookieToken         = "token"

	SessionUserID       = "user-id"
	SessionEmail        = "email"
	SessionProfileImage = "profile-image"
	SessionFirstName    = "first-name"
	SessionLastName     = "last-name"
	SessionEmpFirstName = "emp-first-name"
	SessionEmpLastName  = "emp-last-name"
	SessionDesignation  = "designation"
	SessionRoleID       = "roleID"
	SessionUserName     = "userName"

	SUPERADMIN       = "Super Admin"
	FromName         = "restora-session"
	LimitPerPage     = 10
	InvitationStatus = 3
)

const (
	HomePath  = "/"
	ErrorPath = "/error"

	RegistrationPath = "/registration"
	LoginPath        = "/login"

	ProfilePath            = "/profile"
	ProfileEditPath        = "/profile/edit"
	ResendOtpPath          = "/resend-otp"
	ChangePasswordPath     = "/password"
	ChangePasswordEditPath = "/change/password"
	OTPPasswordPath        = "/otp/password"
	UploadProfileImagePath = "/profile/update/image"
	LoginCallBackPath      = "/oauth2/callback"
	RedirectURLPath        = "/redirect-url"
	ConsentPath            = "/consent"
	LogoutPath             = "/logout"

	DashboardPath = "/dashboard"

	UserListPath = "/user/list"

	CategoryCreatePath = "/category/create"
	CategoryEditPath   = "/category/edit/{id}"
	CategoryListPath   = "/category/list"
	CategoryDeletePath = "/category/delete/{id}"

	BrandCreatePath = "/brand/create"
	BrandEditPath   = "/brand/edit/{id}"
	BrandListPath   = "/brand/list"
	BrandDeletePath = "/brand/delete/{id}"
)

// regex validation
// only text, space is allowed but no number is not allowed
const TextValidation = `^[A-Za-z.-]+(\s*[A-Za-z.-]+)*$`
