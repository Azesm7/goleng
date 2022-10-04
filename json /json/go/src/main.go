package main

import (
	"encoding/json"
	"fmt"
)

type User struct { // создание структуры
	ID       int
	UserName string
	Phone    string
}

var jsonStr = `{"id": 42, "username":"Roman","phone": "223-45"}` // создание примера json файла

func main() {
	data := []byte(jsonStr)             //преобразуем строку с json в слейсу байт
	u := &User{}                        // создание пользователя
	json.Unmarshal(data, u)             //принемает слайс байт и куда нужно записать данные
	fmt.Printf("struct:\n\t%#v\n\n", u) // вывод изначального содержания файла

	u.Phone = "9295551235"         // изменения номера
	result, err := json.Marshal(u) //  обратно запоковать , возращает ошибку и результат
	if err != nil {                // обработка ошибки
		panic(err) // обработка паники
	}
	fmt.Printf("json string \n\t%s\n", string(result)) // вывод результата изменений
}
