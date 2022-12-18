package service_test

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/raymondgitonga/joins-vs-multiple-queries/internal/adapters/db"
	"github.com/raymondgitonga/joins-vs-multiple-queries/internal/core/repository"
	"github.com/raymondgitonga/joins-vs-multiple-queries/internal/core/service"
	"github.com/stretchr/testify/assert"

	"os"
	"testing"
)

func BenchmarkProductService_GetProduct(b *testing.B) {
	err := godotenv.Load("../../../.env")
	assert.NoError(b, err)

	appConfigs, err := db.NewAppConfigs(
		os.Getenv("DB_CONNECTION_URL"),
		os.Getenv("DB_NAME"),
	)
	assert.NoError(b, err)

	dbClient, err := appConfigs.NewClient(context.Background(), appConfigs.DbURL)
	assert.NoError(b, err)

	err = appConfigs.RunMigrations(dbClient)
	assert.NoError(b, err)

	repo, err := repository.NewProductRepository(dbClient)
	assert.NoError(b, err)

	serv := service.NewProductService(repo)

	b.Run("Performance using joins", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			product, err := serv.GetProductJoin(context.Background(), 1)

			assert.NoError(b, err)
			assert.Equal(b, 2, product.CategoryID)
		}
	})

	b.Run("Performance without joins sync", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			product, err := serv.GetProductSync(context.Background(), 1)

			assert.NoError(b, err)
			assert.Equal(b, 2, product.CategoryID)
		}
	})

	b.Run("Performance without joins async", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			product, err := serv.GetProductAsync(context.Background(), 1)

			assert.NoError(b, err)
			assert.Equal(b, 2, product.CategoryID)
		}
	})
}
