package backend

import (
	influx "github.com/influxdata/influxdb/client/v2"
)

type Accumulator interface {
	AddLine(name string, metrics []influx.Result, err error)
	AddBar(name string, metrics []influx.Result, err error)
}
