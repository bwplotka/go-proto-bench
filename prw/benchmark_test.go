package prw

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/bwplotka/go-proto-bench/prw/internal/base"
	"github.com/bwplotka/go-proto-bench/prw/internal/labels"
	"github.com/bwplotka/go-proto-bench/prw/v1testgogo"
	v2 "github.com/bwplotka/go-proto-bench/prw/v2"
	"github.com/bwplotka/go-proto-bench/prw/v2testgogo"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

var (
	testTime       = int64(1701256179047)
	testLabelNames = []string{
		"job",
		"environment",
		"cluster",
		"pod_name",
		"instance",
		"container_name",
		"service_job",
		"reason",
		"code",
		"beta_kubernetes_io_instance_type",
		"beta_kubernetes_io_os",
		"failure_domain_beta_kubernetes_io_region",
		"kubernetes_io_hostname",
	}
)

func newTestWriteRequest(series, lbls, samples int) *base.WriteRequest {
	wr := &base.WriteRequest{Timeseries: make([]base.TimeSeries, series)}

	for i := range wr.Timeseries {
		wr.Timeseries[i].Labels = make([]base.Label, lbls)
		lbls := labels.NewBuilder(labels.EmptyLabels())
		for j := range wr.Timeseries[i].Labels {
			if j == 0 {
				lbls.Set(labels.MetricName, fmt.Sprintf("test_metric_name_prometheus_counter%v_total", i%100))
				continue
			}
			if j >= len(testLabelNames) {
				panic("generate more testLabelNames")
			}
			// Random UUID as value for now to simulate some compressions challenges (:
			lbls.Set(testLabelNames[j-1], uuid.New().String())
		}
		j := 0
		lbls.Labels().Range(func(l labels.Label) {
			wr.Timeseries[i].Labels[j].Name = l.Name
			wr.Timeseries[i].Labels[j].Value = l.Value
			j++
		})

		wr.Timeseries[i].Samples = make([]base.Sample, samples)
		for j := range wr.Timeseries[i].Samples {
			wr.Timeseries[i].Samples[j].Timestamp = testTime + int64(j*1e3)
			wr.Timeseries[i].Samples[j].Value = 1e6 + 4223 + float64(j*1e3)
		}
	}
	return wr
}

var (
	protoChoice = flag.String("proto", "", "name of the proto package from prw/ directory e.g. v2testgogo")
	b           base.Benchmarkable
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func newBenchmarkable(tb testing.TB, wr *base.WriteRequest) base.Benchmarkable {
	if len(*protoChoice) == 0 {
		fmt.Println("use --proto flag to choose the proto package to test/benchmark")
		os.Exit(1)
	}

	switch *protoChoice {
	case "v1testgogo":
		return v1testgogo.NewBenchmarkable(tb, wr)
	case "v2testgogo":
		return v2testgogo.NewBenchmarkable(tb, wr)
	case "v2":
		return v2.NewBenchmarkable(tb, wr)
	default:
		tb.Fatal(*protoChoice, "is not supported. Use --proto flag to choose the proto package to test/benchmark")
		return nil
	}
}

func TestAll(t *testing.T) {
	wr := newTestWriteRequest(2000, 10, 1)

	t.Run("v1testgogo", func(t *testing.T) {
		be := v1testgogo.NewBenchmarkable(t, wr)
		out := be.DeserializeToBase(be.Serialize())
		if diff := cmp.Diff(wr, out); diff != "" {
			t.Fatal("got different base write request after serialization->deserialization", diff)
		}
		// Do it again to ensure nothing is shared.
		out = be.DeserializeToBase(be.Serialize())
		if diff := cmp.Diff(wr, out); diff != "" {
			t.Fatal("got different base write request after serialization->deserialization 2x", diff)
		}
	})

	t.Run("v2testgogo", func(t *testing.T) {
		be := v2testgogo.NewBenchmarkable(t, wr)
		out := be.DeserializeToBase(be.Serialize())
		if diff := cmp.Diff(wr, out); diff != "" {
			t.Fatal("got different base write request after serialization->deserialization", diff)
		}
		// Do it again to ensure nothing is shared.
		out = be.DeserializeToBase(be.Serialize())
		if diff := cmp.Diff(wr, out); diff != "" {
			t.Fatal("got different base write request after serialization->deserialization 2x", diff)
		}
	})

	t.Run("v2", func(t *testing.T) {
		be := v2.NewBenchmarkable(t, wr)
		out := be.DeserializeToBase(be.Serialize())
		if diff := cmp.Diff(wr, out); diff != "" {
			t.Fatal("got different base write request after serialization->deserialization", diff)
		}
		// Do it again to ensure nothing is shared.
		out = be.DeserializeToBase(be.Serialize())
		if diff := cmp.Diff(wr, out); diff != "" {
			t.Fatal("got different base write request after serialization->deserialization 2x", diff)
		}
	})
}

func BenchmarkPRWSerialize(b *testing.B) {
	// 2000 series as 2000 is a default MaxSampleSend option in Prometheus
	// and typically we will see 1 sample sent for each series.
	wr := newTestWriteRequest(2000, 10, 1)
	be := newBenchmarkable(b, wr)

	b.ResetTimer()

	totalSize := 0
	for i := 0; i < b.N; i++ {
		out := be.Serialize()
		totalSize += len(out)
		b.ReportMetric(float64(totalSize)/float64(b.N), "serializedSize/op")
	}
}

var Sink any

func BenchmarkPRWDeserialize(b *testing.B) {
	// 2000 series as 2000 is a default MaxSampleSend option in Prometheus
	// and typically we will see 1 sample sent for each series.
	wr := newTestWriteRequest(2000, 10, 1)
	be := newBenchmarkable(b, wr)
	out := be.Serialize()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sink = be.Deserialize(out)
	}
}

func BenchmarkPRWDeserializeToBase(b *testing.B) {
	// 2000 series as 2000 is a default MaxSampleSend option in Prometheus
	// and typically we will see 1 sample sent for each series.
	wr := newTestWriteRequest(2000, 10, 1)
	be := newBenchmarkable(b, wr)
	out := be.Serialize()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sink = be.DeserializeToBase(out)
	}
}
