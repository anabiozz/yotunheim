package backend

import "github.com/anabiozz/yotunheim/backend/common/datastore"

const (
	// Counter ...
	Counter = "counter"
	// Gauge ...
	Gauge
	// Untyped ...
	Untyped
	// Summary ...
	Summary
	// Histogram ...
	Histogram = "histogram"
	// Table ...
	Table = "table"
)

// Accumulator ...
type Accumulator interface {
	AddMetric(datastore.InfluxMetrics)
	AddTable(datastore.InfluxMetrics)
}
