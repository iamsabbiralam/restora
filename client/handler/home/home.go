package guest

import (
	"net/http"

	"github.com/iamsabbiralam/restora/client/handler/common"
)

func (s *Svc) GetHomeHandler(w http.ResponseWriter, r *http.Request) {
	template := s.Server.LookupTemplate("home.html")
	if template == nil {
		errMsg := "unable to load template"
		s.Server.Logger.Error(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	/* data := common.PublicTemplateData{
		UserInfo: s.GetSessionUser(r),
	} */

	if err := template.Execute(w, nil); err != nil {
		s.Server.Logger.Infof("error with template execution: %+v", err)
		http.Redirect(w, r, common.ErrorPath, http.StatusSeeOther)
	}
}
