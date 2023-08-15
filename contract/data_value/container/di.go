// Code generated by protoc-gen-kresolver. DO NOT EDIT.

package container

import (
	contract "gitlab.com/tmds-io/core-model/hyperion/contract.git/v2/services/contract/contract"
	mock "gitlab.com/tmds-io/core-model/hyperion/contract.git/v2/services/contract/contract/mock"
	container "gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/container"
)

func init() {
	// Registering DataDefinitionHandler service
	container.Register(contract.NewDataDefinitionHandlerServiceWithDefaultClient).
		Alias(contract.DataDefinitionHandlerRegistryServiceName).
		Mock(mock.NewMockDataDefinitionHandlerService).
		Tag("service.client").
		Lazy(true)
}