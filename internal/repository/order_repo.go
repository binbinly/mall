package repository

import (
	"context"
	"project-layout/internal/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IOrder interface {
	CreateOrder(ctx context.Context, tx *gorm.DB, order *model.OmsOrder) error
	GetOrderDetail(ctx context.Context, id, memberID int) (*model.OmsOrder, error)
	GetOrderByID(ctx context.Context, id int) (*model.OmsOrder, error)
	GetOrderByNo(ctx context.Context, orderNo string) (*model.OmsOrder, error)
	GetOrderList(ctx context.Context, memberID int, status, offset, limit int) (list []*model.OmsOrder, err error)
	OrderSave(ctx context.Context, tx *gorm.DB, order *model.OmsOrder) (err error)
	OrderDelete(ctx context.Context, order *model.OmsOrder) error
	CreateOrderRefund(ctx context.Context, tx *gorm.DB, refund *model.OmsOrderRefund) error
	BatchCreateOrderItem(ctx context.Context, tx *gorm.DB, items []*model.OmsOrderItem) error
	CreateOrderBill(ctx context.Context, tx *gorm.DB, bill *model.OmsOrderBill) error
}

// CreateOrder 创建订单
func (r *Repo) CreateOrder(ctx context.Context, tx *gorm.DB, order *model.OmsOrder) error {
	if err := tx.WithContext(ctx).Create(order).Error; err != nil {
		return errors.Wrapf(err, "[repo.order] create")
	}

	return nil
}

// GetOrderByID 获取订单信息
func (r *Repo) GetOrderByID(ctx context.Context, id int) (*model.OmsOrder, error) {
	order := new(model.OmsOrder)
	if err := r.db.WithContext(ctx).First(order, id).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.order] by id: %v", id)
	}

	return order, nil
}

// GetOrderDetail 获取订单详情
func (r *Repo) GetOrderDetail(ctx context.Context, id, memberID int) (*model.OmsOrder, error) {
	order := new(model.OmsOrder)
	err := r.db.WithContext(ctx).Where("id=? and member_id=?", id, memberID).Preload("Items").First(order).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.order] by id: %v", id)
	}

	return order, nil
}

// GetOrderByNo 订单号获取订单信息
func (r *Repo) GetOrderByNo(ctx context.Context, orderNo string) (*model.OmsOrder, error) {
	order := new(model.OmsOrder)
	err := r.db.WithContext(ctx).Where("order_no=?", orderNo).First(order).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.order] by no: %v", orderNo)
	}

	return order, nil
}

// GetOrderList 用户订单列表
func (r *Repo) GetOrderList(ctx context.Context, memberID int, status, offset, limit int) (list []*model.OmsOrder, err error) {
	err = r.db.WithContext(ctx).Where("member_id=?", memberID).Preload("Items").Order(model.DefaultOrder).
		Scopes(orderScopesStatus(status), model.OffsetPage(offset, limit)).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.order] find")
	}

	return
}

// OrderSave 订单保存
func (r *Repo) OrderSave(ctx context.Context, tx *gorm.DB, order *model.OmsOrder) (err error) {
	if tx == nil {
		err = r.db.WithContext(ctx).Save(order).Error
	} else {
		err = tx.WithContext(ctx).Save(order).Error
	}
	if err != nil {
		return errors.Wrapf(err, "[repo.order] save")
	}
	return nil
}

// OrderDelete 订单删除
func (r *Repo) OrderDelete(ctx context.Context, order *model.OmsOrder) error {
	if err := r.db.WithContext(ctx).Delete(order).Error; err != nil {
		return errors.Wrapf(err, "[repo.order] delete")
	}

	return nil
}

// orderScopesStatus 状态筛选
func orderScopesStatus(status int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status == 0 {
			return db
		}
		return db.Where("status=?", status)
	}
}

// CreateOrderRefund 创建订单退款记录
func (r *Repo) CreateOrderRefund(ctx context.Context, tx *gorm.DB, refund *model.OmsOrderRefund) error {
	if err := tx.WithContext(ctx).Create(&refund).Error; err != nil {
		return errors.Wrapf(err, "[repo.order] refund create")
	}

	return nil
}

// BatchCreateOrderItem 批量创建订单子项
func (r *Repo) BatchCreateOrderItem(ctx context.Context, tx *gorm.DB, items []*model.OmsOrderItem) error {
	if err := tx.WithContext(ctx).Model(&model.OmsOrderItem{}).Create(&items).Error; err != nil {
		return errors.Wrapf(err, "[repo.order] item batch create")
	}

	return nil
}

// CreateOrderBill 创建发票
func (r *Repo) CreateOrderBill(ctx context.Context, tx *gorm.DB, bill *model.OmsOrderBill) error {
	if err := tx.WithContext(ctx).Create(bill).Error; err != nil {
		return errors.Wrapf(err, "[repo.order] bill create")
	}

	return nil
}
