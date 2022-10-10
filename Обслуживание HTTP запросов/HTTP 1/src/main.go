package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { // w куда будем записывать наш результат r  параметры запроса
	fmt.Fprintln(w, "heello User")
	w.Write([]byte("!!!!"))
}
func main() {
	http.HandleFunc("/", handler) //  определение оброботчика по адресу и функции
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
