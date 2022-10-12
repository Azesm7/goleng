package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type Post struct { // создание структуры Post
	ID       int
	Text     string
	Author   string
	Comments int
	Time     time.Time
}

func handle(w http.ResponseWriter, req *http.Request) {
	s := ""                     //создание переменой
	for i := 0; i < 1000; i++ { // запуск цыкла
		p := &Post{ID: i, Text: "new post"} // создание нового поста
		s += fmt.Sprintf("%#v", p)          // приводит в текстовую строку
	}
	w.Write([]byte(s)) //вывод масива из байт поста
}
func main() {
	http.HandleFunc("/", handle)
	fmt.Println("starting server :8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}

// Для тестирования надо запустить программу и подать на неё нагрузку, например с помощью стандартной
//утилиты ab.
//go build -o pprof_1.exe pprof_1.go && ./pprof_1.exe
//ab -t 300 -n 1000000000 -c 10 http://127.0.0.1:8080/
//Затем можно снять хип-дамп и профиль CPU, для этого нужно дернуть URL-ы debug/pprof/heap и
//debug/pprof/profile для снятия CPU, при этом мы можем указать в параметре seconds, за сколько времени
//мы хотим снять профиль CPU.
//curl http://127.0.0.1:8080/debug/pprof/heap -o mem_out.txt
//curl http://127.0.0.1:8080/debug/pprof/profile?seconds=5 -o cpu_out.txt
//Далее, используя саму команду pprof, мы можем построить .svg файл по полученным результатам.
//go tool pprof -svg -alloc_objects pprof_1.exe mem_out.txt > mem_ao.svg
//go tool pprof -svg pprof_1.exe cpu_out.txt > cpu.svg
