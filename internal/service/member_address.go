package service

import (
	"context"
	"project-layout/internal/model"

	"github.com/pkg/errors"
)

// MemberAddressAdd 添加收货地址
func (s *Service) MemberAddressAdd(ctx context.Context, uid int, name, phone,
	province, city, county, detail string, areaCode int, isDefault int8) (int, error) {
	addr := &model.UmsMemberAddress{
		MemberID:  uid,
		Name:      name,
		Phone:     phone,
		Province:  province,
		City:      city,
		County:    county,
		AreaCode:  areaCode,
		Detail:    detail,
		IsDefault: isDefault,
	}
	return s.repo.MemberAddressCreate(ctx, addr)
}

// MemberAddressEdit 收货地址修改
func (s *Service) MemberAddressEdit(ctx context.Context, id, MemberID int, name, phone,
	province, city, county, detail string, areaCode int, isDefault int8) error {
	address, err := s.repo.GetMemberAddressByID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.member] address get id: %v", id)
	}
	if address == nil || address.MemberID != MemberID {
		return ErrMemberAddressNotFound
	}
	isEdit := false
	if name != "" && address.Name != name {
		address.Name = name
		isEdit = true
	}
	if phone != "" && address.Phone != phone {
		address.Phone = phone
		isEdit = true
	}
	if province != "" && address.Province != province {
		address.Province = province
		isEdit = true
	}
	if city != "" && address.City != city {
		address.City = city
		isEdit = true
	}
	if county != "" && address.County != county {
		address.County = city
		isEdit = true
	}
	if detail != "" && address.Detail != detail {
		address.Detail = detail
		isEdit = true
	}
	if areaCode != 0 && address.AreaCode != areaCode {
		address.AreaCode = areaCode
		isEdit = true
	}
	if address.IsDefault != isDefault {
		address.IsDefault = isDefault
		isEdit = true
	}
	if isEdit {
		if err = s.repo.MemberAddressSave(ctx, address); err != nil {
			return errors.Wrapf(err, "[service.member] address update id: %v, uid: %v", id, MemberID)
		}
	}

	return nil
}

// MemberAddressList 收货地址列表
func (s *Service) MemberAddressList(ctx context.Context, MemberID int) ([]*model.UmsMemberAddress, error) {
	return s.repo.GetMemberAddressList(ctx, MemberID)
}

// MemberAddressDel 删除用户收货地址
func (s *Service) MemberAddressDel(ctx context.Context, id, uid int) error {
	addr, err := s.repo.GetMemberAddressByID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.member] address check id : %v uid: %v", id, uid)
	}
	if addr == nil || addr.MemberID != uid {
		return ErrMemberAddressNotFound
	}
	return s.repo.MemberAddressDelete(ctx, addr)
}

// GetMemberAddressInfo 获取收货地址信息
func (s *Service) GetMemberAddressInfo(ctx context.Context, id, uid int) (*model.UmsMemberAddress, error) {
	addr, err := s.repo.GetMemberAddressByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.member] address check id : %v uid: %v", id, uid)
	}
	if addr == nil || addr.MemberID != uid {
		return nil, ErrMemberAddressNotFound
	}
	return addr, nil
}
