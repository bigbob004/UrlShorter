package handler

import (
	"UrlShorter/internal/service"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type Hnd struct {
	Serivce service.Storage
}

type ViewData struct {
	Err  error
	Data string
}

func (s *Hnd) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Panicf("/CreateShortUrl, err: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmpl, err := template.ParseFiles("./forms/shorted_url_form.html")
	if err != nil {
		log.Panicf("/CreateShortUrl, err: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	longURL := r.FormValue("url")
	if _, err := http.Get(longURL); err != nil {
		log.Printf("/CreateShortUrl, err: %s", err)
		//TODO: обработка ошибок на форме
		fmt.Fprintf(w, fmt.Sprintf("url: %s isn't clicable", longURL))
	} else {
		short := s.Serivce.Save(longURL)
		tmpl.Execute(w, ViewData{
			Data: fmt.Sprintf("http://localhost:8080/%s", short),
		})
	}

}

func (s *Hnd) GetOgUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	longURL, err := s.Serivce.Get(params["hash"])
	//fmt.Fprintf(w, "OG url : "+longURL)

	if err != nil {
		fmt.Fprintf(w, "Bad short url")
	} else {
		http.Redirect(w, r, longURL, 301)
	}
}
