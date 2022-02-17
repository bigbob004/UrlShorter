package main

import (
	"UrlShorter/pkg/handler"
	"UrlShorter/pkg/repository"
	"UrlShorter/pkg/repository/db"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

var html = "<!DOCTYPE html>" +
	"<form action=\"/create\" method=\"POST\">" +
	"<input type=\"text\" name=\"url\" autofocus placeholder=\"URL\" required>" +
	"<input type=\"submit\" value=\"Получить короткую ссылку\"></form>"

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, html)
	if err != nil {
		return
	}
}

func main() {

	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if !(len(os.Args) == 2) {
		log.Fatal("use flags: -InMemory or -InDb")
	}

	var h handler.Hnd

	router := mux.NewRouter()
	router.HandleFunc("/create", h.CreateShortUrl).Methods("POST")
	router.HandleFunc("/{hash}", h.GetOgUrl).Methods("GET")
	router.HandleFunc("/", welcomeHandler)

	if os.Args[1] == "-InMemory" {
		log.Print("Запуск через встроенное хранилище")
		h.Handler = &repository.MyRepo{ShortToLong: map[string]string{},
			LongToShort: map[string]string{}}
	}

	if os.Args[1] == "-InDb" {
		log.Print("Запуск через БД")
		database, err := db.NewPostgresDB(db.Config{
			Host:     "localhost",
			Port:     "5432",
			Username: "postgres",
			Password: "qwerty",
			DBName:   "postgres",
			SSLMode:  "disable",
		})

		if err != nil {
			log.Fatalf("failed to initialize db: %s", err.Error())
		}

		h.Handler = &db.DBRepo{DB: database}

	}
	log.Print("http://localhost:8080")
	err1 := http.ListenAndServe(viper.GetString("port"), router)
	if err1 != nil {
		log.Print(err1.Error())
	}

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
