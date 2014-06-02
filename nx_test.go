package gonx

import (
	"testing"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New("./Data/Base.nx")
	}
}
