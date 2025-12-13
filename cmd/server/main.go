package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/masaya-nishimura-09/setup-mixer/internal/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/item", handler.ItemHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
