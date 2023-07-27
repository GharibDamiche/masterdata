package processor

import (
	"context"
	"fmt"
	"github.com/antonmedv/expr/ast"
	expr_parser "github.com/antonmedv/expr/parser"
	"github.com/tmds-io/masterdata/service/model/data_value"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/errors"
)

type Parser struct {
	repository data_value.DataValueRepository // Your database interface
}

func (p *Parser) Parse(ctx context.Context, expr string) (*ast.Node, error) {
	tree, err := expr_parser.Parse(expr)
	if err != nil {
		return nil, errors.New("invalid expression syntax")
	}

	// Check if all identifiers exist in the database and if there are circular dependencies.
	keys := ast.GetIdentifiers(tree.Node)
	visited := map[string]bool{}
	for _, key := range keys {
		if err := p.checkKey(ctx, key, visited); err != nil {
			return nil, err
		}
	}

	return &tree, nil
}

func (p *Parser) checkKey(ctx context.Context, key string, visited map[string]bool) error {
	if visited[key] {
		return fmt.Errorf("circular dependency detected at key %s", key)
	}

	dataValue := data_value.NewDataValue(key)
	err := p.repository.Read(ctx, dataValue)
	if err != nil {
		return fmt.Errorf("unknown key %s", key)
	}

	visited[key] = true

	tree, err := expr_parser.Parse(dataValue.Expression)
	if err != nil {
		return errors.New("invalid expression syntax")
	}

	identifiers := ast.GetIdentifiers(tree.Node)
	for _, id := range identifiers {
		if err := p.checkKey(id, visited); err != nil {
			return err
		}
	}

	return nil
}
