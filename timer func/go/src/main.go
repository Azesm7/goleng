package main

import (
	"fmt"
	"time"
)

func Print() {
	fmt.Println("3") //вывод 3
}
func Print2() {
	timer := time.AfterFunc(2*time.Second, Print)
	fmt.Println("2") //вывод 2
	fmt.Scanln()     // ждёт пустого ввода
	timer.Stop()     // отключение таймера
}
func main() {
	timer := time.AfterFunc(2*time.Second, Print2)
	fmt.Println("1") //вывод 1
	fmt.Scanln()     // ждёт пустого ввода
	timer.Stop()     // отключение таймера
}
