package main

import (
	"context"

	"github.com/issueye/build_magic/backend/common/model"
	commonService "github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/logic"
	"github.com/issueye/build_magic/backend/repository"
	"github.com/issueye/build_magic/backend/service"
)

var template *Template

// 数据模型
type Template struct {
	Ctx context.Context
}

// NewApp creates a new App application struct
func GetTemplate() *Template {
	if template == nil {
		template = &Template{}
	}

	return template
}

// 创建数据模型
func (app *Template) Create(data *repository.CreateTemplate) error {
	tplc := new(logic.Template)
	return tplc.CreateNode(&repository.CreateChildScheme{
		Title:      data.Title,
		ParentCode: data.SchemeCode,
		NodeType:   data.NodeType,
		FileType:   data.FileType,
		Icon:       data.Icon,
	})
}

func (app *Template) List(condition repository.QryTemplate, page, size int) ([]*model.CodeTemplate, error) {
	srv := commonService.NewService(&service.Template{})
	qry := model.NewPage(condition)
	qry.PageNum = int64(page)
	qry.PageSize = int64(size)

	return srv.List(qry)
}

func (app *Template) GetCode(id string) (string, error) {
	return new(logic.Template).GetCode(id)
}

func (app *Template) SaveCode(id string, code string) error {
	return new(logic.Template).SaveCodeToFile(id, code)
}

// 创建数据模型
func (app *Template) Delete(id string) error {
	return new(logic.Template).DeleteNode(id)
}

// 创建数据模型
func (app *Template) DeleteByCode(code string) error {
	return new(logic.Template).DeleteNode(code)
}

// Modify
// 修改代码模板信息
func (app *Template) Modify(data *repository.ModifyTemplate) error {
	tplc := new(logic.Template)
	return tplc.ModifyNode(data)
}

// DeleteTreeNode
// 删除树节点
func (app *Template) DeleteTreeNode(code string) error {
	return new(logic.Template).DeleteTreeNode(code)
}

// GetTree
// 获取树
func (app *Template) GetTree() (trees []*repository.SchemeTree, err error) {
	return new(logic.Template).GetTree()
}

// GetTreeByCode
// 通过编码获取树
func (app *Template) GetTreeByCode(code string) (trees []*repository.SchemeTree, err error) {
	return new(logic.Template).GetTreeByCode(code)
}

// GetTpByCode
// 通过编码获取代码模板
func (app *Template) GetTpByCode(code string) (data *model.CodeTemplate, err error) {
	return new(logic.Template).GetTpByCode(code)
}

// ProgrammeList
// 获取方案列表
func (app *Template) ProgrammeList() (data []*model.Scheme, err error) {
	data, err = new(logic.Template).ProgrammeList()
	return
}

func (app *Template) GetChildrenByCode(code string) (list []*model.CodeTemplate, err error) {
	list, err = new(logic.Template).GetChildrenByCode(code)
	return
}
