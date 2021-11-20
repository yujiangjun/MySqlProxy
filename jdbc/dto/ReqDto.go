package dto

// InsertDataForTab 插入数据for 表
type InsertDataForTab struct {
	DatabaseId int `json:"databaseId"`
	Schema string `json:"schema"`
	InsertSql string `json:"insertSql"`
}

type AlertTab struct {
	DatabaseId int    `json:"databaseId"`
	Schema     string `json:"schema"`
	Table      string `json:"table"`
	Columns    []struct {
		ColumnName string `json:"columnName"`
		ColumnType string `json:"columnType"`
		Default string `json:"default"`
		Comment string `json:"comment"`
	} `json:"columns"`
}

//ChangeColReq 修改字段
type ChangeColReq struct {
	DatabaseId int    `json:"databaseId"`
	Schema     string `json:"schema"`
	Table      string `json:"table"`
	Columns    []struct {
		OriColName string `json:"oriColName"`
		NewColName string `json:"newColName"`
		ColumnType string `json:"columnType"`
		Default string `json:"default"`
		Comment string `json:"comment"`
	} `json:"columns"`
}

//type BaseTableReq struct {
//	DatabaseId int `json:"databaseId"`
//	Schema string `json:"schema"`
//	Table string `json:"table"`
//}

//DeleteColForTabReq 删除字段
type DeleteColForTabReq struct {
	DatabaseId *int `json:"databaseId" binding:"required"`
	Schema *string `json:"schema" binding:"required"`
	Table *string `json:"table" binding:"required"`
	Columns *[]string `json:"columns" binding:"required"`
}

//ChangeTabNameReq 修改表名
type ChangeTabNameReq struct {
	DatabaseId int `json:"databaseId" binding:"required"`
	Schema string `json:"schema" binding:"required"`
	OriTab string `json:"oriTab" binding:"required"`
	NewTab string `json:"newTab" binding:"required"`
}

type DropTabReq struct {
	DatabaseId int `json:"databaseId"`
	Schema string `json:"schema"`
	Table string `json:"table"`
}

type ExecSql struct {
	DatabaseId int    `json:"databaseId" binding:"required"`
	Schema     string `json:"schema" binding:"required"`
	Sql        string `json:"sql" binding:"required"`
}