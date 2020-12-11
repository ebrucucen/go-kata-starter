package trialBigOWithElse

import (
	"testing"
)

// from trialBigOWithElse.go
func BenchmarkGetListUpto10WithElse(b *testing.B) {
	// run the GetListUpto10WithElse function b.N times
	for n := 0; n < b.N; n++ {
		GetListUpto10WithElse()
	}
}
