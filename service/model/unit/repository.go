package unit

import (
	"github.com/arangodb/go-driver"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
)

type UnitRepositoryInterface interface {
	arangodb.RepositoryInterface[*Unit]
}

type UnitRepository struct {
	arangodb.Repository[*Unit]
}

// @Service()
func NewUnitRepository(db driver.Database) (UnitRepositoryInterface, error) {
	return &UnitRepository{arangodb.NewRepository[*Unit](db, Collection, nil)}, nil
}
