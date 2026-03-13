package log_benchmarks_test

import (
	"errors"
	"log"
	"testing"
)

func BenchmarkLog(b *testing.B) {
	e := errors.New("error")

	for i := 0; i < b.N; i++ {
		log.Println("debug test")
		log.Println("info", "test", "value")
		log.Println("warn", "test", 3)
		log.Printf("%v", e)
		log.Println(e)
		log.Println("info from log", "<3")
	}
}
