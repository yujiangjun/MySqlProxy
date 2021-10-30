package jdbc

import (
	"database/sql"
)

type Field struct {
	FName string `db:"fName"`
	FDesc string `db:"fDesc"`
	DataType string `db:"dataType"`
	IsNull string `db:"isNull"`
	SLength string `db:"sLength"`
}
/**

 */
type Tables struct {
	TableCateLog string `db:"TABLE_CATALOG"`
	TableSchema string `db:"TABLE_SCHEMA"`
	TableName string `db:"TABLE_NAME"`
	TableType string `db:"TABLE_TYPE"`
	Engine string `db:"ENGINE"`
	Version int8 `db:"VERSION"`
	RowFormat string `db:"ROW_FORMAT"`
	TableRows int16 `db:"TABLE_ROWS"`
	DataLength int64 `db:"DATA_LENGTH"`
	CreateTime sql.NullTime `db:"CREATE_TIME"`
	UpdateTime sql.NullTime `db:"UPDATE_TIME"`
	TableCollation string `db:"TABLE_COLLATION"`
	TableComment string `db:"TABLE_COMMENT"`
}