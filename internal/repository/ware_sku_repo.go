package repository

import (
	"context"
	"project-layout/internal/model"

	"github.com/binbinly/pkg/logger"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IWareSku interface {
	GetWareSkuStock(ctx context.Context, skuID int) (info *model.WareSkuStock, err error)
	BatchGetWareSkusStock(ctx context.Context, skuIds []int) (list []*model.WareSkuStock, err error)
	BatchGetWareSkus(ctx context.Context, skuIds []int) (list []*model.WmsWareSku, err error)
	WareSkuSave(ctx context.Context, tx *gorm.DB, ware *model.WmsWareSku) error
}

// GetWareSkuStock 获取sku库存数量
func (r *Repo) GetWareSkuStock(ctx context.Context, skuID int) (info *model.WareSkuStock, err error) {
	if err = r.QueryCache(ctx, stockCacheKey(skuID), &info, 0, func() (any, error) {
		info = new(model.WareSkuStock)
		if err = r.db.WithContext(ctx).Model(&model.WmsWareSku{}).Select("sku_id,sum(stock) as stock,sum(stock_lock) as stock_lock").
			Where("sku_id=?", skuID).Group("sku_id").Scan(&info).Error; err != nil {
			return nil, err
		}
		return info, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.wareSku] query cache")
	}

	return
}

// BatchGetWareSkusStock 批量获取sku的库存数量
func (r *Repo) BatchGetWareSkusStock(ctx context.Context, skuIds []int) (list []*model.WareSkuStock, err error) {
	keys := make([]string, 0, len(skuIds))
	for _, id := range skuIds {
		keys = append(keys, stockCacheKey(id))
	}
	// 从cache批量获取
	cacheMap := make(map[string]*model.WareSkuStock)
	if err = r.Cache.MultiGet(ctx, keys, cacheMap, func() any {
		return &model.WareSkuStock{}
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.wareSku] multi get cache data")
	}

	// 查询未命中
	for _, id := range skuIds {
		s, ok := cacheMap[stockCacheKey(id)]
		if !ok {
			s, err = r.GetWareSkuStock(ctx, id)
			if err != nil {
				logger.Warnf("[repo.wareSku] find id: %v err: %v", id, err)
				continue
			}
		}
		if s == nil || s.SkuID == 0 {
			continue
		}
		list = append(list, s)
	}

	return
}

// BatchGetWareSkus 批量获取sku库存信息
func (r *Repo) BatchGetWareSkus(ctx context.Context, skuIds []int) (list []*model.WmsWareSku, err error) {
	if err = r.db.WithContext(ctx).Model(&model.WmsWareSku{}).Where("sku_id in ?", skuIds).Find(&list).Error; err != nil {
		return nil, errors.Wrap(err, "[repo.wareSku] batch ware skus")
	}

	return list, nil
}

// WareSkuSave 更新库存信息
func (r *Repo) WareSkuSave(ctx context.Context, tx *gorm.DB, ware *model.WmsWareSku) error {
	if err := tx.WithContext(ctx).Save(ware).Error; err != nil {
		return errors.Wrap(err, "[repo.wareSku] save")
	}

	r.DelCache(ctx, stockCacheKey(ware.SkuID))

	return nil
}
