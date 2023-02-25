package main

import (
	"UrlShorter/internal/handler"
	"UrlShorter/internal/repository"
	"UrlShorter/internal/repository/db"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

const pathOfForm = "forms/short_url_form.html"

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
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, pathOfForm)
	})
	fileServer := http.FileServer(http.Dir("./forms"))
	router.PathPrefix("/res/").Handler(http.StripPrefix("/res/", fileServer))
	http.Handle("/", router)

	if os.Args[1] == "-InMemory" {
		log.Print("Запуск через встроенное хранилище")
		h.Serivce = &repository.MyRepo{ShortToLong: map[string]string{},
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

		h.Serivce = &db.DBRepo{DB: database}

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
