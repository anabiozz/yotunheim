package backend

import "github.com/anabiozz/yotunheim/backend/common/datastore"

type Gatherer interface {
	Gather(c datastore.Datastore, acc Accumulator)
}
