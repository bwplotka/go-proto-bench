package prw

import (
	"testing"
)

var testTime = t

func newTestV1WriteRequest(series, labels, samples int) *v1WriteRequest {
	wr := &v1WriteRequest{Timeseries: make([]v1TimeSeries, series)}

	for i := range wr.Timeseries {
		wr.Timeseries[i].Samples

	}
}

func BenchmarkPRWSerialize(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}

func BenchmarkPRWDeserialize(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
