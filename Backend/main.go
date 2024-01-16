package main

import "C"
import (
	"Backend/Search"
	"Backend/ent"
	"Backend/ent/website"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	e.Static("/static", "assets")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/ranking", getRanking)

	e.Logger.Fatal(e.Start(":1323"))
}
func getRanking(c echo.Context) error {
	//フォームの情報を受け取る
	adult := c.FormValue("adult")
	flashy := c.FormValue("flashy")
	smart := c.FormValue("smart")
	beautiful := c.FormValue("beautiful")
	like := c.FormValue("like")
	bright := c.FormValue("bright")

	//DBのクライアントを生成
	client, err := ent.Open("sqlite3", "file:test.db?_foreign_keys=on")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	targetData := Search.Website{
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
			Bright:    new(float64),
		},
	}

	// フォームの値を代入。
	*targetData.Adult, err = strconv.ParseFloat(adult, 32)
	if err != nil {
		log.Fatalf(err.Error())
	}

	*targetData.Bright, err = strconv.ParseFloat(bright, 32)
	if err != nil {
		log.Fatalf(err.Error())
	}

	*targetData.Beautiful, err = strconv.ParseFloat(beautiful, 32)
	if err != nil {
		log.Fatalf(err.Error())
	}

	*targetData.Smart, err = strconv.ParseFloat(smart, 32)
	if err != nil {
		log.Fatalf(err.Error())
	}

	*targetData.Like, err = strconv.ParseFloat(like, 32)
	if err != nil {
		log.Fatalf(err.Error())
	}

	*targetData.Flashy, err = strconv.ParseFloat(flashy, 32)
	if err != nil {
		log.Fatalf(err.Error())
	}

	result := Search.Ranking(targetData, client)

	websiteDatas := make([]ent.Website, 0, 5)
	for i := 0; i < 4; i++ {
		websiteData, err := client.Website.Query().Where(website.ID(result[i].Id)).Only(context.Background())
		if err != nil {
			log.Fatalf(err.Error())
		}

		websiteDatas = append(websiteDatas, *websiteData)
	}

	return c.JSON(http.StatusOK, websiteDatas)
}
