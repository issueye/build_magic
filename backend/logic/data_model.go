package logic

import (
	commonService "github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/repository"
	"github.com/issueye/build_magic/backend/service"
)

// 数据模型
type DataModel struct {
}

// 创建数据模型
func (lc *DataModel) CreateDataModel(data *repository.CreateDataModel) error {
	srv := commonService.NewService(&service.DataModel{})
	return srv.Create(data)
}
