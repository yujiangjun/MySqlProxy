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

//type BaseTableReq struct {
//	DatabaseId int `json:"databaseId"`
//	Schema string `json:"schema"`
//	Table string `json:"table"`
//}

type DeleteColForTabReq struct {
	DatabaseId *int `json:"databaseId"`
	Schema *string `json:"schema"`
	Table *string `json:"table"`
	Columns *[]string `json:"columns"`
}