package generate

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

type protoParserImpl struct {
	proto          string
	content        []string
	currentType    *TypeDef
	existTypeNames map[string]string
	currentComment []string
	result         []TypeDef
}

func (parser *protoParserImpl) Parse(proto string) ([]TypeDef, error) {
	if len(proto) == 0 {
		return nil, errors.New("Proto content cannot be empty")
	}

	parser.proto = proto

	lineNumber := 0
	var currentLine string

	reader := strings.NewReader(proto)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lineNumber++
		currentLine = trimString(scanner.Text())

		if len(currentLine) == 0 {
			continue
		}

		if strings.Index(currentLine, "#") == 0 {
			parser.processComment(lineNumber, currentLine)
		} else if strings.Index(currentLine, "type") == 0 {
			parser.processTypeStart(lineNumber, currentLine)
		} else if strings.Index(currentLine, "}") == 0 {
			parser.processTypeEnd(lineNumber, currentLine)
		} else {
			parser.processField(lineNumber, currentLine)
		}
	}

	return parser.result, nil
}

func (parser *protoParserImpl) processComment(lineNumber int, line string) {
	comment := strings.TrimPrefix(line, "#")
	parser.currentComment = append(parser.currentComment, comment)
}

func (parser *protoParserImpl) processTypeStart(lineNumber int, line string) {
	if parser.currentType != nil {
		throw(lineNumber, "The type definition [%s] is not finished", parser.currentType.Name)
	}
	if strings.LastIndex(line, "{") != len(line)-1 {
		throw(lineNumber, "Type definition error, missing {")
	}

	arr := strings.Fields(line)

	line = strings.TrimSuffix(line, "{")
	if len(arr) != 2 {
		throw(lineNumber, "Missing keywords in type definition. Use [type xxx {] to define a type")
	}

	typeName := arr[1]
	if _, ok := parser.existTypeNames[typeName]; ok {
		throw(lineNumber, "Already define the type [%s]", typeName)
	}

	parser.existTypeNames[typeName] = typeName

	parser.currentType = &TypeDef{}
	parser.currentType.Name = typeName
	parser.currentType.Comment = parser.getCommentAndClear()
}

func (parser *protoParserImpl) processTypeEnd(lineNumber int, line string) {
	if parser.currentType == nil {
		throw(lineNumber, "Missing type definition")
	}

	parser.result = append(parser.result, *parser.currentType)
	parser.currentType = nil
	parser.currentComment = nil
}

func (parser *protoParserImpl) processField(lineNumber int, line string) {
	if parser.currentType == nil {
		throw(lineNumber, "Missing type definition")
	}

	arr := strings.Fields(line)

	if len(arr) < 2 {
		throw(lineNumber, "Missing keywords in field definition, use [int xxx] to define a field")
	}

	field := FieldDef{}
	fieldType := arr[0]
	if strings.LastIndex(fieldType, "?") == len(fieldType)-1 {
		field.FieldType = strings.TrimSuffix(fieldType, "?")
	} else {
		field.FieldType = fieldType
		field.Required = true
	}
	field.Name = arr[1]
	field.Comment = parser.getCommentAndClear()

	parser.currentType.Fields = append(parser.currentType.Fields, field)
}

func (parser *protoParserImpl) getCommentAndClear() []string {
	comment := parser.currentComment
	parser.currentComment = nil
	return comment
}

func trimString(text string) string {
	text = strings.Trim(text, " ")
	text = strings.Replace(text, "\r", "", -1)
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r\n", "", -1)

	return text
}

func throw(lineNumber int, message string, args ...interface{}) {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args)
	}

	err := fmt.Sprintf(`Parse error at line: %d.
%s`, lineNumber, message)

	panic(err)
}
