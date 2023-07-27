package shared

import "gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db"

type Meta struct {
	// Creation date
	CreatedAt *db.Timestamp `json:"created_at,omitempty"`
	// Last update date
	UpdatedAt *db.Timestamp `json:"updated_at,omitempty"`
	// Deleted at date
	DeletedAt *db.Timestamp `json:"deleted_at,omitempty"`
}

type ActionMadeBy struct {
	UserId    string `json:"user_id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}
