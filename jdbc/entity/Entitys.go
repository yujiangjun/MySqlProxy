package entity

type DataConnect struct {
	Id int `gorm:"primaryKey"`
	DbName string
	DbType int
	Url string
	UserName string
	Password string
}
