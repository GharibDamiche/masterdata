package data_definition

import (
	"github.com/arangodb/go-driver"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
)

type DataDefinitionRepositoryInterface interface {
	arangodb.RepositoryInterface[*DataDefinition]
}

type DataDefinitionRepository struct {
	arangodb.Repository[*DataDefinition]
}

// @Service()
func NewDataDefinitionRepository(db driver.Database) (DataDefinitionRepositoryInterface, error) {
	return &DataDefinitionRepository{arangodb.NewRepository[*DataDefinition](db, Collection, nil)}, nil
}
