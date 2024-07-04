package repository

import (
	"context"
	"project-layout/internal/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IWareTask interface {
	GetTaskByOrderID(ctx context.Context, oid int) (task *model.WmsWareTask, err error)
	UpdateWareTaskStatus(ctx context.Context, tx *gorm.DB, od int, status int) error
	CreateWareTask(ctx context.Context, tx *gorm.DB, task *model.WmsWareTask) error
	BatchCreateWareTaskDetail(ctx context.Context, tx *gorm.DB, items []*model.WmsWareTaskDetail) error
	GetTaskDetailByID(ctx context.Context, id int) (list []*model.WmsWareTaskDetail, err error)
}

// GetTaskByOrderID 获取库存工作单
func (r *Repo) GetTaskByOrderID(ctx context.Context, oid int) (task *model.WmsWareTask, err error) {
	task = new(model.WmsWareTask)
	if err = r.db.WithContext(ctx).Where("order_id=?", oid).First(task).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "[repo.wareTask] by order_id: %v", oid)
	}

	return task, nil
}

// UpdateWareTaskStatus 更新库存工作单状态
func (r *Repo) UpdateWareTaskStatus(ctx context.Context, tx *gorm.DB, od int, status int) error {
	result := tx.WithContext(ctx).Model(&model.WmsWareTask{}).
		Where("order_id=? and status=1", od).Update("status", status)
	if result.Error != nil {
		return errors.Wrapf(result.Error, "[repo.wareTask] update")
	}
	if result.RowsAffected == 0 { //没有记录更新
		return model.ErrRecordNotModified
	}

	return nil
}

// CreateWareTask 创建库存工作单
func (r *Repo) CreateWareTask(ctx context.Context, tx *gorm.DB, task *model.WmsWareTask) error {
	if err := tx.WithContext(ctx).Create(task).Error; err != nil {
		return errors.Wrapf(err, "[repo.wareTask] create")
	}

	return nil
}

// BatchCreateWareTaskDetail 批量创建库存工作单详情
func (r *Repo) BatchCreateWareTaskDetail(ctx context.Context, tx *gorm.DB, items []*model.WmsWareTaskDetail) error {
	if err := tx.WithContext(ctx).Model(&model.WmsWareTaskDetail{}).Create(&items).Error; err != nil {
		return errors.Wrapf(err, "[repo.wareTask] batch detail create")
	}

	return nil
}

// GetTaskDetailByID 获取库存工作单详情
func (r *Repo) GetTaskDetailByID(ctx context.Context, id int) (list []*model.WmsWareTaskDetail, err error) {
	if err = r.db.WithContext(ctx).Model(&model.WmsWareTaskDetail{}).
		Where("task_id=?", id).Find(&list).Error; err != nil {
		return nil, errors.Wrapf(err, "[repo.wareTask] detail list")
	}

	return
}
