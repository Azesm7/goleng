package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 1) // создаём канал
	go func(in chan int) {   //запускаем функцию в отделной go рутине
		val := <-in                             // создаём переменую val в неё передаём значение с канала
		fmt.Println("Go: get from chan", val)   // вывод
		fmt.Println("Go: after read from chan") // вывод
		close(in)                               // закрываем канал
	}(ch1)
	ch1 <- 20                              // передаём значение в канал
	fmt.Println("main: after put to chan") // вывод
}
