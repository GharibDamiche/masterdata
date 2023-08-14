// Code generated by container-gen. DO NOT EDIT.

package main

import (
	import_0 "github.com/tmds-io/masterdata/service/model/category"
	import_1 "github.com/tmds-io/masterdata/service/model/data_definition"
	import_2 "github.com/tmds-io/masterdata/service/model/data_value"
	import_3 "github.com/tmds-io/masterdata/service/model/segment"
	import_4 "github.com/tmds-io/masterdata/service/model/unit"
	import_5 "github.com/tmds-io/masterdata/service/processor"
	_ "gitlab.com/tmds-io/core-model/hyperion/contract.git/v2/services"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/container"
)

func init() {
	container.Register(import_0.NewCategoryRepository)
	container.Register(import_1.NewDataDefinitionRepository)
	container.Register(import_2.NewDataValueRepository)
	container.Register(import_3.NewSegmentRepository)
	container.Register(import_4.NewUnitRepository)
	container.Register(import_5.NewParser).Tag("processor.parser")
	container.Register(import_5.NewResolver).Tag("processor.resolver")
}
