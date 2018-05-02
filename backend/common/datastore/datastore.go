package datastore

import "github.com/anabiozz/yotunheim/backend/common/utility"

// Datastore ...
type Datastore interface{}

// Err ...
type Err struct {
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
			return nil, Err{(utility.WrapError(err, "cannot create new datastore"))}
		}
		return client, nil
	}
	return nil, nil
}
