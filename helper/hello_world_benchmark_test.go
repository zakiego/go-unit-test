package helper

import (
	"testing"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Zaki")
	}
}

func BenchmarkHelloWorldLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Muhammad Zakiyuddin")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Zaki", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Zaki")
		}
	})

	b.Run("Muhammad Zakiyuddin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Muhammad Zakiyuddin")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{name: "HelloWorld(Zaki)", request: "Zaki"},
		{name: "HelloWorld(Muhammad Zakiyuddin)", request: "Muhammad Zakiyuddin"},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
