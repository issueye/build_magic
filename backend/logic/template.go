package logic

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/issueye/build_magic/backend/common/model"
	commonService "github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/global"
	"github.com/issueye/build_magic/backend/pkg/utils"
	"github.com/issueye/build_magic/backend/repository"
	"github.com/issueye/build_magic/backend/service"
)

type Template struct{}

// 保存代码到文件
func (t *Template) SaveCodeToFile(tpId string, code string) error {
	srv := commonService.NewService(&service.Template{})
	info, err := srv.Get(tpId)
	if err != nil {
		return err
	}

	// 获取文件名称和文件类型
	path := filepath.Join("runtime", "static", "code", info.FileName)
	return os.WriteFile(path, []byte(code), 0644)
}

func (t *Template) GetCodeBySchemeID(id string) (string, error) {
	srv := commonService.NewService(&service.Scheme{})
	info, err := srv.Get(id)
	if err != nil {
		return "", err
	}

	// 通过 schemeID 获取模板ID
	tpSrv := commonService.NewService(&service.Template{})
	tpInfo, err := tpSrv.GetBySchemeCode(info.Code)
	if err != nil {
		return "", err
	}

	return t.GetCode(tpInfo.ID)
}

// 获取文件内容
func (t *Template) GetCode(tpId string) (string, error) {
	srv := commonService.NewService(&service.Template{})
	info, err := srv.Get(tpId)
	if err != nil {
		return "", err
	}

	// fileName := fmt.Sprintf("%s.js", info.FileName)
	path := filepath.Join("runtime", "static", "code", info.FileName)
	exist := utils.FileExist(path)
	// fmt.Println("exist", exist)
	if !exist {
		return "", nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (t *Template) GetTree() (trees []*repository.SchemeTree, err error) {
	trees = make([]*repository.SchemeTree, 0)

	schemeSrv := commonService.NewService(&service.Scheme{})
	schemes, err := schemeSrv.List(model.NewPage(repository.QryScheme{}))
	if err != nil {
		return nil, err
	}

	// 先创建根节点
	for _, scheme := range schemes {
		if scheme.Level == 0 {
			tree := t.createNode(scheme)
			children := t.findChildren(scheme.Code, schemes)

			if len(children) > 0 {
				tree.Children = t.createNodes(children, schemes)
			}

			trees = append(trees, tree)
		}
	}

	return
}

func (t *Template) GetTpByCode(code string) (*model.CodeTemplate, error) {
	srv := commonService.NewService(&service.Template{})
	return srv.GetBySchemeCode(code)
}

func (t *Template) GetTreeByCode(code string) (trees []*repository.SchemeTree, err error) {
	trees = make([]*repository.SchemeTree, 0)
	// 获取当前节点信息
	schemeSrv := commonService.NewService(&service.Scheme{})
	scheme, err := schemeSrv.GetByCode(code)
	if err != nil {
		return nil, err
	}

	var schemes []*model.Scheme

	if len(scheme.Nodes) == 0 {
		return nil, nil
	}

	if scheme.Nodes[0] == "001" {
		schemes, err = schemeSrv.GetCommonNode()
	} else {
		schemes, err = schemeSrv.GetNodeByCode(scheme.Nodes, scheme.ParentCode)
	}

	if err != nil {
		return nil, err
	}

	// 先创建根节点
	for _, scheme := range schemes {
		if scheme.Level == 0 {
			tree := t.createNode(scheme)
			children := t.findChildren(scheme.Code, schemes)

			if len(children) > 0 {
				tree.Children = t.createNodes(children, schemes)
			}

			trees = append(trees, tree)
		}
	}

	return
}

func (t *Template) createNodes(list []*model.Scheme, baseList []*model.Scheme) []*repository.SchemeTree {
	trees := make([]*repository.SchemeTree, 0)
	for _, scheme := range list {
		node := t.createNode(scheme)
		// fmt.Println("scheme", scheme.Code, scheme.Title)
		children := t.findChildren(scheme.Code, baseList)
		// fmt.Printf("children: %v\n", children)
		if len(children) > 0 {
			node.Children = t.createNodes(children, baseList)
		}

		trees = append(trees, node)
	}

	return trees
}

func (t *Template) findChildren(code string, schemes []*model.Scheme) []*model.Scheme {
	children := make([]*model.Scheme, 0)
	if code == "" {
		return children
	}

	for _, scheme := range schemes {
		if scheme.ParentCode == code {
			children = append(children, scheme)
		}
	}

	return children
}

func (t *Template) createNode(info *model.Scheme) *repository.SchemeTree {
	return &repository.SchemeTree{
		Code:       info.Code,
		Title:      info.Title,
		Icon:       info.Icon,
		ParentCode: info.ParentCode,
		Type:       info.NodeType,
		Path:       info.Path,
		Children:   make([]*repository.SchemeTree, 0),
	}
}

func (t *Template) DeleteTreeNode(code string) (err error) {
	srv := commonService.NewService(&service.Scheme{})
	srv.OpenTx()

	defer func() {
		if err != nil {
			global.Log.Error("删除节点失败, 数据回滚")
			srv.Rollback()
		} else {
			global.Log.Info("删除节点成功, 数据提交")
			srv.Commit()
		}
	}()

	err = srv.Delete(code)

	if err != nil {
		global.Log.Errorf("删除节点失败: %s", err.Error())
		return
	}

	// 处理代码模板
	tpSrv := commonService.NewService(&service.Template{}, srv.GetContext())
	err = tpSrv.DeleteBySchemeCode(code)
	return
}

// 初始化脚本模板方案
func (t *Template) InitScriptScheme() {
	t.CheckOrCreate("001", "公共代码", "common")
	t.CheckOrCreate("002", "业务方案", "projects")
}

func (t *Template) CheckOrCreate(code string, title string, path string) {
	srv := commonService.NewService(&service.Scheme{})
	info, err := srv.GetByCode(code)
	if err != nil {
		global.Log.Errorf("查询方案失败: %s", err.Error())
		return
	}

	create := func() {
		scheme := &repository.CreateScheme{
			Scheme: *model.NewScheme(&model.SchemeBase{
				Code:       code,
				Title:      title,
				Icon:       "vscode-icons:default-root-folder-opened",
				Level:      0,
				NodeType:   0,
				ParentCode: "",
				Path:       path,
			}),
		}

		scheme.Nodes = append(scheme.Nodes, code)
		_, err := srv.Create(scheme)
		if err != nil {
			global.Log.Errorf("初始化脚本模板方案失败: %s", err.Error())
			return
		}
	}

	if info == nil {
		create()
	}

	if info != nil && info.Code == "" {
		create()
	}
}

func (t *Template) DeleteNode(code string) (err error) {
	srv := commonService.NewService(&service.Template{})

	srv.OpenTx()
	defer func() {
		if err != nil {
			global.Log.Error("修改节点失败, 数据回滚")
			srv.Rollback()
		} else {
			global.Log.Info("修改节点成功, 数据提交")
			srv.Commit()
		}
	}()

	// 更新方案信息
	schemeSrv := commonService.NewService(&service.Scheme{}, srv.GetContext())

	// 查询方案信息
	info, err := schemeSrv.GetByCode(code)
	if err != nil {
		return
	}

	err = schemeSrv.DeleteByCode(code)
	if err != nil {
		return
	}

	// 如果是文件夹，则需要删除子节点
	if info.NodeType == 1 {
		err = schemeSrv.DeleteByParentCode(code)
		if err != nil {
			global.Log.Errorf("删除子节点失败: %s", err.Error())
			return
		}
	}

	err = srv.DeleteBySchemeCode(code)
	if err != nil {
		return
	}

	// 删除对应的文件夹
	err = os.RemoveAll(filepath.Join("runtime", "static", "code", info.Path))
	return
}

func (t *Template) ModifyNode(data *repository.ModifyTemplate) (err error) {
	srv := commonService.NewService(&service.Template{})

	srv.OpenTx()
	defer func() {
		if err != nil {
			global.Log.Error("修改节点失败, 数据回滚")
			srv.Rollback()
		} else {
			global.Log.Info("修改节点成功, 数据提交")
			srv.Commit()
		}
	}()

	err = srv.Modify(data)
	if err != nil {
		return
	}

	// 更新方案信息
	schemeSrv := commonService.NewService(&service.Scheme{}, srv.GetContext())
	err = schemeSrv.ModifyByCode(&repository.ModifyScheme{
		Scheme: *model.NewScheme(&model.SchemeBase{Code: data.CodeTemplateBase.SchemeCode, Title: data.Title, Icon: data.Icon}),
	})

	return
}

func (t *Template) CreateNode(data *repository.CreateChildScheme) (tpId string, err error) {
	// 获取父节点信息
	srv := commonService.NewService(&service.Scheme{})
	var parent *model.Scheme
	parent, err = srv.GetByCode(data.ParentCode)
	if err != nil {
		global.Log.Errorf("获取父节点信息失败: %s", err.Error())
		return
	}

	var path string
	if data.NodeType == 2 {
		// 文件类型
		switch data.FileType {
		case 1:
			path = filepath.Join(parent.Path, data.Title) + ".js"
		case 2:
			path = filepath.Join(parent.Path, data.Title) + ".ts"
		}
	}

	if data.NodeType == 1 {
		// 文件夹类型
		path = filepath.Join(parent.Path, data.Title)
	}

	code := strconv.FormatInt(utils.GenID(), 5)
	level := parent.Level + 1
	info := &repository.CreateScheme{
		Scheme: *model.NewScheme(&model.SchemeBase{
			Code:       code,
			Title:      data.Title,
			Icon:       data.Icon,
			Level:      level,
			NodeType:   data.NodeType,
			ParentCode: data.ParentCode,
			Path:       path,
		}),
	}

	parent.Nodes = append(parent.Nodes, code)
	info.Nodes = parent.Nodes

	_, err = srv.Create(info)
	if err != nil {
		global.Log.Errorf("初始化脚本模板方案失败: %s", err.Error())
		return
	}

	// 如果创建的是文件，则需要创建对应的代码模板
	if data.NodeType == 2 {
		tpSrv := commonService.NewService(&service.Template{})
		tpInfo, err := tpSrv.Create(&repository.CreateTemplate{
			CodeTemplateBase: model.CodeTemplateBase{
				Title:            data.Title,
				FileName:         path,
				SchemeCode:       code,
				SchemeParentCode: parent.Code,
			},
		})
		if err != nil {
			global.Log.Errorf("初始化脚本模板方案失败: %s", err.Error())
			return "", err
		}

		tpId = tpInfo.ID
	}

	info.Nodes = append(info.Nodes, data.ParentCode)

	// 创建文件夹
	if data.NodeType == 1 {
		// 判断父文件夹是否存在
		path := filepath.Join("runtime", "static", "code", info.Path)
		fmt.Println("path", path)
		utils.PathExists(path)
	}

	return
}

// ProgrammeList
// @Description: 方案列表
// @return list
// @return err
func (lc *Template) ProgrammeList() (list []*model.Scheme, err error) {
	srv := commonService.NewService(&service.Scheme{})
	return srv.ProgrammeList()
}

func (lc *Template) GetChildrenByCode(code string) (list []*model.CodeTemplate, err error) {
	srv := commonService.NewService(&service.Template{})
	return srv.GetChildrenByCode(code)
}
