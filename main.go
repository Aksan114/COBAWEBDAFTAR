package main

import (
	"golang/config"
	"golang/controller"
	"net/http"
)

func main() {
	config.GetConnection()

	http.HandleFunc("/", controller.Haldepan)

	http.HandleFunc("/create",controller.Buat)

	http.HandleFunc("/update",controller.Update)

	http.HandleFunc("/delete",controller.Selesai)

	controller.Serverhandlestatic()

	http.ListenAndServe(":8080", nil)
}