package concat

import (
	"testing"
)

func generateTestData() []string {
	strs := make([]string, 30)
	for i := 0; i < 30; i++ {
		strs[i] = "some data 321 "
	}
	return strs
}

func BenchmarkConcat(b *testing.B) {
	strs := generateTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Concat(strs)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	strs := generateTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ConcatBuilder(strs)
	}
}

func BenchmarkConcatCopy(b *testing.B) {
	strs := generateTestData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ConcatCopy(strs)
	}
}
