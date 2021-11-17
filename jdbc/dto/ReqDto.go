package dto

// InsertDataForTab 插入数据for 表
type InsertDataForTab struct {
	DatabaseId int `json:"databaseId"`
	Schema string `json:"schema"`
	InsertSql string `json:"insertSql"`
}
