package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/raymondgitonga/matching-service/internal/adapters/db"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"log"
	"strings"
	"testing"
	"time"
)

func TestRepository_GetPartner(t *testing.T) {
	dbCLient, postgres := SetupTestDatabase()
	defer destroyDB(postgres)

	repository := NewRepository(dbCLient)
	partner, err := repository.GetPartner(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "Cummerata, Wolff and Hauck", partner.Name)
}

func SetupTestDatabase() (*sql.DB, *testcontainers.LocalDockerCompose) {
	log.Println("Starting postgres container...")
	postgres := testcontainers.NewLocalDockerCompose([]string{"/Users/raymondgitonga/Projects/matching-service/docker-compose.yml"},
		strings.ToLower(uuid.New().String()))

	postgres.WithCommand([]string{"up", "-d"}).Invoke()
	time.Sleep(20 * time.Second)

	postgresURL := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	log.Printf("Postgres container started, running at:  %s\n", postgresURL)

	conn, err := db.NewClient(context.Background(), postgresURL)
	if err != nil {
		log.Fatal("connect:", err)
	}

	if err := db.RunMigrations(conn, "postgres"); err != nil {
		log.Fatal("runMigrations:", err)
	}

	return conn, postgres
}

func destroyDB(compose *testcontainers.LocalDockerCompose) {
	compose.Down()
}
