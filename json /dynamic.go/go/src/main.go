package main

import (
	"encoding/json"
	"fmt"
)

var jsonStr = `[
	{"id": 17, "username": "Roman", "phone": "223-45"},
	{"id": "17", "address": "address1", "company": "godbeck"}
]` // пример json файла

func main() {
	data := []byte(jsonStr)                                   // преобразуем json в байты
	var user1 interface{}                                     //создание пременой с пустой интерфейс
	json.Unmarshal(data, &user1)                              // распоковка json в интерфейс
	fmt.Printf("unpacked in emty interface:\n%#v\n\n", user1) // вывод результата
	user2 := map[string]interface{}{                          // создания слайса с интерфейсом
		"id":       42,
		"username": "iivan",
	}
	var user2i interface{} = user2      // создание переменой с интерфейсом и предаётся данные с пременой user2
	result, err := json.Marshal(user2i) // стерилизую
	if err != nil {                     // обработка ошибок
		panic(err)
	}
	fmt.Printf("json string from map:\n %s\n", string(result)) // вывод результата
}
