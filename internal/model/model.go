package model

import (
	"encoding/json"
	"errors"
	"github.com/binbinly/pkg/auth"
	"gorm.io/gorm"
	"time"
)

const (
	// ReleaseYes 已发布
	ReleaseYes = 1
	// DefaultOrder 默认排序
	DefaultOrder     = "id DESC"
	DefaultOrderSort = "sort DESC"
)

const (
	// StatusInit 状态-初始化
	StatusInit = iota
	// StatusSuccess 状态-成功
	StatusSuccess
	// StatusError 状态-失败
	StatusError
)

const (
	// OrderStatusInit 初始化,待付款
	OrderStatusInit = 1 + iota
	// OrderStatusDelivered 已支付，待发货, 可退款
	OrderStatusDelivered
	// OrderStatusShipped 已发货
	OrderStatusShipped
	// OrderStatusReceived 已收货，可退款
	OrderStatusReceived
	// OrderStatusFinish 已完成，可评价
	OrderStatusFinish
	// OrderStatusPendingRefund 待退款
	OrderStatusPendingRefund
	// OrderStatusRefund 已退款
	OrderStatusRefund
)

const (
	// OrderSourceTypeApp 订单来源app
	OrderSourceTypeApp = 1
)

const (
	//TaskStatusLock 锁定
	TaskStatusLock = 1 + iota
	//TaskStatusUnlock 解锁
	TaskStatusUnlock
	//TaskStatusFinish 订单完成扣减库存
	TaskStatusFinish
)

const (
	// AppPageHome 首页
	AppPageHome = 1 + iota
	// AppPageSearch 搜索页
	AppPageSearch
)

const (
	// AppTypeSwiper 轮播图
	AppTypeSwiper = iota + 1
	// AppTypeNav 图标
	AppTypeNav
	// AppTypeThreeAdv 三图广告
	AppTypeThreeAdv
	// AppTypeOneAdv 单图广告
	AppTypeOneAdv
	// AppTypeProduct 商品列表
	AppTypeProduct
)

const (
	// CouponUseTypeAll 全场通用
	CouponUseTypeAll = iota
	// CouponUseTypeCat 指定分类
	CouponUseTypeCat
	// CouponUseTypeSpu 指定商品
	CouponUseTypeSpu
)

const (
	// CouponStatusNotReceive 未领取
	CouponStatusNotReceive = iota
	// CouponStatusInit 未使用
	CouponStatusInit
	// CouponStatusUsed 已使用
	CouponStatusUsed
	// CouponStatusExpire 已过期
	CouponStatusExpire
)

const (
	// CouponGetTypeDraw 主动领取
	CouponGetTypeDraw = 1
)

const (
	// ConfigKeyHomeCat 首页分类键
	ConfigKeyHomeCat = "app_home_cat"
	// ConfigKeyPayList 支付配置
	ConfigKeyPayList = "app_pay_list"
)

var ErrRecordNotModified = errors.New("record not modified")

// Attrs 规格属性结构
type Attrs struct {
	GroupID   int    `json:"group_id"`
	AttrID    int    `json:"attr_id"`
	AttrName  string `json:"attr_name"`
	AttrValue string `json:"attr_value"`
}

// Brand 品牌结构
type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

// WareSkuStock 商品库存对外结构
type WareSkuStock struct {
	SkuID     int `json:"sku_id"`
	Stock     int `json:"stock"`
	StockLock int `json:"stock_lock"`
}

type Coupon struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Amount   int       `json:"amount"`
	MinPoint int       `json:"min_point"`
	StartAt  time.Time `json:"start_at"`
	EndAt    time.Time `json:"end_at"`
	Note     string    `json:"note"`
	Status   int       `json:"status"`
}

// Config 对外配置结构
type Config struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// AppSetting 对外页面配置数据结构
type AppSetting struct {
	Type int8            `json:"type"`
	Data json.RawMessage `json:"data"`
}

// ConfigHomeCat 首页分类配置
type ConfigHomeCat struct {
	ID   int           `json:"id"`
	Name string        `json:"name"`
	List []*AppSetting `json:"list"`
}

// ConfigPayList 支付配置
type ConfigPayList struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (m *UmsMember) Compare(pwd string) (err error) {
	return auth.Compare(m.Password, pwd)
}

// BeforeSave 保存前回调
func (m *UmsMember) BeforeSave(tx *gorm.DB) (err error) {
	m.Password, err = auth.Encrypt(m.Password)
	return err
}

// OffsetPage 分页查询
func OffsetPage(offset, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}

func WhereRelease(db *gorm.DB) *gorm.DB {
	return db.Where("is_release=?", ReleaseYes)
}
