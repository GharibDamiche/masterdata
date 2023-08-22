package handler

import (
	"context"
	proto "github.com/tmds-io/masterdata/contract/segment"
	"github.com/tmds-io/masterdata/service/manager"
	model "github.com/tmds-io/masterdata/service/model/segment"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/utils/handlers"
)

type SegmentHandler struct {
	segments *manager.SegmentManager
}

func NewSegmentHandler(segments *manager.SegmentManager) *SegmentHandler {
	return &SegmentHandler{segments}
}

func (h *SegmentHandler) AddChildSegmentValue(ctx context.Context, request *proto.AddChildSegmentValueRequest, response *proto.AddChildSegmentValueResponse) error {
	//TODO implement me
	panic("implement me")
}

func (h *SegmentHandler) Create(ctx context.Context, request *proto.CreateSegmentRequest, response *proto.CreateSegmentResponse) error {
	dto, result := &model.Segment{}, &proto.Segment{}

	err := handlers.CopyCallCopy(request.GetSegment(), dto, func() error {
		return h.segments.Create(ctx, dto)
	}, result)

	response.Id = result.Id
	return err
}

func (h *SegmentHandler) Read(ctx context.Context, request *proto.ReadSegmentRequest, response *proto.ReadSegmentResponse) error {
	dto, result := model.NewSegment(request.GetId()), &proto.Segment{}

	err := handlers.CallCopy(dto, func() error {
		return h.segments.Read(ctx, dto)
	}, result)

	response.Segment = result
	return err
}

func (h *SegmentHandler) Update(ctx context.Context, request *proto.UpdateSegmentRequest, response *proto.UpdateSegmentResponse) error {
	dto, result := &model.Segment{}, &proto.Segment{}

	err := handlers.CopyCallCopy(request.GetSegment(), dto, func() error {
		return h.segments.Update(ctx, dto)
	}, result)

	response.Id = result.Id
	return err
}

func (h *SegmentHandler) Delete(ctx context.Context, request *proto.DeleteSegmentRequest, response *proto.DeleteSegmentResponse) error {
	return h.segments.Delete(ctx, request.GetId())
}

func (h *SegmentHandler) Search(ctx context.Context, request *proto.SearchSegmentsRequest, response *proto.SearchSegmentsResponse) error {
	var result arangodb.ResultSetInterface[*model.Segment]
	var err error

	result, err = h.segments.Search(ctx, request.GetFilters())

	//Check for error
	if err != nil {
		return err
	}

	response.Segments = []*proto.Segment{}
	if err = result.Copy(&response.Segments); err != nil {
		return err
	}

	response.Pagination = arangodb.NewPaginationResult(result, request.GetFilters())
	return nil
}
