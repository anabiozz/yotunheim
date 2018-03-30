package backend

import (
	influx "github.com/influxdata/influxdb/client/v2"
)

const (
	Counter = "counter"
	Gauge
	Untyped
	Summary
	Histogram = "histogram"
	Table     = "table"
)

type Accumulator interface {
	AddLine(name string, metrics []influx.Result, err error)
	AddBar(name string, metrics []influx.Result, err error)
	AddTable(name string, metrics []influx.Result, err error)
}
