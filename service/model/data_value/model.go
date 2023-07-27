package data_value

import (
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
)

const Collection = "data_values"

type DataValue struct {
	arangodb.Model
	// Unique Key code (used as id in db)
	Key string `json:"key,omitempty"`

	// Expression of the data value
	Expression string `json:"expression,omitempty"`

	// All data-values deeply present expression
	DataValuesInExpression []string `json:"data_values_in_expression,omitempty"`

	// UnitId of the data-value
	UnitId string `json:"unit_id,omitempty"`

	// Type of the data-value
	Type string `json:"type,omitempty"`

	// Data value is valid between From and Till dates
	From *db.Timestamp `json:"from,omitempty"`
	Till *db.Timestamp `json:"till,omitempty"`

	// Date of the last update
	UpdatedAt *db.Timestamp `json:"updated_at,omitempty"`

	// GovernanceId of the data-value
	GovernanceId string `json:"governance_id,omitempty"`

	// All Segments path of the data-value
	SegmentsFullPath map[string]string `json:"segments_full_path,omitempty"`
}

func NewDataValue(id string) *DataValue {
	return &DataValue{Model: arangodb.Model{Id: id}}
}
