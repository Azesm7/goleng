package main

import (
	"fmt"
	"net/http"
)

var loginFormTmp1 = []byte(`
<html> 
	<body> 
	<form action="/" method="post"> 
	  Login: <input type="text" name="login">
	  Password: <input type="password" name="password"> 
	  <input type="submit" value="login"> 
	</form> 
	</body> 
</html> 
`)

func MainPage(w http.ResponseWriter, r *http.Request) { // w куда будем записывать наш результат r  параметры запроса
	if r.Method != http.MethodPost { // проверка  если метод не пуст то
		w.Write(loginFormTmp1) //
		return
	}
	//r.ParseForm()
	// inputlogin := r.Form["login"][0]
	inputlogin := r.FormValue("login") // караткое обращение  вернёт первый нужный параметр
	fmt.Fprintln(w, "you enter: ", inputlogin)

}
func main() {
	http.HandleFunc("/", MainPage)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
