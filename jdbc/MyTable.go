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
	TableCateLog sql.NullString `db:"TABLE_CATALOG"`
	TableSchema sql.NullString `db:"TABLE_SCHEMA"`
	TableName sql.NullString `db:"TABLE_NAME"`
	TableType sql.NullString `db:"TABLE_TYPE"`
	Engine sql.NullString `db:"ENGINE"`
	Version sql.NullInt16 `db:"VERSION"`
	RowFormat sql.NullString `db:"ROW_FORMAT"`
	TableRows sql.NullInt16 `db:"TABLE_ROWS"`
	DataLength sql.NullInt16 `db:"DATA_LENGTH"`
	CreateTime sql.NullTime `db:"CREATE_TIME"`
	UpdateTime sql.NullTime `db:"UPDATE_TIME"`
	TableCollation sql.NullString `db:"TABLE_COLLATION"`
	TableComment sql.NullString `db:"TABLE_COMMENT"`
}