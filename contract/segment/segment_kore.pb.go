// Code generated by protoc-gen-kore. DO NOT EDIT.

package segment

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

// Client API for SegmentHandler service

const SegmentHandlerRegistryServiceName = constant.RegistryServiceName + ".segment_handler"

type SegmentHandlerService interface {
	Create(ctx context.Context, in *CreateSegmentRequest, opts ...client.CallOption) (*CreateSegmentResponse, error)
	Read(ctx context.Context, in *ReadSegmentRequest, opts ...client.CallOption) (*ReadSegmentResponse, error)
	Update(ctx context.Context, in *UpdateSegmentRequest, opts ...client.CallOption) (*UpdateSegmentResponse, error)
	Delete(ctx context.Context, in *DeleteSegmentRequest, opts ...client.CallOption) (*DeleteSegmentResponse, error)
	Search(ctx context.Context, in *SearchSegmentsRequest, opts ...client.CallOption) (*SearchSegmentsResponse, error)
	AddChildSegmentValue(ctx context.Context, in *AddChildSegmentValueRequest, opts ...client.CallOption) (*AddChildSegmentValueResponse, error)
}

type segmentHandlerService struct {
	c    client.Client
	name string
}

func NewSegmentHandlerService(name string, c client.Client) SegmentHandlerService {
	return &segmentHandlerService{
		c:    c,
		name: name,
	}
}

func NewSegmentHandlerServiceDefault(clt ...client.Client) SegmentHandlerService {
	if 0 == len(clt) {
		clt = []client.Client{client.DefaultClient}
	}
	return NewSegmentHandlerService(constant.RegistryServiceName, clt[0])
}

func NewSegmentHandlerServiceWithDefaultClient() SegmentHandlerService {
	return NewSegmentHandlerService(constant.RegistryServiceName, client.DefaultClient)
}

func (c *segmentHandlerService) Create(ctx context.Context, in *CreateSegmentRequest, opts ...client.CallOption) (*CreateSegmentResponse, error) {
	req := c.c.NewRequest(c.name, "SegmentHandler.Create", in)
	out := new(CreateSegmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentHandlerService) Read(ctx context.Context, in *ReadSegmentRequest, opts ...client.CallOption) (*ReadSegmentResponse, error) {
	req := c.c.NewRequest(c.name, "SegmentHandler.Read", in)
	out := new(ReadSegmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentHandlerService) Update(ctx context.Context, in *UpdateSegmentRequest, opts ...client.CallOption) (*UpdateSegmentResponse, error) {
	req := c.c.NewRequest(c.name, "SegmentHandler.Update", in)
	out := new(UpdateSegmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentHandlerService) Delete(ctx context.Context, in *DeleteSegmentRequest, opts ...client.CallOption) (*DeleteSegmentResponse, error) {
	req := c.c.NewRequest(c.name, "SegmentHandler.Delete", in)
	out := new(DeleteSegmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentHandlerService) Search(ctx context.Context, in *SearchSegmentsRequest, opts ...client.CallOption) (*SearchSegmentsResponse, error) {
	req := c.c.NewRequest(c.name, "SegmentHandler.Search", in)
	out := new(SearchSegmentsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentHandlerService) AddChildSegmentValue(ctx context.Context, in *AddChildSegmentValueRequest, opts ...client.CallOption) (*AddChildSegmentValueResponse, error) {
	req := c.c.NewRequest(c.name, "SegmentHandler.AddChildSegmentValue", in)
	out := new(AddChildSegmentValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SegmentHandler service

// SegmentHandlerHandler is the server API for SegmentHandler service.
type SegmentHandlerHandler interface {
	Create(context.Context, *CreateSegmentRequest, *CreateSegmentResponse) error
	Read(context.Context, *ReadSegmentRequest, *ReadSegmentResponse) error
	Update(context.Context, *UpdateSegmentRequest, *UpdateSegmentResponse) error
	Delete(context.Context, *DeleteSegmentRequest, *DeleteSegmentResponse) error
	Search(context.Context, *SearchSegmentsRequest, *SearchSegmentsResponse) error
	AddChildSegmentValue(context.Context, *AddChildSegmentValueRequest, *AddChildSegmentValueResponse) error
}

func RegisterSegmentHandlerHandler(s server.Server, hdlr SegmentHandlerHandler, opts ...server.HandlerOption) error {
	type segmentHandler interface {
		Create(ctx context.Context, in *CreateSegmentRequest, out *CreateSegmentResponse) error
		Read(ctx context.Context, in *ReadSegmentRequest, out *ReadSegmentResponse) error
		Update(ctx context.Context, in *UpdateSegmentRequest, out *UpdateSegmentResponse) error
		Delete(ctx context.Context, in *DeleteSegmentRequest, out *DeleteSegmentResponse) error
		Search(ctx context.Context, in *SearchSegmentsRequest, out *SearchSegmentsResponse) error
		AddChildSegmentValue(ctx context.Context, in *AddChildSegmentValueRequest, out *AddChildSegmentValueResponse) error
	}
	type SegmentHandler struct {
		segmentHandler
	}
	h := &segmentHandlerHandler{hdlr}
	return s.Handle(s.NewHandler(&SegmentHandler{h}, opts...))
}

type segmentHandlerHandler struct {
	SegmentHandlerHandler
}

//go:embed validation/segment_handler.create.yaml
var ValidationTplOfCreateSegmentRequest string

//go:embed mask/segment_handler.create.yaml
var MaskOfCreateSegmentRequest string

func (in *CreateSegmentRequest) Validate(ctx context.Context) error {
	if err := masker.ApplyFromYaml(in, MaskOfCreateSegmentRequest, true); err != nil {
		return err
	}
	return errors.WrapValidation(validator.ValidateSchema(ctx, in, ValidationTplOfCreateSegmentRequest))
}

//go:embed validation/segment_handler.read.yaml
var ValidationTplOfReadSegmentRequest string

//go:embed mask/segment_handler.read.yaml
var MaskOfReadSegmentRequest string

func (in *ReadSegmentRequest) Validate(ctx context.Context) error {
	if err := masker.ApplyFromYaml(in, MaskOfReadSegmentRequest, true); err != nil {
		return err
	}
	return errors.WrapValidation(validator.ValidateSchema(ctx, in, ValidationTplOfReadSegmentRequest))
}

//go:embed validation/segment_handler.update.yaml
var ValidationTplOfUpdateSegmentRequest string

//go:embed mask/segment_handler.update.yaml
var MaskOfUpdateSegmentRequest string

func (in *UpdateSegmentRequest) Validate(ctx context.Context) error {
	if err := masker.ApplyFromYaml(in, MaskOfUpdateSegmentRequest, true); err != nil {
		return err
	}
	return errors.WrapValidation(validator.ValidateSchema(ctx, in, ValidationTplOfUpdateSegmentRequest))
}

//go:embed validation/segment_handler.delete.yaml
var ValidationTplOfDeleteSegmentRequest string

//go:embed mask/segment_handler.delete.yaml
var MaskOfDeleteSegmentRequest string

func (in *DeleteSegmentRequest) Validate(ctx context.Context) error {
	if err := masker.ApplyFromYaml(in, MaskOfDeleteSegmentRequest, true); err != nil {
		return err
	}
	return errors.WrapValidation(validator.ValidateSchema(ctx, in, ValidationTplOfDeleteSegmentRequest))
}

//go:embed validation/segment_handler.search.yaml
var ValidationTplOfSearchSegmentsRequest string

//go:embed mask/segment_handler.search.yaml
var MaskOfSearchSegmentsRequest string

func (in *SearchSegmentsRequest) Validate(ctx context.Context) error {
	if err := masker.ApplyFromYaml(in, MaskOfSearchSegmentsRequest, true); err != nil {
		return err
	}
	return errors.WrapValidation(validator.ValidateSchema(ctx, in, ValidationTplOfSearchSegmentsRequest))
}

//go:embed validation/segment_handler.add_child_segment_value.yaml
var ValidationTplOfAddChildSegmentValueRequest string

//go:embed mask/segment_handler.add_child_segment_value.yaml
var MaskOfAddChildSegmentValueRequest string

func (in *AddChildSegmentValueRequest) Validate(ctx context.Context) error {
	if err := masker.ApplyFromYaml(in, MaskOfAddChildSegmentValueRequest, true); err != nil {
		return err
	}
	return errors.WrapValidation(validator.ValidateSchema(ctx, in, ValidationTplOfAddChildSegmentValueRequest))
}

func (h *segmentHandlerHandler) Create(ctx context.Context, in *CreateSegmentRequest, out *CreateSegmentResponse) error {
	sanitizer.Sanitize(in)
	if err := in.Validate(ctx); err != nil {
		return err
	}
	return h.SegmentHandlerHandler.Create(ctx, in, out)
}

func (h *segmentHandlerHandler) Read(ctx context.Context, in *ReadSegmentRequest, out *ReadSegmentResponse) error {
	sanitizer.Sanitize(in)
	if err := in.Validate(ctx); err != nil {
		return err
	}
	return h.SegmentHandlerHandler.Read(ctx, in, out)
}

func (h *segmentHandlerHandler) Update(ctx context.Context, in *UpdateSegmentRequest, out *UpdateSegmentResponse) error {
	sanitizer.Sanitize(in)
	if err := in.Validate(ctx); err != nil {
		return err
	}
	return h.SegmentHandlerHandler.Update(ctx, in, out)
}

func (h *segmentHandlerHandler) Delete(ctx context.Context, in *DeleteSegmentRequest, out *DeleteSegmentResponse) error {
	sanitizer.Sanitize(in)
	if err := in.Validate(ctx); err != nil {
		return err
	}
	return h.SegmentHandlerHandler.Delete(ctx, in, out)
}

func (h *segmentHandlerHandler) Search(ctx context.Context, in *SearchSegmentsRequest, out *SearchSegmentsResponse) error {
	sanitizer.Sanitize(in)
	if err := in.Validate(ctx); err != nil {
		return err
	}
	return h.SegmentHandlerHandler.Search(ctx, in, out)
}

func (h *segmentHandlerHandler) AddChildSegmentValue(ctx context.Context, in *AddChildSegmentValueRequest, out *AddChildSegmentValueResponse) error {
	sanitizer.Sanitize(in)
	if err := in.Validate(ctx); err != nil {
		return err
	}
	return h.SegmentHandlerHandler.AddChildSegmentValue(ctx, in, out)
}
