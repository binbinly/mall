package service

import (
	"context"
	"project-layout/internal/model"

	"github.com/pkg/errors"
)

// GetSkuStock 获取sku库存数量
func (s *Service) GetSkuStock(ctx context.Context, skuID int) (int, error) {
	stock, err := s.repo.GetWareSkuStock(ctx, skuID)
	if err != nil {
		return 0, errors.Wrapf(err, "[service.ware] sku stock sku_id: %v", skuID)
	}

	return stock.Stock - stock.StockLock, nil
}

// GetSkusStock 批量获取sku的库存数量
func (s *Service) GetSkusStock(ctx context.Context, skuIds []int) (map[int]int, error) {
	stocks, err := s.repo.BatchGetWareSkusStock(ctx, skuIds)
	if err != nil {
		return nil, err
	}

	res := make(map[int]int, len(stocks))
	for _, stock := range stocks {
		res[stock.SkuID] = stock.Stock - stock.StockLock
	}

	return res, nil
}

// SKuStockLock sku库存锁定
func (s *Service) SKuStockLock(ctx context.Context, orderID int, orderNo, consignee, phone, address, note string,
	sku map[int]int) error {

	skuIds := make([]int, 0, len(sku))
	for id := range sku {
		skuIds = append(skuIds, id)
	}

	skus, err := s.repo.BatchGetWareSkus(ctx, skuIds)
	if err != nil {
		return errors.Wrapf(err, "[service.ware] sku batch skus by ids: %v", skuIds)
	}
	if len(skus) == 0 {
		return ErrWareInventoryShortage
	}
	//库存工作单
	task := &model.WmsWareTask{
		OrderID:   orderID,
		OrderNo:   orderNo,
		Consignee: consignee,
		Phone:     phone,
		Address:   address,
		Note:      note,
		Status:    model.TaskStatusLock,
	}

	// 开启事务
	tx := s.repo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	//保存库存工作单
	if err = s.repo.CreateWareTask(ctx, tx, task); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.ware] sku create ware task")
	}

	var taskDetails []*model.WmsWareTaskDetail
	for _, m := range skus {
		if m.Stock-m.StockLock < sku[m.SkuID] { //库存不足
			tx.Rollback()
			return ErrWareInventoryShortage
		}
		m.StockLock += sku[m.SkuID]

		if err = s.repo.WareSkuSave(ctx, tx, m); err != nil {
			tx.Rollback()
			return errors.Wrapf(err, "[service.ware] sku save")
		}
		taskDetails = append(taskDetails, &model.WmsWareTaskDetail{
			TaskID:  task.ID,
			WareID:  m.WareID,
			SkuID:   m.SkuID,
			SkuName: m.SkuName,
			SkuNum:  sku[m.SkuID],
		})
	}

	//批量保存库存工作单详情
	if err = s.repo.BatchCreateWareTaskDetail(ctx, tx, taskDetails); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.ware] sku batch create task detail")
	}

	//提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.ware] sku tx commit lock stock")
	}

	return nil
}

// SkuStockUnlock sku库存解锁 finish 订单是否完成 订单完成直接扣减库存
// 1，下订单成功后，订单自动过期或者用户手动取消，需要解锁库存
// 2，下订单成功，库存锁定成功，接下来业务调用失败，导致订单回滚，之前锁定成功的库存需要解锁
func (s *Service) SkuStockUnlock(ctx context.Context, orderID int, finish bool) error {
	task, err := s.repo.GetTaskByOrderID(ctx, orderID)
	if err != nil {
		return errors.Wrapf(err, "[service.ware] sku task by order_id:%v", orderID)
	}
	//库存工作单不存在或者工作单状态非解锁，无需解锁跳过
	if task.ID == 0 || task.Status != model.TaskStatusLock {
		return nil
	}

	//库存工作单详情
	details, err := s.repo.GetTaskDetailByID(ctx, task.ID)
	if err != nil {
		return errors.Wrapf(err, "[service.ware] sku task detail by task_id: %v", task.ID)
	}
	skuIds := make([]int, 0, len(details))
	skuStock := make(map[int]int, len(details))
	for _, detail := range details {
		skuIds = append(skuIds, detail.SkuID)
		skuStock[detail.SkuID] = detail.SkuNum
	}
	skus, err := s.repo.BatchGetWareSkus(ctx, skuIds)
	if err != nil {
		return errors.Wrapf(err, "[service.ware] sku batch skus by ids: %v", skuIds)
	}

	// 开启事务
	tx := s.repo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//更新库存工作单状态
	status := model.TaskStatusUnlock
	if finish {
		status = model.TaskStatusFinish
	}
	if err = s.repo.UpdateWareTaskStatus(ctx, tx, orderID, status); err != nil {
		tx.Rollback()
		if errors.Is(err, model.ErrRecordNotModified) {
			return nil
		}
		return errors.Wrapf(err, "[service.ware] sku update task")
	}

	for _, m := range skus {
		if m.StockLock < skuStock[m.SkuID] { //库存不足
			continue
		}
		m.StockLock -= skuStock[m.SkuID]
		if finish { //订单完成，库存扣减
			m.Stock -= skuStock[m.SkuID]
		}
		if err = s.repo.WareSkuSave(ctx, tx, m); err != nil {
			tx.Rollback()
			return errors.Wrapf(err, "[service.ware] sku save")
		}
	}

	//提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.ware] sku tx commit unlock stock")
	}

	return nil
}
