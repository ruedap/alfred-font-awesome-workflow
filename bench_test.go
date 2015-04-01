package main

import "testing"

func BenchmarkIcons_NewIcons(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewIcons()
	}
}
