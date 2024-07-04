package service

import (
	"context"
	"fmt"
	"project-layout/internal/model"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// GetCouponList 优惠券列表
func (s *Service) GetCouponList(ctx context.Context, memberID, skuID int) ([]*model.SmsCoupon, []int, error) {
	list, err := s.repo.GetCouponList(ctx, 0, 0)
	if err != nil {
		return nil, nil, errors.Wrap(err, "[service.coupon] get list")
	}
	if len(list) == 0 {
		return []*model.SmsCoupon{}, nil, nil
	}
	//已领取的优惠券
	ids, err := s.repo.GetCouponUserDrawIds(ctx, memberID)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[service.coupon] get draw ids uid: %v", memberID)
	}
	return list, ids, nil
}

// GetMyCouponList 我的优惠券
func (s *Service) GetMyCouponList(ctx context.Context, memberID int) ([]*model.Coupon, error) {
	list, err := s.repo.GetCouponMemberList(ctx, memberID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.coupon] get my list")
	}
	return list, nil
}

// CouponDraw 领取优惠券
func (s *Service) CouponDraw(ctx context.Context, memberID, id int) error {
	coupon, err := s.repo.GetCouponByID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.coupon] find id: %v", id)
	}
	now := time.Now().Unix()
	if coupon == nil || coupon.ID == 0 || coupon.StartAt > now || coupon.EndAt < now {
		return ErrCouponNotFound
	}
	exist, err := s.repo.CheckReceived(ctx, memberID, id)
	if err != nil {
		return errors.Wrapf(err, "[service.coupon] check uid: %v id: %v", memberID, id)
	}
	if exist {
		return ErrCouponReceived
	}
	num, err := s.rdb.Decr(ctx, fmt.Sprintf("coupon_num:%d", id)).Result()
	if errors.Is(err, redis.Nil) {
		return ErrCouponFinished
	} else if err != nil {
		return errors.Wrapf(err, "[service.coupon] redis decr")
	}
	if num < 0 { //已领取完
		return ErrCouponFinished
	}
	couponMember := &model.SmsCouponMember{
		MID:      dbs.MID{MemberID: memberID},
		CouponID: id,
		GetType:  model.CouponGetTypeDraw,
		Status:   model.CouponStatusInit,
	}
	err = l.repo.CouponUserSave(ctx, couponMember)
	if err != nil {
		return errors.Wrapf(err, "[service.coupon] save")
	}
	return nil
}

// CouponUsed 使用优惠券
func (s *Service) CouponUsed(ctx context.Context, memberID, id, orderID int) error {
	mCoupon, err := l.repo.GetCouponMemberByID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.coupon] get coupon user by id: %v", id)
	}
	if mCoupon.ID == 0 || mCoupon.MemberID != memberID || mCoupon.Status != model.CouponStatusInit {
		return errno.ErrCouponNotFound
	}
	err = l.repo.SetCouponMemberUsed(ctx, id, memberID, orderID)
	if errors.Is(err, dbs.ErrRecordNotModified) {
		return errno.ErrCouponNotFound
	} else if err != nil {
		return errors.Wrapf(err, "[service.coupon] set used id: %v uid: %v, oid: %v", id, memberID, orderID)
	}
	return nil
}

// GetCouponInfo 获取优惠券详情
func (s *Service) GetCouponInfo(ctx context.Context, memberID, id int) (*model.SmsCouponMember, *model.SmsCoupon, error) {
	mCoupon, err := s.repo.GetCouponMemberByID(ctx, id)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[service.coupon] get coupon user by id: %v", id)
	}
	if mCoupon.ID == 0 || mCoupon.MemberID != memberID || mCoupon.Status != model.CouponStatusInit {
		return nil, nil, ErrCouponNotFound
	}
	coupon, err := s.repo.GetCouponByID(ctx, mCoupon.CouponID)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[service.coupon] get coupon by id: %v", mCoupon.CouponID)
	}
	now := time.Now()
	if coupon == nil || coupon.ID == 0 || now.Before(coupon.EndAt) || now.After(coupon.StartAt) {
		return nil, nil, ErrCouponNotFound
	}

	return mCoupon, coupon, nil
}
