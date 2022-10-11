package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { // w куда будем записывать наш результат r  параметры запроса
	myParam := r.URL.Query().Get("param") // получить параметры из url
	if myParam != "" {                    // если параметр не пуст
		fmt.Fprintln(w, "`myParam` is", myParam)
	}
	key := r.FormValue("key") // получить параметры из функцию key
	if key != "" {            // если параметр не пуст
		fmt.Fprintln(w, "`key` is", key)
	}

}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

//http://127.0.0.1:8080/?param=1&key=2
