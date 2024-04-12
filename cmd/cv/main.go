package main

import (
	"github.com/edkliff/cv-canva/src/models/config"
	"github.com/edkliff/cv-canva/src/modules/logger"
	"github.com/edkliff/cv-canva/src/services/api"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
)

func main() {
	l := logger.InitLog(0)
	confPath, ok := os.LookupEnv("CONFIG")
	if !ok {
		confPath = "config-example.yaml"
	}
	c, err := config.ReadConfig(confPath)
	if err != nil {
		l.Fatal(err)
	}
	apiserver, err := api.NewAPIServer(c)
	if err != nil {
		apiserver.Logger.Fatal(err)
	}
	apiserver.Logger.Printf("api-server: %+v", *apiserver)
	r := chi.NewRouter()
	fs := http.FileServer(http.Dir("./static"))
	imgs := http.FileServer(http.Dir("./img"))

	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.Handle("/img/*", http.StripPrefix("/img/", imgs))
	r.Get("/", apiserver.CV)
	r.HandleFunc("/favicon.ico", apiserver.FaviconHandler)
	apiserver.Logger.Fatal(http.ListenAndServe(apiserver.Host, r))
}
