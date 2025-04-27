package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	//"github.com/go-delve/delve/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/szaluzhanskaya/Innopolis/chain-service/config"
	v1 "github.com/szaluzhanskaya/Innopolis/chain-service/internal/controller/http"
	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/repo/postgres"
	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/usecase"
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

	// загрузка конфигурации
	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Fatal("ошибка загрузки конфигурации", err)
	}

	connURL := "postgres://" + cfg.DB.Username + ":" + cfg.DB.Password + "@localhost:" + strconv.Itoa(cfg.DB.Port) + "/" + cfg.DB.DBname

	dbPool, err := initDb(connURL)
	if err != nil {
		log.Fatal("ошибка подключения к базе данных", err)
	}
	defer dbPool.Close()

	// port := "8080" //TODO: create ENV variable

	repo := postgres.New(dbPool)
	service := usecase.New(repo)
	handler := v1.New(service)

	http.HandleFunc("/ping", v1.PingHandler)
	http.HandleFunc("/delete-chain/{uuid}", handler.DeleteMessageChain)
	http.HandleFunc("/create-chain", handler.CreateMessageChain)

	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":"+cfg.AppConfig.Port, nil))

}

func initDb(connURL string) (*pgxpool.Pool, error) {

	dbPool, err := pgxpool.New(context.Background(), connURL)
	if err != nil {
		return nil, err
	}
	return dbPool, nil
}
