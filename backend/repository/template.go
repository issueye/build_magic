package repository

import "github.com/issueye/build_magic/backend/common/model"

type ModifyTemplate struct {
	model.CodeTemplate
}

type CreateTemplate struct {
	model.CodeTemplateBase
}

// 查询请求数据
type QryTemplate struct {
	Condition  string   `json:"condition"`  // 条件
	ParentCode string   `json:"parentCode"` // 父级代码
	Ids        []string `json:"ids"`        // 模板id
}

type PushCode struct {
	Id          string `json:"id"`           // 模板id
	CodeContent string `json:"code_content"` // 代码内容
}
