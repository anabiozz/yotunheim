package backend

import (
	"github.com/anabiozz/yotunheim/backend/common/datastore"
)

// Gatherer ...
type Gatherer interface {
	Gather(c datastore.Datastore, acc Accumulator, getherTime string, groupby string)
}
