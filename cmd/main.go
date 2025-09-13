package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/wsb777/involta-test/internal/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("[WARN] Load env failed")
	}
	handler, err := app.StartApp()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[INFO] Server is running on :8080")

	log.Fatal(http.ListenAndServe(":8080", handler))
}
