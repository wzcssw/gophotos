package controllers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// Templates 模板全局变量
var Templates map[string]*template.Template

// Intialize Intialize
func Intialize() {
	allPath, _ := filepath.Glob("./pages/*/*.html")
	Templates = make(map[string]*template.Template)
	for _, t := range allPath {
		temp := template.Must(template.ParseFiles("./" + t))
		Templates["./"+t] = temp
	}
}

func isPresent(r *http.Request, str string) bool {
	param := r.URL.Query()[str]
	return len(param) > 0
}
