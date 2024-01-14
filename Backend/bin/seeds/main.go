package main

import (
	"Backend/ent"
	"context"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
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

	website, err := createWebSite(context.Background(), client)
	if err != nil {

	}
	log.Println("website was created: ", website)
}

func createWebSite(ctx context.Context, client *ent.Client) (*ent.Website, error) {
	website, err := client.Website.Create().
		SetName("株式会社 CirKit").SetFilePath("/cirkit.png").
		SetURL("https://cirkit.jp/").SetFlashy(3).SetBeautiful(2).
		SetAdult(1).SetSmart(2).SetLike(1).Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating website: %w", err)
	}

	return website, nil
}
