package datastore

import (
	"heimdall_project/yotunheim/backend/common/utility"

	influx "github.com/influxdata/influxdb/client/v2"
)

// InfluxMetrics ...
type InfluxMetrics struct {
	Metrics   map[string][]interface{}
	ChartType []string
}

// InfluxErr ...
type InfluxErr struct {
	error
}

// InfluxMetricItem ...
type InfluxMetricItem struct {
	Timestamp interface{} `json:"timestamp"`
	Payload   interface{} `json:"payload"`
	Type      interface{} `json:"tipe"`
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

// CPUUsageInfluxQuery return cpu usage responce
func CPUUsageInfluxQuery(c Datastore) (res []influx.Result, err error) {
	res, err = queryDB(c.(influx.Client), "SELECT (100 - usage_idle) as cpu_usage from cpu WHERE time >= now() - 60s AND cpu = 'cpu-total'")
	if err != nil {
		return nil, utility.WrapError(err, err.Error())
	}
	return res, nil
}

// MemUsageInfluxQuery return cpu usage responce
func MemUsageInfluxQuery(c Datastore) (res []influx.Result, err error) {
	res, err = queryDB(c.(influx.Client), "SELECT used_percent as mem_usage from mem WHERE time >= now() - 60s")
	if err != nil {
		return nil, utility.WrapError(err, err.Error())
	}
	return res, nil
}

// DiskUsageInfluxQuery return cpu usage responce
func DiskUsageInfluxQuery(c Datastore) (res []influx.Result, err error) {
	res, err = queryDB(c.(influx.Client), "SELECT mean(used_percent) as disk_usage from disk WHERE time >= now() - 60s and device = 'sda2'")
	if err != nil {
		return nil, utility.WrapError(err, err.Error())
	}
	return res, nil
}

// queryDB convenience function to query the database
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
