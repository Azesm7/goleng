package main

import (
	"fmt"
	"net/http"
	"time"
)

func MainPage(w http.ResponseWriter, r *http.Request) { // w куда будем записывать наш результат r  параметры запроса
	session, err := r.Cookie("session_id") //получает метод куку
	loggedIn := (err != http.ErrNoCookie)  // если кука есть мы считаем что пользователь зарег. если нет надо автор.

	if loggedIn { //если он залагинен
		fmt.Fprintln(w, `<a href="/logout">logout</a>`)
		fmt.Fprintln(w, "Welcome, "+session.Value)
	} else { //если нет
		fmt.Fprintln(w, `<a href="/login">login</a>`)
		fmt.Fprintln(w, "You need to login")
	}
}
func loginPage(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{ // создание новой куки
		Name:    "session_id",
		Value:   "Roman",
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)                 // поставить куку
	http.Redirect(w, r, "/", http.StatusFound) //редирект на главную страницу
}
func logoutPage(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id") //попытка получить куку
	if err == http.ErrNoCookie {           // если нет то сразу возращаемся на главную
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1) //если она есть то ставлю её в прошлую
	http.SetCookie(w, session)                     //поставить куку в нове значение
	http.Redirect(w, r, "/", http.StatusFound)     //редирект на главную страницу
}
func main() {
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/logout", logoutPage)
	http.HandleFunc("/", MainPage)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
