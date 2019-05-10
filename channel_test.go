package golangtest

import (
	"sync"
	"testing"
)

/*
	windows 10
	x64
	i7-7700 CPU @ 3.60GHz 3.60GHz
	RMA 16GB

	go version go1.12.4 windows/amd64

	PS C:\Users\lya79_liou\go\src\lib> go test -bench="^BenchmarkChannel" -benchmem -benchtime="5s"
	BenchmarkChannel-8        500000             16123 ns/op               0 B/op          0 allocs/op
	BenchmarkChannel2-8       300000             21404 ns/op               3 B/op          0 allocs/op
	BenchmarkChannel3-8       500000             18837 ns/op             461 B/op          1 allocs/op
	BenchmarkChannel4-8       500000             18998 ns/op             474 B/op          1 allocs/op
	BenchmarkChannel5-8       300000             21776 ns/op             506 B/op          1 allocs/op
*/

func BenchmarkChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		quit := make(chan int)
		lock := sync.Mutex{}
		sum := 0
		answer := 5050
		b.StartTimer()

		for k := 0; k <= 100; k++ {
			go func(k int) {
				lock.Lock()
				defer lock.Unlock()
				sum += k
				if sum >= answer {
					quit <- 1
				}
			}(k)
		}

		b.StopTimer()
		<-quit
		b.StartTimer()
	}
}

func BenchmarkChannel2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		quit := make(chan int)
		lock := sync.Mutex{}
		sum := 0
		answer := 5050
		b.StartTimer()

		for k := 0; k <= 100; k++ {
			go func(k int) {
				lock.Lock()
				sum += k
				lock.Unlock()
				if sum >= answer {
					quit <- 1
				}
			}(k)
		}

		b.StopTimer()
		<-quit
		b.StartTimer()
	}
}

func BenchmarkChannel3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		quit := make(chan int)
		c := make(chan int)
		sum := 0
		answer := 5050
		b.StartTimer()

		go func() {
			for {
				sum += <-c
				if sum >= answer {
					quit <- 1
				}
			}
		}()
		for k := 0; k <= 100; k++ {
			go func(k int) {
				c <- k
			}(k)
		}

		b.StopTimer()
		<-quit
		b.StartTimer()
	}
}

func BenchmarkChannel4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		quit := make(chan int)
		c := make(chan int, 50)
		sum := 0
		answer := 5050
		b.StartTimer()

		go func() {
			for {
				sum += <-c
				if sum >= answer {
					quit <- 1
				}
			}
		}()
		for k := 0; k <= 100; k++ {
			go func(k int) {
				c <- k
			}(k)
		}

		b.StopTimer()
		<-quit
		b.StartTimer()
	}
}

func BenchmarkChannel5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		quit := make(chan int)
		c := make(chan int, 105)
		sum := 0
		answer := 5050
		b.StartTimer()

		go func() {
			for {
				sum += <-c
				if sum >= answer {
					quit <- 1
				}
			}
		}()
		for k := 0; k <= 100; k++ {
			go func(k int) {
				c <- k
			}(k)
		}

		b.StopTimer()
		<-quit
		b.StartTimer()
	}
}
