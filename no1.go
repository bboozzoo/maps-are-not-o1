package no1

import "sort"

type tester interface {
	has(int) bool
	max() int
}

type mapTester struct {
	m map[int]bool
	v int
}

func (m *mapTester) max() int {
	return m.v
}

func (m *mapTester) has(x int) bool {
	_, ok := m.m[x]
	return ok
}

type sliceTester struct {
	s []int
	v int
}

func (s *sliceTester) max() int {
	return s.v
}

func (s *sliceTester) has(x int) bool {
	i := sort.Search(len(s.s), func(i int) bool { return s.s[i] >= x })
	return i < len(s.s) && s.s[i] == x
}

func genMap(size int) tester {
	m := map[int]bool{}
	for i := 0; i < size; i++ {
		m[i] = true
	}

	return &mapTester{m: m, v: size}
}

func genSlice(size int) tester {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = i
	}

	return &sliceTester{s: s, v: size}
}
