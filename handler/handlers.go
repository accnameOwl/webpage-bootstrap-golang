package handlers

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var htmltmpl *template.Template

func init() {
	htmltmpl = template.Must(template.ParseGlob("./web/*.html"))
}

//Homepage respond homepage location to http.ResponseWriter
func Homepage(response http.ResponseWriter, request *http.Request, Param httprouter.Params) {
	htmltmpl.ExecuteTemplate(response, "index.html", nil)
}

func About(response http.ResponseWriter, request *http.Request, Param httprouter.Params) {
	htmltmpl.ExecuteTemplate(response, "about.html", nil)
}
