package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

const (
	dbConnString   = "user=ismaldts dbname=employee sslmode=disable"
	dbMaxIdleConns = 10
	dbMaxConns     = 50
	totalWorker    = 10
	csvFile        = "majestic_million.csv"
	batchSize      = 1000
)

var dataHeaders = make([]string, 0)

func openDbConnection() (*sql.DB, error) {
	log.Println("=> open db connection")

	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbMaxConns)
	db.SetMaxIdleConns(dbMaxIdleConns)

	return db, nil
}

func openCsvFile() (*csv.Reader, *os.File, error) {
	log.Println("=> open csv file")

	f, err := os.Open(csvFile)
	if err != nil {
		return nil, nil, err
	}

	reader := csv.NewReader(f)
	return reader, f, nil
}

func dispatchWorkers(db *sql.DB, jobs <-chan [][]interface{}, wg *sync.WaitGroup) {
	for workerIndex := 0; workerIndex < totalWorker; workerIndex++ {
		go func(workerIndex int, db *sql.DB, jobs <-chan [][]interface{}, wg *sync.WaitGroup) {
			counter := 0

			for job := range jobs {
				doTheJob(workerIndex, counter, db, job)
				wg.Done()
				counter++
			}
		}(workerIndex, db, jobs, wg)
	}
}

func readCsvFilePerLineThenSendToWorker(csvReader *csv.Reader, jobs chan<- [][]interface{}, wg *sync.WaitGroup) {
	batch := make([][]interface{}, 0, batchSize)

	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if len(dataHeaders) == 0 {
			dataHeaders = row
			continue
		}

		rowOrdered := make([]interface{}, len(row))
		for i, each := range row {
			rowOrdered[i] = each
		}

		batch = append(batch, rowOrdered)
		if len(batch) >= batchSize {
			wg.Add(1)
			jobs <- batch
			batch = make([][]interface{}, 0, batchSize)
		}
	}

	if len(batch) > 0 {
		wg.Add(1)
		jobs <- batch
	}

	close(jobs)
}

func doTheJob(workerIndex, counter int, db *sql.DB, values [][]interface{}) {
	for {
		var outerError error
		func(outerError *error) {
			defer func() {
				if err := recover(); err != nil {
					*outerError = fmt.Errorf("%v", err)
				}
			}()

			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err.Error())
			}

			query := fmt.Sprintf("INSERT INTO millions (%s) VALUES %s",
				strings.Join(dataHeaders, ","),
				strings.Join(generateQuestionsMarks(len(values), len(dataHeaders)), ","),
			)

			args := make([]interface{}, 0, len(values)*len(dataHeaders))
			for _, row := range values {
				args = append(args, row...)
			}

			_, err = tx.ExecContext(context.Background(), query, args...)
			if err != nil {
				tx.Rollback()
				log.Fatal(err.Error())
			}

			err = tx.Commit()
			if err != nil {
				log.Fatal(err.Error())
			}
		}(&outerError)
		if outerError == nil {
			break
		}
	}

	if counter%10 == 0 {
		log.Println("=> worker", workerIndex, "inserted", counter*batchSize, "rows")
	}
}

func generateQuestionsMarks(batchSize, colSize int) []string {
	s := make([]string, batchSize)
	for i := 0; i < batchSize; i++ {
		placeholders := make([]string, colSize)
		for j := 0; j < colSize; j++ {
			placeholders[j] = fmt.Sprintf("$%d", i*colSize+j+1)
		}
		s[i] = fmt.Sprintf("(%s)", strings.Join(placeholders, ","))
	}
	return s
}

func main() {
	start := time.Now()

	db, err := openDbConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	csvReader, csvFile, err := openCsvFile()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csvFile.Close()

	jobs := make(chan [][]interface{}, totalWorker)
	wg := new(sync.WaitGroup)

	go dispatchWorkers(db, jobs, wg)
	readCsvFilePerLineThenSendToWorker(csvReader, jobs, wg)

	wg.Wait()

	duration := time.Since(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
}
