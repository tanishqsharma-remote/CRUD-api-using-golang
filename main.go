package main

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"sum/database"
	"sum/handler"
)

func main() {

	/*rows, err := db.Query("SELECT * FROM emp")
	fmt.Println(rows)*/
	db := database.DBconnect()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	m.Up()

	http.HandleFunc("/", handler.GETemp)
	http.HandleFunc("/post", handler.POSTemp)
	http.HandleFunc("/del", handler.DELemp)
	http.HandleFunc("/put", handler.PUTemp)
	err1 := http.ListenAndServe(":8080", nil)
	if err1 != nil {
		log.Fatalf("Error")
		return
	}
}
