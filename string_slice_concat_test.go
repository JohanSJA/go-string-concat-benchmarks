// Package stringconcat exists only to provide benchmarks for the different approaches
// to string concatenation in Go.
package stringconcat

import (
	"bytes"
	"strings"
	"testing"
)

// benchmarkNaiveConcat provides a benchmark for basic built-in
// Go string concatenation. Because strings are immutable in Go,
// it performs the worst of the tested methods. The time taken to
// set up the array that is appended is not counted towards the
// time for naive concatenation.
func benchmarkSliceNaiveConcat(b *testing.B, numConcat int) {
	slice := make([]string, numConcat)
	for i := 0; i < numConcat; i++ {
		slice[i] = nextString()()
	}

	b.ResetTimer()
	// Reports memory allocations
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ns := ""
		for i := range slice {
			ns += slice[i]
		}
		// we assign to a global variable to make sure compiler
		// or runtime optimizations don't skip over the operations
		// we were benchmarking. This might be unnecessary, but it's
		// safe.
		global = ns
	}
}

func BenchmarkSliceNaiveConcat10(b *testing.B) {
	benchmarkNaiveConcat(b, 10)
}

func BenchmarkSliceNaiveConcat100(b *testing.B) {
	benchmarkNaiveConcat(b, 100)
}

func BenchmarkSliceNaiveConcat1000(b *testing.B) {
	benchmarkNaiveConcat(b, 1000)
}

func BenchmarkSliceNaiveConcat10000(b *testing.B) {
	benchmarkNaiveConcat(b, 10000)
}

// benchmarkByteSlice provides a benchmark for the time it takes
// to repeatedly append returned strings to a byte slice, and
// finally casting the byte slice to string type.
func benchmarkSliceByteSlice(b *testing.B, numConcat int) {
	slice := make([]string, numConcat)
	for i := 0; i < numConcat; i++ {
		slice[i] = nextString()()
	}

	b.ResetTimer()
	// Reports memory allocations
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b := []byte{}
		for i := range slice {
			b = append(b, []byte(slice[i])...)
		}
		global = string(b)
	}
}

func BenchmarkSliceByteSlice10(b *testing.B) {
	benchmarkByteSlice(b, 10)
}

func BenchmarkSliceByteSlice100(b *testing.B) {
	benchmarkByteSlice(b, 100)
}

func BenchmarkSliceByteSlice1000(b *testing.B) {
	benchmarkByteSlice(b, 1000)
}

func BenchmarkSliceByteSlice10000(b *testing.B) {
	benchmarkByteSlice(b, 10000)
}

// benchmarkJoin provides a benchmark for the time it takes to set
// up an array with strings, and calling strings.Join on that array
// to get a fully concatenated string.
func benchmarkSliceJoin(b *testing.B, numConcat int) {
	slice := make([]string, numConcat)
	for i := 0; i < numConcat; i++ {
		slice[i] = nextString()()
	}

	b.ResetTimer()
	// Reports memory allocations
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ns := strings.Join(slice, "")
		global = ns
	}
}

func BenchmarkSliceJoin10(b *testing.B) {
	benchmarkJoin(b, 10)
}

func BenchmarkSliceJoin100(b *testing.B) {
	benchmarkJoin(b, 100)
}

func BenchmarkSliceJoin1000(b *testing.B) {
	benchmarkJoin(b, 1000)
}

func BenchmarkSliceJoin10000(b *testing.B) {
	benchmarkJoin(b, 10000)
}

// benchmarkBufferString
func benchmarkSliceBufferString(b *testing.B, numConcat int) {
	slice := make([]string, numConcat)
	for i := 0; i < numConcat; i++ {
		slice[i] = nextString()()
	}

	b.ResetTimer()
	// Reports memory allocations
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buffer := bytes.NewBufferString("")
		for i := range slice {
			buffer.WriteString(slice[i])
		}
		global = buffer.String()
	}
}

func BenchmarkSliceBufferString10(b *testing.B) {
	benchmarkBufferString(b, 10)
}

func BenchmarkSliceBufferString100(b *testing.B) {
	benchmarkBufferString(b, 100)
}

func BenchmarkSliceBufferString1000(b *testing.B) {
	benchmarkBufferString(b, 1000)
}

func BenchmarkSliceBufferString10000(b *testing.B) {
	benchmarkBufferString(b, 10000)
}
