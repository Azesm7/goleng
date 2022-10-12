package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 2) // создаём канал
	ch1 <- 55                // предаём значение в канал
	ch1 <- 65                // предаём значение в канал
	ch2 := make(chan int, 1) // создаём канал
	ch2 <- 20                // предаём значение в канал
	for {                    // запускаем цикл
		select {
		case val := <-ch1: // если в val передоётся значение с канала
			fmt.Println("ch1 val", val)
		case ch2 <- 1: // если во 2 канал передаётся значение
			fmt.Println("put val to ch2")
		default: // если не чего не сработало
			break // выход

		}
	}
}
