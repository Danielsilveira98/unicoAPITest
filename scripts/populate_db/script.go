package main

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/Danielsilveira98/unicoAPITest/internal/domain"
	"github.com/Danielsilveira98/unicoAPITest/internal/pkg/repository"
	"github.com/Danielsilveira98/unicoAPITest/internal/pkg/streetmarket"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	ctx := context.Background()

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
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, os.Getenv("MIGRATIONS_PATH")); err != nil {
		panic(err)
	}

	repo := repository.NewStreetMarketRepository(db)
	srv := streetmarket.NewWriter(repo, uuid.NewString)

	dataPath := os.Getenv("DATA_PATH")
	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Printf("Processing file %s\n", file.Name())
		sms, err := processFile(fmt.Sprintf("%s/%s", dataPath, file.Name()))
		if err != nil {
			fmt.Printf("Error processing file %s. Err: %v\n", file.Name(), err)
			continue
		}

		for _, sm := range sms {
			if err := sm.Validate(); err != nil {
				fmt.Println(err)
			} else {
				id, err := srv.Create(ctx, sm)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(id)
				}
			}
		}
	}
}

func processFile(path string) ([]domain.StreetMarketCreateInput, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	// Skip first row (line)
	row1, err := bufio.NewReader(csvFile).ReadSlice('\n')
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	_, err = csvFile.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	reader := csv.NewReader(csvFile)
	reader.LazyQuotes = true

	csvLines, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	fmt.Println("Successfully Read CSV file")

	sms := []domain.StreetMarketCreateInput{}

	for _, line := range csvLines {
		if len(line) <= 1 {
			continue
		}

		var long, lat float64

		if long, err = strconv.ParseFloat(line[1], 32); line[1] != "" && err != nil {
			fmt.Printf("long : %s \n", line[1])
			fmt.Println(err)
			return nil, fmt.Errorf("%w", err)
		}
		if lat, err = strconv.ParseFloat(line[2], 32); line[2] != "" && err != nil {
			fmt.Printf("lat : %s \n", line[2])
			fmt.Println(err)
			return nil, fmt.Errorf("%w", err)
		}

		sm := domain.StreetMarketCreateInput{
			Long:          long,
			Lat:           lat,
			SectCens:      line[3],
			Area:          line[4],
			IDdist:        line[6],
			District:      line[6],
			IDSubTH:       line[7],
			SubTownHall:   line[8],
			Region5:       line[9],
			Region8:       line[10],
			Name:          line[11],
			Register:      line[12],
			Street:        line[13],
			Number:        line[14],
			Neighborhood:  line[15],
			AddrExtraInfo: line[16],
		}
		sms = append(sms, sm)
	}

	return sms, nil
}
