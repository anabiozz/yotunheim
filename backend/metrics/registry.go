package metrics

import (
	"github.com/anabiozz/yotunheim/backend"
)

type Creator func() backend.Gatherer

var Metrics = map[string]Creator{}

func Add(name string, creator Creator) {
	Metrics[name] = creator
}
