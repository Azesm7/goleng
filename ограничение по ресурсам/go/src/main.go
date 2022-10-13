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
	quotaLimit    = 5 // лимит буфера для канала
)

func StartWorker(in int, wg *sync.WaitGroup, qoutach chan struct{}) {
	qoutach <- struct{}{} //берём свободный слот
	defer wg.Done()       // именьшаем счётчик на 1
	for j := 0; j < InerationsNum; j++ {
		fmt.Println("Gorutines", in, "Intertions", j) // вывод номер гоу рутины и интерации
		runtime.Gosched()                             // даём пороботать другим гоу рутинам
		if j%2 == 0 {
			<-qoutach             // возращаем слот
			qoutach <- struct{}{} // берём слот
		}
		<-qoutach // возращаем слот
	}
}
func main() {
	wg := &sync.WaitGroup{}                    // инцилизируем группу
	qoutach := make(chan struct{}, quotaLimit) // создаём канал
	for i := 0; i < GorutinesNum; i++ {
		wg.Add(1)                      // запускаем счётчик воркеров и добавляем ему 1 значение
		go StartWorker(i, wg, qoutach) // в отдельной гоу рутине запускаем функцию и передаём в неё номер интерации, значение счётчика и канал
	}
	time.Sleep(1 * time.Second) //создаём таймер
	wg.Wait()                   // ожидае пока defer wg.Done() не преведёт счётчик к нулю
}
