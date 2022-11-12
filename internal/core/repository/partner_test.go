package repository_test

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/raymondgitonga/matching-service/internal/adapters/db"
	"github.com/raymondgitonga/matching-service/internal/core/repository"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestRepository_GetPartner(t *testing.T) {
	dbCLient, postgres := SetupTestDatabase()
	defer destroyDB(postgres)

	repo := repository.NewPartnerRepository(dbCLient)
	partner, err := repo.GetPartner(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "Cummerata, Wolff and Hauck", partner.Name)
}

func TestRepository_GetPartners(t *testing.T) {
	dbCLient, postgres := SetupTestDatabase()
	defer destroyDB(postgres)

	testCases := []struct {
		speciality   string
		partnersSize int
	}{
		{
			speciality:   "wood",
			partnersSize: 9,
		},
		{
			speciality:   "tiles",
			partnersSize: 7,
		},
		{
			speciality:   "carpet",
			partnersSize: 10,
		},
	}

	for _, tc := range testCases {
		t.Run("Test partner number by speciality: "+tc.speciality, func(t *testing.T) {
			repo := repository.NewPartnerRepository(dbCLient)
			partners, err := repo.GetPartners(context.Background(), tc.speciality)

			assert.NoError(t, err)
			assert.Equal(t, tc.partnersSize, len(*partners))
		})
	}

}

func SetupTestDatabase() (*sql.DB, *testcontainers.LocalDockerCompose) {
	postgres := testcontainers.NewLocalDockerCompose([]string{getRootDir() + "/docker-compose-test.yml"},
		strings.ToLower(uuid.New().String()))

	postgres.WithCommand([]string{"up", "-d"}).Invoke()

	postgresURL := "postgres://postgres:postgres@localhost:9876/postgres?sslmode=disable"
	log.Printf("Postgres container started, running at:  %s\n", postgresURL)

	conn, err := db.NewClient(context.Background(), postgresURL)
	if err != nil {
		log.Fatal("connect:", err)
	}

	err = db.RunMigrations(conn, "postgres")
	if err != nil {
		log.Fatal("error migrating:", err)
	}
	return conn, postgres
}

func destroyDB(compose *testcontainers.LocalDockerCompose) {
	compose.Down()
}

func getRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../../..")
}
