package trialBigOWithoutElse

import (
	"testing"
)

// from trialBigOWithoutElse.go
func BenchmarkGetListUpto10WithoutElse(b *testing.B) {
	// run the GetListUpto10WithoutElse function b.N times
	for n := 0; n < b.N; n++ {
		GetListUpto10WithoutElse()
	}
}

