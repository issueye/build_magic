package repository

import "github.com/issueye/build_magic/backend/common/model"

type CreateDataSource struct {
	model.DataSourceBase
}

type UpdateDataSource struct {
	model.DataSource
}

type QueryDataSource struct {
	Condition string `json:"condition"` // 模糊查询条件
}
