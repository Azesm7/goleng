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

func IsUserOdd(u *User) bool { //евляется ли индификатор пользователя чётным числом
	return u.ID%2 != 0 // возврат bool значение
}
func main() {
	tmplFunc := template.FuncMap{ //создание карты функции
		"OddUser": IsUserOdd,
	}
	tmpl, err := template. // создание шаблонизации
				New("").
				Funcs(tmplFunc).        // вызов функции
				ParseFiles("Func.html") // парсинг файла
	if err != nil {
		panic(err)
	}
	users := []User{ // создание users
		User{1, "Roman", true},
		User{2, "Dmitriy", false},
		User{1, "Ivan", true},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "Func.html", //нужно сделать шаблон
			struct { // создание структуры
				Users []User
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
