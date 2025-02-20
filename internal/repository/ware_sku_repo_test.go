package repository

import (
	"context"
	"testing"
)

func TestRepo_GetWareSkuStock(t *testing.T) {
	type args struct {
		ctx   context.Context
		skuID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetWareSkuStock",
			args: args{
				ctx:   context.TODO(),
				skuID: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStock, err := r.GetWareSkuStock(tt.args.ctx, tt.args.skuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWareSkuStock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("stock: %v", gotStock)
		})
	}
}

func TestRepo_BatchGetWareSkuStocks(t *testing.T) {
	list, err := r.BatchGetWareSkusStock(context.TODO(), []int{13, 15})
	if err != nil {
		t.Errorf("BatchGetWareSkuStocks() error = %v", err)
		return
	}
	t.Logf("list len:%v", len(list))
}

func TestRepo_BatchGetWareSkus(t *testing.T) {
	list, err := r.BatchGetWareSkus(context.TODO(), []int{13, 15})
	if err != nil {
		t.Errorf("BatchGetWareSkus() error = %v", err)
		return
	}
	t.Logf("list len:%v", len(list))
}
