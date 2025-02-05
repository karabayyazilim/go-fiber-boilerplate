package controllers

import (
	"github.com/gofiber/fiber/v2"
	"karabayyazilim/src/internal/models"
	"karabayyazilim/src/internal/services"
	"karabayyazilim/src/pkg/paginate"
	"karabayyazilim/src/pkg/redis"
	"strconv"
	"time"
)

var userService service.UserService

func init() {
	userService = new(service.User)
}

func UserList(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))

	paginator := paginate.Paginator{Page: page, PageSize: pageSize}
	cacheKey := "user_list_" + strconv.Itoa(page) + "_" + strconv.Itoa(pageSize)

	users, _ := redis.Remember(c.Context(), cacheKey, 10*time.Minute, func() (interface{}, error) {
		users := userService.List(paginator)
		return users, nil
	})

	return c.JSON(fiber.Map{"data": users})
}

func UserCreate(c *fiber.Ctx) error {
	var user models.User
	user = userService.Create(c)
	return c.Status(fiber.StatusCreated).JSON(user)
}

func UserFindById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, err := userService.FindById(id)

	if err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(fiber.Map{
		"message": user,
	})
}

func UserUpdate(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user := userService.Update(id, c)
	return c.JSON(user)
}

func UserDelete(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user := userService.Delete(id)
	return c.JSON(fiber.Map{
		"message": user,
	})
}
