package app

import (
	"net/http"

	"github.com/Chayanon/gonews/pkg/model"
)

// Mound mounts handlers to mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index)
	mux.Handle("/upload/", http.StripPrefix("/upload", http.FileServer(http.Dir("upload")))) // list all news
	mux.Handle("/news/", http.StripPrefix("/news", http.HandlerFunc(newsView)))              // /news/:path
	mux.HandleFunc("/login", adminLogin)                                                     // /admin/login
	mux.HandleFunc("/register", adminRegister)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/logout", adminLogout)
	adminMux.HandleFunc("/list", adminList)     // /admin/list
	adminMux.HandleFunc("/create", adminCreate) // /admin/create
	adminMux.HandleFunc("/edit", adminEdit)     // /admin/edit

	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))
}

func onlyAdmin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("user")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			// http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		userID := cookie.Value
		ok, err := model.CheckUserID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !ok {
			http.Redirect(w, r, "/", http.StatusFound)
			// http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}
