package popcount_test

import "testing"
import . "../popcount"

func benchmark(b *testing.B, f func(uint64) int) {
  for i := 0; i < b.N; i++ {
    f(uint64(i))
  }
}

func BenchmarkOrig(b *testing.B) {
  benchmark(b, PopCount)
}

func BenchmarkLoop(b *testing.B) {
  benchmark(b, PopCountLoop)
}

func BenchmarkBitshift(b *testing.B) {
  benchmark(b, PopCountBitshift)
}

func BenchmarkBitclear(b *testing.B) {
  benchmark(b, PopCountBitclear)
}


