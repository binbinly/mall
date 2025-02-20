// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSmsSkuLadder = "sms_sku_ladder"

// SmsSkuLadder mapped from table <sms_sku_ladder>
type SmsSkuLadder struct {
	ID        int   `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	SkuID     int64 `gorm:"column:sku_id;not null;comment:sku_id" json:"sku_id"`          // sku_id
	FullCount int64 `gorm:"column:full_count;not null;comment:满几件" json:"full_count"`     // 满几件
	Discount  int64 `gorm:"column:discount;not null;comment:打几折" json:"discount"`         // 打几折
	Price     int64 `gorm:"column:price;not null;comment:折后价" json:"price"`               // 折后价
	IsSuper   int8  `gorm:"column:is_super;not null;comment:是否叠加优惠" json:"is_super"`      // 是否叠加优惠
}

// TableName SmsSkuLadder's table name
func (*SmsSkuLadder) TableName() string {
	return TableNameSmsSkuLadder
}
