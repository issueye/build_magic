package model

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"
)

type Page[T any] struct {
	PageNum   int64 `json:"pageNum" form:"pageNum"`   // 页数
	PageSize  int64 `json:"pageSize" form:"pageSize"` // 页码
	Total     int64 `json:"total"`                    // 总数  由服务器返回回去
	Condition T     `json:"condition"`                // 条件
}

func NewPage[T any](condition T) *Page[T] {
	return &Page[T]{Condition: condition}
}

type Base struct {
	ID        string    `gorm:"column:id;size:100;comment:编码;primaryKey;autoIncrement:false;" json:"id"` // 编码
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;" json:"createdAt"`          // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;comment:修改时间;" json:"updatedAt"`          // 修改时间
}

// 数组数据类型 数据中以逗号分隔的字符串
type Arr []string

// 存入数据库
func (arr Arr) Value() (driver.Value, error) {
	if len(arr) > 0 {
		str := arr[0]
		for _, v := range arr[1:] {
			str += "," + v
		}
		return str, nil
	} else {
		return "", nil
	}
}

// 从数据库取数据
func (arr *Arr) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("不匹配的数据类型")
	}
	*arr = strings.Split(string(str), ",")
	return nil
}

func (arr *Arr) UnmarshalJSON(data []byte) error {
	d := string(data)

	if d == "null" {
		*arr = []string{}
		return nil
	}

	// 处理 ["983193","983173"] 这种情况
	if strings.HasPrefix(d, "[") || strings.HasPrefix(d, "]") {
		d = strings.Trim(d, "[")
		d = strings.Trim(d, "]")
	}

	if d != "" {
		arrData := strings.Split(d, ",")
		for _, data := range arrData {
			data = strings.Trim(data, "\"")
			*arr = append(*arr, data)
		}
	} else {
		*arr = []string{}
	}

	return nil
}
