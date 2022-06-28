package main

import (
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Qty   int    `json:"qty"`
}

type Place struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

var items []Item
var places []Place

var db *gorm.DB

func main() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:root00@172.16.16.225:5432/testDB"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Item{}, Place{})
	// if err != nil {
	// 	panic(err)
	// }
	router := http.NewServeMux()
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte(`"{"Message": "sukses"}"`))
	// })

	router.HandleFunc("/get-item", GetItem)
	router.HandleFunc("/update-item", UpdateItem)
	router.HandleFunc("/create-item", CreateItem)
	router.HandleFunc("/delete-item", DeleteItem)

	http.ListenAndServe(":4000", router)
}
