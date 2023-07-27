package processor

import (
	"context"
	"fmt"
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/ast"
	"github.com/antonmedv/expr/vm"
	"github.com/tmds-io/masterdata/service/model/data_value"
)

type Resolver struct {
	repository data_value.DataValueRepositoryInterface
	parser     Parser
}

// @Service(tag="processor.resolver")
func NewResolver(repository data_value.DataValueRepositoryInterface, parser Parser) *Resolver {
	return &Resolver{repository: repository, parser: parser}
}

func (r *Resolver) Resolve(ctx context.Context, dataValue data_value.DataValue) (interface{}, error) {
	env := make(map[string]interface{})
	tree, _ := r.parser.Parse(dataValue.Expression)

	ast.Walk(tree.Node, func(node ast.Node) bool {
		if n, ok := node.(*ast.IdentifierNode); ok {
			value, err := r.Resolve(n.Name)
			if err != nil {
				return false
			}
			env[n.Name] = value
		}
		return true
	})

	program, err := vm.Compile(tree.Node, expr.Env(env))
	if err != nil {
		return nil, fmt.Errorf("failed to compile expression: %v", err)
	}

	result, err := expr.Run(program, env)
	if err != nil {
		return nil, err
	}

	return result, nil
}
