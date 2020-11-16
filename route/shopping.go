package route

import (
	"github.com/ilhamrobyana/bts-test/mwcustom"
	"github.com/ilhamrobyana/bts-test/shopping"
	"github.com/labstack/echo"
)

func shoppingRoute(e *echo.Echo) {
	g := e.Group("/api/shopping")
	g.Use(mwcustom.Authorization)
	g.POST("/", shopping.Create)
	g.GET("/", shopping.GetAll)
	g.GET("/:id", shopping.GetByID)
	g.PUT("/:id", shopping.Update)
	g.DELETE("/:id", shopping.Delete)
}
