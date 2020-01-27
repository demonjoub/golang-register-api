package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/mystyle/db"
	"github.com/mystyle/schema"
)

func PostRegisterPersona(c echo.Context, h db.PersonaHandler) error {
	persona := schema.Personas{CreateAt: time.Now()}
	if err := c.Bind(&persona); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&persona).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, persona)
}

func GetPersonas(c echo.Context, h db.PersonaHandler) error {
	personas := []schema.Personas{}
	h.DB.Find(&personas)
	return c.JSON(http.StatusOK, personas)
}
