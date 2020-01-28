package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/mystyle/db"
	"github.com/mystyle/schema"
)

func main() {
	// initia db
	h := db.PersonaHandler{}
	h.Initialize(&schema.Personas{})
	// router
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.GET("/personas", func(c echo.Context) error {
		return HandlerGetPersonas(c, h)
	})

	e.POST("/register", func(c echo.Context) error {
		return HandlerPostRegisterPersona(c, h)
	})

	e.PUT("/register/:id", func(c echo.Context) error {
		return HandlerUpdatePersona(c, h)
	})

	e.DELETE("/register/:id", func(c echo.Context) error {
		return HandlerDeletePersona(c, h)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
