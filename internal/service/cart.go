package service

import (
	"context"
	"project-layout/internal/model"
	"sort"

	"github.com/pkg/errors"
)

// AddCart 添加购物车
func (s *Service) AddCart(ctx context.Context, uid int, num int, sku *model.Sku) error {
	// 获取当前sku_id 的购物车信息
	cart, err := s.repo.GetCartByID(ctx, uid, sku.ID)
	if err != nil {
		return errors.Wrapf(err, "[service.cart] get by uid: %v id: %v", uid, sku.ID)
	}

	if cart.SkuID != 0 { //商品已存在,加数量即可
		cart.Num += num
	} else {
		setCart(sku, num, cart)
	}

	if err = s.repo.AddCart(ctx, uid, cart); err != nil {
		return errors.Wrapf(err, "[service.cart] add cart")
	}

	return nil
}

// EditCart 更新购物车
func (s *Service) EditCart(ctx context.Context, uid, oldSkuID int, sku *model.Sku, num int) error {
	cart, err := s.repo.GetCartByID(ctx, uid, sku.ID)
	if err != nil {
		return errors.Wrapf(err, "[service.cart] get by uid: %v, id: %v", uid, sku.ID)
	}
	if cart.SkuID == sku.ID && cart.Num == num { //购物车未更新，直接返回
		return nil
	}
	if cart.SkuID == 0 { // 商品不存在，直接添加
		setCart(sku, num, cart)
	} else {
		cart.Num = num
	}

	if err = s.repo.EditCart(ctx, uid, oldSkuID, cart); err != nil {
		return errors.Wrapf(err, "[service.cart] edit sku")
	}

	return nil
}

// EditCartNum 修改商品数量
func (s *Service) EditCartNum(ctx context.Context, uid int, sku *model.Sku, num int) error {
	cart, err := s.repo.GetCartByID(ctx, uid, sku.ID)
	if err != nil {
		return errors.Wrapf(err, "[service.cart] get by uid: %v, id: %v", uid, sku.ID)
	}
	if cart.SkuID == 0 { // 商品不存在，直接添加
		setCart(sku, num, cart)
	}
	if cart.Num == num { // 数量未修改，直接返回
		return nil
	} else {
		cart.Num = num
	}

	if err = s.repo.EditCart(ctx, uid, 0, cart); err != nil {
		return errors.Wrapf(err, "[service.cart] edit num")
	}

	return nil
}

// DelCart 删除购物车
func (s *Service) DelCart(ctx context.Context, uid int, skuIds []int) error {
	return s.repo.DelCart(ctx, uid, skuIds)
}

// ClearCart 清空购物车
func (s *Service) ClearCart(ctx context.Context, uid int) error {
	return s.repo.EmptyCart(ctx, uid)
}

// CartList 购物车
func (s *Service) CartList(ctx context.Context, uid int) ([]*model.CartModel, error) {
	list, err := s.repo.CartList(ctx, uid)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.cart] list uid: %v", uid)
	}
	sort.Sort(model.CartSort(list))

	return list, nil
}

// BatchGetCarts 批量获取购物车信息
func (s *Service) BatchGetCarts(ctx context.Context, uid int, skuIds []int) ([]*model.CartModel, error) {
	return s.repo.GetCartsByIds(ctx, uid, skuIds)
}

func setCart(sku *model.Sku, num int, cart *model.CartModel) {
	cart.SkuID = sku.ID
	cart.Num = num
	cart.Cover = sku.Cover
	cart.Price = int(sku.Price)
	cart.Title = sku.Title
	cart.SkuAttr = sku.AttrValue
}
