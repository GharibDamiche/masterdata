package data_definition

import (
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
)

const Collection = "data_definitions"

type DataDefinition struct {
	arangodb.Model
	// Unique Key code (used as id in db)
	Key string `json:"key,omitempty"`
	// Name of the data definition
	Name string `json:"name,omitempty"`
	// Category of the data definition
	CategoryId string `json:"category_id,omitempty"`
	// Type of the data definition by default
	DefaultType string `json:"default_type,omitempty"`
	// Segment path of the data definition by default
	DefaultSegmentsFullPath map[string]string `json:"default_zone_path,omitempty"`
	// Unit of the data definition by default
	DefaultUnitId string `json:"default_unit_id,omitempty"`
	// Should the data definition be visible
	Visibility string `json:"visibility,omitempty"`

	UpdatedAt *db.Timestamp `json:"updated_at,omitempty"`
}

func NewDataDefiniton(id string) *DataDefinition {
	return &DataDefinition{Model: arangodb.Model{Id: id}}
}
