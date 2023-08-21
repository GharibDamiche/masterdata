// Code generated by protoc-gen-kresolver. DO NOT EDIT.

package container

import (
	data_value "github.com/tmds-io/masterdata/contract/data_value"
	mock "github.com/tmds-io/masterdata/contract/data_value/mock"
	container "gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/container"
)

func init() {
	// Registering DataValueHandler service
	container.Register(data_value.NewDataValueHandlerServiceWithDefaultClient).
		Alias(data_value.DataValueHandlerRegistryServiceName).
		Mock(mock.NewMockDataValueHandlerService).
		Tag("service.client").
		Lazy(true)
}