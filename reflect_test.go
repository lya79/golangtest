package golangtest

import (
	"reflect"
	"testing"
)

/*
	windows 10
	x64
	i7-7700 CPU @ 3.60GHz 3.60GHz
	RMA 16GB

	go version go1.12.4 windows/amd64

	PS C:\Users\lya79_liou\go\src\lib> go test -bench="^BenchmarkReflect" -benchmem -benchtime="5s"
	BenchmarkReflect-8      10000000               686 ns/op             272 B/op          7 allocs/op
	BenchmarkReflect2-8     10000000000              0.25 ns/op            0 B/op          0 allocs/op
*/

type T struct{}

func (t *T) Do(a int) (string, error) {
	return "hello", nil
}

func Test_reflect(t *testing.T) {
	v := 1111

	name := "Do"
	ty := &T{}

	a := reflect.ValueOf(v)
	in := []reflect.Value{a}

	ret := reflect.ValueOf(ty).MethodByName(name).Call(in)

	if ret[1].Interface() == nil {
		t.Log("success", ret[0].Interface().(string))
	} else if ret[1].Interface().(error) != nil {
		t.Fatal("err:", ret[1].Interface().(error))
	}
}

func BenchmarkReflect(b *testing.B) {
	v := 1111
	name := "Do"
	t := &T{}
	a := reflect.ValueOf(v)
	in := []reflect.Value{a}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ret := reflect.ValueOf(t).MethodByName(name).Call(in)
		if ret[1].Interface() == nil || ret[0].Interface().(string) == "" {

		}
	}
}

func BenchmarkReflect2(b *testing.B) {
	v := 1111
	t := &T{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		str, err := t.Do(v)
		if err != nil || str == "" {

		}
	}
}
