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

const testPartner = `insert into partner (id, name, location, speciality, radius, rating)
values (nextval('partner_seq'), 'Cummerata, Wolff and Hauck', point(51.73212999999999, -1.0831176441976451), '{
"wood": false,
"carpet": true,
"tiles": true
}', 1, 4.5);`

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
	time.Sleep(5 * time.Second)

	postgresURL := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	log.Printf("Postgres container started, running at:  %s\n", postgresURL)

	conn, err := db.NewClient(context.Background(), postgresURL)
	if err != nil {
		log.Fatal("connect:", err)
	}
	_, err = conn.Exec(testPartner)

	if err != nil {
		log.Printf("error executing query %s", err)
	}

	return conn, postgres
}

func destroyDB(compose *testcontainers.LocalDockerCompose) {
	compose.Down()
}
