package main

import "testing"

func BenchmarkIcons_newIcons(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newIcons()
	}
}

func BenchmarkIcons_find(b *testing.B) {
	ics := newIcons()
	terms := []string{"a", "b", "c"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ics.find(terms)
	}
}
