package repository

import (
	"context"
	"github.com/binbinly/gin-pkg/app"
	"github.com/binbinly/gin-pkg/config"
	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/repo"
	"github.com/spf13/viper"
	"project-layout/internal/dal"
	"project-layout/internal/model"
	"reflect"
	"testing"
)

var r *Repo

func TestMain(m *testing.M) {
	c := config.New(config.WithDir("../../configs"))
	if err := c.Load("app", app.Conf, func(v *viper.Viper) {
		app.SetDefaultConf(v)
	}); err != nil {
		panic(err)
	}
	rdb := app.InitRedis()
	r = New(app.InitDB(), cache.NewRedisCache(rdb), rdb)
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestRepo_GetMemberByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name       string
		args       args
		wantMember *model.UmsMember
		wantErr    bool
	}{
		{
			name: "GetMemberByID",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantMember: &model.UmsMember{
				ID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMember, err := r.GetMemberByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMemberByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMember.ID, tt.wantMember.ID) {
				t.Errorf("GetMemberByID() gotMember = %v, want %v", gotMember, tt.wantMember)
			}
		})
	}
}

func TestRepo_GetMemberByPhone(t *testing.T) {
	type args struct {
		ctx   context.Context
		phone int64
	}
	tests := []struct {
		name    string
		args    args
		want    *model.UmsMember
		wantErr bool
	}{
		{
			name: "GetMemberByPhone",
			args: args{
				ctx:   context.Background(),
				phone: 12345678901,
			},
			want: &model.UmsMember{
				ID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetMemberByPhone(tt.args.ctx, tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMemberByPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("GetMemberByPhone() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_MemberCreate(t *testing.T) {
	type args struct {
		ctx    context.Context
		member *model.UmsMember
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ctx: context.Background(),
				member: &model.UmsMember{
					ID:       1,
					Username: "test",
					Password: "test",
					Nickname: "test",
					Phone:    12345678901,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				Repo:  r.Repo,
				Query: r.Query,
			}
			if err := r.MemberCreate(tt.args.ctx, tt.args.member); (err != nil) != tt.wantErr {
				t.Errorf("MemberCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_MemberDelete(t *testing.T) {
	type fields struct {
		Repo  *repo.Repo
		Query *dal.Query
	}
	type args struct {
		ctx    context.Context
		member *model.UmsMember
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "memberDelete",
			args: args{
				ctx:    context.Background(),
				member: &model.UmsMember{ID: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.MemberDelete(tt.args.ctx, tt.args.member); (err != nil) != tt.wantErr {
				t.Errorf("MemberDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
