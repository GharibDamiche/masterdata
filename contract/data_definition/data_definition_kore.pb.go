// Code generated by protoc-gen-kore. DO NOT EDIT.

package data_definition

import (
	context "context"
	_ "embed"
	constant "github.com/tmds-io/masterdata/contract/constant"
	masker "gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/masker"
	sanitizer "gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/sanitizer"
	validator "gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/validator"
	client "gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/service/client"
	errors "gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/service/errors"
	server "gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/service/server"
)

// Client API for DataDefinitionHandler service

const DataDefinitionHandlerRegistryServiceName = constant.RegistryServiceName + ".data_definition_handler"

type DataDefinitionHandlerService interface {
	Create(ctx context.Context, in *CreateDataDefinitionRequest, opts ...client.CallOption) (*CreateDataDefinitionResponse, error)
	Read(ctx context.Context, in *ReadDataDefinitionRequest, opts ...client.CallOption) (*ReadDataDefinitionResponse, error)
	Update(ctx context.Context, in *UpdateDataDefinitionRequest, opts ...client.CallOption) (*UpdateDataDefinitionResponse, error)
	Delete(ctx context.Context, in *DeleteDataDefinitionRequest, opts ...client.CallOption) (*DeleteDataDefinitionResponse, error)
	Search(ctx context.Context, in *SearchDataDefinitionRequest, opts ...client.CallOption) (*SearchDataDefinitionResponse, error)
}

type dataDefinitionHandlerService struct {
	c    client.Client
	name string
}

func NewDataDefinitionHandlerService(name string, c client.Client) DataDefinitionHandlerService {
	return &dataDefinitionHandlerService{
		c:    c,
		name: name,
	}
}

func NewDataDefinitionHandlerServiceDefault(clt ...client.Client) DataDefinitionHandlerService {
	if 0 == len(clt) {
		clt = []client.Client{client.DefaultClient}
	}
	return NewDataDefinitionHandlerService(constant.RegistryServiceName, clt[0])
}

func NewDataDefinitionHandlerServiceWithDefaultClient() DataDefinitionHandlerService {
	return NewDataDefinitionHandlerService(constant.RegistryServiceName, client.DefaultClient)
}

func (c *dataDefinitionHandlerService) Create(ctx context.Context, in *CreateDataDefinitionRequest, opts ...client.CallOption) (*CreateDataDefinitionResponse, error) {
	req := c.c.NewRequest(c.name, "DataDefinitionHandler.Create", in)
	out := new(CreateDataDefinitionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataDefinitionHandlerService) Read(ctx context.Context, in *ReadDataDefinitionRequest, opts ...client.CallOption) (*ReadDataDefinitionResponse, error) {
	req := c.c.NewRequest(c.name, "DataDefinitionHandler.Read", in)
	out := new(ReadDataDefinitionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataDefinitionHandlerService) Update(ctx context.Context, in *UpdateDataDefinitionRequest, opts ...client.CallOption) (*UpdateDataDefinitionResponse, error) {
	req := c.c.NewRequest(c.name, "DataDefinitionHandler.Update", in)
	out := new(UpdateDataDefinitionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataDefinitionHandlerService) Delete(ctx context.Context, in *DeleteDataDefinitionRequest, opts ...client.CallOption) (*DeleteDataDefinitionResponse, error) {
	req := c.c.NewRequest(c.name, "DataDefinitionHandler.Delete", in)
	out := new(DeleteDataDefinitionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataDefinitionHandlerService) Search(ctx context.Context, in *SearchDataDefinitionRequest, opts ...client.CallOption) (*SearchDataDefinitionResponse, error) {
	req := c.c.NewRequest(c.name, "DataDefinitionHandler.Search", in)
	out := new(SearchDataDefinitionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataDefinitionHandler service

// DataDefinitionHandlerHandler is the server API for DataDefinitionHandler service.
type DataDefinitionHandlerHandler interface {
	Create(context.Context, *CreateDataDefinitionRequest, *CreateDataDefinitionResponse) error
	Read(context.Context, *ReadDataDefinitionRequest, *ReadDataDefinitionResponse) error
	Update(context.Context, *UpdateDataDefinitionRequest, *UpdateDataDefinitionResponse) error
	Delete(context.Context, *DeleteDataDefinitionRequest, *DeleteDataDefinitionResponse) error
	Search(context.Context, *SearchDataDefinitionRequest, *SearchDataDefinitionResponse) error
}

func RegisterDataDefinitionHandlerHandler(s server.Server, hdlr DataDefinitionHandlerHandler, opts ...server.HandlerOption) error {
	type dataDefinitionHandler interface {
		Create(ctx context.Context, in *CreateDataDefinitionRequest, out *CreateDataDefinitionResponse) error
		Read(ctx context.Context, in *ReadDataDefinitionRequest, out *ReadDataDefinitionResponse) error
		Update(ctx context.Context, in *UpdateDataDefinitionRequest, out *UpdateDataDefinitionResponse) error
		Delete(ctx context.Context, in *DeleteDataDefinitionRequest, out *DeleteDataDefinitionResponse) error
		Search(ctx context.Context, in *SearchDataDefinitionRequest, out *SearchDataDefinitionResponse) error
	}
	type DataDefinitionHandler struct {
		dataDefinitionHandler
	}
	h := &dataDefinitionHandlerHandler{hdlr}
	return s.Handle(s.NewHandler(&DataDefinitionHandler{h}, opts...))
}

type dataDefinitionHandlerHandler struct {
	DataDefinitionHandlerHandler
}

//go:embed validation/data_definition_handler.create.yaml
var ValidationTplOfCreateDataDefinitionRequest string

//go:embed mask/data_definition_handler.create.yaml
var MaskOfCreateDataDefinitionRequest string

func (in *CreateDataDefinitionRequest) Validate(ctx context.Context) error {
	if err := masker.ApplyFromYaml(in, MaskOfCreateDataDefinitionRequest, true); err != nil {
		return err
	}
	return errors.WrapValidation(validator.ValidateSchema(ctx, in, ValidationTplOfCreateDataDefinitionRequest))
}

//go:embed validation/data_definition_handler.read.yaml
var ValidationTplOfReadDataDefinitionRequest string

//go:embed mask/data_definition_handler.read.yaml
var MaskOfReadDataDefinitionRequest string

func (in *ReadDataDefinitionRequest) Validate(ctx context.Context) error {
	if err := masker.ApplyFromYaml(in, MaskOfReadDataDefinitionRequest, true); err != nil {
		return err
	}
	return errors.WrapValidation(validator.ValidateSchema(ctx, in, ValidationTplOfReadDataDefinitionRequest))
}

//go:embed validation/data_definition_handler.update.yaml
var ValidationTplOfUpdateDataDefinitionRequest string

//go:embed mask/data_definition_handler.update.yaml
var MaskOfUpdateDataDefinitionRequest string

func (in *UpdateDataDefinitionRequest) Validate(ctx context.Context) error {
	if err := masker.ApplyFromYaml(in, MaskOfUpdateDataDefinitionRequest, true); err != nil {
		return err
	}
	return errors.WrapValidation(validator.ValidateSchema(ctx, in, ValidationTplOfUpdateDataDefinitionRequest))
}

//go:embed validation/data_definition_handler.delete.yaml
var ValidationTplOfDeleteDataDefinitionRequest string

//go:embed mask/data_definition_handler.delete.yaml
var MaskOfDeleteDataDefinitionRequest string

func (in *DeleteDataDefinitionRequest) Validate(ctx context.Context) error {
	if err := masker.ApplyFromYaml(in, MaskOfDeleteDataDefinitionRequest, true); err != nil {
		return err
	}
	return errors.WrapValidation(validator.ValidateSchema(ctx, in, ValidationTplOfDeleteDataDefinitionRequest))
}

//go:embed validation/data_definition_handler.search.yaml
var ValidationTplOfSearchDataDefinitionRequest string

//go:embed mask/data_definition_handler.search.yaml
var MaskOfSearchDataDefinitionRequest string

func (in *SearchDataDefinitionRequest) Validate(ctx context.Context) error {
	if err := masker.ApplyFromYaml(in, MaskOfSearchDataDefinitionRequest, true); err != nil {
		return err
	}
	return errors.WrapValidation(validator.ValidateSchema(ctx, in, ValidationTplOfSearchDataDefinitionRequest))
}

func (h *dataDefinitionHandlerHandler) Create(ctx context.Context, in *CreateDataDefinitionRequest, out *CreateDataDefinitionResponse) error {
	sanitizer.Sanitize(in)
	if err := in.Validate(ctx); err != nil {
		return err
	}
	return h.DataDefinitionHandlerHandler.Create(ctx, in, out)
}

func (h *dataDefinitionHandlerHandler) Read(ctx context.Context, in *ReadDataDefinitionRequest, out *ReadDataDefinitionResponse) error {
	sanitizer.Sanitize(in)
	if err := in.Validate(ctx); err != nil {
		return err
	}
	return h.DataDefinitionHandlerHandler.Read(ctx, in, out)
}

func (h *dataDefinitionHandlerHandler) Update(ctx context.Context, in *UpdateDataDefinitionRequest, out *UpdateDataDefinitionResponse) error {
	sanitizer.Sanitize(in)
	if err := in.Validate(ctx); err != nil {
		return err
	}
	return h.DataDefinitionHandlerHandler.Update(ctx, in, out)
}

func (h *dataDefinitionHandlerHandler) Delete(ctx context.Context, in *DeleteDataDefinitionRequest, out *DeleteDataDefinitionResponse) error {
	sanitizer.Sanitize(in)
	if err := in.Validate(ctx); err != nil {
		return err
	}
	return h.DataDefinitionHandlerHandler.Delete(ctx, in, out)
}

func (h *dataDefinitionHandlerHandler) Search(ctx context.Context, in *SearchDataDefinitionRequest, out *SearchDataDefinitionResponse) error {
	sanitizer.Sanitize(in)
	if err := in.Validate(ctx); err != nil {
		return err
	}
	return h.DataDefinitionHandlerHandler.Search(ctx, in, out)
}
