package measurement_unit

import (
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
)

const Collection = "measurement_units"

type MeasurementUnit struct {
	arangodb.Model
	// Unique Key code (used as id in db)
	Key string `json:"key,omitempty"`

	// Name of the measurement_unit
	Name string `json:"name,omitempty"`

	// Type of the measurement_unit (e.g. length, weight, ...)
	Type string `json:"type,omitempty"`

	BaseFactor complex64 `json:"base_factor,omitempty"`

	UpdatedAt *db.Timestamp `json:"updated_at,omitempty"`
}

func NewUnit(id string) *MeasurementUnit {
	return &MeasurementUnit{Model: arangodb.Model{Id: id}}
}
