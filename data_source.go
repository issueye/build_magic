package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/issueye/build_magic/backend/common/model"
	commonService "github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/global"
	"github.com/issueye/build_magic/backend/pkg/db"
	"github.com/issueye/build_magic/backend/repository"
	"github.com/issueye/build_magic/backend/service"
	"gorm.io/gorm"
)

var dataSource *DataSource

// 数据模型
type DataSource struct {
	Ctx context.Context
	db  *gorm.DB
}

// NewApp creates a new App application struct
func GeDataSource() *DataSource {
	if dataSource == nil {
		dataSource = &DataSource{Ctx: context.Background()}
	}

	return dataSource
}

// 创建数据模型
func (lc *DataSource) Create(data *repository.CreateDataSource) error {
	srv := commonService.NewService(&service.DataSource{})
	return srv.Create(data)
}

// 创建数据模型
func (lc *DataSource) Delete(id string) error {
	srv := commonService.NewService(&service.DataSource{})
	return srv.Delete(id)
}

// 创建数据模型
func (lc *DataSource) Modify(data *repository.UpdateDataSource) error {
	srv := commonService.NewService(&service.DataSource{})
	return srv.Modify(data)
}

func (lc *DataSource) List(condition string, page, size int) ([]*model.DataSource, error) {
	srv := commonService.NewService(&service.DataSource{})
	qry := model.NewPage(repository.QueryDataSource{})
	qry.Condition.Condition = condition
	qry.PageNum = int64(page)
	qry.PageSize = int64(size)

	return srv.List(qry)
}

type TablelInfo struct {
	TableName string `json:"table_name" gorm:"column:table_name"` // 表名称
}

func (lc *DataSource) GetTableList(id string) (list []*TablelInfo, err error) {
	srv := commonService.NewService(&service.DataSource{})
	info, err := srv.GetById(id)
	if err != nil {
		return nil, err
	}

	fmt.Println("info", info.DbType)

	list = make([]*TablelInfo, 0)

	switch info.DbType {
	case model.MYSQL:
		{
			sqlStr := "select table_name from information_schema.tables where table_schema = ?"
			res := lc.db.Raw(sqlStr, info.Database)
			if res.Error != nil {
				return nil, err
			}

			err = res.Scan(&list).Error
			if err != nil {
				return nil, err
			}
		}
	case model.SQLSERVER:
		{
			sqlStr := "select name as table_name from sys.tables"
			fmt.Println("sqlStr", sqlStr)
			res := lc.db.Raw(sqlStr)
			if res.Error != nil {
				global.Log.Errorf("获取表列表失败 %s", res.Error.Error())
				return nil, err
			}

			err = res.Scan(&list).Error
			if err != nil {
				global.Log.Errorf("获取表列表失败 %s", err.Error())
				return nil, err
			}

			// global.Log.Info("list", list)
		}
	case model.SQLITE3:
		{

			sqlStr := "select name as table_name from sqlite_master where type = 'table'"
			res := lc.db.Raw(sqlStr)
			if res.Error != nil {
				return nil, err
			}

			err = res.Scan(&list).Error
			if err != nil {
				return nil, err
			}
		}

	default:
		return nil, errors.New("不支持的数据库类型")
	}

	return
}

type ColumnInfo struct {
	Name string `json:"name" gorm:"column:name"` // 列名称
	Type string `json:"type" gorm:"column:type"` // 列类型
	Size int    `json:"size" gorm:"column:size"` // 列大小
}

func (lc *DataSource) GetColumns(id string, tableName string) (list []*ColumnInfo, err error) {
	srv := commonService.NewService(&service.DataSource{})
	info, err := srv.GetById(id)
	if err != nil {
		return nil, err
	}

	list = make([]*ColumnInfo, 0)

	switch info.DbType {
	case model.MYSQL:
		{
			sqlStr := "select column_name as name, data_type as type, character_maximum_length as size from information_schema.columns where table_schema = ? and table_name = ?"
			res := lc.db.Raw(sqlStr, info.Database, tableName)
			if res.Error != nil {
				return nil, res.Error
			}

			err = res.Scan(&list).Error
			if err != nil {
				return nil, err
			}
		}
	case model.SQLSERVER:
		{
			sqlStr := "select COLUMN_NAME as name, DATA_TYPE as type, CHARACTER_MAXIMUM_LENGTH as size from INFORMATION_SCHEMA.COLUMNS where TABLE_NAME = ?"
			global.Log.Infof("sqlStr: %s", sqlStr)
			// sqlStr = fmt.Sprintf(sqlStr, tableName)
			res := lc.db.Raw(sqlStr, tableName)
			if res.Error != nil {
				return nil, res.Error
			}

			err = res.Scan(&list).Error
			if err != nil {
				return nil, err
			}
		}
	case model.SQLITE3:
		{
			sqlStr := "pragma table_info(" + tableName + ")"
			res := lc.db.Raw(sqlStr)
			if res.Error != nil {
				return nil, res.Error
			}

			err = res.Scan(&list).Error
			if err != nil {
				return nil, err
			}
		}
	default:
		return nil, errors.New("不支持的数据库类型")
	}

	// fmt.Println("list", list)

	return
}

func (lc *DataSource) LinkDataSource(id string) error {
	srv := commonService.NewService(&service.DataSource{})
	info, err := srv.GetById(id)
	if err != nil {
		return err
	}

	switch info.DbType {
	case model.MYSQL:
		{
			lc.db, err = db.InitMysql(&db.Config{
				Host:     info.Host,
				Port:     info.Port,
				Username: info.UserName,
				Password: info.Password,
				Database: info.Database,
			}, global.Logger)

		}
	case model.SQLSERVER:
		{
			lc.db, err = db.InitSqlServer(&db.Config{
				Host:     info.Host,
				Port:     info.Port,
				Username: info.UserName,
				Password: info.Password,
				Database: info.Database,
			}, global.Logger)
		}
	case model.SQLITE3:
		lc.db, err = db.InitSqlite(info.Path, global.Logger)
	}

	return err
}
