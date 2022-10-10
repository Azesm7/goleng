package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) { // w куда будем записывать наш результат r  параметры запроса
	w.Header().Set("RequestID", "d41d8cd98f00b204")          //создание хедара
	fmt.Fprintln(w, "You brower is", r.UserAgent())          // получение данных
	fmt.Fprintln(w, "You brower is", r.Header.Get("Accept")) // получение данных
}

func main() {
	http.HandleFunc("/", Handler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
