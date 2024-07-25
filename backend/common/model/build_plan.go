package model

import (
	"strconv"

	"github.com/issueye/build_magic/backend/pkg/utils"
)

// 数据模型
type BuildPlan struct {
	Base
	BuildPlanBase
}

// TableName
// 表名称
func (BuildPlan) TableName() string {
	return "build_plan"
}

type BuildPlanBase struct {
	Title string `binding:"required" label:"名称" gorm:"column:title;size:300;comment:标题;" json:"title"` // 标题
	Mark  string `binding:"required" label:"备注" gorm:"column:mark;size:300;comment:备注;" json:"mark"`   // 备注
}

func NewBuildPlan(base *BuildPlanBase) *BuildPlan {
	db := &BuildPlan{
		BuildPlanBase: *base,
	}
	db.ID = strconv.FormatInt(utils.GenID(), 10)

	return db
}
