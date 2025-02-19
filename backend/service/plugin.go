package service

import (
	"github.com/issueye/build_magic/backend/common/model"
	"github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/repository"
)

type Plugin struct {
	service.BaseService
}

func (owner *Plugin) Self() *Plugin {
	return owner
}

func (owner *Plugin) SetBase(base service.BaseService) {
	owner.BaseService = base
}

func NewPlugin(args ...service.ServiceContext) *Plugin {
	return service.NewService(&Plugin{}, args...)
}

// Create
// 创建信息
func (s *Plugin) Create(data *repository.CreatePlugin) error {
	info := model.PluginInfo{}.New()
	info.Name = data.Name
	info.Path = data.Path
	info.Version = data.Version
	info.Key = data.Key
	info.Value = data.Value

	return s.GetDB().Model(info).Create(info).Error
}
