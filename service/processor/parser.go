package processor

import (
	"context"
	"github.com/Knetic/govaluate"
	model "github.com/tmds-io/masterdata/service/model/data_value"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/errors"
	"regexp"
	"strings"
)

type Parser struct {
	repository model.DataValueRepository // Your database interface
}

// @Service(tag="processor.parser")
func NewParser(repository model.DataValueRepository) *Parser {
	return &Parser{
		repository: repository,
	}
}

func (p *Parser) Parse(ctx context.Context, dataValue *model.DataValue) (string, error) {
	err := p.checkSyntax(dataValue)
	if err != nil {
		return "", err
	}

	// Extract keys and check for circular dependencies
	parsedExpression, err := p.doParse(ctx, dataValue.Key, dataValue.Expression)
	if err != nil {
		return "", err
	}

	return parsedExpression, nil
}

func (p *Parser) checkSyntax(dataValue *model.DataValue) error {
	// Syntax Check
	_, err := govaluate.NewEvaluableExpression(dataValue.Expression)

	return err
}

func (p *Parser) doParse(ctx context.Context, originalKey string, expression string) (string, error) {
	// Extract rawKeys from expression
	regex := regexp.MustCompile(`\$\w+\$`)
	rawKeys := regex.FindAllString(expression, -1)

	for _, rawKey := range rawKeys {
		// Remove $ symbols from rawKeys
		key := rawKey[1 : len(rawKey)-1]
		if key == originalKey {
			return "", errors.New("circular dependency detected: " + key)
		}

		childDataValue, err := p.repository.GetByKey(ctx, key)
		if err != nil || childDataValue == nil {
			return "", errors.New("key not found: " + key)
		}

		err = p.checkSyntax(childDataValue)
		if err != nil {
			return "", err
		}

		// Recursively doParse dependencies
		parsedChildExpression, err := p.doParse(ctx, originalKey, childDataValue.Expression)
		if err != nil {
			return "", err
		}

		expression = strings.Replace(expression, "$"+key+"$", parsedChildExpression, -1)

	}

	return expression, nil
}
