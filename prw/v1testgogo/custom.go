package v1testgogo

import (
	"testing"

	"github.com/bwplotka/go-proto-bench/prw/internal/base"
	gogoproto "github.com/gogo/protobuf/proto"
)

type b struct {
	tb testing.TB

	src *WriteRequest
}

func NewBenchmarkable(tb testing.TB, wr *base.WriteRequest) base.Benchmarkable {
	return &b{tb: tb, src: fromBase(wr)}
}

func (b *b) Serialize() []byte {
	out, err := gogoproto.Marshal(b.src)
	if err != nil {
		b.tb.Fatal(err)
	}
	return out
}

func (b *b) Deserialize(in []byte) base.Sizer {
	obj := &WriteRequest{}
	if err := gogoproto.Unmarshal(in, obj); err != nil {
		b.tb.Fatal(err)
	}
	return obj
}

func (b *b) DeserializeToBase(in []byte) *base.WriteRequest {
	obj := &WriteRequest{}
	if err := gogoproto.Unmarshal(in, obj); err != nil {
		b.tb.Fatal(err)
	}
	return toBase(obj)
}

func fromBase(wr *base.WriteRequest) *WriteRequest {
	ret := &WriteRequest{Timeseries: make([]TimeSeries, len(wr.Timeseries))}
	for i := range wr.Timeseries {
		lbls := make([]Label, len(wr.Timeseries[i].Labels))
		for j := range wr.Timeseries[i].Labels {
			lbls[j].Name = wr.Timeseries[i].Labels[j].Name
			lbls[j].Value = wr.Timeseries[i].Labels[j].Value
		}
		ret.Timeseries[i].Labels = lbls

		s := make([]Sample, len(wr.Timeseries[i].Samples))
		for j := range wr.Timeseries[i].Samples {
			s[j].Value = wr.Timeseries[i].Samples[j].Value
			s[j].Timestamp = wr.Timeseries[i].Samples[j].Timestamp
		}
		ret.Timeseries[i].Samples = s
	}
	return ret
}

func toBase(wr *WriteRequest) *base.WriteRequest {
	ret := &base.WriteRequest{Timeseries: make([]base.TimeSeries, len(wr.Timeseries))}
	for i := range wr.Timeseries {
		lbls := make([]base.Label, len(wr.Timeseries[i].Labels))
		for j := range wr.Timeseries[i].Labels {
			lbls[j].Name = wr.Timeseries[i].Labels[j].Name
			lbls[j].Value = wr.Timeseries[i].Labels[j].Value
		}
		ret.Timeseries[i].Labels = lbls

		s := make([]base.Sample, len(wr.Timeseries[i].Samples))
		for j := range wr.Timeseries[i].Samples {
			s[j].Value = wr.Timeseries[i].Samples[j].Value
			s[j].Timestamp = wr.Timeseries[i].Samples[j].Timestamp
		}
		ret.Timeseries[i].Samples = s
	}
	return ret
}
