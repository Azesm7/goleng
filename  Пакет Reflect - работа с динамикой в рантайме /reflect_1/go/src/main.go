package main

import (
	"fmt"
	"reflect"
)

type User struct { // создание структуры
	ID       int    `"user_id,string"`
	UserName string `unpack:"-"`
	Phone    string `",omitempty"`
	Login    string
	Flag     int
}

func PrintReflect(u interface{}) error {
	val := reflect.ValueOf(u).Elem()                     //  объект из пакета reflect который представляет структуру
	fmt.Printf("%T have %d field:\n", u, val.NumField()) // вывод  количество палей структуры
	for i := 0; i < val.NumField(); i++ {                // обходим поля
		valueField := val.Field(i)        // получаем значение поля
		typeFaielf := val.Type().Field(i) // получаем тип этого поля
		fmt.Printf("\tnam=%v, type=%v, value=%v, tag=`%v`\n",
			typeFaielf.Name,        // имя поля
			typeFaielf.Type.Kind(), // тип поля
			valueField,             // значение
			typeFaielf.Tag,         //teg
		)
	}
	return nil
}

func main() {
	u := &User{ // значение полей структуры
		ID:       55,
		UserName: "Roman",
		Phone:    "555-444",
		Login:    "fegi",
		Flag:     11,
	}
	err := PrintReflect(u) // возращение ошибки
	if err != nil {        // обработка ошибки
		panic(err)
	}

}
