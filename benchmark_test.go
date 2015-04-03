package main

import "testing"

func BenchmarkIcons_newIcons(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newIcons()
	}
}
