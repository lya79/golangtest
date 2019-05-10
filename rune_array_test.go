package golangtest

import (
	"log"
	"testing"
)

/*
	windows 10
	x64
	i7-7700 CPU @ 3.60GHz 3.60GHz
	RMA 16GB

	go version go1.12.4 windows/amd64

	PS C:\Users\lya79_liou\go\src\golangtest> go test -bench="^BenchmarkRune" -benchmem -benchtime="5s"
	BenchmarkRune-8         10000000000              0.25 ns/op            0 B/op          0 allocs/op
	BenchmarkRune2-8        10000000000              0.25 ns/op            0 B/op          0 allocs/op
*/

const str = "中文"
const str2 = "ABC"

func Test_rune(t *testing.T) {
	log.Println(str, len([]byte(str)))   // 中文 6
	log.Println(str2, len([]byte(str2))) // ABC 3

	log.Println(str, len([]rune(str)))   // 中文 2
	log.Println(str2, len([]rune(str2))) // ABC 3
}
func BenchmarkRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if len([]rune(str)) > 0 {

		}
	}
}

func BenchmarkRune2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if len([]byte(str)) > 0 {

		}
	}
}
