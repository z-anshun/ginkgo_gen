package model

type FieldType int

const (
	Normal FieldType = iota
	Pointer
	Array
)

type Model struct {
	Name     string      `json:"name"` // 结构体名
	Fs       []Field     `json:"fs"`   // 字段
	Pkg      string      `json:"pkg"`
	Describe string      `json:"describe"`
	Contexts []ContextIt `json:"contexts"`
	Service  string      `json:"service"` // Tracing
}

type Content struct {
	Describe string      `json:"describe"`
	Contexts []ContextIt `json:"contexts"`
}

type ContextIt struct {
	Context string `json:"context"`
	It      string `json:"it"`
}

type Field struct {
	Name     []string `json:"name"` // 字段名
	Type     FieldType
	TypeName string `json:"type"` // 类型
	Tag      string `json:"tag"`  // tag
}
