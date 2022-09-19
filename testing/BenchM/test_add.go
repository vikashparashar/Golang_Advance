package testing

import (
	"testing"

	add "github.com/gic-vikash/golang-advance/testing/Add"
)

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add.Addition(1, 2)
	}
}

// Lets Try Becnchmarking