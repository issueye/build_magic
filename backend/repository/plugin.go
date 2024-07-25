package repository

import "github.com/issueye/build_magic/backend/common/model"

type CreatePlugin struct {
	model.PluginBase
}

type ModifyPlugin struct {
	model.PluginInfo
}
