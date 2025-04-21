package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	//"github.com/go-delve/delve/pkg/config"
	"github.com/joho/godotenv"
	"github.com/szaluzhanskaya/Innopolis/chain-service/config"
	v1 "github.com/szaluzhanskaya/Innopolis/chain-service/internal/controller/http"
)

func main() {
	// загрузка значений из .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	// получаем значение переменной APP_ENVIRONMENT
	env := os.Getenv("APP_ENVIRONMENT")
	if env == "" {
		env = "local"
	}

	// загрузка кофнигурации
	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Fatal("ошибка загрузки конфигурации", err)
	}

	// port := "8080" //TODO: create ENV variable

	// Registers a handler for the /ping route
	http.HandleFunc("/ping", v1.PingHandler)

	// Starts the HTTP server on the port:8080
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":"+cfg.AppConfig.Port, nil))

}
