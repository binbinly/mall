package repository

import (
	"context"
	"testing"
)

func TestRepo_GetSkuByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetSkuByID",
			args: args{
				ctx: context.Background(),
				id:  819440,
			},
		},
		{
			name: "GetSkuByID_empty",
			args: args{
				ctx: context.Background(),
				id:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSku, err := r.GetSkuByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSkuByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSku == nil {
				t.Logf("%d empty", tt.args.id)
				return
			}
			t.Logf("id:%v", gotSku.ID)
		})
	}
}

func TestRepo_GetSkusBySpuID(t *testing.T) {
	list, err := r.GetSkusBySpuID(context.Background(), 102908)
	if err != nil {
		t.Errorf("GetSkusBySpuID() error = %v", err)
		return
	}
	t.Logf("got len: %v", len(list))
}

func TestRepo_GetSkusByIds(t *testing.T) {
	list, err := r.GetSkusByIds(context.Background(), []int{819440, 12901, 3})
	if err != nil {
		t.Errorf("GetSkusByIds() error = %v", err)
		return
	}
	t.Logf("got len: %v", len(list))
}

func TestRepo_GetSkuAttrsBySpuID(t *testing.T) {
	type args struct {
		ctx   context.Context
		spuID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetSkuAttrsBySpuID",
			args: args{
				ctx:   context.Background(),
				spuID: 102908,
			},
		},
		{
			name: "GetSkuAttrsBySpuID_empty",
			args: args{
				ctx:   context.Background(),
				spuID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GetSkuAttrsBySpuID(tt.args.ctx, tt.args.spuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSkuAttrsBySpuID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got len:%v", len(gotList))
		})
	}
}
