package categories

import (
	"html/template"
	"net/http"
	"regexp"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"

	"github.com/iamsabbiralam/restora/client/handler/common"
	"github.com/iamsabbiralam/restora/client/handler/paginator"
	catG "github.com/iamsabbiralam/restora/proto/v1/server/category"
)

type Svc struct {
	*common.Server
}

type Category struct {
	ID        string
	Name      string
	Status    catG.Status
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

type CategoryTempData struct {
	CSRFField      template.HTML
	Form           Category
	Data           []Category
	FormAction     string
	FormErrors     map[string]string
	FormMessage    map[string]string
	SearchTerm     string
	PramStatus     string
	PaginationData paginator.Paginator
	SortByColumn   string
	StartDate      string
	EndDate        string
	SortBy         string
	Status         []common.Status
	FilterFormErr  string
}

func Register(h *common.Server, mw *mux.Router) (*mux.Router, error) {
	s := &Svc{
		Server: h,
	}

	mw.HandleFunc(common.CategoryCreatePath, s.createCategoryHandler).Methods("GET").Name("category.create")
	mw.HandleFunc(common.CategoryCreatePath, s.storeCategoryHandler).Methods("POST").Name("category.store")
	mw.HandleFunc(common.CategoryUpdatePath, s.editCategoryHandler).Methods("GET").Name("category.edit")
	mw.HandleFunc(common.CategoryUpdatePath, s.updateCategoryHandler).Methods("POST").Name("category.update")
	mw.HandleFunc(common.CategoryDeletePath, s.deleteCategoryHandler).Methods("GET").Name("category.delete")
	mw.HandleFunc(common.CategoryListPath, s.listCategoryHandler).Methods("GET").Name("category.list")
	return mw, nil
}

func (s *Svc) loadCategoryTemplate(w http.ResponseWriter, r *http.Request, data CategoryTempData, htmlFile string) {
	log := s.Logger.WithField("method", "handler.admin.category.loadCategoryTemplate")
	tmpl := s.LookupTemplate(htmlFile)
	if tmpl == nil {
		log.Error("unable to load template")
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Errorf("unable to execute template: %s", err)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
		return
	}
}

func (c *Category) ValidateCategory(server *common.Server, r *http.Request, id string) error {
	vre := validation.Required.Error
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, vre("The category name is required"), validation.Length(3, 100), validation.Match(regexp.MustCompile(common.TextValidation)).Error("Category name is invalid, only text is allowed.")),
		validation.Field(&c.Status, vre("The status is required"), validation.Min(1).Error("Status is Invalid"), validation.Max(2).Error("Status is Invalid")),
	)
}

func (s *Svc) validateMsg(w http.ResponseWriter, r *http.Request, err error, form Category, errEmp map[string]string, tmpl string) error {
	vErrs := map[string]string{}
	if e, ok := err.(validation.Errors); ok {
		if len(e) > 0 {
			for key, value := range e {
				vErrs[key] = value.Error()
			}
		}
	}

	if errEmp != nil {
		vErrs = errEmp
	}

	data := CategoryTempData{
		CSRFField:  csrf.TemplateField(r),
		Form:       form,
		FormErrors: vErrs,
		FormAction: common.CategoryCreatePath,
	}

	s.loadCategoryTemplate(w, r, data, tmpl)
	return nil
}

func getVErrs(err string) map[string]string {
	vErrs := map[string]string{}
	for _, v := range strings.Split(err, "; ") {
		val := strings.Split(v, ": ")
		vErrs[strings.Title(val[0])] = val[1]
	}

	return vErrs
}
