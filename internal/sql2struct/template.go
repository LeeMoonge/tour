package sql2struct

import (
	"fmt"
	"github.com/go-programming-tour-book/tour/internal/word"
	"os"
	"text/template"
)

// 预定模板
const structTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}} {{ $length := len .Comment}} {{ if gt $length 0 }}//
{{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}
	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	// 用来存储转换后的GO结构体中的所有字段信息
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	// 用来存储最终用于渲染的模板对象信息
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	// 对通过查询COLUMNS表所组装的道德tbColumns进行进一步的分解和转换
	// 例如：数据库类型到Go结构体的转换和对JSON Tag的处理，都在这一层完成了
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeTostructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns: tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}