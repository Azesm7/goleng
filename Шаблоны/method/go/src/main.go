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

func (u *User) PrintActive() string {
	if !u.Active { //если user не активен
		return ""
	}
	return "method says user " + u.Name + " active" //если user активен
}
func main() {
	tmpl, err := template.New("").ParseFiles("users.html") // парсим шаблон
	if err != nil {
		panic(err)
	}
	users := []User{ // создание users
		User{1, "Roman", true},
		User{2, "Ivan", false},
		User{3, "Dmitry", true},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "users.html", // выполняем шаблон
			struct { //нужно сделать шаблон
				Users []User // создание структуры
			}{
				users, // создание экземпляра структуры
			})
		if err != nil {
			panic(err)
		}
	})
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
