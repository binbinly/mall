package service

import (
	"context"
	"project-layout/internal/app"
	"project-layout/internal/model"
	"strings"
	"time"

	"github.com/binbinly/pkg/logger"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

// OrderSubmit 购物车提交订单
func (s *Service) OrderSubmit(ctx context.Context, uid, addressID, couponID int, skuIds []int, note string) error {
	// 前往购物车服务获取需要购买的商品信息
	carts, err := s.BatchGetCarts(ctx, uid, skuIds)
	if err != nil {
		return err
	}
	if len(carts) == 0 {
		return ErrOrderSkuEmpty
	}

	// 商品总金额
	var total int
	items := make([]*model.OmsOrderItem, 0, len(carts))
	for _, c := range carts {
		items = append(items, buildOrderItem(c.SkuID, c.Price, c.Num, c.Title, c.Cover, c.SkuAttr))
		total += c.Price * c.Num
	}

	order, err := s.buildOrder(ctx, addressID, uid, couponID, total, note, "")
	if err != nil {
		return err
	}

	if err = s.orderSubmit(ctx, order, items); err != nil {
		return err
	}

	// 购物车服务删除已购买的sku，不抛出错误，只记录日志
	if err = s.DelCart(ctx, uid, skuIds); err != nil {
		logger.Warnf("[service.order] del cart by ids: %v err: %v", skuIds, err)
	}

	return nil
}

// SubmitSkuOrder 不通过购物车，直接在商品页面提交sku订单
func (s *Service) SubmitSkuOrder(ctx context.Context, uid, skuID, addressID, couponID int, num int, note string) error {
	// 产品服务获取sku信息
	sku, err := s.GetSkuByID(ctx, skuID)
	if err != nil {
		return err
	}

	total := sku.Price * num
	items := []*model.OmsOrderItem{
		buildOrderItem(sku.Sku.Id, int(sku.Sku.Price), num, sku.Sku.Title, sku.Sku.Cover, sku.Sku.AttrValue),
	}

	order, err := l.buildOrder(ctx, addressID, uid, couponID, total, note, "")
	if err != nil {
		return err
	}

	if err = l.orderSubmit(ctx, order, items); err != nil {
		return err
	}

	return nil
}

// SubmitKillOrder 秒杀订单提交
func (s *Service) SubmitKillOrder(ctx context.Context, memberID, skuID, addressID int, price, num int, orderNo string) error {
	// 产品服务获取sku信息
	sku, err := l.productService.GetSkuByID(ctx, &cpb.SkuIDReq{Id: skuID})
	if err != nil {
		return err
	}

	total := price * num
	items := []*model.OrderItemModel{
		buildOrderItem(sku.Sku.Id, price, num, sku.Sku.Title, sku.Sku.Cover, sku.Sku.AttrValue),
	}

	order, err := l.buildOrder(ctx, addressID, memberID, 0, total, "", orderNo)
	if err != nil {
		return err
	}

	if err = l.orderSubmit(ctx, order, items); err != nil {
		return err
	}

	return nil
}

// OrderDetail 订单详情
func (s *Service) OrderDetail(ctx context.Context, mid, id int) (*model.OmsOrder, error) {
	order, err := s.repo.GetOrderDetail(ctx, id, mid)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.order] find by id: %v", id)
	}
	if order.ID == 0 {
		return nil, ErrOrderNotFound
	}

	return order, nil
}

// OrderCancel 取消订单
func (s *Service) OrderCancel(ctx context.Context, mid, id int) error {
	order, err := s.orderInfo(ctx, id)
	if err != nil {
		return err
	}
	if order.MemberID != mid || order.Status != model.OrderStatusInit {
		return ErrOrderNotFound
	}

	if err = s.repo.OrderDelete(ctx, order); err != nil {
		return errors.Wrapf(err, "[service.order] delete")
	}

	// TODO 订单取消，发送订单消息，其他服务订阅处理

	return nil
}

// MyOrderList 订单列表
func (s *Service) MyOrderList(ctx context.Context, mid int, status int, page int) ([]*model.OrderModel, error) {
	return s.repo.GetOrderList(ctx, mid, status, app.GetPageOffset(page), app.DefaultLimit)
}

// OrderPayNotify 支付成功回调处理
func (s *Service) OrderPayNotify(ctx context.Context, mid int, amount int, pType int, orderNo, tradeNo, transHash string) error {
	order, err := s.repo.GetOrderByNo(ctx, orderNo)
	if err != nil {
		return errors.Wrapf(err, "[service.order] get by no: %v", orderNo)
	}
	if order.MemberID != mid || order.Status != model.OrderStatusInit {
		return ErrOrderNotFound
	}

	// 营销服务获取支付配置
	pays, err := l.marketService.GetPayConfig(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	// 获取当前支付的以太坊合约地址
	var address string
	for _, pay := range pays.Data {
		if pay.Id == int32(pType) {
			address = pay.Address
		}
	}
	if address == "" {
		return ErrPayActionInvalid
	}

	// 第三方服务中检查 当前订单号是否已支付
	if _, err = l.thirdService.CheckETHPay(ctx, &core.ETHPayReq{
		Id:      pType,
		Address: address,
		OrderNo: orderNo,
	}); err != nil {
		return err
	}

	order.Status = model.OrderStatusDelivered
	order.PayAt = time.Now()
	order.PayType = int8(pType)
	order.PayAmount = amount
	order.TradeNo = tradeNo
	order.TransHash = transHash
	if err = s.repo.OrderSave(ctx, nil, order); err != nil {
		return errors.Wrap(err, "[service.order] save")
	}

	return nil
}

// OrderConfirmReceipt 确认收货
func (s *Service) OrderConfirmReceipt(ctx context.Context, mid, id int) error {
	order, err := s.orderInfo(ctx, id)
	if err != nil {
		return err
	}
	//已支付，已发货，才可以确认收货
	if order.MemberID != mid || (order.Status != model.OrderStatusDelivered && order.Status != model.OrderStatusShipped) {
		return ErrOrderNotFound
	}

	order.Status = model.OrderStatusReceived
	order.ReceiveAt = time.Now()
	order.IsConfirm = 1
	if err = s.repo.OrderSave(ctx, nil, order); err != nil {
		return errors.Wrapf(err, "[service.order] save")
	}

	return nil
}

// OrderRefund 退款
func (s *Service) OrderRefund(ctx context.Context, mid, id int, content string) error {
	order, err := s.orderInfo(ctx, id)
	if err != nil {
		return err
	}
	//已支付，已收货状态下方可退款
	if order.MemberID != mid || (order.Status != model.OrderStatusReceived && order.Status != model.OrderStatusDelivered) {
		return ErrOrderNotFound
	}

	// 开启事务
	tx := s.repo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	order.Status = model.OrderStatusPendingRefund
	if err = s.repo.OrderSave(ctx, tx, order); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.order] save")
	}

	if err = s.repo.CreateOrderRefund(ctx, tx, &model.OmsOrderRefund{
		ID:      order.ID,
		Amount:  order.PayAmount,
		Content: content,
	}); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.order] create refund")
	}

	//提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.order] tx commit refund")
	}

	return nil
}

// OrderComment 评价
func (s *Service) OrderComment(ctx context.Context, mid, id int, skuIds []int, star int, content, resources string) error {
	order, err := s.orderInfo(ctx, id)
	if err != nil {
		return err
	}
	//已收货方可评价
	if order.MemberID != mid || order.Status != model.OrderStatusReceived {
		return ErrOrderNotFound
	}

	// 开启事务
	tx := s.repo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	order.Status = model.OrderStatusFinish
	order.CommentAt = time.Now()
	if err = s.repo.OrderSave(ctx, tx, order); err != nil {
		return errors.Wrapf(err, "[service.order] save")
	}

	// 产品服务 add sku comment record
	if err = s.SpuComment(ctx, skuIds, mid, id, int8(star), content, resources); err != nil {
		return err
	}

	//提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.order] tx commit refund")
	}

	// TODO 订单完成，发送订单完成任务，其他服务监听异步处理，比如库存服务监听扣减库存

	return nil
}

// OrderInfo 订单详情
func (s *Service) OrderInfo(ctx context.Context, id int) (*model.OmsOrder, error) {
	return s.orderInfo(ctx, id)
}

func (s *Service) orderInfo(ctx context.Context, id int) (*model.OmsOrder, error) {
	order, err := s.repo.GetOrderByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.order] find by id: %v", id)
	}
	if order.ID == 0 {
		return nil, ErrOrderNotFound
	}
	return order, nil
}

// orderSubmit 商品不通过购物车，直接提交订单
// 分布式事务 - 柔性事务-可靠消息 + 最终一致性方案->rabbitmq延迟队列
func (s *Service) orderSubmit(ctx context.Context, order *model.OmsOrder, items []*model.OmsOrderItem) error {
	// 开启事务
	tx := s.repo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 创建订单
	if err := s.repo.CreateOrder(ctx, tx, order); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.order] create")
	}

	// 批量创建子订单
	if err := s.repo.BatchCreateOrderItem(ctx, tx, items); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.order] create items")
	}

	if err := s.submitBefore(ctx, order, items); err != nil {
		tx.Rollback()
		return err
	}

	//提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.order] tx commit")
	}

	return nil
}

// submitBefore 订单提交之前执行其他操作，锁定库存，优惠券使用
func (s *Service) submitBefore(ctx context.Context, order *model.OmsOrder, items []*model.OmsOrderItem) error {
	// TODO 各种，优惠券，积分的使用，扣除

	skuNums := getSkuNums(order, items)
	// 仓储服务 锁定库存
	if err := s.SKuStockLock(ctx, order.ID, order.OrderNo, order.AddressName, order.AddressPhone,
		strings.Join([]string{order.AddressProvince, order.AddressCity, order.AddressCounty, order.AddressDetail}, " "),
		order.Note, skuNums); err != nil {
		return err
	}
	return nil
}

// buildOrder 生成订单
// 1，订单信息
// 2，商品spu信息（暂不处理）
// 3，商品sku信息
// 4，优惠信息
// 5，积分信息,与价格同步 1元 = 1积分，1成长值
func (s *Service) buildOrder(ctx context.Context, addressId, uid, couponID int, total int, note, orderNo string) (*model.OmsOrder, error) {
	// 生成订单号
	if orderNo == "" {
		orderNo = app.GenOrderNo()
	}

	// 会员服务获取收货地址信息
	address, err := s.GetMemberAddressInfo(ctx, addressId, uid)
	if err != nil {
		return nil, err
	}

	// 实际应该支付总金额
	amount := total
	// TODO 各种优惠券，积分，成长值逻辑

	return &model.OmsOrder{
		MemberID:        uid,
		OrderNo:         orderNo,
		CouponID:        couponID,
		TotalAmount:     total,
		CouponAmount:    0,
		Amount:          amount,
		SourceType:      model.OrderSourceTypeApp,
		Status:          model.OrderStatusInit,
		Integration:     total / 100,
		Growth:          total / 100,
		AddressName:     address.Name,
		AddressPhone:    address.Phone,
		AddressProvince: address.Province,
		AddressCity:     address.City,
		AddressCounty:   address.County,
		AddressDetail:   address.Detail,
		Note:            note,
	}, nil
}

// buildOrderItem 构建 order items
// 积分信息,与价格同步 1元 = 1积分，1成长值
func buildOrderItem(skuId int, price, num int, title, cover, skuAttrs string) *model.OmsOrderItem {
	amount := price * num
	return &model.OmsOrderItem{
		SkuID:           skuId,
		SkuName:         title,
		SkuImg:          cover,
		SkuPrice:        amount,
		SkuAttrs:        skuAttrs,
		SkuNum:          num,
		RealAmount:      amount,
		GiveIntegration: amount / 100,
		GiveGrowth:      amount / 100,
	}
}

func getSkuNums(order *model.OmsOrder, items []*model.OmsOrderItem) map[int]int {
	skuNums := make(map[int]int, len(items))
	// 子订单批量创建
	for _, item := range items {
		item.OrderID = order.ID
		item.OrderNo = order.OrderNo
		skuNums[item.SkuID] = item.SkuNum
	}

	return skuNums
}
