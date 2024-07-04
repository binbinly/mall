package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"project-layout/internal/model"
	"time"

	"github.com/pkg/errors"
)

// ICoupon 优惠券接口
type ICoupon interface {
	GetCouponList(ctx context.Context, catID, spuID int) (list []*model.SmsCoupon, err error)
	GetCouponByID(ctx context.Context, id int) (coupon *model.SmsCoupon, err error)
	GetCouponMemberList(ctx context.Context, memberID int) (list []*model.Coupon, err error)
	GetCouponMemberByID(ctx context.Context, id int) (coupon *model.SmsCouponMember, err error)
	SetCouponMemberUsed(ctx context.Context, id, memberID, orderID int) error
	CouponUserSave(ctx context.Context, coupon *model.SmsCouponMember) error
	GetCouponUserDrawIds(ctx context.Context, memberID int) (ids []int, err error)
	CheckReceived(ctx context.Context, memberID, couponID int) (bool, error)
}

// GetCouponList 可以被领取的优惠券列表
func (r *Repo) GetCouponList(ctx context.Context, catID, spuID int) (list []*model.SmsCoupon, err error) {
	now := time.Now().Unix()
	err = r.db.WithContext(ctx).Model(&model.SmsCoupon{}).Scopes(model.WhereRelease).
		Where("enable_start_at<=?", now).Where("enable_end_at>=?", now).
		Where("use_type=?", model.CouponUseTypeAll).
		Order("amount desc,id desc").Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.coupon] by db list")
	}
	return
}

// GetCouponByID 获取优惠券详情
func (r *Repo) GetCouponByID(ctx context.Context, id int) (coupon *model.SmsCoupon, err error) {
	doKey := fmt.Sprintf("coupon:%d", id)
	if err = r.QueryCache(ctx, doKey, &coupon, 0, func() (any, error) {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).First(&coupon, id).Error; err != nil {
			return nil, err
		}
		return coupon, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.coupon] query cache")
	}

	return coupon, nil
}

// GetCouponMemberList 获取用户领取的优惠券列表
func (r *Repo) GetCouponMemberList(ctx context.Context, memberID int) (list []*model.Coupon, err error) {
	err = r.db.WithContext(ctx).Table("sms_coupon_member as m").
		Select("`name`,`amount`,min_point,start_at,end_at,`note`,m.id,m.status").
		Scopes(model.WhereRelease).Joins("left join sms_coupon as c on m.coupon_id = c.id").
		Where("member_id=?", memberID).Order("amount desc").Scan(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.coupon] member by db uid: %v", memberID)
	}
	return
}

// GetCouponMemberByID 会员优惠券详情
func (r *Repo) GetCouponMemberByID(ctx context.Context, id int) (coupon *model.SmsCouponMember, err error) {
	coupon = new(model.SmsCouponMember)
	if err = r.db.WithContext(ctx).First(coupon, id).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.coupon] member first db")
	}

	return coupon, nil
}

// SetCouponMemberUsed 设置优惠券已使用
func (r *Repo) SetCouponMemberUsed(ctx context.Context, id, memberID, orderID int) error {
	result := r.db.WithContext(ctx).Model(&model.SmsCouponMember{}).
		Where("id=? and member_id=? and status=?", id, memberID, model.CouponStatusInit).
		Updates(map[string]any{
			"status":   model.CouponStatusUsed,
			"used_at":  time.Now().Unix(),
			"order_id": orderID,
		})
	if result.Error != nil {
		return errors.Wrapf(result.Error, "[repo.coupon] member set used")
	}
	if result.RowsAffected == 0 { //没有记录更新
		return model.ErrRecordNotModified
	}

	return nil
}

// CouponUserSave 保存记录
func (r *Repo) CouponUserSave(ctx context.Context, coupon *model.SmsCouponMember) error {
	if err := r.db.WithContext(ctx).Save(coupon).Error; err != nil {
		return errors.Wrapf(err, "[repo.coupon] member save db")
	}
	return nil
}

// GetCouponUserDrawIds 已领取的优惠券id列表
func (r *Repo) GetCouponUserDrawIds(ctx context.Context, memberID int) (ids []int, err error) {
	err = r.db.WithContext(ctx).Model(&model.SmsCouponMember{}).Where("member_id=?", memberID).Pluck("coupon_id", &ids).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.coupon] member draw pluck db")
	}
	return
}

// CheckReceived 检查是否已领取过
func (r *Repo) CheckReceived(ctx context.Context, memberID, couponID int) (bool, error) {
	var c int64
	err := r.db.WithContext(ctx).Model(&model.SmsCouponMember{}).Where("member_id=? && coupon_id=?", memberID, couponID).Count(&c).Error
	if err != nil {
		return false, errors.Wrapf(err, "[repo.coupon] member count")
	}
	return c > 0, nil
}
