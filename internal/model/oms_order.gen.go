// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameOmsOrder = "oms_order"

// OmsOrder mapped from table <oms_order>
type OmsOrder struct {
	ID                int            `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                        // ID
	MemberID          int            `gorm:"column:member_id;not null;comment:用户id" json:"member_id"`                             // 用户id
	OrderNo           string         `gorm:"column:order_no;not null;comment:订单号" json:"order_no"`                                // 订单号
	CouponID          int            `gorm:"column:coupon_id;not null;comment:优惠券id" json:"coupon_id"`                            // 优惠券id
	Username          string         `gorm:"column:username;not null;comment:用户名" json:"username"`                                // 用户名
	TotalAmount       int            `gorm:"column:total_amount;not null;comment:订单总额/分" json:"total_amount"`                     // 订单总额/分
	FreightAmount     int            `gorm:"column:freight_amount;not null;comment:运费/分" json:"freight_amount"`                   // 运费/分
	PromotionAmount   int            `gorm:"column:promotion_amount;not null;comment:促销优惠金额（促销价、满减、阶梯价）" json:"promotion_amount"` // 促销优惠金额（促销价、满减、阶梯价）
	IntegrationAmount int            `gorm:"column:integration_amount;not null;comment:积分抵扣金额" json:"integration_amount"`         // 积分抵扣金额
	CouponAmount      int            `gorm:"column:coupon_amount;not null;comment:优惠券折扣金额" json:"coupon_amount"`                  // 优惠券折扣金额
	DiscountAmount    int            `gorm:"column:discount_amount;not null;comment:后台调整订单使用的折扣金额" json:"discount_amount"`        // 后台调整订单使用的折扣金额
	Amount            int            `gorm:"column:amount;not null;comment:优惠后的真实金额" json:"amount"`                               // 优惠后的真实金额
	PayAmount         int            `gorm:"column:pay_amount;not null;comment:支付金额/分" json:"pay_amount"`                         // 支付金额/分
	PayType           int8           `gorm:"column:pay_type;not null;comment:支付方式" json:"pay_type"`                               // 支付方式
	SourceType        int8           `gorm:"column:source_type;not null;comment:订单来源[0->PC订单；1->app订单]" json:"source_type"`       // 订单来源[0->PC订单；1->app订单]
	Status            int8           `gorm:"column:status;not null;comment:订单状态" json:"status"`                                   // 订单状态
	DeliveryCompany   string         `gorm:"column:delivery_company;not null;comment:物流公司(配送方式)" json:"delivery_company"`         // 物流公司(配送方式)
	DeliveryNo        string         `gorm:"column:delivery_no;not null;comment:物流单号" json:"delivery_no"`                         // 物流单号
	Integration       int            `gorm:"column:integration;not null;comment:可以获得的积分" json:"integration"`                      // 可以获得的积分
	Growth            int            `gorm:"column:growth;not null;comment:可以获得的成长值" json:"growth"`                               // 可以获得的成长值
	AddressName       string         `gorm:"column:address_name;not null;comment:收货人姓名" json:"address_name"`                      // 收货人姓名
	AddressPhone      string         `gorm:"column:address_phone;not null;comment:收货人手机" json:"address_phone"`                    // 收货人手机
	AddressProvince   string         `gorm:"column:address_province;not null;comment:省" json:"address_province"`                  // 省
	AddressCity       string         `gorm:"column:address_city;not null;comment:市" json:"address_city"`                          // 市
	AddressCounty     string         `gorm:"column:address_county;not null;comment:县/区" json:"address_county"`                    // 县/区
	AddressDetail     string         `gorm:"column:address_detail;not null;comment:详情" json:"address_detail"`                     // 详情
	Note              string         `gorm:"column:note;not null;comment:备注" json:"note"`                                         // 备注
	IsConfirm         int8           `gorm:"column:is_confirm;not null;comment:是否确认收货" json:"is_confirm"`                         // 是否确认收货
	UseIntegration    int64          `gorm:"column:use_integration;not null;comment:下单时使用的积分" json:"use_integration"`             // 下单时使用的积分
	TradeNo           string         `gorm:"column:trade_no;not null;comment:交易号" json:"trade_no"`                                // 交易号
	TransHash         string         `gorm:"column:trans_hash;not null;comment:交易hash" json:"trans_hash"`                         // 交易hash
	PayAt             time.Time      `gorm:"column:pay_at;comment:支付时间" json:"pay_at"`                                            // 支付时间
	DeliveryAt        time.Time      `gorm:"column:delivery_at;comment:发货时间" json:"delivery_at"`                                  // 发货时间
	ReceiveAt         time.Time      `gorm:"column:receive_at;comment:确认收货时间" json:"receive_at"`                                  // 确认收货时间
	CommentAt         time.Time      `gorm:"column:comment_at;comment:评价时间" json:"comment_at"`                                    // 评价时间
	CreatedAt         time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt         time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                                    // 删除时间
}

// TableName OmsOrder's table name
func (*OmsOrder) TableName() string {
	return TableNameOmsOrder
}
