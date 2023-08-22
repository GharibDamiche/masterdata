package handler

import (
	"context"
	proto "github.com/tmds-io/masterdata/contract/measurement_unit"
	"github.com/tmds-io/masterdata/service/manager"
	model "github.com/tmds-io/masterdata/service/model/measurement_unit"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/utils/handlers"
)

type MeasurementUnitHandler struct {
	measurementUnits *manager.MeasurementUnitManager
}

func NewMeasurementUnitHandler(measurementUnits *manager.MeasurementUnitManager) *MeasurementUnitHandler {
	return &MeasurementUnitHandler{measurementUnits}
}

func (h *MeasurementUnitHandler) Create(ctx context.Context, request *proto.CreateMeasurementUnitRequest, response *proto.CreateMeasurementUnitResponse) error {
	dto, result := &model.MeasurementUnit{}, &proto.MeasurementUnit{}

	err := handlers.CopyCallCopy(request.GetMeasurementUnit(), dto, func() error {
		return h.measurementUnits.Create(ctx, dto)
	}, result)

	response.Id = result.Id
	return err
}

func (h *MeasurementUnitHandler) Read(ctx context.Context, request *proto.ReadMeasurementUnitRequest, response *proto.ReadMeasurementUnitResponse) error {
	dto, result := model.NewMeasurementUnit(request.GetId()), &proto.MeasurementUnit{}

	err := handlers.CallCopy(dto, func() error {
		return h.measurementUnits.Read(ctx, dto)
	}, result)

	response.MeasurementUnit = result
	return err
}

func (h *MeasurementUnitHandler) Update(ctx context.Context, request *proto.UpdateMeasurementUnitRequest, response *proto.UpdateMeasurementUnitResponse) error {
	dto, result := &model.MeasurementUnit{}, &proto.MeasurementUnit{}

	err := handlers.CopyCallCopy(request.GetMeasurementUnit(), dto, func() error {
		return h.measurementUnits.Update(ctx, dto)
	}, result)

	response.MeasurementUnit = result
	return err
}

func (h *MeasurementUnitHandler) Delete(ctx context.Context, request *proto.DeleteMeasurementUnitRequest, response *proto.DeleteMeasurementUnitResponse) error {
	return h.measurementUnits.Delete(ctx, request.GetId())
}

func (h *MeasurementUnitHandler) Search(ctx context.Context, request *proto.SearchMeasurementUnitsRequest, response *proto.SearchMeasurementUnitsResponse) error {
	var result arangodb.ResultSetInterface[*model.MeasurementUnit]
	var err error

	result, err = h.measurementUnits.Search(ctx, request.GetFilters())

	//Check for error
	if err != nil {
		return err
	}

	response.MeasurementUnits = []*proto.MeasurementUnit{}
	if err = result.Copy(&response.MeasurementUnits); err != nil {
		return err
	}

	response.Pagination = arangodb.NewPaginationResult(result, request.GetFilters())
	return nil
}
