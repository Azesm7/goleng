package main

import (
	"fmt"
	"runtime"
	"time"
)

const groutinesnum = 3 // количество гоу рутин

func StartWorker(WorkerNum int, in <-chan string) { //функци которая принемает число интераций и данные из канала
	for input := range in { //пребераем данные из канала
		fmt.Println("WorkerNum", WorkerNum, "Input", input) //вывод
		runtime.Gosched()                                   // октивно отказаться  от управления гоу рутиной
	}
}
func main() {
	WorkerInput := make(chan string, 2) //создаём канал
	for i := 0; i < groutinesnum; i++ { //запускаем цикал
		go StartWorker(i, WorkerInput) // в отдельной гоу рутине запускаем функцию в неё передаём номер интерации и канал
	}
	Munth := []string{"январь", "февраль", "март", // слайс
		"апрель", "май", "июнь", "июль", "август", "сентябрь",
		"октябрь", "ноябрь", "декабрь",
	}
	for _, MaunthName := range Munth { // пребор элементов слайса
		WorkerInput <- MaunthName // отправляем в канал элемент слайса
	}
	close(WorkerInput)          // закрываем канал (закрываем все воркеры)
	time.Sleep(1 * time.Second) // создание таймера
}
