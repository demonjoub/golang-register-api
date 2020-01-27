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
		return GetPersonas(c, h)
	})

	e.POST("/register", func(c echo.Context) error {
		return PostRegisterPersona(c, h)
	})

	e.PUT("/register/:id", func(c echo.Context) error {
		id := c.Param("id")
		persona := schema.Personas{}

		if err := h.DB.Find(&persona, id).Error; err != nil {
			return c.NoContent(http.StatusNotFound)
		}

		if err := c.Bind(&persona); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		if err := h.DB.Save(&persona).Error; err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, persona)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
