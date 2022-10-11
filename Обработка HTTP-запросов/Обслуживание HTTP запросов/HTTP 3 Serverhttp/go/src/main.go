package main

import (
	"fmt"
	"net/http"
)

type handlerType struct {
	Name string
}

func (h *handlerType) ServeHTTP(w http.ResponseWriter, r *http.Request) { // w куда будем записывать наш результат r  параметры запроса
	fmt.Fprintln(w, "Name:", h.Name, "URL:", r.URL.String())
}
func main() {
	testHandler := &handlerType{Name: "test"} // 1 оброботчик
	http.Handle("/test/", testHandler)
	rootHandler := &handlerType{Name: "root"} // 2 оброботчик
	http.Handle("/", rootHandler)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
