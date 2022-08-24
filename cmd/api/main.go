package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Danielsilveira98/unicoAPITest/internal/app/httphandler"
	"github.com/Danielsilveira98/unicoAPITest/internal/pkg/logger"
	"github.com/Danielsilveira98/unicoAPITest/internal/pkg/repository"
	"github.com/Danielsilveira98/unicoAPITest/internal/pkg/streetmarket"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func main() {
	db, err := setupDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = runMigrations(db)
	if err != nil {
		panic(err)
	}

	logger := logger.NewLogger(os.Stdout, true)

	streetMarketRepository := repository.NewStreetMarketRepository(db)

	writer := streetmarket.NewWriter(streetMarketRepository, uuid.NewString)
	eraser := streetmarket.NewEraser(streetMarketRepository)
	reader := streetmarket.NewReader(streetMarketRepository)

	pingHandler := httphandler.NewPingHandler()
	streeMarketEditHandler := httphandler.NewStreetMarketEditHandler(writer, logger)
	streeMarketCreateHandler := httphandler.NewStreetMarketCreateHandler(writer, logger)
	streeMarketDeleteHandler := httphandler.NewStreetMarketDeleteHandler(eraser, logger)
	streeMarketListHandler := httphandler.NewStreetMarketListHandler(reader, logger)

	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler.Handle).Methods(http.MethodGet)
	r.HandleFunc("/street_market", streeMarketListHandler.Handle).Methods(http.MethodGet)
	r.HandleFunc("/street_market", streeMarketCreateHandler.Handle).Methods(http.MethodPost)
	r.HandleFunc("/street_market/{street-market-id}", streeMarketDeleteHandler.Handle).Methods(http.MethodDelete)
	r.HandleFunc("/street_market/{street-market-id}", streeMarketEditHandler.Handle).Methods(http.MethodPatch)

	log.Fatal(http.ListenAndServe(":8000", r))
}

func setupDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open(os.Getenv("DB_DIALECT"), psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return db, nil
}

func runMigrations(db *sql.DB) error {
	if err := goose.SetDialect(os.Getenv("DB_DIALECT")); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := goose.Up(db, os.Getenv("MIGRATIONS_PATH")); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
