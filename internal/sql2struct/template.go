package sql2struct

import (
	"fmt"
	"os"
	"text/template"

	"github.com/xielizyh/goprj-tour/internal/word"
)

// structTpl 转换为结构体模板
const strcutTpl = `type {{.TableName | ToCamelCase}} struct {
	{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
		{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
	{{end}}}
	
func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	strcutTpl string
}

// StructColumn 存储转换后的结构体中的所有字段信息
type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

// StructTemplateDB 存储最终用于渲染的模板对象信息
type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{strcutTpl: strcutTpl}
}

// AssemblyColumns 转换成结构体信息
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn{
			// 字段的名字
			Name: column.ColumnName,
			// 字段的类型
			Type: DBTypeToStructType[column.DataType],
			// 字段的TAG
			Tag: fmt.Sprintf(" `json:"+"%s"+"`", column.ColumnName),
			// 字段的注释
			Comment: column.ColumnComment,
		})
	}

	return tplColumns
}

// Generate 执行渲染和处理
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.strcutTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}
