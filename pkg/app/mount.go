package app

import (
	"net/http"
)

// Mound mounts handlers to mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index)
	mux.Handle("/upload/", http.StripPrefix("/upload", http.FileServer(http.Dir("upload")))) // list all news
	mux.Handle("/news/", http.StripPrefix("/news", http.HandlerFunc(newsView)))              // /news/:path

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/login", adminLogin)   // /admin/login
	adminMux.HandleFunc("/list", adminList)     // /admin/list
	adminMux.HandleFunc("/create", adminCreate) // /admin/create
	adminMux.HandleFunc("/edit", adminEdit)     // /admin/edit

	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))
}

func onlyAdmin(h http.Handler) http.Handler {
	return h
}
