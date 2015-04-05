package main

import "testing"

func BenchmarkIcons_newIcons(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newIcons()
	}
}

func BenchmarkIcons_iconsReadYaml(b *testing.B) {
	p := iconsYamlPath()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		iconsReadYaml(p)
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

func BenchmarkIcons_Sort(b *testing.B) {
	p := iconsYamlPath()
	by, _ := iconsReadYaml(p)
	y, _ := iconsUnmarshalYaml(by)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		y.Icons.Sort()
	}
}
