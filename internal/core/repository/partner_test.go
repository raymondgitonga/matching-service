package repository_test

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/raymondgitonga/matching-service/internal/adapters/db"
	"github.com/raymondgitonga/matching-service/internal/core/repository"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"strings"
	"testing"
)

func TestRepository_GetPartner(t *testing.T) {
	dbCLient, postgres := setupTestDatabase(t)
	defer destroyDB(postgres)

	repo, err := repository.NewPartnerRepository(dbCLient)
	assert.NoError(t, err)

	partner, err := repo.GetPartner(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "Cummerata, Wolff and Hauck", partner.Name)
}

func TestRepository_GetPartners(t *testing.T) {
	dbCLient, postgres := setupTestDatabase(t)
	defer destroyDB(postgres)

	testCases := []struct {
		material     string
		partnersSize int
	}{
		{
			material:     "wood",
			partnersSize: 9,
		},
		{
			material:     "tiles",
			partnersSize: 7,
		},
		{
			material:     "carpet",
			partnersSize: 10,
		},
	}

	for _, tc := range testCases {
		t.Run("Test partner number by material: "+tc.material, func(t *testing.T) {
			repo, err := repository.NewPartnerRepository(dbCLient)
			assert.NoError(t, err)

			partners, err := repo.GetPartners(context.Background(), tc.material)
			assert.NoError(t, err)
			assert.Equal(t, tc.partnersSize, len(*partners))
		})
	}

}

func setupTestDatabase(t *testing.T) (*sql.DB, *testcontainers.LocalDockerCompose) {
	postgres := testcontainers.NewLocalDockerCompose([]string{"../../../docker-compose-test.yml"},
		strings.ToLower(uuid.New().String()))
	postgres.WithCommand([]string{"up", "-d"}).Invoke()
	postgresURL := "postgres://postgres:postgres@localhost:9876/postgres?sslmode=disable"

	conn, err := db.NewClient(context.Background(), postgresURL)
	assert.NoError(t, err)

	err = db.RunMigrations(conn, "postgres")
	assert.NoError(t, err)

	return conn, postgres
}

func destroyDB(compose *testcontainers.LocalDockerCompose) {
	compose.Down()
}
