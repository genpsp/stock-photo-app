package routes

import (
	"stock-photo-api/src/handlers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, h handlers.Handler) {
	app := e.Group("/api")

	images := app.Group("/images")
	images.POST("/upload", h.Image.Upload)

}
