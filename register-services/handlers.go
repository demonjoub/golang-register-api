package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/mystyle/db"
	"github.com/mystyle/schema"
)

func HandlerPostRegisterPersona(c echo.Context, h db.PersonaHandler) error {
	persona := schema.Personas{CreateAt: time.Now()}
	if err := c.Bind(&persona); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&persona).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, persona)
}

func HandlerGetPersonas(c echo.Context, h db.PersonaHandler) error {
	personas := []schema.Personas{}
	h.DB.Find(&personas)
	return c.JSON(http.StatusOK, personas)
}

func HandlerUpdatePersona(c echo.Context, h db.PersonaHandler) error {
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
}

func HandlerDeletePersona(c echo.Context, h db.PersonaHandler) error {
	id := c.Param("id")
	persona := schema.Personas{}

	if err := h.DB.Find(&persona, id).Error; err != nil {
		errMessage := schema.MyError{Code: http.StatusNotFound, Message: err.Error()}
		return c.JSON(http.StatusNotFound, errMessage)
	}

	if err := h.DB.Delete(&persona).Error; err != nil {
		errMessage := schema.MyError{Code: http.StatusNotFound, Message: err.Error()}
		return c.JSON(http.StatusNoContent, errMessage)
	}

	return c.JSON(http.StatusOK, persona)
}
