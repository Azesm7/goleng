package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

type User struct {
	ID       int
	RealName string `unpack:"-"`
	Login    string
	Flags    int
}

func UnpackReflect(u interface{}, data []byte) error {
	r := bytes.NewReader(data)       //считывание из бинарных данных date
	val := reflect.ValueOf(u).Elem() // получаем внутрению структуру
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)              // получаем значение
		typeField := val.Type().Field(i)        // получаем тип
		if typeField.Tag.Get("unpack") == "-" { // если есть тег unpack
			continue
		}
		switch typeField.Type.Kind() { // предстовление типа
		case reflect.Int: // если содержит тип int
			var value uint32                            // создаём переменую
			binary.Read(r, binary.LittleEndian, &value) // считываем данные из буфера
			valueField.Set(reflect.ValueOf(int(value))) // устанавливаем зпредставление тех данных которые нам нужны
		case reflect.String: // если содержит тип string
			var lenRaw uint32                             // создание переменой
			binary.Read(r, binary.LittleEndian, &lenRaw)  // считываем данные из буфера
			dataRaw := make([]byte, lenRaw)               //создание слайса байт
			binary.Read(r, binary.LittleEndian, &dataRaw) // читаю в него данные
			valueField.SetString(string(dataRaw))         //присвоение значение в поле структуры
		default:
			return fmt.Errorf("bad type: %v for field %v", typeField.Type.Kind(), typeField.Name)
		}
	}
	return nil
}
func main() {
	/*
		parl - E '$b = pack("L L/a* L", 1_123_456, "v.romanov", 16);
		    print map { ord.", " } split{"", $b}; '
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
