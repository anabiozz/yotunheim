package endpoints

import (
	"github.com/anabiozz/yotunheim/backend/common"
	"github.com/anabiozz/yotunheim/backend/common/datastore"
	"github.com/anabiozz/yotunheim/backend/internal/config"
	"github.com/anabiozz/yotunheim/backend/metrics"

	"github.com/kataras/iris"
)

// GetJSONnEndpoint ...
func GetJSONnEndpoint(e *common.Env, newConfig *config.Config) iris.Handler {
	return func(ctx iris.Context) {

		influxMetrics := datastore.InfluxResult{}

		metricChannel := make(chan datastore.InfluxMetrics, 100)

		for _, input := range newConfig.Inputs {
			acc := metrics.NewAccumulator(input, metricChannel)
			input.Metrics.Gather(e.DB, acc)
			result := <-metricChannel
			influxMetrics.Metrics = append(influxMetrics.Metrics, result)
		}

		if _, err := ctx.JSON(influxMetrics); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	}
}
