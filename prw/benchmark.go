package prw

type v1WriteRequest struct {
	Timeseries []v1TimeSeries
}

type v1TimeSeries struct {
	Labels  []v1Label
	Samples []Sample
}

type v1Label struct {
	Name  string
	Value string
}

type Sample struct {
	Value     float64
	Timestamp int64
}

type v2WriteRequest struct {
	// The symbols table for all strings used in WriteRequest message.
	Symbols    []string
	Timeseries []*v2TimeSeries
}

type v2TimeSeries struct {
	// Sorted list of label name-value pair references. This list's len is always multiple of 2,
	// packing tuples of (label name ref, label value ref) refs to WriteRequests.symbols.
	LabelSymbols []uint32
	// Sorted by time, oldest sample first.
	Samples []*Sample
}
