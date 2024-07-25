package model

// 表信息
type SqliteTableInfo struct {
	TableName string                  `json:"name"`   // 表名
	Fields    []*SqliteTableFieldInfo `json:"fields"` // 字段信息
}

// 字段信息
type SqliteTableFieldInfo struct {
	Cid          int    `gorm:"column:cid" json:"cid"`           // 列序号
	Name         string `gorm:"column:name" json:"name"`         // 列名
	Type         string `gorm:"column:type" json:"type"`         // 列类型
	IsPrimaryKey bool   `gorm:"column:pk" json:"is_primary_key"` // 是否是主键
}
