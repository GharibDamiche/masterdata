package manager

import (
	"context"
	model "github.com/tmds-io/masterdata/service/model/segment"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/errors"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/service/telemetry/logger"
)

type SegmentManager struct {
	segments model.SegmentRepositoryInterface
}

// @Service()
func NewSegmentManager(
	segments model.SegmentRepositoryInterface,
	// Add other dependencies if needed
) *SegmentManager {
	return &SegmentManager{segments}
}

// ****************************************************** Public ******************************************************

// Create updates a measurement unit and returns the error logged if it happened
func (m *SegmentManager) Create(ctx context.Context, model *model.Segment) error {
	if err := m.segments.Create(ctx, model); err != nil {
		err = errors.Wrap(err, 400) // Update the error code accordingly
		logger.Error(err)
		return err
	}
	return nil
}

// Read reads a measurement unit.
func (m *SegmentManager) Read(ctx context.Context, model *model.Segment) error {
	err := m.segments.Read(ctx, model)
	if err != nil {
		return errors.Wrap(err, 400) // Update the error codes accordingly
	}
	return nil
}

// Update creates or updates a measurement unit
func (m *SegmentManager) Update(ctx context.Context, model *model.Segment) error {
	err := m.segments.Update(ctx, model)
	if err != nil {
		return errors.Wrap(err, 400) // Update the error code accordingly
	}
	return nil
}

// Delete deletes a measurement unit
func (m *SegmentManager) Delete(ctx context.Context, id string) error {
	if err := m.segments.Delete(ctx, id); err != nil {
		return errors.Wrap(err, 400) // Update the error code accordingly
	}
	return nil
}

// Search searches for measurement units.
func (m *SegmentManager) Search(ctx context.Context, filters *arangodb.Filters) (arangodb.ResultSetInterface[*model.Segment], error) {
	rs, err := m.segments.Search(ctx, filters)
	if err != nil {
		return nil, errors.Wrap(err, 400)
	}

	return rs, nil
}
