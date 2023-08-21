package measurement_unit

import (
	"github.com/arangodb/go-driver"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
)

type MeasurementUnitRepositoryInterface interface {
	arangodb.RepositoryInterface[*MeasurementUnit]
}

type MeasurmentUnitRepository struct {
	arangodb.Repository[*MeasurementUnit]
}

// @Service()
func NewMeasurementUnitRepository(db driver.Database) (MeasurementUnitRepositoryInterface, error) {
	return &MeasurmentUnitRepository{arangodb.NewRepository[*MeasurementUnit](db, Collection, nil)}, nil
}
