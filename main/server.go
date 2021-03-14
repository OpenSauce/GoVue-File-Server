package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/avaliablespace", avaliablespaceHandler)

	fs := http.FileServer(http.Dir("../frontend/dist/"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func avaliablespaceHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	stats := GetHDDStats("/")

	fmt.Printf("Recieved the message %f ", GetHDDStats("/").Percentage)
	fmt.Fprintf(w, `{ "avaliablespace": "%f",
	"totalSpace": "%s",
	"freeSpace": "%s",
	"pcName": "%s" }`, stats.Percentage, stats.TotalSpace , stats.FreeSpace, GetHostname())
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}