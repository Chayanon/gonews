package app

import (
	"net/http"

	"github.com/Chayanon/gonews/pkg/model"
	"github.com/Chayanon/gonews/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	list, err := model.ListNews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.Index(w, &view.IndexData{
		List: list,
	})
}
