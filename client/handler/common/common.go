package common

import (
	"errors"
	"html/template"

	"github.com/benbjohnson/hashfs"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	HomePath               = "/"
	ErrorPath              = "/error"

	RegistrationPathPath   = "/registration"
	LoginInPath            = "/login"
	
	ProfilePath            = "/profile"
	ResendOtpPath          = "/resend-otp"
	ChangePasswordPath     = "/change/password"
	OTPPasswordPath        = "/otp/password"
	ProfileEditPath        = "/profile/edit"
	UploadProfileImagePath = "/profile/update/image"
	LoginCallBackPath      = "/oauth2/callback"
	RedirectURLPath        = "/redirect-url"
	ConsentPath            = "/consent"
	LogoutPath             = "/logout"

	AdminDashboardPath    = "/admin/dashboard"
	EmployeeDashboardPath = "/employee/dashboard"

	DesignationListPath   = "/admin/designations"
	DesignationCreatePath = "/admin/designations/create"
	DesignationUpdatePath = "/admin/designations/update/{id}"
	DesignationDeletePath = "/admin/designations/delete/{id}"

	DepartmentListPath   = "/admin/departments"
	DepartmentCreatePath = "/admin/departments/create"
	DepartmentUpdatePath = "/admin/departments/update/{id}"
	DepartmentDeletePath = "/admin/departments/delete/{id}"

	SettingPath = "/admin/settings/{groupSetting}"

	CreateRolePath       = "/admin/roles/create"
	UpdateRolePath       = "/admin/roles/update/{id}"
	UpdateRoleStatusPath = "/admin/roles/update-status/{id}"
	RoleListPath         = "/admin/roles"
	DeleteRolePath       = "/admin/roles/delete/{id}"
)

// regex validation
// only text, space is allowed but no number is not allowed
const TextValidation = `^[A-Za-z.-]+(\s*[A-Za-z.-]+)*$`
