package main

import "C"
import (
	"Backend/Search"
	"Backend/ent"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	e.Static("/static", "assets")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ranking", getRanking)

	e.Logger.Fatal(e.Start(":1323"))
}
func getRanking(c echo.Context) error {
	result := Search.Ranking(Search.Website{
		ent.Website{
			ID:        0,
			Name:      "",
			FilePath:  "",
			URL:       "",
			Flashy:    new(float64),
			Adult:     new(float64),
			Smart:     new(float64),
			Beautiful: new(float64),
			Like:      new(float64),
		},
	})

	return c.JSON(http.StatusOK, result[0:4])
}
