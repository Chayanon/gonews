package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Chayanon/gonews/pkg/app"
	"github.com/Chayanon/gonews/pkg/model"
)

const (
	// port     = ":8080"
	// mongoURL = "mongodb://127.0.0.1:27017"
	mongoURL = "mongodb://gonews:chayanon@ds131854.mlab.com:31854/gonews"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	app.Mount(mux)
	err := model.Init(mongoURL)
	if err != nil {
		log.Fatalf("can not init model; %v", err)
	}
	http.ListenAndServe(":"+port, mux)
}
