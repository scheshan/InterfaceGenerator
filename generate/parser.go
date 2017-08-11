package generate

//ProtoParser 定义协议解析的接口
type ProtoParser interface {
	Parse(proto string) ([]TypeDef, error)
}

//NewProtoParser 返回一个默认的协议解析接口的实例
func NewProtoParser() ProtoParser {
	parser := &protoParserImpl{}
	parser.existTypeNames = make(map[string]string)

	return parser
}
