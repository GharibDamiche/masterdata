package data_value

import (
	"context"
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

func (r DataValueRepository) GetByKey(ctx context.Context, key string) (*DataValue, error) {
	dataValue := &DataValue{Model: arangodb.Model{Id: key}}

	err := r.Read(ctx, dataValue)
	if err != nil {
		return nil, err
	}

	return dataValue, nil
}
