package conf

var MainTpl = `package {{.Package}}

import (
    "github.com/woaijssss/godbx"
    {{if .HasType}}"github.com/woaijssss/godbx/ttypes"{{end}}
    {{if .HasDecimal}}"github.com/shopspring/decimal"{{end}}
)

var {{.GoTable}}Fields = struct {
   {{range .Columns}}{{.GoName}} string
   {{end}}
}{
    {{range .Columns}}"{{.DbName}}",
    {{end}}
}

var  {{.GoTable}}Meta = &godbx.TableMeta[{{.GoTable}}]{
    Table: "{{.TableName}}",
    Columns: []string {
        {{range .Columns}}"{{.DbName}}",
        {{end}}
    },
    AutoColumn: "{{.AutoColumn}}",
    LookupFieldFunc: func(columnName string,ins *{{.GoTable}},point bool) any {
        {{range .Columns}}if "{{.DbName}}" == columnName {
            if point {
                 return &ins.{{.GoName}}
            }
            return ins.{{.GoName}}
        }
        {{end}}
        return nil
    },
}

var {{.GoTable}}Dao godbx.QuickDao[{{.GoTable}}] = &struct {
	godbx.QuickDao[{{.GoTable}}]
}{
	godbx.NewBaseQuickDao({{.GoTable}}Meta),
}

type {{.GoTable}} struct {
    {{range .Columns}}{{.GoName}} {{.GoType}}
    {{end}}
}
`
