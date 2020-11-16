package route

import (
	"github.com/ilhamrobyana/bts-test/auth"
	"github.com/ilhamrobyana/bts-test/mwcustom"
	"github.com/labstack/echo"
)

func authRoute(e *echo.Echo) {
	g := e.Group("/api/users")
	g.POST("/signup", auth.Signup)
	g.POST("/signin", auth.Signin)
	g.GET("/", auth.GetAll, mwcustom.Authorization)
}
