package v2testvtproto

import (
	"testing"

	"github.com/bwplotka/go-proto-bench/prw/internal/base"
)

type b struct {
	tb testing.TB

	src *WriteRequest
}

func (x *WriteRequest) Size() int {
	return x.SizeVT()
}

// TODO(bwplotka): Try with pooling
func NewBenchmarkable(tb testing.TB, wr *base.WriteRequest) base.Benchmarkable {
	return &b{tb: tb, src: fromBase(wr)}
}

func (b *b) Serialize() []byte {
	out, err := b.src.MarshalVT()
	if err != nil {
		b.tb.Fatal(err)
	}
	return out
}

func (b *b) Deserialize(in []byte) base.Sizer {
	obj := &WriteRequest{}
	if err := obj.UnmarshalVTUnsafe(in); err != nil {
		b.tb.Fatal(err)
	}
	return obj
}

func (b *b) DeserializeToBase(in []byte) *base.WriteRequest {
	obj := &WriteRequest{}
	if err := obj.UnmarshalVTUnsafe(in); err != nil {
		b.tb.Fatal(err)
	}
	return toBase(obj)
}

func fromBase(wr *base.WriteRequest) *WriteRequest {
	symbolsMap := map[string]uint32{}

	ret := &WriteRequest{Timeseries: make([]*TimeSeries, len(wr.Timeseries))}
	for i := range wr.Timeseries {
		lbls := make([]uint32, 0, 2*len(wr.Timeseries[i].Labels))
		for j := range wr.Timeseries[i].Labels {
			str := wr.Timeseries[i].Labels[j].Name
			ref, ok := symbolsMap[str]
			if !ok {
				ret.Symbols = append(ret.Symbols, str)
				ref = uint32(len(ret.Symbols) - 1)
				symbolsMap[str] = ref
			}
			lbls = append(lbls, ref)

			str2 := wr.Timeseries[i].Labels[j].Value
			ref2, ok := symbolsMap[str2]
			if !ok {
				ret.Symbols = append(ret.Symbols, str2)
				ref2 = uint32(len(ret.Symbols) - 1)
				symbolsMap[str2] = ref2
			}
			lbls = append(lbls, ref2)
		}

		s := make([]*Sample, len(wr.Timeseries[i].Samples))
		for j := range wr.Timeseries[i].Samples {
			s[j] = &Sample{Value: wr.Timeseries[i].Samples[j].Value, Timestamp: wr.Timeseries[i].Samples[j].Timestamp}
		}

		ret.Timeseries[i] = &TimeSeries{LabelSymbols: lbls, Samples: s}
	}
	return ret
}

func toBase(wr *WriteRequest) *base.WriteRequest {
	ret := &base.WriteRequest{Timeseries: make([]base.TimeSeries, len(wr.Timeseries))}
	for i := range wr.Timeseries {
		lbls := make([]base.Label, len(wr.Timeseries[i].LabelSymbols)/2)

		sID := 0
		for j := range lbls {
			lbls[j].Name = wr.Symbols[wr.Timeseries[i].LabelSymbols[sID]]
			sID++
			lbls[j].Value = wr.Symbols[wr.Timeseries[i].LabelSymbols[sID]]
			sID++
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
