package model

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"strings"

	"github.com/issueye/build_magic/backend/pkg/utils"
)

type DataSource struct {
	Base
	DataSourceBase
}

type DB_TYPE_ENUM string

const (
	SQLSERVER  DB_TYPE_ENUM = "sqlserver"
	MYSQL      DB_TYPE_ENUM = "mysql"
	POSTGRESQL DB_TYPE_ENUM = "postgresql"
	SQLITE3    DB_TYPE_ENUM = "sqlite3"
)

type DataSourceBase struct {
	Title    string       `binding:"required" label:"名称" gorm:"column:title;size:300;comment:标题;" json:"title"`             // 标题
	Host     string       `binding:"required" label:"主机" gorm:"column:host;size:300;comment:主机;" json:"host"`               // 主机
	Port     int          `binding:"required" label:"端口" gorm:"column:port;size:11;comment:端口;" json:"port"`                // 端口
	Database string       `binding:"required" label:"数据库名称" gorm:"column:database;size:300;comment:数据库名称;" json:"database"` // 数据库名称
	UserName string       `binding:"required" label:"用户名" gorm:"column:user_name;size:300;comment:用户名;" json:"user_name"`   // 用户名
	Password string       `binding:"required" label:"密码" gorm:"column:password;size:300;comment:密码;" json:"password"`       // 密码
	DbType   DB_TYPE_ENUM `binding:"required" label:"数据库类型" gorm:"column:db_type;size:300;comment:数据库类型;" json:"db_type"`   // 数据库类型 sqlserver mysql postgresql sqlite3
	Schema   string       `binding:"required" label:"数据库名称" gorm:"column:schema;size:300;comment:数据库名称;" json:"schema"`     // 数据库名称
	Path     string       `binding:"required" label:"路径" gorm:"column:path;size:1000;comment:路径;" json:"path"`              // 路径
	Mark     string       `binding:"required" label:"备注" gorm:"column:mark;size:1000;comment:备注;" json:"mark"`              // 备注
}

func (enum *DB_TYPE_ENUM) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("不匹配的数据类型")
	}

	*enum = DB_TYPE_ENUM(str)
	return nil
}

func (enum DB_TYPE_ENUM) Value() (driver.Value, error) {
	data := string(enum)

	return data, nil
}

func (enum *DB_TYPE_ENUM) UnmarshalJSON(data []byte) error {
	val := string(data)
	val = strings.Trim(val, "\"")
	*enum = DB_TYPE_ENUM(val)
	return nil
}

func NewDataSource(base *DataSourceBase) *DataSource {
	db := &DataSource{
		DataSourceBase: *base,
	}
	db.ID = strconv.FormatInt(utils.GenID(), 10)

	return db
}

func (DataSource) TableName() string {
	return "data_source"
}
