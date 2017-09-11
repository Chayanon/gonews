package app

import (
	"net/http"

	"github.com/Chayanon/gonews/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	view.Index(w, nil)
}
