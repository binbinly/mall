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

func newSysExpress(db *gorm.DB, opts ...gen.DOOption) sysExpress {
	_sysExpress := sysExpress{}

	_sysExpress.sysExpressDo.UseDB(db, opts...)
	_sysExpress.sysExpressDo.UseModel(&model.SysExpress{})

	tableName := _sysExpress.sysExpressDo.TableName()
	_sysExpress.ALL = field.NewAsterisk(tableName)
	_sysExpress.ID = field.NewInt(tableName, "id")
	_sysExpress.Code = field.NewString(tableName, "code")
	_sysExpress.Name = field.NewString(tableName, "name")
	_sysExpress.Sort = field.NewInt(tableName, "sort")
	_sysExpress.IsShow = field.NewBool(tableName, "is_show")
	_sysExpress.Status = field.NewBool(tableName, "status")

	_sysExpress.fillFieldMap()

	return _sysExpress
}

// sysExpress 快递公司表
type sysExpress struct {
	sysExpressDo

	ALL    field.Asterisk
	ID     field.Int    // 快递公司id
	Code   field.String // 快递公司简称
	Name   field.String // 快递公司全称
	Sort   field.Int    // 排序
	IsShow field.Bool   // 是否显示
	Status field.Bool   // 是否可用

	fieldMap map[string]field.Expr
}

func (s sysExpress) Table(newTableName string) *sysExpress {
	s.sysExpressDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysExpress) As(alias string) *sysExpress {
	s.sysExpressDo.DO = *(s.sysExpressDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysExpress) updateTableName(table string) *sysExpress {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt(table, "id")
	s.Code = field.NewString(table, "code")
	s.Name = field.NewString(table, "name")
	s.Sort = field.NewInt(table, "sort")
	s.IsShow = field.NewBool(table, "is_show")
	s.Status = field.NewBool(table, "status")

	s.fillFieldMap()

	return s
}

func (s *sysExpress) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysExpress) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id"] = s.ID
	s.fieldMap["code"] = s.Code
	s.fieldMap["name"] = s.Name
	s.fieldMap["sort"] = s.Sort
	s.fieldMap["is_show"] = s.IsShow
	s.fieldMap["status"] = s.Status
}

func (s sysExpress) clone(db *gorm.DB) sysExpress {
	s.sysExpressDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysExpress) replaceDB(db *gorm.DB) sysExpress {
	s.sysExpressDo.ReplaceDB(db)
	return s
}

type sysExpressDo struct{ gen.DO }

type ISysExpressDo interface {
	gen.SubQuery
	Debug() ISysExpressDo
	WithContext(ctx context.Context) ISysExpressDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysExpressDo
	WriteDB() ISysExpressDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysExpressDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysExpressDo
	Not(conds ...gen.Condition) ISysExpressDo
	Or(conds ...gen.Condition) ISysExpressDo
	Select(conds ...field.Expr) ISysExpressDo
	Where(conds ...gen.Condition) ISysExpressDo
	Order(conds ...field.Expr) ISysExpressDo
	Distinct(cols ...field.Expr) ISysExpressDo
	Omit(cols ...field.Expr) ISysExpressDo
	Join(table schema.Tabler, on ...field.Expr) ISysExpressDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysExpressDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysExpressDo
	Group(cols ...field.Expr) ISysExpressDo
	Having(conds ...gen.Condition) ISysExpressDo
	Limit(limit int) ISysExpressDo
	Offset(offset int) ISysExpressDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysExpressDo
	Unscoped() ISysExpressDo
	Create(values ...*model.SysExpress) error
	CreateInBatches(values []*model.SysExpress, batchSize int) error
	Save(values ...*model.SysExpress) error
	First() (*model.SysExpress, error)
	Take() (*model.SysExpress, error)
	Last() (*model.SysExpress, error)
	Find() ([]*model.SysExpress, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysExpress, err error)
	FindInBatches(result *[]*model.SysExpress, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysExpress) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysExpressDo
	Assign(attrs ...field.AssignExpr) ISysExpressDo
	Joins(fields ...field.RelationField) ISysExpressDo
	Preload(fields ...field.RelationField) ISysExpressDo
	FirstOrInit() (*model.SysExpress, error)
	FirstOrCreate() (*model.SysExpress, error)
	FindByPage(offset int, limit int) (result []*model.SysExpress, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysExpressDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysExpressDo) Debug() ISysExpressDo {
	return s.withDO(s.DO.Debug())
}

func (s sysExpressDo) WithContext(ctx context.Context) ISysExpressDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysExpressDo) ReadDB() ISysExpressDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysExpressDo) WriteDB() ISysExpressDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysExpressDo) Session(config *gorm.Session) ISysExpressDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysExpressDo) Clauses(conds ...clause.Expression) ISysExpressDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysExpressDo) Returning(value interface{}, columns ...string) ISysExpressDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysExpressDo) Not(conds ...gen.Condition) ISysExpressDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysExpressDo) Or(conds ...gen.Condition) ISysExpressDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysExpressDo) Select(conds ...field.Expr) ISysExpressDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysExpressDo) Where(conds ...gen.Condition) ISysExpressDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysExpressDo) Order(conds ...field.Expr) ISysExpressDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysExpressDo) Distinct(cols ...field.Expr) ISysExpressDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysExpressDo) Omit(cols ...field.Expr) ISysExpressDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysExpressDo) Join(table schema.Tabler, on ...field.Expr) ISysExpressDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysExpressDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysExpressDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysExpressDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysExpressDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysExpressDo) Group(cols ...field.Expr) ISysExpressDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysExpressDo) Having(conds ...gen.Condition) ISysExpressDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysExpressDo) Limit(limit int) ISysExpressDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysExpressDo) Offset(offset int) ISysExpressDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysExpressDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysExpressDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysExpressDo) Unscoped() ISysExpressDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysExpressDo) Create(values ...*model.SysExpress) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysExpressDo) CreateInBatches(values []*model.SysExpress, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysExpressDo) Save(values ...*model.SysExpress) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysExpressDo) First() (*model.SysExpress, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysExpress), nil
	}
}

func (s sysExpressDo) Take() (*model.SysExpress, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysExpress), nil
	}
}

func (s sysExpressDo) Last() (*model.SysExpress, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysExpress), nil
	}
}

func (s sysExpressDo) Find() ([]*model.SysExpress, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysExpress), err
}

func (s sysExpressDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysExpress, err error) {
	buf := make([]*model.SysExpress, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysExpressDo) FindInBatches(result *[]*model.SysExpress, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysExpressDo) Attrs(attrs ...field.AssignExpr) ISysExpressDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysExpressDo) Assign(attrs ...field.AssignExpr) ISysExpressDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysExpressDo) Joins(fields ...field.RelationField) ISysExpressDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysExpressDo) Preload(fields ...field.RelationField) ISysExpressDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysExpressDo) FirstOrInit() (*model.SysExpress, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysExpress), nil
	}
}

func (s sysExpressDo) FirstOrCreate() (*model.SysExpress, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysExpress), nil
	}
}

func (s sysExpressDo) FindByPage(offset int, limit int) (result []*model.SysExpress, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysExpressDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysExpressDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysExpressDo) Delete(models ...*model.SysExpress) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysExpressDo) withDO(do gen.Dao) *sysExpressDo {
	s.DO = *do.(*gen.DO)
	return s
}
