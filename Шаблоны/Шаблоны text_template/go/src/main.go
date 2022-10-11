package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type tplParams struct { // создание структуры параметров
	Url     string
	Browser string
}

// создание шаблона (пременные вводится({{.Имя переменой}}))
const EXAMPLE = ` 
Browser: {{.Browser}}
you at {{.Url}}
`

func handle(w http.ResponseWriter, r *http.Request) {
	tmp1 := template.New(`example`) //компиляция шаблона
	tmp1, _ = tmp1.Parse(EXAMPLE)
	Params := tplParams{ // создание параметров
		Url:     r.URL.String(),
		Browser: r.UserAgent(),
	}
	tmp1.Execute(w, Params) // запуск шаблона
}
func main() {
	http.HandleFunc("/", handle)
	fmt.Println("starting server :8080")
	http.ListenAndServe(":8080", nil)

}
