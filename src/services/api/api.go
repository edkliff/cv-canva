package api

import (
	"github.com/edkliff/cv-canva/src/models/config"
	"github.com/edkliff/cv-canva/src/modules/logger"
	"github.com/sirupsen/logrus"
	"html/template"
	"log"
	"net/http"
)

type APIServer struct {
	Logger *logrus.Entry
	Host   string
}

func NewAPIServer(c *config.Config) (*APIServer, error) {
	return &APIServer{
		Logger: logger.InitLog(0),
	}, nil
}

func (a *APIServer) FaviconHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./img/favicon.ico")
}

func (a *APIServer) CV(res http.ResponseWriter, req *http.Request) {
	a.Logger.Print("OP!")

	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	content := struct{}{}

	err = t.Execute(res, content)
	if err != nil {
		log.Println(err)
	}
}
