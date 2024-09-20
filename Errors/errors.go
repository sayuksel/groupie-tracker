package Errors

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func Error500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	Error_500 := filepath.Join("..", "templates", "500.html")

	Error_500_parse, err := template.ParseFiles(Error_500)

	if err != nil {
		fmt.Fprint(w, "500 error")
		return
	}

	Error_500_parse.ExecuteTemplate(w, "500.html", nil)
}
func Error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	Error_400 := filepath.Join("..", "templates", "404.html")

	Error_400_parse, err := template.ParseFiles(Error_400)

	if err != nil {
		fmt.Fprint(w, "404 error")
		return
	}

	Error_400_parse.ExecuteTemplate(w, "404.html", nil)
}
