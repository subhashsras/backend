package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type Item struct {
// 	ID   uint64 `json:"id"`
// 	Name string `json:"name"`
// }

// var ConnectionURI = fmt.Sprintf(
// 	"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 	"root",
// 	"12345678",
// 	"localhost",
// 	"3306",
// 	"crud",
// )

// var dbconn, err = gorm.Open(
// 	mysql.Open(ConnectionURI),
// 	&gorm.Config{},
// )

// func GetItems(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var items []Item
// 	err := dbconn.Table("items").Find(&items)
	
// 	if err.Error != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("Cannot get item"))
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(items)
// }

// func GetItem(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	idToGet, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	var item Item
// 	err := dbconn.Table("items").Where("id = ?", idToGet).Find(&item)
// 	if err.Error != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("Cannot get item"))
// 		return
// 	}
// 	if item.ID == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte("Item not found"))
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(item)

// }

// func CreateItem(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var item Item
// 	err := json.NewDecoder(r.Body).Decode(&item)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Cannot create item"))
// 		return
// 	}

// 	dbconn.Create(item)

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Item created successfully"))

// }

// // func DeleteItem(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "application/json")

// // 	idToGet, _ := strconv.Atoi(mux.Vars(r)["id"])

// // 	for i, item := range items {
// // 		if uint64(idToGet) == item.ID {
// // 			items = append(items[:i], items[i+1:]...)
// // 			w.WriteHeader(http.StatusOK)
// // 			w.Write([]byte("Item deleted successfully"))
// // 			return
// // 		}
// // 	}
// // 	w.WriteHeader(http.StatusNotFound)
// // 	w.Write([]byte("Item not found"))
// // }

// func main() {

// 	router := mux.NewRouter()

// 	dbconn.AutoMigrate(&Item{})
// 	router.HandleFunc("/items", GetItems).Methods("GET")
// 	router.HandleFunc("/item/{id}", GetItem).Methods("GET")
// 	router.HandleFunc("/item", CreateItem).Methods("POST")
// 	//router.HandleFunc("/item/{id}", DeleteItem).Methods("DELETE")

// 	http.ListenAndServe(":8000", router)
// }
