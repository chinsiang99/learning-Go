package iteration

import "testing"

func TestRepeat(t *testing.T) {
	result := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// To run the benchmarks do go test -bench=. -benchmem(or if you're in Windows Powershell go test -bench=".")
// We can use BenchmarkRepeat to confirm that strings.Builder significantly improves performance.
// Run go test -bench=. -benchmem:

// Copy
// goos: darwin
// goarch: amd64
// pkg: github.com/quii/learn-go-with-tests/for/v4
// 10000000           25.70 ns/op           8 B/op           1 allocs/op
// PASS
// The -benchmem flag reports information about memory allocations:

// B/op: the number of bytes allocated per iteration

// allocs/op: the number of memory allocations per iteration
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}
