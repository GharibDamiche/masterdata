package data_value

import (
	"github.com/arangodb/go-driver"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
)

type DataValueRepositoryInterface interface {
	arangodb.RepositoryInterface[*DataValue]
}

type DataValueRepository struct {
	arangodb.Repository[*DataValue]
}

// @Service()
func NewDataValueRepository(db driver.Database) (DataValueRepositoryInterface, error) {
	return &DataValueRepository{arangodb.NewRepository[*DataValue](db, Collection, nil)}, nil
}
