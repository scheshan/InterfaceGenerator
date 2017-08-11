package generate

import (
	"errors"
	"strings"
)

type protoParserImpl struct {
	proto          string
	content        []string
	currentType    *TypeDef
	existTypeNames []string
	currentComment []string
}

func (parser *protoParserImpl) Parse(proto string) ([]TypeDef, error) {
	if len(proto) == 0 {
		return nil, errors.New("Proto content cannot be empty")
	}

	reader := strings.NewReader(proto)

	parser.proto = proto

	return nil, nil
}
