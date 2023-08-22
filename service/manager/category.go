package manager

import (
	"context"
	model "github.com/tmds-io/masterdata/service/model/category"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/errors"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/service/telemetry/logger"
)

type CategoryManager struct {
	categories model.CategoryRepositoryInterface
}

// @Service()
func NewCategoryManager(
	categories model.CategoryRepositoryInterface,
	// Add other dependencies if needed
) *CategoryManager {
	return &CategoryManager{categories}
}

// ****************************************************** Public ******************************************************

// Create updates a measurement unit and returns the error logged if it happened
func (m *CategoryManager) Create(ctx context.Context, model *model.Category) error {
	if err := m.categories.Create(ctx, model); err != nil {
		err = errors.Wrap(err, 400) // Update the error code accordingly
		logger.Error(err)
		return err
	}
	return nil
}

// Read reads a measurement unit.
func (m *CategoryManager) Read(ctx context.Context, model *model.Category) error {
	err := m.categories.Read(ctx, model)
	if err != nil {
		return errors.Wrap(err, 400) // Update the error codes accordingly
	}
	return nil
}

// Update creates or updates a measurement unit
func (m *CategoryManager) Update(ctx context.Context, model *model.Category) error {
	err := m.categories.Update(ctx, model)
	if err != nil {
		return errors.Wrap(err, 400) // Update the error code accordingly
	}
	return nil
}

// Delete deletes a measurement unit
func (m *CategoryManager) Delete(ctx context.Context, id string) error {
	if err := m.categories.Delete(ctx, id); err != nil {
		return errors.Wrap(err, 400) // Update the error code accordingly
	}
	return nil
}

// Search searches for measurement units.
func (m *CategoryManager) Search(ctx context.Context, filters *arangodb.Filters) (arangodb.ResultSetInterface[*model.Category], error) {
	rs, err := m.categories.Search(ctx, filters)
	if err != nil {
		return nil, errors.Wrap(err, 400)
	}

	return rs, nil
}
