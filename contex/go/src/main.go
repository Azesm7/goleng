package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func GoContext(ctx context.Context, WorkerNum int, out chan<- int) { // обевляем функцию и передаём в неё данные
	WailtTime := time.Duration(rand.Intn(100)+10) * time.Millisecond // создаём таймер с рандомным диапозоне до 100 и добавлем ему 10 милисекунд
	select {
	case <-ctx.Done(): // если контекст закрыт
		return // возрат
	case <-time.After(WailtTime): // если таймер работает то
		fmt.Println("worker", WorkerNum, "done") // вывод
		out <- WorkerNum                         //созраняем в поток данные интерации
	}
}

func main() {
	ctx, Finish := context.WithCancel(context.Background()) // создание контекста для закрытия
	Result := make(chan int, 1)                             // создание буферозированого канала
	for i := 0; i <= 10; i++ {                              // запуск цыкла
		go GoContext(ctx, i, Result) // запуск в отделной го рутине функции и передаём в неё контекст интерацию и канал
	}
	FoundBy := <-Result                     // передаём канал в переменую
	fmt.Println("result found by", FoundBy) //вывод
	Finish()                                //закрываем контекст
	time.Sleep(time.Second)                 // время работы
}
