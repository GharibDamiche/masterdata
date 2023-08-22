package handler

import (
	"context"
	proto "github.com/tmds-io/masterdata/contract/data_value"
	"github.com/tmds-io/masterdata/service/manager"
	model "github.com/tmds-io/masterdata/service/model/data_value"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/utils/handlers"
)

type DataValueHandler struct {
	dataValues *manager.DataValueManager
}

func NewDataValueHandler(dataValues *manager.DataValueManager) *DataValueHandler {
	return &DataValueHandler{dataValues}
}

func (h *DataValueHandler) GenerateDataValueVariations(ctx context.Context, request *proto.GenerateDataValueVariationsRequest, response *proto.GenerateDataValueVariationsResponse) error {
	//TODO implement me
	panic("implement me")
}

func (h *DataValueHandler) Configure(ctx context.Context, request *proto.ConfigureDataValueRequest, response *proto.ConfigureDataValueResponse) error {
	dto, result := &model.DataValue{}, &proto.DataValue{}

	err := handlers.CopyCallCopy(request.GetDataValue(), dto, func() error {
		return h.dataValues.Configure(ctx, dto)
	}, result)

	response.Id = result.Id
	return err
}

func (h *DataValueHandler) Read(ctx context.Context, request *proto.ReadDataValueRequest, response *proto.ReadDataValueResponse) error {
	dto, result := model.NewDataValue(request.GetId()), &proto.DataValue{}

	err := handlers.CallCopy(dto, func() error {
		return h.dataValues.Read(ctx, dto)
	}, result)

	response.DataValue = result
	return err
}

func (h *DataValueHandler) Update(ctx context.Context, request *proto.UpdateDataValueRequest, response *proto.UpdateDataValueResponse) error {
	dto, result := &model.DataValue{}, &proto.DataValue{}

	err := handlers.CopyCallCopy(request.GetDataValue(), dto, func() error {
		return h.dataValues.Update(ctx, dto)
	}, result)

	response.Id = result.Id
	return err
}

func (h *DataValueHandler) Delete(ctx context.Context, request *proto.DeleteDataValueRequest, response *proto.DeleteDataValueResponse) error {
	return h.dataValues.Delete(ctx, request.GetId())
}

func (h *DataValueHandler) Search(ctx context.Context, request *proto.SearchDataValueRequest, response *proto.SearchDataValueResponse) error {
	var result arangodb.ResultSetInterface[*model.DataValue]
	var err error

	result, err = h.dataValues.Search(ctx, request.GetFilters())

	//Check for error
	if err != nil {
		return err
	}

	response.DataValues = []*proto.DataValue{}
	if err = result.Copy(&response.DataValues); err != nil {
		return err
	}

	response.Pagination = arangodb.NewPaginationResult(result, request.GetFilters())
	return nil
}
