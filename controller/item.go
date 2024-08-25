package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/subhashsras/backend_service/models"
	"github.com/subhashsras/backend_service/resources"
	"github.com/subhashsras/backend_service/service"
)

type ItemController struct {
	itemService service.ItemService
}

func NewItemController(itemService service.ItemService) ItemController {
	return ItemController{
		itemService: itemService,
	}
}

func (i ItemController) GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	items, err := i.itemService.GetItems()
	if err != nil {
		log.Println("Error in get items:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error in get items"))
		return
	}

	itemBytes, _ := json.Marshal(items)

	w.WriteHeader(http.StatusOK)
	w.Write(itemBytes)
}

func (i ItemController) GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemId, _ := strconv.Atoi(mux.Vars(r)["id"])

	item, err := i.itemService.GetItem(uint64(itemId))
	if err != nil {
		log.Println("Error in get item:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error in get item"))
		return
	}

	itemBytes, _ := json.Marshal(item)

	w.WriteHeader(http.StatusOK)
	w.Write(itemBytes)
}

func (i ItemController) CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var itemRes resources.Item
	err := json.NewDecoder(r.Body).Decode(&itemRes)
	if err != nil {
		log.Println("Error in decoding given payload:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error in decoding given payload"))
		return
	}

	err = i.itemService.CreateItem(models.Item(itemRes))
	if err != nil {
		log.Println("Error in create item:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error in create item"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Item created"))
}
