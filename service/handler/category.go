package handler

import (
	"context"
	proto "github.com/tmds-io/masterdata/contract/category"
	"github.com/tmds-io/masterdata/service/manager"
	model "github.com/tmds-io/masterdata/service/model/category"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/utils/handlers"
)

type CategoryHandler struct {
	categories *manager.CategoryManager
}

func NewCategoryHandler(categories *manager.CategoryManager) *CategoryHandler {
	return &CategoryHandler{categories}
}

func (h *CategoryHandler) AddChildCategoryValue(ctx context.Context, request *proto.AddChildCategoryValueRequest, response *proto.AddChildCategoryValueResponse) error {
	//TODO implement me
	panic("implement me")
}

func (h *CategoryHandler) Create(ctx context.Context, request *proto.CreateCategoryRequest, response *proto.CreateCategoryResponse) error {
	dto, result := &model.Category{}, &proto.Category{}

	err := handlers.CopyCallCopy(request.GetCategory(), dto, func() error {
		return h.categories.Create(ctx, dto)
	}, result)

	response.Id = result.Id
	return err
}

func (h *CategoryHandler) Read(ctx context.Context, request *proto.ReadCategoryRequest, response *proto.ReadCategoryResponse) error {
	dto, result := model.NewCategory(request.GetId()), &proto.Category{}

	err := handlers.CallCopy(dto, func() error {
		return h.categories.Read(ctx, dto)
	}, result)

	response.Category = result
	return err
}

func (h *CategoryHandler) Update(ctx context.Context, request *proto.UpdateCategoryRequest, response *proto.UpdateCategoryResponse) error {
	dto, result := &model.Category{}, &proto.Category{}

	err := handlers.CopyCallCopy(request.GetCategory(), dto, func() error {
		return h.categories.Update(ctx, dto)
	}, result)

	response.Id = result.Id
	return err
}

func (h *CategoryHandler) Delete(ctx context.Context, request *proto.DeleteCategoryRequest, response *proto.DeleteCategoryResponse) error {
	return h.categories.Delete(ctx, request.GetId())
}

func (h *CategoryHandler) Search(ctx context.Context, request *proto.SearchCategoriesRequest, response *proto.SearchCategoriesResponse) error {
	var result arangodb.ResultSetInterface[*model.Category]
	var err error

	result, err = h.categories.Search(ctx, request.GetFilters())

	//Check for error
	if err != nil {
		return err
	}

	response.Categories = []*proto.Category{}
	if err = result.Copy(&response.Categories); err != nil {
		return err
	}

	response.Pagination = arangodb.NewPaginationResult(result, request.GetFilters())
	return nil
}
