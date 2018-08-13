package main

import (
	"io/ioutil"
	"testing"

	"./complex"
)

func BenchmarkC64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		complex.MainC64(ioutil.Discard)
	}
}

func BenchmarkC128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		complex.MainC128(ioutil.Discard)
	}
}

//func BenchmarkBigFloat(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		complex.MainBigFloat(ioutil.Discard)
//	}
//}

//func BenchmarkRat(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		complex.MainRat(ioutil.Discard)
//	}
//}
