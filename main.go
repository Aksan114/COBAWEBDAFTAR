package main

import (
	"golang/config"
	"golang/controller"
	"net/http"
)

func main() {
	config.GetConnection()

	http.HandleFunc("/", controller.Index)

	http.HandleFunc("/create",controller.Create)

	http.HandleFunc("/update",controller.Update)

	http.HandleFunc("/delete",controller.Delete)

	fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}