package repository

import (
	"context"

	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/repo"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"project-layout/internal/dal"
	"project-layout/internal/model"
)

var _ IRepo = (*Repo)(nil)

type IRepo interface {
	IMember
	IMemberAddress
	IOrder
	ISpu
	ISku
	ICart
	ICategory
	IWareSku
	IWareTask
	IConfig
	ICoupon

	GetAttrGroupByCatID(ctx context.Context, cid int) (list []*model.PmsAttrGroup, err error)
	GetAttrsBySpuID(ctx context.Context, spuID int) (list []*model.Attrs, err error)
	GetBrandByID(ctx context.Context, id int) (brand *model.Brand, err error)
	GetBrandsByIds(ctx context.Context, ids []int) (brands map[int]*model.Brand, err error)

	GetDB() *gorm.DB
}

// Repo dbs struct
type Repo struct {
	*repo.Repo
	*dal.Query
	db  *gorm.DB
	rdb *redis.Client
}

// New new a Dao and return
func New(db *gorm.DB, c cache.Cache, rdb *redis.Client) *Repo {
	return &Repo{
		repo.New(c),
		dal.Use(db),
		db,
		rdb}
}

func (r *Repo) GetDB() *gorm.DB {
	return r.db
}
