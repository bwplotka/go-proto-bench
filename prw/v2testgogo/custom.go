package v2testgogo

import (
	"testing"

	"github.com/bwplotka/go-proto-bench/prw/internal/base"
	gogoproto "github.com/gogo/protobuf/proto"

	"golang.org/x/exp/slices"
)

type b struct {
	tb testing.TB

	src *WriteRequest
}

func NewBenchmarkable(tb testing.TB, wr *base.WriteRequest) base.Benchmarkable {
	return &b{tb: tb, src: fromBase(wr)}
}

func NewOptimizedBenchmarkable(tb testing.TB, wr *base.WriteRequest) base.Benchmarkable {
	return &optimizedB{b: b{tb: tb, src: fromBase(wr)}}
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
	symbolsMap := map[string]uint32{}

	ret := &WriteRequest{Timeseries: make([]TimeSeries, len(wr.Timeseries))}
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
		ret.Timeseries[i].LabelSymbols = lbls

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

type optimizedB struct {
	b
}

func (b *optimizedB) Serialize() []byte {
	out, err := b.src.OptimizedMarshal()
	if err != nil {
		b.tb.Fatal(err)
	}
	return out
}

/*
OptimizedMarshal is an optimized marshal to avoid extra allocs:
github.com/bwplotka/go-proto-bench/prw/v2testgogo.(*TimeSeries).MarshalToSizedBuffer
/Users/bwplotka/Repos/go-proto-bench/prw/v2testgogo/write.pb.go (allocs)

	Total:   328529070  328529070 (flat, cum) 99.55%
	  313            .          .           			i--
	  314            .          .           			dAtA[i] = 0x12
	  315            .          .           		}
	  316            .          .           	}
	  317            .          .           	if len(m.LabelSymbols) > 0 {
	  318    328529070  328529070           		dAtA2 := make([]byte, len(m.LabelSymbols)*10)
	  319            .          .           		var j1 int
	  320            .          .           		for _, num := range m.LabelSymbols {
	  321            .          .           			for num >= 1<<7 {
	  322            .          .           				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
	  323            .          .           				num >>= 7
*/
func (m *WriteRequest) OptimizedMarshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.OptimizedMarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WriteRequest) OptimizedMarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Timeseries) > 0 {
		for iNdEx := len(m.Timeseries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Timeseries[iNdEx].OptimizedMarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWrite(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Symbols) > 0 {
		for iNdEx := len(m.Symbols) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Symbols[iNdEx])
			copy(dAtA[i:], m.Symbols[iNdEx])
			i = encodeVarintWrite(dAtA, i, uint64(len(m.Symbols[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TimeSeries) OptimizedMarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Samples) > 0 {
		for iNdEx := len(m.Samples) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Samples[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWrite(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.LabelSymbols) > 0 {
		// This is the trick: encode the varints in reverse order to make it easier
		// to do it in place. Then reverse the whole thing.
		var j10 int
		start := i
		for _, num := range m.LabelSymbols {
			for num >= 1<<7 {
				dAtA[i-1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				i--
				j10++
			}
			dAtA[i-1] = uint8(num)
			i--
			j10++
		}
		slices.Reverse(dAtA[i:start])
		// --- end of trick

		i = encodeVarintWrite(dAtA, i, uint64(j10))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
