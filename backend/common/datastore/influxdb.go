package datastore

import (
	"github.com/anabiozz/yotunheim/backend/common/utility"

	influx "github.com/influxdata/influxdb/client/v2"
)

// InfluxMetrics ...
type InfluxMetrics struct {
	Metric    []TableMetrics
	ChartType string
	ChartName string
}

// TableMetrics
type TableMetrics struct {
	Titles []string
	Value  [][]interface{}
}

// InfluxErr ...
type InfluxErr struct {
	error
}

// Influx ...
type Influx struct {
	Client influx.Client
}

// NewInfluxDatastore ctreate new influxdb client
func NewInfluxDatastore(dbConnectionString string) (client influx.Client, err error) {
	client, err = influx.NewHTTPClient(influx.HTTPConfig{
		Addr: dbConnectionString,
	})
	if err != nil {
		return nil, utility.WrapError(err, err.Error())
	}
	return client, nil

}

// QueryDB convenience function to query the database
func QueryDB(clnt influx.Client, cmd string) (res []influx.Result, err error) {
	q := influx.Query{
		Command:  cmd,
		Database: "telegraf",
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, utility.WrapError(response.Error(), response.Err)
		}
		res = response.Results
	} else {
		return res, utility.WrapError(err, err.Error())
	}
	return res, nil
}
