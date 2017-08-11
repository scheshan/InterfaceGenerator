package generate

//AttributeDef 特性实体
type AttributeDef struct {
	Name  string
	Value string
}

//FieldDef 字段实体
type FieldDef struct {
	Name       string
	FieldType  string
	Required   bool
	Comment    string
	Attributes []AttributeDef
}

//TypeDef 类型实体
type TypeDef struct {
	Name       string
	Comment    string
	Fields     []FieldDef
	Attributes []AttributeDef
}
