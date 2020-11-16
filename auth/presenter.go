package auth

import (
	"net/http"

	"github.com/ilhamrobyana/bts-test/entity"
	"github.com/ilhamrobyana/bts-test/storage"
	"github.com/labstack/echo"
)

var coreInstance *core

func Signup(c echo.Context) error {
	r := new(entity.RegisterRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}
	user := r.User

	authCore := getCore()
	createdUser, err := authCore.signup(user)

	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	return c.JSON(http.StatusCreated, createdUser)
}

func Signin(c echo.Context) error {
	r := new(entity.LoginRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	if err := checkLogin(*r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
	}

	authCore := getCore()
	response, err := authCore.signin(r.Email, r.Password)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}
	return c.JSON(http.StatusOK, response)
}

func GetAll(c echo.Context) error {
	authCore := getCore()
	userList, err := authCore.getAll()
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}
	return c.JSON(http.StatusOK, userList)
}

func getCore() (c *core) {
	c = coreInstance

	if c == nil {
		c = new(core)
		userStorage, _ := storage.GetUserStorage(storage.Postgre)
		tokenStorage, _ := storage.GetTokenStorage(storage.Redis)

		c.userStorage = userStorage
		c.tokenStorage = tokenStorage
		coreInstance = c
	}

	return
}
