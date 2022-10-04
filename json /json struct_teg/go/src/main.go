package main

import (
	"encoding/json"
	"fmt"
)

type User struct { // создание структуры
	ID       int `json:"user_id,string"` // тег структуры  к чему это относиться: имя , тип данных в который он должен записать этот тип
	UserName string
	Address  string `json:",omitempty"` //тег структуры  к чему это относиться: если поле пустое то его не нужно писать
	Comnpay  string `json:"-"`          //тег структуры  к чему это относиться:  поле не нужно стерилизовать и дистеризововать при json
}

func main() {
	u := &User{
		ID:       42,
		UserName: "Roman",
		Address:  "address",
		Comnpay:  "redceit",
	} // создание пользователя
	result, err := json.Marshal(u) //  запоковать , возращает ошибку и результат
	if err != nil {                // обработка ошибки
		panic(err) // обработка паники
	}
	fmt.Printf("json string %s\n", string(result)) // вывод результата изменений
}
