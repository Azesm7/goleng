package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	InerationsNum = 2 //количество интераций
	GorutinesNum  = 5 //количество гоу рутин
)

func StartWorker(in int, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшаем счётчик воркеров на 1
	for j := 0; j < InerationsNum; j++ {
		fmt.Println("Gorutines", in, "Interations", j) // выводи номер гоу рутины и интерации
		runtime.Gosched()                              // закрываем цыкал
	}
}

func main() {
	wg := &sync.WaitGroup{} //wailt_2 инцилизируем группу
	for i := 0; i < GorutinesNum; i++ {
		wg.Add(1)             //wailt_2  добавляем воркер
		go StartWorker(i, wg) // в отдельной гоу рутине запускаем функцию и передаёь в неё номер гоу рутины
	}
	time.Sleep(1 * time.Second) // создание таймера
	wg.Wait()                   //wailt_2 ожидае пока defer wg.Done() не преведёт счётчик к нулю
}
