package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	// Подключаемся к сокету
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Print("not connection server")
	}
	for {
		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin) // ждём  ввода
		fmt.Print("enter: ")
		text, err := reader.ReadString('\n')// считываем ввод
		if err != nil {
			panic(err)
		}
		if text == "Exit" {
			os.Exit(1)
		}
		// Отправляем в socket
		fmt.Fprintf(conn, text+"\n")
		// Прослушиваем ответ
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Print("not server " + message)
			panic(err)
		}
		fmt.Print("not server " + message)
	}
}
