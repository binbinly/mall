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

func newPmsAttrValue(db *gorm.DB, opts ...gen.DOOption) pmsAttrValue {
	_pmsAttrValue := pmsAttrValue{}

	_pmsAttrValue.pmsAttrValueDo.UseDB(db, opts...)
	_pmsAttrValue.pmsAttrValueDo.UseModel(&model.PmsAttrValue{})

	tableName := _pmsAttrValue.pmsAttrValueDo.TableName()
	_pmsAttrValue.ALL = field.NewAsterisk(tableName)
	_pmsAttrValue.ID = field.NewInt(tableName, "id")
	_pmsAttrValue.SpuID = field.NewInt(tableName, "spu_id")
	_pmsAttrValue.AttrID = field.NewInt(tableName, "attr_id")
	_pmsAttrValue.AttrName = field.NewString(tableName, "attr_name")
	_pmsAttrValue.AttrValue = field.NewString(tableName, "attr_value")
	_pmsAttrValue.IsShow = field.NewInt8(tableName, "is_show")
	_pmsAttrValue.Sort = field.NewInt32(tableName, "sort")

	_pmsAttrValue.fillFieldMap()

	return _pmsAttrValue
}

type pmsAttrValue struct {
	pmsAttrValueDo

	ALL       field.Asterisk
	ID        field.Int    // ID
	SpuID     field.Int    // spu_id
	AttrID    field.Int    // 属性id
	AttrName  field.String // 属性名
	AttrValue field.String // 属性值
	IsShow    field.Int8   // 是否展示在介绍上
	Sort      field.Int32  // 排序

	fieldMap map[string]field.Expr
}

func (p pmsAttrValue) Table(newTableName string) *pmsAttrValue {
	p.pmsAttrValueDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsAttrValue) As(alias string) *pmsAttrValue {
	p.pmsAttrValueDo.DO = *(p.pmsAttrValueDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsAttrValue) updateTableName(table string) *pmsAttrValue {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt(table, "id")
	p.SpuID = field.NewInt(table, "spu_id")
	p.AttrID = field.NewInt(table, "attr_id")
	p.AttrName = field.NewString(table, "attr_name")
	p.AttrValue = field.NewString(table, "attr_value")
	p.IsShow = field.NewInt8(table, "is_show")
	p.Sort = field.NewInt32(table, "sort")

	p.fillFieldMap()

	return p
}

func (p *pmsAttrValue) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsAttrValue) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["id"] = p.ID
	p.fieldMap["spu_id"] = p.SpuID
	p.fieldMap["attr_id"] = p.AttrID
	p.fieldMap["attr_name"] = p.AttrName
	p.fieldMap["attr_value"] = p.AttrValue
	p.fieldMap["is_show"] = p.IsShow
	p.fieldMap["sort"] = p.Sort
}

func (p pmsAttrValue) clone(db *gorm.DB) pmsAttrValue {
	p.pmsAttrValueDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsAttrValue) replaceDB(db *gorm.DB) pmsAttrValue {
	p.pmsAttrValueDo.ReplaceDB(db)
	return p
}

type pmsAttrValueDo struct{ gen.DO }

type IPmsAttrValueDo interface {
	gen.SubQuery
	Debug() IPmsAttrValueDo
	WithContext(ctx context.Context) IPmsAttrValueDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPmsAttrValueDo
	WriteDB() IPmsAttrValueDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPmsAttrValueDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPmsAttrValueDo
	Not(conds ...gen.Condition) IPmsAttrValueDo
	Or(conds ...gen.Condition) IPmsAttrValueDo
	Select(conds ...field.Expr) IPmsAttrValueDo
	Where(conds ...gen.Condition) IPmsAttrValueDo
	Order(conds ...field.Expr) IPmsAttrValueDo
	Distinct(cols ...field.Expr) IPmsAttrValueDo
	Omit(cols ...field.Expr) IPmsAttrValueDo
	Join(table schema.Tabler, on ...field.Expr) IPmsAttrValueDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPmsAttrValueDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPmsAttrValueDo
	Group(cols ...field.Expr) IPmsAttrValueDo
	Having(conds ...gen.Condition) IPmsAttrValueDo
	Limit(limit int) IPmsAttrValueDo
	Offset(offset int) IPmsAttrValueDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsAttrValueDo
	Unscoped() IPmsAttrValueDo
	Create(values ...*model.PmsAttrValue) error
	CreateInBatches(values []*model.PmsAttrValue, batchSize int) error
	Save(values ...*model.PmsAttrValue) error
	First() (*model.PmsAttrValue, error)
	Take() (*model.PmsAttrValue, error)
	Last() (*model.PmsAttrValue, error)
	Find() ([]*model.PmsAttrValue, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsAttrValue, err error)
	FindInBatches(result *[]*model.PmsAttrValue, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PmsAttrValue) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPmsAttrValueDo
	Assign(attrs ...field.AssignExpr) IPmsAttrValueDo
	Joins(fields ...field.RelationField) IPmsAttrValueDo
	Preload(fields ...field.RelationField) IPmsAttrValueDo
	FirstOrInit() (*model.PmsAttrValue, error)
	FirstOrCreate() (*model.PmsAttrValue, error)
	FindByPage(offset int, limit int) (result []*model.PmsAttrValue, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPmsAttrValueDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pmsAttrValueDo) Debug() IPmsAttrValueDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsAttrValueDo) WithContext(ctx context.Context) IPmsAttrValueDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsAttrValueDo) ReadDB() IPmsAttrValueDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsAttrValueDo) WriteDB() IPmsAttrValueDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsAttrValueDo) Session(config *gorm.Session) IPmsAttrValueDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsAttrValueDo) Clauses(conds ...clause.Expression) IPmsAttrValueDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsAttrValueDo) Returning(value interface{}, columns ...string) IPmsAttrValueDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsAttrValueDo) Not(conds ...gen.Condition) IPmsAttrValueDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsAttrValueDo) Or(conds ...gen.Condition) IPmsAttrValueDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsAttrValueDo) Select(conds ...field.Expr) IPmsAttrValueDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsAttrValueDo) Where(conds ...gen.Condition) IPmsAttrValueDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsAttrValueDo) Order(conds ...field.Expr) IPmsAttrValueDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsAttrValueDo) Distinct(cols ...field.Expr) IPmsAttrValueDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsAttrValueDo) Omit(cols ...field.Expr) IPmsAttrValueDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsAttrValueDo) Join(table schema.Tabler, on ...field.Expr) IPmsAttrValueDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsAttrValueDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPmsAttrValueDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsAttrValueDo) RightJoin(table schema.Tabler, on ...field.Expr) IPmsAttrValueDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsAttrValueDo) Group(cols ...field.Expr) IPmsAttrValueDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsAttrValueDo) Having(conds ...gen.Condition) IPmsAttrValueDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsAttrValueDo) Limit(limit int) IPmsAttrValueDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsAttrValueDo) Offset(offset int) IPmsAttrValueDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsAttrValueDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsAttrValueDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsAttrValueDo) Unscoped() IPmsAttrValueDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsAttrValueDo) Create(values ...*model.PmsAttrValue) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsAttrValueDo) CreateInBatches(values []*model.PmsAttrValue, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsAttrValueDo) Save(values ...*model.PmsAttrValue) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsAttrValueDo) First() (*model.PmsAttrValue, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAttrValue), nil
	}
}

func (p pmsAttrValueDo) Take() (*model.PmsAttrValue, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAttrValue), nil
	}
}

func (p pmsAttrValueDo) Last() (*model.PmsAttrValue, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAttrValue), nil
	}
}

func (p pmsAttrValueDo) Find() ([]*model.PmsAttrValue, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsAttrValue), err
}

func (p pmsAttrValueDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsAttrValue, err error) {
	buf := make([]*model.PmsAttrValue, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsAttrValueDo) FindInBatches(result *[]*model.PmsAttrValue, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsAttrValueDo) Attrs(attrs ...field.AssignExpr) IPmsAttrValueDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsAttrValueDo) Assign(attrs ...field.AssignExpr) IPmsAttrValueDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsAttrValueDo) Joins(fields ...field.RelationField) IPmsAttrValueDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsAttrValueDo) Preload(fields ...field.RelationField) IPmsAttrValueDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsAttrValueDo) FirstOrInit() (*model.PmsAttrValue, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAttrValue), nil
	}
}

func (p pmsAttrValueDo) FirstOrCreate() (*model.PmsAttrValue, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsAttrValue), nil
	}
}

func (p pmsAttrValueDo) FindByPage(offset int, limit int) (result []*model.PmsAttrValue, count int64, err error) {
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

func (p pmsAttrValueDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsAttrValueDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsAttrValueDo) Delete(models ...*model.PmsAttrValue) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsAttrValueDo) withDO(do gen.Dao) *pmsAttrValueDo {
	p.DO = *do.(*gen.DO)
	return p
}
