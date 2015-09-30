package no1

import (
	"math/rand"
	"runtime"
	"testing"
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

func shuffledRange(start, end int) []int {
	s := make([]int, end-start)
	for i := 0; start+i < end; i++ {
		s[i] = i
	}
	for i := range s {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func benchmark(r tester, b *testing.B) {
	runtime.GC()
	max := r.max()
	is := shuffledRange(0, max)
	b.ResetTimer()

	var has bool
	for i := 0; i < b.N; i++ {
		for _, i := range is {
			has = r.has(i)
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
