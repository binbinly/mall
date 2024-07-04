package repository

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"project-layout/internal/model"
)

// GetAttrGroupByCatID 获取当前分类下的所有属性分组
func (r *Repo) GetAttrGroupByCatID(ctx context.Context, cid int) (list []*model.PmsAttrGroup, err error) {
	doKey := fmt.Sprintf("pms:attr_group:%d", cid)
	if err = r.QueryCache(ctx, doKey, &list, 0, func() (any, error) {
		// 从数据库中获取
		list, err = r.PmsAttrGroup.WithContext(ctx).Where(r.PmsAttrGroup.CatID.Eq(cid)).Find()
		if err != nil {
			return nil, err
		}
		if len(list) == 0 {
			return nil, gorm.ErrEmptySlice
		}
		return list, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.attr] query group cache")
	}
	return
}

// GetAttrsBySpuID 批量获取spu下所有属性信息
func (r *Repo) GetAttrsBySpuID(ctx context.Context, spuID int) (list []*model.Attrs, err error) {
	doKey := fmt.Sprintf("pms:attrs:%v", spuID)
	if err = r.QueryCache(ctx, doKey, &list, 0, func() (any, error) {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.PmsAttrValue{}).Select("pms_attr_value.*, g.group_id").
			Joins("left join pms_attr_rel_group as g on pms_attr_value.attr_id=g.attr_id").Where("spu_id=?", spuID).
			Scan(&list).Error; err != nil {
			return nil, err
		}
		if len(list) == 0 {
			return nil, gorm.ErrEmptySlice
		}
		return list, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.attr] query value cache")
	}
	return
}
