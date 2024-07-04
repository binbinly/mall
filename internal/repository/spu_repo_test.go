package repository

import (
	"context"
	"testing"
)

func TestRepo_GetSpuByID(t *testing.T) {
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
			name: "GetSpuByID",
			args: args{
				ctx: context.Background(),
				id:  102908,
			},
		},
		{
			name: "GetSpuByID_empty",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSpu, err := r.GetSpuByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSpuByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSpu == nil {
				t.Logf("%d empty", tt.args.id)
				return
			}
			t.Logf("id:%v", gotSpu.ID)
		})
	}
}

func TestRepo_GetSpuDescBySpuID(t *testing.T) {
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
			name: "GetSpuDescBySpuID",
			args: args{
				ctx:   context.Background(),
				spuID: 151001,
			},
		},
		{
			name: "GetSpuDescBySpuID_empty",
			args: args{
				ctx:   context.Background(),
				spuID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDesc, err := r.GetSpuDescBySpuID(tt.args.ctx, tt.args.spuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSpuDescBySpuID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDesc == nil {
				t.Logf("%d empty", tt.args.spuID)
				return
			}
			t.Logf("id:%v", gotDesc.SpuID)
		})
	}
}
