package category

import (
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
)

const Collection = "category"

type Category struct {
	arangodb.Model
	// Unique Key code (used as id in db)
	Key string `json:"key,omitempty"`

	// Name of the category
	Name string `json:"name,omitempty"`

	// CategoryPath of the category
	CategoryPathIds []string `json:"category_path_ids,omitempty"`

	UpdatedAt *db.Timestamp `json:"updated_at,omitempty"`
}

func NewCategory(id string) *Category {
	return &Category{Model: arangodb.Model{Id: id}}
}
