package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

//var Client = initClient()

func initClient() *pgx.Conn {
	client, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer func(client *pgx.Conn, ctx context.Context) {
		err := client.Close(ctx)
		if err != nil {
			log.Fatalf("Unable to close connection: %v\n", err)
		}
	}(client, context.Background())

	return client
}
