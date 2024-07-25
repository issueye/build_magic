package repository

import "github.com/issueye/build_magic/backend/common/model"

type CreateDataModel struct {
	model.DataModelBase
}

type RequestModifyDataModel struct {
	model.DataModel
}

// 查询请求数据
type RequestDataModelQuery struct {
	Condition string `json:"condition"` // 条件
}

// 查询请求数据详情
type RequestModelInfoSave struct {
	model.ModelInfo
}
