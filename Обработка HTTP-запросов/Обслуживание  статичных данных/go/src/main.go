package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) { // w куда будем записывать наш результат r  параметры запроса
	w.Header().Set("Content-Type", "text/html") // вывод
	w.Write([]byte(` 
Hello Word! <br />
<img src="/data/img/gopher.png" />
`))
}

func main() {
	http.HandleFunc("/", Handler)
	staticHandler := http.StripPrefix( // обработчик событий статик
		"/data/",
		http.FileServer(http.Dir("./static")),
	)
	http.Handle("/data/", staticHandler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
