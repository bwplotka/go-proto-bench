package base

type WriteRequest struct {
	Timeseries []TimeSeries
}

type TimeSeries struct {
	Labels  []Label
	Samples []Sample
}

type Label struct {
	Name  string
	Value string
}

type Sample struct {
	Value     float64
	Timestamp int64
}

type Sizer interface {
	Size() int
}

type Benchmarkable interface {
	Serialize() []byte
	Deserialize([]byte) Sizer
	DeserializeToBase([]byte) *WriteRequest
}
