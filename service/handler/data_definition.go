package handler

import (
	"context"
	proto "github.com/tmds-io/masterdata/contract/data_definition"
	"github.com/tmds-io/masterdata/service/manager"
	model "github.com/tmds-io/masterdata/service/model/data_definition"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/utils/handlers"
)

type DataDefinitionHandler struct {
	dataDefinitions *manager.DataDefinitionManager
}

func NewDataDefinitionHandler(dataDefinitions *manager.DataDefinitionManager) *DataDefinitionHandler {
	return &DataDefinitionHandler{dataDefinitions}
}

func (h *DataDefinitionHandler) Create(ctx context.Context, request *proto.CreateDataDefinitionRequest, response *proto.CreateDataDefinitionResponse) error {
	dto, result := &model.DataDefinition{}, &proto.DataDefinition{}

	err := handlers.CopyCallCopy(request.GetDataDefinition(), dto, func() error {
		return h.dataDefinitions.Create(ctx, dto)
	}, result)

	response.Id = result.Id
	return err
}

func (h *DataDefinitionHandler) Read(ctx context.Context, request *proto.ReadDataDefinitionRequest, response *proto.ReadDataDefinitionResponse) error {
	dto, result := model.NewDataDefiniton(request.GetId()), &proto.DataDefinition{}

	err := handlers.CallCopy(dto, func() error {
		return h.dataDefinitions.Read(ctx, dto)
	}, result)

	response.DataDefinition = result
	return err
}

func (h *DataDefinitionHandler) Update(ctx context.Context, request *proto.UpdateDataDefinitionRequest, response *proto.UpdateDataDefinitionResponse) error {
	dto, result := &model.DataDefinition{}, &proto.DataDefinition{}

	err := handlers.CopyCallCopy(request.GetDataDefinition(), dto, func() error {
		return h.dataDefinitions.Update(ctx, dto)
	}, result)

	response.Id = result.Id
	return err
}

func (h *DataDefinitionHandler) Delete(ctx context.Context, request *proto.DeleteDataDefinitionRequest, response *proto.DeleteDataDefinitionResponse) error {
	return h.dataDefinitions.Delete(ctx, request.GetId())
}

func (h *DataDefinitionHandler) Search(ctx context.Context, request *proto.SearchDataDefinitionRequest, response *proto.SearchDataDefinitionResponse) error {
	var result arangodb.ResultSetInterface[*model.DataDefinition]
	var err error

	result, err = h.dataDefinitions.Search(ctx, request.GetFilters())

	//Check for error
	if err != nil {
		return err
	}

	response.DataDefinitions = []*proto.DataDefinition{}
	if err = result.Copy(&response.DataDefinitions); err != nil {
		return err
	}

	response.Pagination = arangodb.NewPaginationResult(result, request.GetFilters())
	return nil
}
