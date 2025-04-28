package main

import (
    "log"
    "net/http"
    "junior/internal/config"
    "junior/internal/handler"
    "junior/pkg/logger"

    "github.com/gorilla/mux"
)

func main() {
    config.LoadConfig()
    logger.InitLogger()

    r := mux.NewRouter()

    handler.InitRoutes(r)

    log.Println("Server started on port", config.GetEnv("API_PORT"))
    http.ListenAndServe(":"+config.GetEnv("API_PORT"), r)
}
