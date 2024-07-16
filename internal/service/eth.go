package service

import (
	"context"
)

// CheckPay 检查订单是否已支付
func (s *Service) CheckPay(ctx context.Context, id int, address, orderNo string) error {
	// TODO 合约逻辑
	return nil
}
