package controllers

import (
	"bufio"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func Register(templates *template.Template) {

	router := mux.NewRouter()

	hc := new(homeController)
	hc.template = templates.Lookup("home.html")
	hc.loginTemplate = templates.Lookup("login.html")
	router.HandleFunc("/home", hc.get)
	router.HandleFunc("/login", hc.login)

	cc := new(categoriesController)
	cc.template = templates.Lookup("categories.html")
	router.HandleFunc("/categories", cc.get)

	http.Handle("/", router)

	http.HandleFunc("/js", serveResource)
	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
