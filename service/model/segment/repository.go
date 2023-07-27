package segment

import (
	"github.com/arangodb/go-driver"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
)

type SegmentRepositoryInterface interface {
	arangodb.RepositoryInterface[*Segment]
}

type SegmentRepository struct {
	arangodb.Repository[*Segment]
}

// @Service()
func NewSegmentRepository(db driver.Database) (SegmentRepositoryInterface, error) {
	return &SegmentRepository{arangodb.NewRepository[*Segment](db, Collection, nil)}, nil
}
