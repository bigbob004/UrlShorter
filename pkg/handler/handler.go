package handler

import (
	"UrlShorter/pkg/service"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Hnd struct {
	Handler service.Storage
}

func (s *Hnd) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	longURL := r.PostForm.Get("url")
	if _, err := http.Get(longURL); err != nil {
		fmt.Fprintf(w, longURL)
	} else {
		short := s.Handler.Save(longURL)
		fmt.Fprintf(w, "Short url: "+"http://localhost:8080/"+short)
	}

}

func (s *Hnd) GetOgUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	longURL, err := s.Handler.Get(params["hash"])
	fmt.Fprintf(w, "OG url : "+longURL)

	if err != nil {
		fmt.Fprintf(w, "Bad short url")
	} /*else {
		http.Redirect(w, r, longURL, 301)
	}*/
}
