package main

import (
	"bytes"
	"encoding/json"
	"sync"
	"testing"
)

const interNum = 1000

type PublicPage struct {
	ID          int
	Name        string
	Url         string
	OwnerID     int
	ImageUr1    string
	Tags        []string
	Description string
	Rules       string
}

var CoolGolangPublic = PublicPage{
	ID:          1,
	Name:        "Goleng122123",
	Url:         "http://example.com",
	OwnerID:     1000500,
	ImageUr1:    "http://exqmple.com/img.png",
	Tags:        []string{"Programming", "go", "goleng"},
	Description: "Best page about golang programing",
	Rules:       "",
}
var Pages = []PublicPage{
	CoolGolangPublic,
	CoolGolangPublic,
	CoolGolangPublic,
}

func BenchmarkAllocNew(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data := bytes.NewBuffer(make([]byte, 0, 64))
			_ = json.NewEncoder(data).Encode(Pages)
		}
	})
}

var dataPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 64))

	},
}

func BenchmarkAllocPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data := dataPool.Get().(*bytes.Buffer)
			_ = json.NewEncoder(data).Encode(Pages)
			data.Reset()       // сбрасываем данные
			dataPool.Put(data) // возращаем данные
		}
	})
}
