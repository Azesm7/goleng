package main

import (
	"fmt"
	"net/http"
)

func runserver(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) { // оброботчик запросов
			fmt.Fprintln(w, "Addr:", addr, "URL:", r.URL.String())
		})
	server := http.Server{ // создание сервера
		Addr:    addr,
		Handler: mux,
	}
	fmt.Println("starting server at ", addr) // оброботчик запросов
	server.ListenAndServe()                  // слушаем сервер
}
func main() {
	go runserver(":8081")
	runserver(":8080")
}
