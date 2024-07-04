package repository

import (
	"context"
	"fmt"
	"github.com/binbinly/pkg/logger"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"project-layout/internal/model"
)

type ISku interface {
	GetSkuAttrsBySpuID(ctx context.Context, spuID int) (list []*model.PmsSkuAttr, err error)
	GetSkuByID(ctx context.Context, id int) (sku *model.PmsSku, err error)
	GetSkusByIds(ctx context.Context, ids []int) (list []*model.PmsSku, err error)
	GetSkusBySpuID(ctx context.Context, spuID int) (list []*model.PmsSku, err error)
}

// GetSkuAttrsBySpuID 批量获取spu下的所有销售属性
func (r *Repo) GetSkuAttrsBySpuID(ctx context.Context, spuID int) (list []*model.PmsSkuAttr, err error) {
	doKey := fmt.Sprintf("mall:skuAttrs:%v", spuID)
	if err = r.QueryCache(ctx, doKey, &list, 0, func() (any, error) {
		// 从数据库中获取
		// EXPLAIN select a.* from pms_sku_attr as a LEFT JOIN pms_sku as s on s.id=a.sku_id where s.spu_id=17664;
		// EXPLAIN select * from pms_sku_attr where sku_id in (select id from pms_sku where spu_id=17664);
		// 此处使用子查询
		if err = r.db.WithContext(ctx).Model(&model.PmsSkuAttr{}).
			Where("sku_id in (select id from pms_sku where spu_id=?)", spuID).
			Order(model.DefaultOrderSort).Find(&list).Error; err != nil {
			return nil, err
		}
		if len(list) == 0 {
			return nil, gorm.ErrEmptySlice
		}
		return list, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.sku] query attrs cache")
	}
	return
}

// GetSkuByID 获取sku信息
func (r *Repo) GetSkuByID(ctx context.Context, id int) (sku *model.PmsSku, err error) {
	if err = r.QueryCache(ctx, buildSkuCacheKey(id), &sku, 0, func() (any, error) {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).First(&sku, id).Error; err != nil {
			return nil, err
		}
		return sku, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.sku] query cache")
	}

	return
}

// GetSkusByIds 批量获取ksu信息
func (r *Repo) GetSkusByIds(ctx context.Context, ids []int) (list []*model.PmsSku, err error) {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, buildSkuCacheKey(id))
	}
	// 从cache批量获取
	cacheMap := make(map[string]*model.PmsSku)
	if err = r.Cache.MultiGet(ctx, keys, cacheMap, func() any {
		return &model.PmsSku{}
	}); err != nil {
		return nil, errors.Wrapf(err, "[r.sku] multi get sku cache data err")
	}

	// 查询未命中
	for _, id := range ids {
		sku, ok := cacheMap[buildSkuCacheKey(id)]
		if !ok {
			sku, err = r.GetSkuByID(ctx, id)
			if err != nil {
				logger.Warnf("[repo.sku] get sku model err: %v", err)
				continue
			}
		}
		if sku.ID == 0 {
			continue
		}
		list = append(list, sku)
	}
	return list, nil
}

// GetSkusBySpuID 获取spu下所有sku
func (r *Repo) GetSkusBySpuID(ctx context.Context, spuID int) (list []*model.PmsSku, err error) {
	doKey := fmt.Sprintf("mall:skus:%d", spuID)
	if err = r.QueryCache(ctx, doKey, &list, 0, func() (any, error) {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.PmsSku{}).Where("spu_id=?", spuID).
			Find(&list).Error; err != nil {
			return nil, err
		}
		if len(list) == 0 {
			return nil, gorm.ErrEmptySlice
		}
		return list, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.sku] query skus cache")
	}
	return
}
