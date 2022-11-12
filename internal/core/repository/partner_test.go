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

const insertPartner = `insert into partner (id, name, location, speciality, radius, rating)
values (1, 'Cummerata, Wolff and Hauck', point(51.73212999999999, -1.0831176441976451), '{
"wood": false,
"carpet": true,
"tiles": true
}', 1, 4.5);`
const createPartnerTable = `CREATE TABLE IF NOT EXISTS partner
(
    id         int PRIMARY KEY,
    name       varchar,
    location   point,
    speciality jsonb,
    radius     int,
    rating     decimal
);`

func TestRepository_GetPartner(t *testing.T) {
	dbCLient, postgres := SetupTestDatabase()
	defer destroyDB(postgres)

	repo := repository.NewRepository(dbCLient)
	partner, err := repo.GetPartner(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "Cummerata, Wolff and Hauck", partner.Name)
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

	if err != nil {
		log.Fatal("error migrating:", err)
	}

	_, err = conn.Exec(createPartnerTable)

	if err != nil {
		log.Printf("error executing query %s", err)
	}

	_, err = conn.Exec(insertPartner)

	if err != nil {
		log.Printf("error executing query %s", err)
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
