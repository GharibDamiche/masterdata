// Code generated by protoc-gen-kresolver. DO NOT EDIT.

package container

import (
	data_definition "github.com/tmds-io/masterdata/contract/data_definition"
	mock "github.com/tmds-io/masterdata/contract/data_definition/mock"
	container "gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/container"
)

func init() {
	// Registering DataDefinitionHandler service
	container.Register(data_definition.NewDataDefinitionHandlerServiceWithDefaultClient).
		Alias(data_definition.DataDefinitionHandlerRegistryServiceName).
		Mock(mock.NewMockDataDefinitionHandlerService).
		Tag("service.client").
		Lazy(true)
}
