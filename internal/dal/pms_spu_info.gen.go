// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"project-layout/internal/model"
)

func newPmsSpuInfo(db *gorm.DB, opts ...gen.DOOption) pmsSpuInfo {
	_pmsSpuInfo := pmsSpuInfo{}

	_pmsSpuInfo.pmsSpuInfoDo.UseDB(db, opts...)
	_pmsSpuInfo.pmsSpuInfoDo.UseModel(&model.PmsSpuInfo{})

	tableName := _pmsSpuInfo.pmsSpuInfoDo.TableName()
	_pmsSpuInfo.ALL = field.NewAsterisk(tableName)
	_pmsSpuInfo.ID = field.NewInt(tableName, "id")
	_pmsSpuInfo.SpuID = field.NewInt(tableName, "spu_id")
	_pmsSpuInfo.Content = field.NewString(tableName, "content")

	_pmsSpuInfo.fillFieldMap()

	return _pmsSpuInfo
}

type pmsSpuInfo struct {
	pmsSpuInfoDo

	ALL     field.Asterisk
	ID      field.Int    // ID
	SpuID   field.Int    // spu_id
	Content field.String // 商品介绍

	fieldMap map[string]field.Expr
}

func (p pmsSpuInfo) Table(newTableName string) *pmsSpuInfo {
	p.pmsSpuInfoDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsSpuInfo) As(alias string) *pmsSpuInfo {
	p.pmsSpuInfoDo.DO = *(p.pmsSpuInfoDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsSpuInfo) updateTableName(table string) *pmsSpuInfo {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt(table, "id")
	p.SpuID = field.NewInt(table, "spu_id")
	p.Content = field.NewString(table, "content")

	p.fillFieldMap()

	return p
}

func (p *pmsSpuInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsSpuInfo) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 3)
	p.fieldMap["id"] = p.ID
	p.fieldMap["spu_id"] = p.SpuID
	p.fieldMap["content"] = p.Content
}

func (p pmsSpuInfo) clone(db *gorm.DB) pmsSpuInfo {
	p.pmsSpuInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsSpuInfo) replaceDB(db *gorm.DB) pmsSpuInfo {
	p.pmsSpuInfoDo.ReplaceDB(db)
	return p
}

type pmsSpuInfoDo struct{ gen.DO }

type IPmsSpuInfoDo interface {
	gen.SubQuery
	Debug() IPmsSpuInfoDo
	WithContext(ctx context.Context) IPmsSpuInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPmsSpuInfoDo
	WriteDB() IPmsSpuInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPmsSpuInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPmsSpuInfoDo
	Not(conds ...gen.Condition) IPmsSpuInfoDo
	Or(conds ...gen.Condition) IPmsSpuInfoDo
	Select(conds ...field.Expr) IPmsSpuInfoDo
	Where(conds ...gen.Condition) IPmsSpuInfoDo
	Order(conds ...field.Expr) IPmsSpuInfoDo
	Distinct(cols ...field.Expr) IPmsSpuInfoDo
	Omit(cols ...field.Expr) IPmsSpuInfoDo
	Join(table schema.Tabler, on ...field.Expr) IPmsSpuInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPmsSpuInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPmsSpuInfoDo
	Group(cols ...field.Expr) IPmsSpuInfoDo
	Having(conds ...gen.Condition) IPmsSpuInfoDo
	Limit(limit int) IPmsSpuInfoDo
	Offset(offset int) IPmsSpuInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsSpuInfoDo
	Unscoped() IPmsSpuInfoDo
	Create(values ...*model.PmsSpuInfo) error
	CreateInBatches(values []*model.PmsSpuInfo, batchSize int) error
	Save(values ...*model.PmsSpuInfo) error
	First() (*model.PmsSpuInfo, error)
	Take() (*model.PmsSpuInfo, error)
	Last() (*model.PmsSpuInfo, error)
	Find() ([]*model.PmsSpuInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsSpuInfo, err error)
	FindInBatches(result *[]*model.PmsSpuInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PmsSpuInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPmsSpuInfoDo
	Assign(attrs ...field.AssignExpr) IPmsSpuInfoDo
	Joins(fields ...field.RelationField) IPmsSpuInfoDo
	Preload(fields ...field.RelationField) IPmsSpuInfoDo
	FirstOrInit() (*model.PmsSpuInfo, error)
	FirstOrCreate() (*model.PmsSpuInfo, error)
	FindByPage(offset int, limit int) (result []*model.PmsSpuInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPmsSpuInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pmsSpuInfoDo) Debug() IPmsSpuInfoDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsSpuInfoDo) WithContext(ctx context.Context) IPmsSpuInfoDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsSpuInfoDo) ReadDB() IPmsSpuInfoDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsSpuInfoDo) WriteDB() IPmsSpuInfoDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsSpuInfoDo) Session(config *gorm.Session) IPmsSpuInfoDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsSpuInfoDo) Clauses(conds ...clause.Expression) IPmsSpuInfoDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsSpuInfoDo) Returning(value interface{}, columns ...string) IPmsSpuInfoDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsSpuInfoDo) Not(conds ...gen.Condition) IPmsSpuInfoDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsSpuInfoDo) Or(conds ...gen.Condition) IPmsSpuInfoDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsSpuInfoDo) Select(conds ...field.Expr) IPmsSpuInfoDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsSpuInfoDo) Where(conds ...gen.Condition) IPmsSpuInfoDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsSpuInfoDo) Order(conds ...field.Expr) IPmsSpuInfoDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsSpuInfoDo) Distinct(cols ...field.Expr) IPmsSpuInfoDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsSpuInfoDo) Omit(cols ...field.Expr) IPmsSpuInfoDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsSpuInfoDo) Join(table schema.Tabler, on ...field.Expr) IPmsSpuInfoDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsSpuInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPmsSpuInfoDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsSpuInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IPmsSpuInfoDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsSpuInfoDo) Group(cols ...field.Expr) IPmsSpuInfoDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsSpuInfoDo) Having(conds ...gen.Condition) IPmsSpuInfoDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsSpuInfoDo) Limit(limit int) IPmsSpuInfoDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsSpuInfoDo) Offset(offset int) IPmsSpuInfoDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsSpuInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsSpuInfoDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsSpuInfoDo) Unscoped() IPmsSpuInfoDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsSpuInfoDo) Create(values ...*model.PmsSpuInfo) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsSpuInfoDo) CreateInBatches(values []*model.PmsSpuInfo, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsSpuInfoDo) Save(values ...*model.PmsSpuInfo) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsSpuInfoDo) First() (*model.PmsSpuInfo, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSpuInfo), nil
	}
}

func (p pmsSpuInfoDo) Take() (*model.PmsSpuInfo, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSpuInfo), nil
	}
}

func (p pmsSpuInfoDo) Last() (*model.PmsSpuInfo, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSpuInfo), nil
	}
}

func (p pmsSpuInfoDo) Find() ([]*model.PmsSpuInfo, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsSpuInfo), err
}

func (p pmsSpuInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsSpuInfo, err error) {
	buf := make([]*model.PmsSpuInfo, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsSpuInfoDo) FindInBatches(result *[]*model.PmsSpuInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsSpuInfoDo) Attrs(attrs ...field.AssignExpr) IPmsSpuInfoDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsSpuInfoDo) Assign(attrs ...field.AssignExpr) IPmsSpuInfoDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsSpuInfoDo) Joins(fields ...field.RelationField) IPmsSpuInfoDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsSpuInfoDo) Preload(fields ...field.RelationField) IPmsSpuInfoDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsSpuInfoDo) FirstOrInit() (*model.PmsSpuInfo, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSpuInfo), nil
	}
}

func (p pmsSpuInfoDo) FirstOrCreate() (*model.PmsSpuInfo, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSpuInfo), nil
	}
}

func (p pmsSpuInfoDo) FindByPage(offset int, limit int) (result []*model.PmsSpuInfo, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pmsSpuInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsSpuInfoDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsSpuInfoDo) Delete(models ...*model.PmsSpuInfo) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsSpuInfoDo) withDO(do gen.Dao) *pmsSpuInfoDo {
	p.DO = *do.(*gen.DO)
	return p
}
