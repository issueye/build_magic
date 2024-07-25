package repository

import "github.com/issueye/build_magic/backend/common/model"

type ModifyScheme struct {
	model.Scheme
}

type CreateScheme struct {
	model.Scheme
}

// 查询请求数据
type QryScheme struct {
	Condition string   `json:"condition"` // 条件
	Ids       []string `json:"ids"`       // 模板id
}

type SchemeTree struct {
	Code       string        `json:"id"`         // code
	Title      string        `json:"title"`      // 名称
	Icon       string        `json:"icon"`       // 图标
	Type       int           `json:"type"`       // 0 跟节点 1 文件夹节点 2 文件节点
	ParentCode string        `json:"parentCode"` // 父节点code
	Path       string        `json:"path"`       // 路径
	Children   []*SchemeTree `json:"children"`   //
}

type CreateChildScheme struct {
	Title      string `binding:"required" label:"名称" gorm:"column:title;size:300;comment:标题;" json:"title"`              // 标题
	ParentCode string `binding:"required" label:"父编码" gorm:"column:parent_code;size:100;comment:父编码;" json:"parentCode"` // 父编码
	Icon       string `binding:"required" label:"图标" gorm:"column:icon;size:500;comment:图标;" json:"icon"`                // 图标  flat-color-icons:folder  flat-color-icons:file vscode-icons:default-root-folder-opened
	NodeType   int    `binding:"required" label:"节点类型" gorm:"column:node_type;size:10;comment:节点类型;" json:"nodeType"`    // 节点类型 0 根节点 1 文件夹节点 2 文件节点
	FileType   int    `binding:"required" label:"文件类型" gorm:"column:file_type;size:10;comment:文件类型;" json:"fileType"`    // 文件类型 0 文件夹 1 文件
}
