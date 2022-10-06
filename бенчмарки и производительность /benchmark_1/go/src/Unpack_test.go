package test_Pack

import (
	"testing" // подклучаем пакет для тестирования
)

func Benchmark_test(b *testing.B) { // созание функции бенчмакрка
	for i := 0; i < b.N; i++ { // пишем цикол
		data := make([]int, 0)
		for j := 0; j < 10; j++ {
			data = append(data, j)
		}
	}
}

// go test -bench . -benchmem ./Unpack_test.go (ввод в консоли )
