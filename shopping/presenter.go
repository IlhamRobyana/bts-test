package shopping

import (
	"net/http"
	"strconv"

	"github.com/ilhamrobyana/bts-test/entity"
	"github.com/ilhamrobyana/bts-test/storage"
	"github.com/labstack/echo"
)

var coreInstance *core

func Create(c echo.Context) error {
	r := new(entity.ShoppingCreateRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	shopping := r.Shopping

	shoppingCore := getCore()
	createdShopping, err := shoppingCore.create(shopping)

	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	return c.JSON(http.StatusCreated, createdShopping)
}

func GetAll(c echo.Context) error {
	shoppingCore := getCore()
	shoppingList, err := shoppingCore.getAll()
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	return c.JSON(http.StatusOK, shoppingList)
}

func GetByID(c echo.Context) error {
	shoppingID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	shoppingCore := getCore()
	shopping, err := shoppingCore.getByID(shoppingID)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	return c.JSON(http.StatusOK, shopping)
}

func Update(c echo.Context) error {
	shoppingID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	r := new(entity.ShoppingCreateRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}
	shopping := r.Shopping

	shoppingCore := getCore()
	updatedShopping, err := shoppingCore.update(shoppingID, shopping)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	return c.JSON(http.StatusOK, updatedShopping)
}

func Delete(c echo.Context) error {
	shoppingID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	shoppingCore := getCore()
	err := shoppingCore.delete(shoppingID)

	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func getCore() (c *core) {
	c = coreInstance

	if c == nil {
		c = new(core)
		shoppingStorage, _ := storage.GetShoppingStorage(storage.Postgre)

		c.shoppingStorage = shoppingStorage
		coreInstance = c
	}

	return
}
