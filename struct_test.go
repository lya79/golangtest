package golangtest

import "testing"

/*
	windows 10
	x64
	i7-7700 CPU @ 3.60GHz 3.60GHz
	RMA 16GB

	go version go1.12.4 windows/amd64

	PS C:\Users\lya79_liou\go\src\golangtest> go test -bench="^BenchmarkStruct" -benchmem -benchtime="5s"
	BenchmarkStruct-8       10000000000              0.50 ns/op            0 B/op          0 allocs/op
	BenchmarkStruct2-8      10000000000              0.25 ns/op            0 B/op          0 allocs/op
*/

func BenchmarkStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var t = []struct {
			s    string
			rune rune
			out  int
		}{
			{"a A x", 'A', 2},
			{"some_text=some_value", '=', 9},
			{"☺a", 'a', 3},
			{"a☻☺b", '☺', 4},
		}
		if t == nil {

		}
	}
}

type ss []struct {
	s    string
	rune rune
	out  int
}

func BenchmarkStruct2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var t = ss{
			{"a A x", 'A', 2},
			{"some_text=some_value", '=', 9},
			{"☺a", 'a', 3},
			{"a☻☺b", '☺', 4},
		}
		if t == nil {

		}
	}
}
