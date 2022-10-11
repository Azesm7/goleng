package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct { // создание структуры параметров
	ID     int
	Name   string
	Active bool
}

func main() {
	tmp1 := template.Must(template.ParseFiles("users.html")) // парсим шаблон
	users := []User{                                         // создание users
		User{1, "Roman", true},
		User{2, "hadler", false},
		User{3, "Dmitry", true},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp1.Execute(w, //нужно сделать шаблон
			struct { // создание структуры
				Users []User
			}{
				users, // создание экземпляра структуры
			})
	})
	fmt.Println("staring server :8080")
	http.ListenAndServe(":8080", nil)
}
