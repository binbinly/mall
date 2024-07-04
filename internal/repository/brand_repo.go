package repository

import (
	"context"
	"github.com/binbinly/pkg/logger"
	"project-layout/internal/model"

	"github.com/pkg/errors"
)

// GetBrandByID 获取品牌
func (r *Repo) GetBrandByID(ctx context.Context, id int) (brand *model.Brand, err error) {
	if err = r.QueryCache(ctx, buildBrandCacheKey(id), &brand, 0, func() (any, error) {
		// 从数据库中获取
		brand = new(model.Brand)
		if err = r.db.WithContext(ctx).Model(&model.PmsBrand{}).Where("id=?", id).Scan(&brand).Error; err != nil {
			return nil, err
		}

		return brand, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.brand] query cache")
	}

	return
}

// GetBrandsByIds 批量获取品牌列表
func (r *Repo) GetBrandsByIds(ctx context.Context, ids []int) (brands map[int]*model.Brand, err error) {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, buildBrandCacheKey(id))
	}
	// 从cache批量获取
	cacheMap := make(map[string]*model.Brand)
	if err = r.Cache.MultiGet(ctx, keys, cacheMap, func() any {
		return &model.Brand{}
	}); err != nil {
		return nil, errors.Wrapf(err, "[r.brand] multi get brand cache data err")
	}

	brands = make(map[int]*model.Brand, len(ids))
	// 查询未命中
	for _, id := range ids {
		brand, ok := cacheMap[buildBrandCacheKey(id)]
		if !ok {
			brand, err = r.GetBrandByID(ctx, id)
			if err != nil {
				logger.Warnf("[repo.brand] get brand model err: %v", err)
				continue
			}
			if brand.ID == 0 {
				continue
			}
		}
		brands[id] = brand
	}
	return brands, nil
}
