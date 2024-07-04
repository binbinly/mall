package repository

import (
	"context"
	"encoding/json"
	"github.com/binbinly/pkg/logger"
	"project-layout/internal/model"
	"strconv"
	"time"

	"github.com/binbinly/pkg/storage/redis"
	"github.com/pkg/errors"
)

type ICart interface {
	AddCart(ctx context.Context, uid int, cart *model.CartModel) error
	EditCart(ctx context.Context, uid int, oldID int, cart *model.CartModel) error
	GetCartByID(ctx context.Context, uid int, skuID int) (*model.CartModel, error)
	GetCartsByIds(ctx context.Context, uid int, ids []int) ([]*model.CartModel, error)
	DelCart(ctx context.Context, uid int, ids []int) error
	EmptyCart(ctx context.Context, uid int) error
	CartList(ctx context.Context, uid int) ([]*model.CartModel, error)
}

// AddCart 添加购物车
func (r *Repo) AddCart(ctx context.Context, uid int, cart *model.CartModel) error {
	cart.UTime = time.Now().Unix()
	data, err := json.Marshal(cart)
	if err != nil {
		return errors.Wrap(err, "[repo.cart] json marshal")
	}
	return r.rdb.HSet(ctx, BuildCartKey(uid), parseID(cart.SkuID), data).Err()
}

// EditCart 更新购物车
func (r *Repo) EditCart(ctx context.Context, uid int, oldID int, cart *model.CartModel) error {
	cart.UTime = time.Now().Unix()
	data, err := json.Marshal(cart)
	if err != nil {
		return errors.Wrap(err, "[repo.cart] json marshal")
	}
	if oldID == 0 { // 未修改商品
		err = r.rdb.HSet(ctx, BuildCartKey(uid), cart.SkuID, data).Err()
	} else {
		pipe := r.rdb.Pipeline()
		pipe.HDel(ctx, BuildCartKey(uid), parseID(oldID))
		pipe.HSet(ctx, BuildCartKey(uid), cart.SkuID, data)
		_, err = pipe.Exec(ctx)
	}
	if err != nil {
		return errors.Wrap(err, "[repo.cart] pipe exec")
	}
	return nil
}

// GetCartByID 获取购物车数据
func (r *Repo) GetCartByID(ctx context.Context, uid int, skuID int) (*model.CartModel, error) {
	data, err := r.rdb.HGet(ctx, BuildCartKey(uid), parseID(skuID)).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, errors.Wrapf(err, "[repo.cart] hget uid:%d skuID:%d", uid, skuID)
	}
	cart := new(model.CartModel)
	if errors.Is(err, redis.Nil) {
		return cart, nil
	}
	err = json.Unmarshal([]byte(data), cart)
	if err != nil {
		return nil, errors.Wrap(err, "[repo.cart] json unmarshal")
	}
	return cart, nil
}

// GetCartsByIds 批量获取购物车数据
func (r *Repo) GetCartsByIds(ctx context.Context, uid int, ids []int) ([]*model.CartModel, error) {
	data, err := r.rdb.HMGet(ctx, BuildCartKey(uid), parseIds(ids)...).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, errors.Wrapf(err, "[repo.cart] hmget db %d", uid)
	}
	var carts []*model.CartModel
	for _, datum := range data {
		if datum == nil {
			continue
		}
		cart := &model.CartModel{}
		err = json.Unmarshal([]byte(datum.(string)), cart)
		if err != nil {
			logger.Warnf("[repo.cart] json.unmarshal err: %v", err)
			continue
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

// DelCart 移除购物车
func (r *Repo) DelCart(ctx context.Context, uid int, ids []int) error {
	return r.rdb.HDel(ctx, BuildCartKey(uid), parseIds(ids)...).Err()
}

// EmptyCart 清空购物车
func (r *Repo) EmptyCart(ctx context.Context, uid int) error {
	return r.rdb.Del(ctx, BuildCartKey(uid)).Err()
}

// CartList 我的购物车
func (r *Repo) CartList(ctx context.Context, uid int) ([]*model.CartModel, error) {
	data, err := r.rdb.HGetAll(ctx, BuildCartKey(uid)).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, errors.Wrapf(err, "[repo.cart] json marshal")
	}
	carts := make([]*model.CartModel, 0, len(data))
	for _, v := range data {
		cart := &model.CartModel{}
		err = json.Unmarshal([]byte(v), cart)
		if err != nil {
			logger.Warnf("[repo.cart] list json unmarshal err:%v", err)
			continue
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

func parseID(skuID int) string {
	return strconv.Itoa(skuID)
}

func parseIds(skuIds []int) (res []string) {
	for _, id := range skuIds {
		res = append(res, parseID(id))
	}
	return
}
