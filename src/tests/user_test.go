package tests

import (
	"encoding/json"
	"github.com/bxcodec/faker/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"karabayyazilim/src/config"
	"karabayyazilim/src/models"
	"karabayyazilim/src/routes"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

type User struct {
	Name  string `faker:"name"`
	Email string `faker:"email"`
}

func NewUser() User {
	user := User{}
	err := faker.FakeData(&user)
	if err != nil {
		panic(err)
	}
	return user
}

var apiEndpoint = "/api/users"
var app *fiber.App

func TestMain(m *testing.M) {
	config.AppConfig()
	app = fiber.New()
	routes.ApiRoute(app)
	m.Run()
}

func createUser() models.User {
	user := NewUser()
	userData, err := json.Marshal(user)
	if err != nil {
		panic("Failed to marshal user: " + err.Error())
	}

	userReader := strings.NewReader(string(userData))

	req := httptest.NewRequest("POST", apiEndpoint, userReader)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != 201 {
		panic("Failed to create user")

	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	var responseData models.User

	err = json.Unmarshal(bodyBytes, &responseData)

	return responseData
}

func TestUserList(t *testing.T) {
	req := httptest.NewRequest("GET", apiEndpoint+"?page=1&pageSize=10", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestUserCreate(t *testing.T) {
	user := NewUser()
	userData, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal user: %v", err)
	}

	userReader := strings.NewReader(string(userData))

	req := httptest.NewRequest("POST", apiEndpoint, userReader)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, 201, resp.StatusCode)
}

func TestUserFindById(t *testing.T) {
	user := createUser()

	userId := strconv.Itoa(int(user.ID))

	req := httptest.NewRequest("GET", apiEndpoint+"/"+userId, nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestUserUpdate(t *testing.T) {
	user := createUser()

	if user.ID != 0 {
		log.Debug("User ID: ", user.ID)
	}

	userId := strconv.Itoa(int(user.ID))

	newUser := models.User{
		Name: "Updated Name",
	}

	userData, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("Failed to marshal user: %v", err)
	}

	userReader := strings.NewReader(string(userData))

	req := httptest.NewRequest("PUT", apiEndpoint+"/"+userId, userReader)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestUserDelete(t *testing.T) {
	user := createUser()

	userId := strconv.Itoa(int(user.ID))

	req := httptest.NewRequest("DELETE", apiEndpoint+"/"+userId, nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}
