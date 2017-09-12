package app

import (
	"net/http"

	"github.com/Chayanon/gonews/pkg/model"
	"github.com/Chayanon/gonews/pkg/view"
)

func newsView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	n, err := model.GetNews(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.News(w, n)
}
