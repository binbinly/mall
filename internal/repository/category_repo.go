package repository

import (
	"context"
	"github.com/binbinly/pkg/logger"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"project-layout/internal/model"
)

type ICategory interface {
	GetCategoryNameByID(ctx context.Context, id int) (name string, err error)
	GetCategoryNamesByIds(ctx context.Context, ids []int) (names map[int]string, err error)
	GetCategoryChild(ctx context.Context, id int) ([]int, error)
	CategoryAll(ctx context.Context) (list []*model.PmsCategory, err error)
}

// GetCategoryNameByID 获取分类名
func (r *Repo) GetCategoryNameByID(ctx context.Context, id int) (name string, err error) {
	var cat *model.PmsCategory
	if err = r.QueryCache(ctx, buildCategoryCacheKey(id), &cat, 0, func() (any, error) {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).First(&cat, id).Error; err != nil {
			return "", err
		}

		return cat, nil
	}); err != nil {
		return "", errors.Wrapf(err, "[repo.category] query cache")
	}

	return cat.Name, nil
}

// GetCategoryNamesByIds 批量获取分类名
func (r *Repo) GetCategoryNamesByIds(ctx context.Context, ids []int) (names map[int]string, err error) {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, buildCategoryCacheKey(id))
	}
	// 从cache批量获取
	cacheMap := make(map[string]*model.PmsCategory)
	if err = r.Cache.MultiGet(ctx, keys, cacheMap, func() any {
		return &model.PmsCategory{}
	}); err != nil {
		return nil, errors.Wrapf(err, "[r.category] multi get catName cache data err")
	}

	names = make(map[int]string, len(ids))
	// 查询未命中
	for _, id := range ids {
		cat, ok := cacheMap[buildCategoryCacheKey(id)]
		var name string
		if !ok {
			name, err = r.GetCategoryNameByID(ctx, id)
			if err != nil {
				logger.Warnf("[repo.category] get catName model err: %v", err)
				continue
			}
		} else {
			name = cat.Name
		}
		if name == "" {
			continue
		}
		names[id] = name
	}
	return names, nil
}

// GetCategoryChild 获取分类下所有子分类
func (r *Repo) GetCategoryChild(ctx context.Context, id int) ([]int, error) {
	list, err := r.CategoryAll(ctx)
	if err != nil {
		return nil, err
	}
	return makeChild(id, list), nil
}

// CategoryAll 获取全部分裂
func (r *Repo) CategoryAll(ctx context.Context) (list []*model.PmsCategory, err error) {
	if err = r.QueryCache(ctx, "category_all", &list, 0, func() (any, error) {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.PmsCategory{}).Order(model.DefaultOrderSort).Find(&list).Error; err != nil {
			return nil, err
		}
		if len(list) == 0 {
			return nil, gorm.ErrEmptySlice
		}
		return list, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.category] query cache")
	}
	return
}

// makeChild 递归所有子分类
func makeChild(id int, list []*model.PmsCategory) (ids []int) {
	for _, categoryModel := range list {
		if id == 0 { //所有子分类
			ids = append(ids, categoryModel.ID)
			continue
		}
		if categoryModel.ParentID == id { //递归查找
			ids = append(ids, makeChild(categoryModel.ID, list)...)
		}
	}
	if id > 0 {
		ids = append(ids, id)
	}
	return ids
}
