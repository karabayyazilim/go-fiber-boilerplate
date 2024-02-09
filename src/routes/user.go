package routes

import (
	"github.com/gofiber/fiber/v2"
	"karabayyazilim/src/controllers"
)

func UserRoutes(route fiber.Router) {
	user := route.Group("/users")

	user.Get("/", controllers.UserList)
	user.Post("/", controllers.UserCreate)
	user.Get("/:id", controllers.UserFindById)
	user.Put("/:id", controllers.UserUpdate)
	user.Delete("/:id", controllers.UserDelete)
}
