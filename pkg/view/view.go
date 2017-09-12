package view

import (
	"net/http"

	"github.com/Chayanon/gonews/pkg/model"
)

type IndexData struct {
	List []*model.News
}

// Index renders index view
func Index(w http.ResponseWriter, data *IndexData) {
	render(tpIndex, w, data)
}

// News renders news view
func News(w http.ResponseWriter, data *model.News) {
	render(tpNews, w, data)
}

// AdminLogin reders admin login view
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

type AdminListData struct {
	List []*model.News
}

// AdminList reders admin login view
func AdminList(w http.ResponseWriter, data *AdminListData) {
	render(tpAdminList, w, data)
}

// AdminCreate reders admin login view
func AdminCreate(w http.ResponseWriter, data interface{}) {
	render(tpAdminCreate, w, data)
}

// AdminEdit reders admin login view
func AdminEdit(w http.ResponseWriter, data *model.News) {
	render(tpAdminEdit, w, data)
}
