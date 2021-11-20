package _const

type DataType struct {
	name string
}

func DataTypeList() []string {
	return []string{
		"TINYINT",
		"SMALLINT",
		"MEDIUMINT",
		"INT",
		"BIGINT",
		"FLOAT",
		"DOUBLE",
		"DECIMAL",
		"DATE",
		"TIME",
		"YEAR",
		"DATETIME",
		"TIMESTAMP",
		"CHAR",
		"VARCHAR",
		"VARCHAR",
		"TINYTEXT",
		"BLOB",
		"TEXT",
		"TEXT",
		"MEDIUMTEXT",
		"LONGBLOB",
		"LONGTEXT",
	}
}