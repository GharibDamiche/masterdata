// Code generated by protoc-gen-kresolver. DO NOT EDIT.

package container

import (
	data_definition "gitlab.com/tmds-io/core-model/hyperion/contract.git/v2/services/data_definition/data_definition"
	mock "gitlab.com/tmds-io/core-model/hyperion/contract.git/v2/services/data_definition/data_definition/mock"
	container "gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/container"
)

func init() {
	// Registering CategoryHandler service
	container.Register(data_definition.NewCategoryHandlerServiceWithDefaultClient).
		Alias(data_definition.CategoryHandlerRegistryServiceName).
		Mock(mock.NewMockCategoryHandlerService).
		Tag("service.client").
		Lazy(true)

	// Registering DataDefinitionHandler service
	container.Register(data_definition.NewDataDefinitionHandlerServiceWithDefaultClient).
		Alias(data_definition.DataDefinitionHandlerRegistryServiceName).
		Mock(mock.NewMockDataDefinitionHandlerService).
		Tag("service.client").
		Lazy(true)
}