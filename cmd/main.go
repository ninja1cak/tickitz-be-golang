package main

import (
	"log"
	"ninja1cak/coffeshop-be/internal/routers"
	"ninja1cak/coffeshop-be/pkg"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database, err := pkg.PgDatabase()
	if err != nil {
		log.Fatal(err)
	}
	router := routers.New(database)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
