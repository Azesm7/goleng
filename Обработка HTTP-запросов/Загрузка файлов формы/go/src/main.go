package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// форма
var uploadFormTmp = []byte(` 
<html>
    <body>
	<form action="/upload" method="post"
	enctype="multipart/form-data">
	    Image: <input type = "file" name="my_fail">
		<input type="submit" value="Upload">
	</form>
	</body>
</html>		
`)

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Write((uploadFormTmp)) // вывод формы
}
func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/upload", UploadPage)
	http.HandleFunc("/raw_body", UploadRawBOdy)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
func UploadPage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 * 1024 * 1025)       //парсинг формы расарсить 5 мб если что то остаётся записывается в временые файлы
	file, handler, err := r.FormFile("my_fail") //обращение к файлу и получаем данные
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() //закрыть файл

	fmt.Fprintf(w, "handler.Filename %v\n", handler.Filename) //вывод имя
	fmt.Fprintf(w, "handler.Header %v\n", handler.Header)     //вывод маин заголовка

	hasher := md5.New()   //создание нового объекта
	io.Copy(hasher, file) //отправление файла

	fmt.Fprintf(w, "md5 %x\n", hasher.Sum(nil)) // вывод
}

type Parms struct {
	ID   int
	User string
}

//...................................................................................................................................................................
// если нужно спарсить не через main

/*
ввод в терминал
curl -v -X POST -H "Content: application/json" -d '{"id":2,
"User": "Roman"}' http://localhost:8080/raw_body
*/
func UploadRawBOdy(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body) // прсинг на примую, считывает всё что есть в Body
	defer r.Body.Close()                // закрываем

	p := &Parms{}                 // создание структуры
	err = json.Unmarshal(body, p) // парсинг json в эту струтуру
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "content-type %#v\n",
		r.Header.Get("Content-Type"))
	fmt.Fprintf(w, "params %#v\n", p) //вывод парсинга

}
