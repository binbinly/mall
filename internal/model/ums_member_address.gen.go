// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUmsMemberAddress = "ums_member_address"

// UmsMemberAddress mapped from table <ums_member_address>
type UmsMemberAddress struct {
	ID        int            `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                        // ID
	MemberID  int            `gorm:"column:member_id;not null;comment:用户id" json:"member_id"`                             // 用户id
	Name      string         `gorm:"column:name;not null;comment:收货人姓名" json:"name"`                                      // 收货人姓名
	Phone     string         `gorm:"column:phone;not null;comment:收货人手机号" json:"phone"`                                   // 收货人手机号
	Province  string         `gorm:"column:province;not null;comment:省" json:"province"`                                  // 省
	City      string         `gorm:"column:city;not null;comment:市" json:"city"`                                          // 市
	County    string         `gorm:"column:county;not null;comment:区/县" json:"county"`                                    // 区/县
	AreaCode  int            `gorm:"column:area_code;not null;comment:最后一级地区编码" json:"area_code"`                         // 最后一级地区编码
	Detail    string         `gorm:"column:detail;not null;comment:详细地址" json:"detail"`                                   // 详细地址
	IsDefault int8           `gorm:"column:is_default;not null;comment:是否默认" json:"is_default"`                           // 是否默认
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                                    // 删除时间
}

// TableName UmsMemberAddress's table name
func (*UmsMemberAddress) TableName() string {
	return TableNameUmsMemberAddress
}
