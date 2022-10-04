package main

import (
	"fmt"
	"sync"
)

func main() {
	var conuters = map[int]int{}
	mu := &sync.Mutex{} //  создание mutex как сылка на объект
	for i := 0; i < 5; i++ {
		go func(conuteres map[int]int, th int, mu *sync.Mutex) {
			for j := 0; j < 5; j++ {
				mu.Lock() // блокируем данные для других
				conuters[th*10+j]++
				mu.Unlock() // снятие блока на данные

			}
		}(conuters, i, mu)
	}
	fmt.Scanln()
	mu.Lock() // блокируем данные
	fmt.Println("1111sja", conuters)
	mu.Unlock() // снятие блока на данные
}
