package backend

import "heimdall_project/yotunheim/backend/common/datastore"

type Gatherer interface {
	Gather(c datastore.Datastore, acc Accumulator)
}