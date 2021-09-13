package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConnString   = "root@/db_test"
	dbMaxIdleConns = 4
	dbMaxConns     = 100
	totalWorker    = 100
)

func main() {
	start := time.Now()
	db, err := openDBCon()
	if err != nil {
		log.Fatal(err.Error())
	}

	totals := 5000
	worker := 100
	buffer := totals + 1
	jobs := make(chan int, buffer)
	results := make(chan int, buffer)
	wg := &sync.WaitGroup{}

	for w := 0; w < worker; w++ {
		wg.Add(1)
		go doJob(w, jobs, results, db, wg)
	}

	for i := 0; i < totals; i++ {
		jobs <- i
	}
	close(jobs)
	for a := 0; a < totals; a++ {
		result := <-results
		log.Println("Data Telah Disimpan :", result)
	}
	wg.Wait()

	elsapsed := time.Since(start)
	fmt.Println("Waktu :", elsapsed)
}
func openDBCon() (*sql.DB, error) {
	log.Println("=> Open Connection")

	db, err := sql.Open("mysql", dbConnString)

	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(dbMaxConns)
	db.SetMaxIdleConns(dbMaxIdleConns)
	return db, nil
}
func doJob(wId int, jobs <-chan int, results chan<- int, con *sql.DB, wg *sync.WaitGroup) {
	for no := range jobs {
		log.Printf("=======Worker ID: %d ===========\n", wId)
		for {
			var outerError error

			func(outerError *error) {
				defer func() {
					if err := recover(); err != nil {
						*outerError = fmt.Errorf("%v", err)
					}
				}()
				log.Println("Terima data ", no)
				stmt, err := con.Prepare(`
						insert into
						tbl_test(
							nomor
						) values (
							?
						)
					`)
				if err != nil {
					log.Println(err.Error())
				}
				defer stmt.Close()
				res, err := stmt.Exec(no)
				if err != nil {
					log.Println(err.Error())
				}
				insert, _ := res.RowsAffected()
				if insert > 0 {
					results <- no
				}
			}(&outerError)
			if outerError == nil {
				break
			}
		}
	}
	wg.Done()
}
