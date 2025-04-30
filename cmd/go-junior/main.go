package main

import (
	"junior/internal/config"
	"junior/internal/db"
	"junior/internal/handler"
	"junior/internal/repository"
	"junior/internal/service"
	"junior/pkg/logger"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()
	dbConn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	logger.InitLogger()

	personRepo := repository.NewPersonRepository(dbConn)
	personService := service.NewPersonService(personRepo)

	r := mux.NewRouter()
	handler := handler.NewHandler(personService)
	handler.InitRoutes(r)

	log.Println("Server started on port", cfg.APIPort)
	log.Fatal(http.ListenAndServe(":"+cfg.APIPort, r))
}
