package service

import (
	"github.com/gofiber/fiber/v2"
	"karabayyazilim/src/internal/config"
	"karabayyazilim/src/internal/models"
	"karabayyazilim/src/pkg/paginate"
)

type UserService interface {
	List(paginator paginate.Paginator) []models.User
	Create(c *fiber.Ctx) models.User
	FindById(id int) (models.User, error)
	Update(id int, c *fiber.Ctx) models.User
	Delete(id int) bool
}

type User struct {
}

var db = config.Database()

func (u *User) List(paginator paginate.Paginator) []models.User {
	var users []models.User
	/*	paginator.DB = db
		db = paginator.Paginate()*/
	db.Find(&users)

	return users
}

func (u *User) Create(c *fiber.Ctx) models.User {
	user := models.User{}

	err := c.BodyParser(&user)
	if err != nil {
		return models.User{}
	}

	db.Create(&user)

	return user
}

func (u *User) FindById(id int) (models.User, error) {
	var user models.User
	res := db.First(&user, id)

	if res.Error != nil {
		return models.User{}, fiber.ErrNotFound
	}

	return user, nil
}

func (u *User) Update(id int, c *fiber.Ctx) models.User {
	user := models.User{}
	db.First(&user, id)

	err := c.BodyParser(&user)
	if err != nil {
		return models.User{}
	}

	db.Save(&user)

	return user
}

func (u *User) Delete(id int) bool {
	user := models.User{}
	db.Unscoped().Delete(&user, id)

	return true
}
