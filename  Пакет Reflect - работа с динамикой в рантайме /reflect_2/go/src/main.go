package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

type User struct {
	Id       int
	UserName string `unpack:"-"`
	Login    string
	flog     int
}

func UnpackReflect(u interface{}, data []byte) error {
	r := bytes.NewReader(data)       //считывание из бинарных данных date
	val := reflect.ValueOf(u).Elem() // получаем внутрению структуру
	for i := 0; i < val.NumField(); i++ {
		valField := val.Field(i)                // получаем значение
		typeFaild := val.Type().Field(i)        // получаем тип
		if typeFaild.Tag.Get("unpack") == "-" { // если есть тег unpack
			continue
		}
		switch typeFaild.Type.Kind() { // предстовление типа
		case reflect.Int: // если содержит тип int
			var value int32                             // создаём переменую
			binary.Read(r, binary.LittleEndian, &value) // считываем данные из буфера
			valField.Set(reflect.ValueOf(int(value)))   // устанавливаем зпредставление тех данных которые нам нужны
		case reflect.String: // если содержит тип string
			var lenRaw uint32                             // создание переменой
			binary.Read(r, binary.LittleEndian, &data)    // считываем данные из буфера
			dataRaw := make([]byte, lenRaw)               //создание слайса байт
			binary.Read(r, binary.LittleEndian, &dataRaw) // читаю в него данные
			valField.SetString(string(dataRaw))           //присвоение значение в поле структуры
		default:
			return fmt.Errorf("bad type: %v for faild %v", typeFaild.Type.Kind(), typeFaild.Name)
		}

	}
	return nil
}
func main() {
	/*
		parl -E =pack("L L/a", 1_123_456, "v.romanov", 16);
		    Print map {ord."," } split{"", $b};
	*/
	data := []byte{ //байтовое  пристовление запокованого
		128, 36, 17, 0,
		9, 0, 0, 0,
		118, 46, 114, 111, 109, 97, 110, 111, 118,

		16, 0, 0, 0,
	}
	u := new(User)
	err := UnpackReflect(u, data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", u)
}
