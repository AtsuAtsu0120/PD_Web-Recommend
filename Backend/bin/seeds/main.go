package main

import (
	"Backend/ent"
	"context"
	"encoding/csv"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
)

func main() {
	client, err := ent.Open("sqlite3", "file:test.db?_foreign_keys=on")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	csvToDb(context.Background(), client)
}

func csvToDb(ctx context.Context, client *ent.Client) {
	file, err := os.Open("file.csv")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, v := range rows {
		flashy, err := strconv.ParseFloat(v[2], 32)
		if err != nil {
			log.Fatalf(err.Error())
		}

		adult, err := strconv.ParseFloat(v[3], 32)
		if err != nil {
			log.Fatalf(err.Error())
		}

		bright, err := strconv.ParseFloat(v[4], 32)
		if err != nil {
			log.Fatalf(err.Error())
		}
		smart, err := strconv.ParseFloat(v[5], 32)
		if err != nil {
			log.Fatalf(err.Error())
		}

		beatiful, err := strconv.ParseFloat(v[6], 32)
		if err != nil {
			log.Fatalf(err.Error())
		}

		like, err := strconv.ParseFloat(v[7], 32)
		if err != nil {
			log.Fatalf(err.Error())
		}
		client.Website.Create().
			SetName(v[0]).SetFilePath("/" + v[0] + ".png").
			SetURL(v[1]).SetFlashy(flashy).SetBeautiful(beatiful).
			SetAdult(adult).SetSmart(smart).SetLike(like).SetBright(bright).Save(ctx)
	}
}
