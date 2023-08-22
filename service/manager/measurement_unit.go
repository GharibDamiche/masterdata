package manager

import (
	"context"
	model "github.com/tmds-io/masterdata/service/model/measurement_unit"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/errors"
)

type MeasurementUnitManager struct {
	measurementUnits model.MeasurementUnitRepositoryInterface
}

// @Service()
func NewMeasurementUnitManager(measurementUnits model.MeasurementUnitRepositoryInterface) *MeasurementUnitManager {
	return &MeasurementUnitManager{measurementUnits}
}

// ****************************************************** Public ******************************************************

// Create updates a measurement unit and returns the error logged if it happened
func (m *MeasurementUnitManager) Create(ctx context.Context, model *model.MeasurementUnit) error {
	if err := m.measurementUnits.Create(ctx, model); err != nil {
		err = errors.Wrap(err, 400) // Update the error code accordingly
		return err
	}
	return nil
}

// Read reads a measurement unit.
func (m *MeasurementUnitManager) Read(ctx context.Context, model *model.MeasurementUnit) error {
	err := m.measurementUnits.Read(ctx, model)
	if err != nil {
		return errors.Wrap(err, 400) // Update the error codes accordingly
	}
	return nil
}

// Update creates or updates a measurement unit
func (m *MeasurementUnitManager) Update(ctx context.Context, model *model.MeasurementUnit) error {
	err := m.measurementUnits.Update(ctx, model)
	if err != nil {
		return errors.Wrap(err, 400) // Update the error code accordingly
	}
	return nil
}

// Delete deletes a measurement unit
func (m *MeasurementUnitManager) Delete(ctx context.Context, id string) error {
	if err := m.measurementUnits.Delete(ctx, id); err != nil {
		return errors.Wrap(err, 400) // Update the error code accordingly
	}
	return nil
}

// Search searches for measurement units.
func (m *MeasurementUnitManager) Search(ctx context.Context, filters *arangodb.Filters) (arangodb.ResultSetInterface[*model.MeasurementUnit], error) {
	rs, err := m.measurementUnits.Search(ctx, filters)
	if err != nil {
		return nil, errors.Wrap(err, 400)
	}

	return rs, nil
}
