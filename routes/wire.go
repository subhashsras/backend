package routes

import (
	"github.com/subhashsras/backend_service/controller"
	"github.com/subhashsras/backend_service/db"
	"github.com/subhashsras/backend_service/service"
)

// Repos
var (
	itemRepo db.ItemRepo
)

func InitRepos() {
	itemRepo = db.NewItemRepo()
}

// Services
var (
	itemService service.ItemService
)

func InitServices() {
	itemService = service.NewItemService(itemRepo)
}

// Handlers
var (
	itemController controller.ItemController
)

func InitHandlers() {
	InitRepos()
	InitServices()
	itemController = controller.NewItemController(itemService)
}