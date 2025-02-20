// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsCategory = "pms_category"

// PmsCategory mapped from table <pms_category>
type PmsCategory struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`  // ID
	Name        string `gorm:"column:name;not null;comment:分类名" json:"name"`                  // 分类名
	ParentID    int    `gorm:"column:parent_id;not null;comment:上级分类" json:"parent_id"`       // 上级分类
	Level       int8   `gorm:"column:level;not null;default:1;comment:层级" json:"level"`       // 层级
	Icon        string `gorm:"column:icon;not null;comment:图标" json:"icon"`                   // 图标
	ProductUnit string `gorm:"column:product_unit;not null;comment:计量单位" json:"product_unit"` // 计量单位
	IsRelease   int8   `gorm:"column:is_release;not null;comment:是否发布上线" json:"is_release"`   // 是否发布上线
	Sort        int32  `gorm:"column:sort;not null;default:50;comment:排序" json:"sort"`        // 排序
}

// TableName PmsCategory's table name
func (*PmsCategory) TableName() string {
	return TableNamePmsCategory
}
