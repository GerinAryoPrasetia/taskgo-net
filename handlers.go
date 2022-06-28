package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetItem(w http.ResponseWriter, r *http.Request) { //G besar jika ingin dibaca diluar directory
	if r.Method == "GET" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(items)
		// iData, err := json.Marshal(items)
		// if err != nil {
		// 	w.Write([]byte("ERROR"))
		// }
		// w.Write(iData)
	}
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if r.Method == "POST" {
		json.NewDecoder(r.Body).Decode(&item)
		// items = append(items, item)
		result := db.Create(&item)
		fmt.Println(result)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(item)
	} else {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))
	if r.Method == "PUT" {
		for index, item := range items {
			json.NewDecoder(r.Body).Decode(&item)
			if item.ID == id {
				items[index].ID = item.ID
				items[index].Name = item.Name
				w.Write([]byte("Success to update item"))
			}
		}
	} else {
		w.Header().Set("Allow", "PUT")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := strconv.Atoi(query.Get("id"))

	if r.Method == "DELETE" {
		for index, item := range items {
			if item.ID == id {
				items = append(items[:index], items[index+1:]...)
				w.Write([]byte("Success to delete item"))
			}
		}
	} else {
		w.Header().Set("Allow", "DELETE")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}
}
