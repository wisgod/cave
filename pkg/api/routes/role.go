package routes

import (
	"github.com/cave/pkg/api/controllers/role"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes Role
func RoleRoutes(api fiber.Router) {

	route := api.Group("/role")
	route.Get("/", role.GetAll)
	route.Post("/", role.CreateNew)
	route.Delete("/:id", role.DeleteSingle)
	route.Delete("/", role.DeleteAll)
	route.Get("/:id", role.GetSingle)
	route.Put("/:id", role.UpdateSingle)
	
}
