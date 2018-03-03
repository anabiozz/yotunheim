package datastore

import "heimdall_project/yotunheim/backend/common/utility"

//
type Datastore interface {
}

//
type DatastoreErr struct {
	error
}

const (
	//
	INFLUXDB = iota
)

//
func NewDatastore(datastoreType int, dbConnectionString string) (Datastore, error) {
	switch datastoreType {
	case INFLUXDB:
		client, err := NewInfluxDatastore(dbConnectionString)
		if err != nil {
			return nil, DatastoreErr{(utility.WrapError(err, "cannot create new datastore"))}
		}
		return client, nil
	}
	return nil, nil
}
