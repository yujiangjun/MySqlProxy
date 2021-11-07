package jdbc

import (
	"database/sql"
)

type Field struct {
	FName sql.NullString `db:"fName"`
	FDesc sql.NullString `db:"fDesc"`
	DataType sql.NullString `db:"dataType"`
	IsNull sql.NullString `db:"isNull"`
	SLength sql.NullString `db:"sLength"`
}
/**

 */
type Tables struct {
	TableCateLog string `db:"TABLE_CATALOG"`
	TableSchema string `db:"TABLE_SCHEMA"`
	TableName string `db:"TABLE_NAME"`
	//TableType sql.NullString `db:"TABLE_TYPE"`
	//Engine sql.NullString `db:"ENGINE"`
	//Version sql.NullInt16 `db:"VERSION"`
	//RowFormat sql.NullString `db:"ROW_FORMAT"`
	//TableRows sql.NullInt16 `db:"TABLE_ROWS"`
	//CreateTime sql.NullTime `db:"CREATE_TIME"`
	//UpdateTime sql.NullTime `db:"UPDATE_TIME"`
	//TableCollation sql.NullString `db:"TABLE_COLLATION"`
	//TableComment sql.NullString `db:"TABLE_COMMENT"`
}