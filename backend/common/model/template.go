package model

import (
	"strconv"

	"github.com/issueye/build_magic/backend/pkg/utils"
)

// 数据模型
type CodeTemplate struct {
	Base
	CodeTemplateBase
}

// TableName
// 表名称
func (CodeTemplate) TableName() string {
	return "code_template"
}

type CodeTemplateBase struct {
	Title            string `binding:"required" label:"名称" gorm:"column:title;size:300;comment:标题;" json:"title"`                                // 标题
	FileName         string `binding:"required" label:"文件名" gorm:"column:file_name;size:300;comment:文件名;" json:"fileName"`                       // 文件名
	SchemeCode       string `binding:"required" label:"模板代码" gorm:"column:scheme_code;type:text;comment:模板代码;" json:"schemeCode"`                // 模板代码
	SchemeParentCode string `binding:"required" label:"父模板代码" gorm:"column:scheme_parent_code;type:text;comment:父模板代码;" json:"schemeParentCode"` // 父模板代码
	Mark             string `binding:"required" label:"备注" gorm:"column:mark;size:300;comment:备注;" json:"mark"`                                  // 备注
	NodeType         int    `binding:"required" label:"节点类型" gorm:"-" json:"nodeType"`                                                           // 节点类型
	FileType         int    `binding:"required" label:"文件类型" gorm:"-" json:"fileType"`                                                           // 文件类型 0 js 1 ts
	Icon             string `binding:"required" label:"图标" gorm:"-" json:"icon"`                                                                 // 图标
}

func NewCodeTemplate(base *CodeTemplateBase) *CodeTemplate {
	db := &CodeTemplate{
		CodeTemplateBase: *base,
	}
	db.ID = strconv.FormatInt(utils.GenID(), 10)

	return db
}
