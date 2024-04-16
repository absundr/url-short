package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

type UrlDetails struct {
	Url string
}

var shortenedUrls map[string]string

func init() {
	shortenedUrls = make(map[string]string)
}

func main() {
	tmpl := template.Must(template.ParseFiles("form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := UrlDetails{
			Url: r.FormValue("url"),
		}

		id := uuid.NewString()
		shortenedUrls[id] = details.Url
		href := "http://localhost/" + id
		tmpl.Execute(w, struct{ ShortUrl string }{ShortUrl: href})
	})

	http.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		uid := r.URL.Path[1:]
		url, ok := shortenedUrls[uid]
		if !ok {
			http.NotFound(w, r)
		}

		http.Redirect(w, r, fmt.Sprintf("%v", url), http.StatusSeeOther)
	})

	http.ListenAndServe(":80", nil)
}
