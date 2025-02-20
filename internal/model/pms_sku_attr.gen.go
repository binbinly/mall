// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsSkuAttr = "pms_sku_attr"

// PmsSkuAttr mapped from table <pms_sku_attr>
type PmsSkuAttr struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	SkuID     int    `gorm:"column:sku_id;not null;comment:sku_id" json:"sku_id"`          // sku_id
	AttrID    int    `gorm:"column:attr_id;not null;comment:属性id" json:"attr_id"`          // 属性id
	AttrName  string `gorm:"column:attr_name;not null;comment:销售属性名" json:"attr_name"`     // 销售属性名
	AttrValue string `gorm:"column:attr_value;not null;comment:销售属性值" json:"attr_value"`   // 销售属性值
	Sort      int32  `gorm:"column:sort;not null;default:50;comment:排序" json:"sort"`       // 排序
}

// TableName PmsSkuAttr's table name
func (*PmsSkuAttr) TableName() string {
	return TableNamePmsSkuAttr
}
