package main

import (
	"golang/config"
	"golang/controller"
	"net/http"
)

func main() {
	config.GetConnection()

	http.HandleFunc("/index", controller.Haldepan)

	http.HandleFunc("/", controller.Home)

	http.HandleFunc("/information-pro", controller.Information)

	http.HandleFunc("/information-jdwl", controller.Informationjdwl)

	http.HandleFunc("/information-pggmn", controller.Informationpnggmn)

	http.HandleFunc("/create",controller.Buat)

	http.HandleFunc("/delete",controller.Selesai)

	http.HandleFunc("/image", controller.ShowImage)

	controller.Serverhandlestatic()

	http.ListenAndServe(":8080", nil)
}