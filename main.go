package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/sushicombo/redrain-api/db/mysql"
	"github.com/sushicombo/redrain-api/repository/events"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Panic("error loading godotenv")
	}
	// Initiate database connection
	mysqlConfig, err := mysql.NewConfig()
	if err != nil {
		log.Panic("error: intiate database config")
	}

	db_conn, err := mysql.NewConnection(mysqlConfig)
	if err != nil {
		log.Panic("error: database connection")
	}

	eventRepo := events.InitEventRepo(db_conn)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if err := eventRepo.UpdateEventCounter(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte("data update successfully"))
		}
	})

	http.ListenAndServe(":8080", r)
}
