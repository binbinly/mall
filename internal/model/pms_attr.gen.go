// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePmsAttr = "pms_attr"

// PmsAttr mapped from table <pms_attr>
type PmsAttr struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                        // ID
	CatID     int       `gorm:"column:cat_id;not null;comment:产品分类" json:"cat_id"`                                   // 产品分类
	Name      string    `gorm:"column:name;not null;comment:属性名" json:"name"`                                        // 属性名
	Icon      string    `gorm:"column:icon;not null;comment:属性图标" json:"icon"`                                       // 属性图标
	Values    string    `gorm:"column:values;not null;comment:可选值多个逗号连接" json:"values"`                              // 可选值多个逗号连接
	Type      int8      `gorm:"column:type;not null;comment:属性类型[0-销售属性，1-基本属性，2-既是销售属性又是基本属性]" json:"type"`         // 属性类型[0-销售属性，1-基本属性，2-既是销售属性又是基本属性]
	IsMany    int8      `gorm:"column:is_many;not null;comment:值类型，是否多个值" json:"is_many"`                            // 值类型，是否多个值
	IsSearch  int8      `gorm:"column:is_search;not null;comment:是否需要搜索" json:"is_search"`                           // 是否需要搜索
	IsShow    int8      `gorm:"column:is_show;not null;comment:是否展示在介绍上" json:"is_show"`                             // 是否展示在介绍上
	IsRelease int8      `gorm:"column:is_release;not null;comment:是否发布上线" json:"is_release"`                         // 是否发布上线
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
}

// TableName PmsAttr's table name
func (*PmsAttr) TableName() string {
	return TableNamePmsAttr
}
