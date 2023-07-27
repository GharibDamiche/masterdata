package category

import (
	"github.com/arangodb/go-driver"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/db/arangodb"
)

type CategoryRepositoryInterface interface {
	arangodb.RepositoryInterface[*Category]
}

type CategoryRepository struct {
	arangodb.Repository[*Category]
}

// @Service()
func NewCategoryRepository(db driver.Database) (CategoryRepositoryInterface, error) {
	return &CategoryRepository{arangodb.NewRepository[*Category](db, Collection, nil)}, nil
}
