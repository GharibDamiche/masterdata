package manager

import (
	"context"
	model "github.com/tmds-io/masterdata/service/model/data_value"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/errors"
)

type DataValueManager struct {
	dataValues model.DataValueRepositoryInterface
}

func NewDataValueManager(dataValues model.DataValueRepositoryInterface) *DataValueManager {
	return &DataValueManager{dataValues}
}

func (m *DataValueManager) Configure(ctx context.Context, model *model.DataValue) error {
	//TODO: Configure data value
	return nil
}

func (m *DataValueManager) Read(ctx context.Context, model *model.DataValue) error {
	err := m.dataValues.Read(ctx, model)
	if err != nil {
		return errors.Wrap(err, 400)
	}
	return nil
}

func (m *DataValueManager) Update(ctx context.Context, model *model.DataValue) error {
	if err := m.dataValues.Update(ctx, model); err != nil {
		return errors.Wrap(err, 400)
	}
	return nil
}

func (m *DataValueManager) Delete(ctx context.Context, id string) error {
	if err := m.dataValues.Delete(ctx, id); err != nil {
		return errors.Wrap(err, 400)
	}
	return nil
}

// Search searches for measurement units.
func (m *DataValueManager) Search(ctx context.Context, filters *arangodb.Filters) (arangodb.ResultSetInterface[*model.DataValue], error) {
	rs, err := m.dataValues.Search(ctx, filters)
	if err != nil {
		return nil, errors.Wrap(err, 400)
	}

	return rs, nil
}
