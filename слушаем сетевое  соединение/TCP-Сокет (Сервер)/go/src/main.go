package main

import (
	"bufio"
	"fmt"
	"net"
)

func conect(conn net.Conn) {
	name := conn.RemoteAddr().String()            // получаем имя удолёного соединения  RemoteAdder
	fmt.Printf("%+v connected\n", name)           // выводим что кто то соединился
	conn.Write([]byte("Hello, " + name + "\n\r")) //приветствие
	defer conn.Close()                            // указываем что соединение закрыть
	scanner := bufio.NewScanner(conn)             // создаём сканер который ожидает ввод пустой строки
	for scanner.Scan() {                          // пока что то приходит то
		text := scanner.Text() //  смотрим что нам приходит и обрабатываем это
		if text == "Exit" {    // если текст равен Exit то
			conn.Write([]byte("Bye\n\r"))     // прощаемся
			fmt.Println(name, "disconnected") //вывод
			break                             // закрываем соединение
		} else if text != "" { // если не пустой текст то
			fmt.Println(name, "enters", text)               // вывод
			conn.Write([]byte("you enter" + text + "\n\r")) // отправка сообщение
		}

	}

}
func main() {
	listner, err := net.Listen("tcp", ":8080") // создаём объект который будет слушать сокет
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listner.Accept() // обработка соединения
		if err != nil {
			panic(err)
		}
		go conect(conn)
	}
}
