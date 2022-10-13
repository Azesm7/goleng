package main

import (
	"fmt"
	"time"
)

func GetCommit() chan string { // функция котороя возращает канал
	result := make(chan string, 1) // создание буфероного канала
	go func(out chan<- string) {   //функция которая принемает канал
		time.Sleep(1 * time.Second)                         // вызов таймера
		fmt.Println("async operation read,return comments") // вывод
		out <- "32 Commit"                                  // предаём в канал строку
	}(result) // вызов функции
	return result //возрат канала
}

func main() {
	resulth := GetCommit()                       // вызов функции
	fmt.Println("get related aeticles")          // 1 вывод
	commentsDate := <-resulth                    // передаём в переменую даные полученые из канала
	fmt.Println("main goroutnine", commentsDate) //вывод данных которые хронятся в потоке
}
