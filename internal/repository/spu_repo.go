package repository

import (
	"context"
	"fmt"
	"project-layout/internal/model"

	"github.com/pkg/errors"
)

type ISpu interface {
	GetSpuByID(ctx context.Context, id int) (spu *model.PmsSpu, err error)
	GetSpuDescBySpuID(ctx context.Context, id int) (desc *model.PmsSpuDesc, err error)
	CreateSpuComment(ctx context.Context, comments []*model.PmsSpuComment) error
}

// GetSpuByID 获取spu信息
func (r *Repo) GetSpuByID(ctx context.Context, id int) (spu *model.PmsSpu, err error) {
	doKey := fmt.Sprintf("mall:spu:%d", id)
	if err = r.QueryCache(ctx, doKey, &spu, 0, func() (any, error) {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).First(&spu, id).Error; err != nil {
			return nil, err
		}
		return spu, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.spu] query cache")
	}

	return
}

// GetSpuDescBySpuID 获取spu介绍
func (r *Repo) GetSpuDescBySpuID(ctx context.Context, id int) (desc *model.PmsSpuDesc, err error) {
	doKey := fmt.Sprintf("mall:spu_desc:%d", id)
	if err = r.QueryCache(ctx, doKey, &desc, 0, func() (any, error) {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Where("spu_id=?", id).First(&desc).Error; err != nil {
			return nil, err
		}
		return desc, nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.spu] query desc cache")
	}

	return
}

// CreateSpuComment 创建商品评价
func (r *Repo) CreateSpuComment(ctx context.Context, comments []*model.PmsSpuComment) error {
	if err := r.db.WithContext(ctx).Create(&comments).Error; err != nil {
		return errors.Wrapf(err, "[repo.spu] batch comment create")
	}
	return nil
}
