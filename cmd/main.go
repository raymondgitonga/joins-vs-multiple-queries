package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/raymondgitonga/joins-vs-multiple-queries/internal/adapters/db"
	"github.com/raymondgitonga/joins-vs-multiple-queries/internal/core/repository"
	service2 "github.com/raymondgitonga/joins-vs-multiple-queries/internal/core/service"
	"log"
	"os"
)

type AppConfigs struct {
	dbURL  string
	dbName string
}

func NewAppConfigs(dbURL, dbName string) (*AppConfigs, error) {
	if dbURL == "" {
		return nil, fmt.Errorf("kindly provide dbURL")
	}
	if dbName == "" {
		return nil, fmt.Errorf("kindly provide dbName")
	}
	return &AppConfigs{
		dbURL:  dbURL,
		dbName: dbName,
	}, nil
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println(err)
		return
	}
	appConfigs, err := NewAppConfigs(
		os.Getenv("DB_CONNECTION_URL"),
		os.Getenv("DB_NAME"),
	)

	dbClient, err := db.NewClient(context.Background(), appConfigs.dbURL)
	if err != nil {
		log.Println(err)
		return
	}

	err = db.RunMigrations(dbClient, appConfigs.dbName)
	if err != nil {
		log.Println(err)
		return
	}

	repo, err := repository.NewProductRepository(dbClient)
	if err != nil {
		log.Println(err)
		return
	}

	sv := service2.NewProductService(repo)
	product, err := sv.GetProductJoin(context.Background(), 1)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("JOIN: ", product)

	product, err = sv.GetProductSync(context.Background(), 1)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("SYNC", product)
}
