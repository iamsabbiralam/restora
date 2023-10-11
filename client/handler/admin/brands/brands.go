package brands

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
	braG "github.com/iamsabbiralam/restora/proto/v1/server/brand"
)

type Svc struct {
	*common.Server
}

type Brand struct {
	ID        string
	Name      string
	Status    braG.Status
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

type BrandTempData struct {
	CSRFField      template.HTML
	Form           Brand
	Data           []Brand
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

	mw.HandleFunc(common.BrandCreatePath, s.createBrandHandler).Methods("GET").Name("brand.create")
	mw.HandleFunc(common.BrandCreatePath, s.storeBrandHandler).Methods("POST").Name("brand.store")
	mw.HandleFunc(common.BrandEditPath, s.editBrandHandler).Methods("GET").Name("brand.edit")
	mw.HandleFunc(common.BrandEditPath, s.updateBrandHandler).Methods("POST").Name("brand.update")
	mw.HandleFunc(common.BrandDeletePath, s.deleteBrandHandler).Methods("GET").Name("brand.delete")
	mw.HandleFunc(common.BrandListPath, s.listBrandHandler).Methods("GET").Name("brand.list")
	return mw, nil
}

func (s *Svc) loadBrandTemplate(w http.ResponseWriter, r *http.Request, data BrandTempData, htmlFile string) {
	log := s.Logger.WithField("method", "handler.admin.brand.loadBrandTemplate")
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

func (c Brand) ValidateBrand(server *common.Server, r *http.Request, id string) error {
	vre := validation.Required.Error
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, vre("The brand name is required"), validation.Length(3, 100), validation.Match(regexp.MustCompile(common.TextValidation)).Error("Brand name is invalid, only text is allowed.")),
		validation.Field(&c.Status, vre("The status is required"), validation.Min(1).Error("Status is Invalid"), validation.Max(2).Error("Status is Invalid")),
	)
}

func getVErrs(err string) map[string]string {
	vErrs := map[string]string{}
	for _, v := range strings.Split(err, "; ") {
		val := strings.Split(v, ": ")
		vErrs[strings.Title(val[0])] = val[1]
	}

	return vErrs
}

func (s *Svc) validateMsg(w http.ResponseWriter, r *http.Request, err error, form Brand, errEmp map[string]string, tmpl string) error {
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

	data := BrandTempData{
		CSRFField:  csrf.TemplateField(r),
		Form:       form,
		FormErrors: vErrs,
		FormAction: common.CategoryCreatePath,
	}

	s.loadBrandTemplate(w, r, data, tmpl)
	return nil
}
