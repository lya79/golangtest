package golangtest

import (
	"strconv"
	"testing"
)

/*
	windows 10
	x64
	i7-7700 CPU @ 3.60GHz 3.60GHz
	RMA 16GB

	go version go1.12.4 windows/amd64

	PS C:\Users\lya79_liou\go\src\golangtest> go test -bench="^BenchmarkByte" -benchmem -benchtime="5s"
	BenchmarkByteValue-8    20000000               439 ns/op               0 B/op          0 allocs/op
	BenchmarkByteIndex-8    30000000               250 ns/op               0 B/op          0 allocs/op
	BenchmarkByte-8         2000000000               5.05 ns/op            0 B/op          0 allocs/op
	BenchmarkByte2-8        2000000000               5.05 ns/op            0 B/op          0 allocs/op
	BenchmarkByte3-8        2000000000               4.96 ns/op            0 B/op          0 allocs/op
*/

func initSlice() []string {
	var routes []string
	len := 1000
	routes = make([]string, len)
	for i := 0; i < len; i++ {
		routes[i] = strconv.Itoa(i)
	}
	return routes
}

func BenchmarkByteValue(b *testing.B) {
	routes := initSlice()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, rt := range routes {
			if rt == "" {
			}
		}
	}
}

func BenchmarkByteIndex(b *testing.B) {
	routes := initSlice()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for index, _ := range routes {
			if routes[index] == "" {
			}
		}
	}
}

const URL = "/get/"
const SUB_URL = URL + "helloabc1/"

func BenchmarkByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		URL := SUB_URL
		lenURL := len(URL)
		x := len(URL)
		y := lenURL // len=y-x
		z := lenURL // cap=z-x
		sub := []byte(URL)[x:y:z]

		if len(sub) > 0 {
		}
	}
}

func BenchmarkByte2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		URL := SUB_URL
		x := len(URL)
		sub := []byte(URL)[x:]
		if len(sub) > 0 {
		}
	}
}

func BenchmarkByte3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		URL := SUB_URL
		sub := []byte(URL)
		if len(sub) > 0 {
		}
	}
}
