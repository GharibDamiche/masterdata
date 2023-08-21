package manager

import (
	"context"
	model "github.com/tmds-io/masterdata/service/model/measurement_unit"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/errors"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/service/telemetry/logger"
)

type MeasurementUnitManager struct {
	measurementUnits model.MeasurementUnitRepositoryInterface
}

// @Service()
func NewMeasurementUnitManager(
	measurementUnits model.MeasurementUnitRepositoryInterface,
	// Add other dependencies if needed
) *MeasurementUnitManager {
	return &MeasurementUnitManager{measurementUnits}
}

// ****************************************************** Public ******************************************************

// Create updates a measurement unit and returns the error logged if it happened
func (m *MeasurementUnitManager) Create(ctx context.Context, model *model.MeasurementUnit) error {
	if err := m.measurementUnits.Create(ctx, model); err != nil {
		err = errors.Wrap(err, 400) // Update the error code accordingly
		logger.Error(err)
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
func (m *MeasurementUnitManager) Search(ctx context.Context, filters interface{}) ([]*model.MeasurementUnit, error) {
	// Implement the search logic based on your database and filters
	return nil, nil
}
