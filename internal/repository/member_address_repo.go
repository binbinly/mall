package repository

import (
	"context"
	"fmt"
	"project-layout/internal/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// IMemberAddress 用户收货地址接口
type IMemberAddress interface {
	MemberAddressCreate(ctx context.Context, addr *model.UmsMemberAddress) (int, error)
	MemberAddressSave(ctx context.Context, addr *model.UmsMemberAddress) error
	GetMemberAddressList(ctx context.Context, memberID int) (list []*model.UmsMemberAddress, err error)
	GetMemberAddressByID(ctx context.Context, id int) (address *model.UmsMemberAddress, err error)
	MemberAddressDelete(ctx context.Context, addr *model.UmsMemberAddress) error
}

// MemberAddressCreate 创建收货地址
func (r *Repo) MemberAddressCreate(ctx context.Context, addr *model.UmsMemberAddress) (int, error) {
	if addr.IsDefault == 1 {
		if err := r.cancelDefaultAddress(ctx, addr.MemberID); err != nil {
			return 0, err
		}
	}

	if err := r.UmsMemberAddress.WithContext(ctx).Create(addr); err != nil {
		return 0, errors.Wrapf(err, "[repo.memberAddress] create")
	}
	r.delAddressCache(ctx, addr.ID, addr.MemberID)

	return addr.ID, nil
}

// MemberAddressSave 更新用户收货地址
func (r *Repo) MemberAddressSave(ctx context.Context, addr *model.UmsMemberAddress) error {
	if addr.IsDefault == 1 {
		if err := r.cancelDefaultAddress(ctx, addr.MemberID); err != nil {
			return err
		}
	}

	if err := r.UmsMemberAddress.WithContext(ctx).Save(addr); err != nil {
		return errors.Wrapf(err, "[repo.memberAddress] save")
	}
	r.delAddressCache(ctx, addr.ID, addr.MemberID)

	return nil
}

// GetMemberAddressList 用户收货地址列表
func (r *Repo) GetMemberAddressList(ctx context.Context, memberID int) (list []*model.UmsMemberAddress, err error) {
	doKey := fmt.Sprintf("member_address_all:%d", memberID)
	if err = r.QueryCache(ctx, doKey, &list, 0, func() (any, error) {
		// 从数据库中获取
		list, err = r.UmsMemberAddress.WithContext(ctx).Where(r.UmsMemberAddress.MemberID.Eq(memberID)).
			Order(r.UmsMemberAddress.MemberID.Desc()).Find()
		if err != nil {
			return nil, err
		}
		if len(list) == 0 {
			return nil, gorm.ErrEmptySlice
		}
		return list, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.memberAddress] query list cache")
	}

	return list, nil
}

// GetMemberAddressByID 用户收货地址详情
func (r *Repo) GetMemberAddressByID(ctx context.Context, id int) (address *model.UmsMemberAddress, err error) {
	doKey := buildAddressAllCacheKey(id)
	if err = r.QueryCache(ctx, doKey, &address, 0, func() (any, error) {
		// 从数据库中获取
		data, err := r.UmsMemberAddress.WithContext(ctx).Where(r.UmsMemberAddress.ID.Eq(id)).First()
		if err != nil {
			return nil, err
		}
		return data, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.memberAddress] query cache")
	}

	return address, nil
}

// MemberAddressDelete 删除收货地址
func (r *Repo) MemberAddressDelete(ctx context.Context, addr *model.UmsMemberAddress) error {
	if _, err := r.UmsMemberAddress.WithContext(ctx).Delete(addr); err != nil {
		return errors.Wrapf(err, "[repo.memberAddress] delete")
	}
	r.delAddressCache(ctx, addr.ID, addr.MemberID)

	return nil
}

// cancelDefaultAddress 取消默认地址
func (r *Repo) cancelDefaultAddress(ctx context.Context, id int) error {
	if _, err := r.UmsMemberAddress.WithContext(ctx).Where(r.UmsMemberAddress.MemberID.Eq(id)).
		Where(r.UmsMemberAddress.IsDefault.Eq(1)).Update(r.UmsMemberAddress.IsDefault, 0); err != nil {
		return errors.Wrapf(err, "[repo.memberAddress] set other not defualt by id: %v", id)
	}
	return nil
}

// delAddressCache 删除售后地址缓存
func (r *Repo) delAddressCache(ctx context.Context, id, mid int) {
	r.DelCache(ctx, buildAddressAllCacheKey(mid))
	r.DelCache(ctx, buildAddressCacheKey(id))
}
