package segment

import (
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
)

const Collection = "segments"

type Segment struct {
	arangodb.Model
	// Unique Key code (used as id in db)
	Key string `json:"key,omitempty"`

	// Name of the segment
	Name string `json:"name,omitempty"`

	// Segment path of the segment
	SegmentPathIds []string `json:"segment_path_ids,omitempty"`

	// Values of the segment
	Values []string `json:"values,omitempty"`

	UpdatedAt *db.Timestamp `json:"updated_at,omitempty"`
}

func NewSegment(id string) *Segment {
	return &Segment{Model: arangodb.Model{Id: id}}
}
