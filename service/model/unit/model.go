package unit

import (
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
)

const Collection = "units"

type Unit struct {
	arangodb.Model
	// Unique Key code (used as id in db)
	Key string `json:"key,omitempty"`

	// Name of the unit
	Name string `json:"name,omitempty"`

	// Type of the unit (e.g. length, weight, ...)
	Type string `json:"type,omitempty"`

	// Unit path of the unit
	BaseUnitFactor complex64 `json:"base_unit_factor,omitempty"`

	UpdatedAt *db.Timestamp `json:"updated_at,omitempty"`
}

func NewUnit(id string) *Unit {
	return &Unit{Model: arangodb.Model{Id: id}}
}
