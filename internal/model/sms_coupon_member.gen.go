// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSmsCouponMember = "sms_coupon_member"

// SmsCouponMember mapped from table <sms_coupon_member>
type SmsCouponMember struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`            // ID
	MemberID  int       `gorm:"column:member_id;not null;comment:用户id" json:"member_id"`                 // 用户id
	CouponID  int       `gorm:"column:coupon_id;not null;comment:优惠券id" json:"coupon_id"`                // 优惠券id
	Nickname  string    `gorm:"column:nickname;not null;comment:会员昵称" json:"nickname"`                   // 会员昵称
	GetType   int8      `gorm:"column:get_type;not null;comment:获取方式[0->后台赠送；1->主动领取]" json:"get_type"`  // 获取方式[0->后台赠送；1->主动领取]
	Status    int8      `gorm:"column:status;not null;comment:使用状态[0->未使用；1->已使用；2->已过期]" json:"status"` // 使用状态[0->未使用；1->已使用；2->已过期]
	OrderID   int       `gorm:"column:order_id;not null;comment:订单id" json:"order_id"`                   // 订单id
	OrderNo   string    `gorm:"column:order_no;not null;comment:订单号" json:"order_no"`                    // 订单号
	UsedAt    time.Time `gorm:"column:used_at;comment:使用时间" json:"used_at"`                              // 使用时间
	CreatedAt time.Time `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`               // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;comment:更新时间" json:"updated_at"`               // 更新时间
}

// TableName SmsCouponMember's table name
func (*SmsCouponMember) TableName() string {
	return TableNameSmsCouponMember
}
