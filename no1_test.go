package no1

import (
	"runtime"
	"testing"
	"math/rand"
)

func testTester(r tester, max int, t *testing.T) {
	if r.max() != max {
		t.Fail()
	}

	for i := 0; i < max; i++ {
		if !r.has(i) {
			t.Fail()
		}
	}

	for i := max; i < max*2; i++ {
		if r.has(i) {
			t.Fail()
		}
	}
}

func TestSlice(t *testing.T) {
	max := 10
	testTester(genSlice(max), max, t)
}

func TestMap(t *testing.T) {
	max := 10
	testTester(genMap(max), max, t)
}

func benchmark(r tester, b *testing.B) {
	runtime.GC()
	max := r.max()
	b.ResetTimer()

	var has bool
	for i := 0; i < b.N; i++ {
		for i := 0; i < max; i++ {
			has = r.has(i)
		}
	}
	_ = has
}

func benchmarkRand(r tester, b *testing.B) {
	runtime.GC()
	max := r.max()
	b.ResetTimer()

	var has bool
	for i := 0; i < b.N; i++ {
		for i := 0; i < max; i++ {
			has = r.has(rand.Int() % max)
		}
	}
	_ = has
}

func BenchmarkMap100(b *testing.B)         { benchmark(genMap(100), b) }
func BenchmarkSlice100(b *testing.B)       { benchmark(genSlice(100), b) }
func BenchmarkMap1000(b *testing.B)        { benchmark(genMap(1000), b) }
func BenchmarkSlice1000(b *testing.B)      { benchmark(genSlice(1000), b) }
func BenchmarkMap10000(b *testing.B)       { benchmark(genMap(10000), b) }
func BenchmarkSlice10000(b *testing.B)     { benchmark(genSlice(10000), b) }
func BenchmarkMap100000(b *testing.B)      { benchmark(genMap(100000), b) }
func BenchmarkSlice100000(b *testing.B)    { benchmark(genSlice(100000), b) }
func BenchmarkMap1000000(b *testing.B)     { benchmark(genMap(1000000), b) }
func BenchmarkSlice1000000(b *testing.B)   { benchmark(genSlice(1000000), b) }
func BenchmarkMap10000000(b *testing.B)    { benchmark(genMap(10000000), b) }
func BenchmarkSlice10000000(b *testing.B)  { benchmark(genSlice(10000000), b) }
func BenchmarkMap100000000(b *testing.B)   { benchmark(genMap(100000000), b) }
func BenchmarkSlice100000000(b *testing.B) { benchmark(genSlice(100000000), b) }

func BenchmarkRandMap100(b *testing.B)         { benchmarkRand(genMap(100), b) }
func BenchmarkRandSlice100(b *testing.B)       { benchmarkRand(genSlice(100), b) }
func BenchmarkRandMap1000(b *testing.B)        { benchmarkRand(genMap(1000), b) }
func BenchmarkRandSlice1000(b *testing.B)      { benchmarkRand(genSlice(1000), b) }
func BenchmarkRandMap10000(b *testing.B)       { benchmarkRand(genMap(10000), b) }
func BenchmarkRandSlice10000(b *testing.B)     { benchmarkRand(genSlice(10000), b) }
func BenchmarkRandMap100000(b *testing.B)      { benchmarkRand(genMap(100000), b) }
func BenchmarkRandSlice100000(b *testing.B)    { benchmarkRand(genSlice(100000), b) }
func BenchmarkRandMap1000000(b *testing.B)     { benchmarkRand(genMap(1000000), b) }
func BenchmarkRandSlice1000000(b *testing.B)   { benchmarkRand(genSlice(1000000), b) }
func BenchmarkRandMap10000000(b *testing.B)    { benchmarkRand(genMap(10000000), b) }
func BenchmarkRandSlice10000000(b *testing.B)  { benchmarkRand(genSlice(10000000), b) }
func BenchmarkRandMap100000000(b *testing.B)   { benchmarkRand(genMap(100000000), b) }
func BenchmarkRandSlice100000000(b *testing.B) { benchmarkRand(genSlice(100000000), b) }
