package manager

import (
	"context"
	model "github.com/tmds-io/masterdata/service/model/data_definition"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/errors"
)

type DataDefinitionManager struct {
	dataDefinitions model.DataDefinitionRepositoryInterface
}

func NewDataDefinitionManager(dataDefinitions model.DataDefinitionRepositoryInterface) *DataDefinitionManager {
	return &DataDefinitionManager{dataDefinitions}
}

func (m *DataDefinitionManager) Create(ctx context.Context, model *model.DataDefinition) error {
	if err := m.dataDefinitions.Create(ctx, model); err != nil {
		return errors.Wrap(err, 400)
	}
	return nil
}

func (m *DataDefinitionManager) Read(ctx context.Context, model *model.DataDefinition) error {
	err := m.dataDefinitions.Read(ctx, model)
	if err != nil {
		return errors.Wrap(err, 400)
	}
	return nil
}

func (m *DataDefinitionManager) Update(ctx context.Context, model *model.DataDefinition) error {
	if err := m.dataDefinitions.Update(ctx, model); err != nil {
		return errors.Wrap(err, 400)
	}
	return nil
}

func (m *DataDefinitionManager) Delete(ctx context.Context, id string) error {
	if err := m.dataDefinitions.Delete(ctx, id); err != nil {
		return errors.Wrap(err, 400)
	}
	return nil
}

// Search searches for measurement units.
func (m *DataDefinitionManager) Search(ctx context.Context, filters *arangodb.Filters) (arangodb.ResultSetInterface[*model.DataDefinition], error) {
	rs, err := m.dataDefinitions.Search(ctx, filters)
	if err != nil {
		return nil, errors.Wrap(err, 400)
	}

	return rs, nil
}
