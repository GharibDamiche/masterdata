package processor

import (
	"context"
	"github.com/Knetic/govaluate"
	model "github.com/tmds-io/masterdata/service/model/data_value"
)

type Resolver struct {
	parser *Parser
}

// @Service(tag="processor.resolver")
func NewResolver(parser *Parser) *Resolver {
	return &Resolver{
		parser: parser,
	}
}

func (r *Resolver) Resolve(ctx context.Context, dataValue *model.DataValue) (float64, error) {
	expression, err := r.parser.Parse(ctx, dataValue)
	if err != nil {
		return 0, err
	}

	return r.resolveExpression(expression)
}

func (r *Resolver) resolveExpression(expression string) (float64, error) {

	// Evaluate the expression
	evaluableExpression, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, err // Syntax error
	}

	result, err := evaluableExpression.Evaluate(nil)
	if err != nil {
		return 0, err // Evaluation error
	}

	return result.(float64), nil
}
