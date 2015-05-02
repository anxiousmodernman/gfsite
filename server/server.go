package server

import (
	"github.com/anxiousmodernman/gfsite/controllers"
	"net/http"
	"os"
	"text/template"
)

func main() {
	templates := loadTemplates()

	controllers.Register(templates)

	http.ListenAndServe(":8000", nil)
}

func loadTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths,
				basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}
