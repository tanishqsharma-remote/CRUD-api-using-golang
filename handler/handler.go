package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sum/database"
	"sum/model"
)

func POSTemp(w http.ResponseWriter, r *http.Request) {
	db := database.DBconnect()
	var item model.Emp
	json.NewDecoder(r.Body).Decode(&item)

	querr := "Insert into emp(id,name,age) values($1,$2,$3)"

	_, err := db.Exec(querr, item.Id, item.Name, item.Age)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
}

func GETemp(w http.ResponseWriter, r *http.Request) {
	db := database.DBconnect()

	rows, err := db.Query("SELECT * FROM emp order by id")
	if err != nil {
		log.Fatal(err)
	}

	var items []model.Emp

	for rows.Next() {
		var item model.Emp
		rows.Scan(&item.Id, &item.Name, &item.Age)
		items = append(items, item)
	}

	itemsBytes, _ := json.MarshalIndent(items, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(itemsBytes)

	defer rows.Close()
	//defer db.Close()
}

func PUTemp(w http.ResponseWriter, r *http.Request) {
	db := database.DBconnect()
	var item model.Emp
	json.NewDecoder(r.Body).Decode(&item)
	querr := "update emp set name=$1,age=$2 where id=$3"
	_, err := db.Exec(querr, item.Name, item.Age, item.Id)
	if err != nil {
		log.Fatal(err)
	}

}

func DELemp(w http.ResponseWriter, r *http.Request) {
	db := database.DBconnect()

	//Name := chi.URLParam(r, "Name")
	var item model.Emp
	json.NewDecoder(r.Body).Decode(&item)
	querr := "delete from emp where id=$1"
	_, err := db.Exec(querr, item.Id)
	if err != nil {
		log.Fatal(err)
	}

}
