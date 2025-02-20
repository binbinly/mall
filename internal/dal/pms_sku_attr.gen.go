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

func newPmsSkuAttr(db *gorm.DB, opts ...gen.DOOption) pmsSkuAttr {
	_pmsSkuAttr := pmsSkuAttr{}

	_pmsSkuAttr.pmsSkuAttrDo.UseDB(db, opts...)
	_pmsSkuAttr.pmsSkuAttrDo.UseModel(&model.PmsSkuAttr{})

	tableName := _pmsSkuAttr.pmsSkuAttrDo.TableName()
	_pmsSkuAttr.ALL = field.NewAsterisk(tableName)
	_pmsSkuAttr.ID = field.NewInt(tableName, "id")
	_pmsSkuAttr.SkuID = field.NewInt(tableName, "sku_id")
	_pmsSkuAttr.AttrID = field.NewInt(tableName, "attr_id")
	_pmsSkuAttr.AttrName = field.NewString(tableName, "attr_name")
	_pmsSkuAttr.AttrValue = field.NewString(tableName, "attr_value")
	_pmsSkuAttr.Sort = field.NewInt32(tableName, "sort")

	_pmsSkuAttr.fillFieldMap()

	return _pmsSkuAttr
}

type pmsSkuAttr struct {
	pmsSkuAttrDo

	ALL       field.Asterisk
	ID        field.Int    // ID
	SkuID     field.Int    // sku_id
	AttrID    field.Int    // 属性id
	AttrName  field.String // 销售属性名
	AttrValue field.String // 销售属性值
	Sort      field.Int32  // 排序

	fieldMap map[string]field.Expr
}

func (p pmsSkuAttr) Table(newTableName string) *pmsSkuAttr {
	p.pmsSkuAttrDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsSkuAttr) As(alias string) *pmsSkuAttr {
	p.pmsSkuAttrDo.DO = *(p.pmsSkuAttrDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsSkuAttr) updateTableName(table string) *pmsSkuAttr {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt(table, "id")
	p.SkuID = field.NewInt(table, "sku_id")
	p.AttrID = field.NewInt(table, "attr_id")
	p.AttrName = field.NewString(table, "attr_name")
	p.AttrValue = field.NewString(table, "attr_value")
	p.Sort = field.NewInt32(table, "sort")

	p.fillFieldMap()

	return p
}

func (p *pmsSkuAttr) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsSkuAttr) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["sku_id"] = p.SkuID
	p.fieldMap["attr_id"] = p.AttrID
	p.fieldMap["attr_name"] = p.AttrName
	p.fieldMap["attr_value"] = p.AttrValue
	p.fieldMap["sort"] = p.Sort
}

func (p pmsSkuAttr) clone(db *gorm.DB) pmsSkuAttr {
	p.pmsSkuAttrDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsSkuAttr) replaceDB(db *gorm.DB) pmsSkuAttr {
	p.pmsSkuAttrDo.ReplaceDB(db)
	return p
}

type pmsSkuAttrDo struct{ gen.DO }

type IPmsSkuAttrDo interface {
	gen.SubQuery
	Debug() IPmsSkuAttrDo
	WithContext(ctx context.Context) IPmsSkuAttrDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPmsSkuAttrDo
	WriteDB() IPmsSkuAttrDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPmsSkuAttrDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPmsSkuAttrDo
	Not(conds ...gen.Condition) IPmsSkuAttrDo
	Or(conds ...gen.Condition) IPmsSkuAttrDo
	Select(conds ...field.Expr) IPmsSkuAttrDo
	Where(conds ...gen.Condition) IPmsSkuAttrDo
	Order(conds ...field.Expr) IPmsSkuAttrDo
	Distinct(cols ...field.Expr) IPmsSkuAttrDo
	Omit(cols ...field.Expr) IPmsSkuAttrDo
	Join(table schema.Tabler, on ...field.Expr) IPmsSkuAttrDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPmsSkuAttrDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPmsSkuAttrDo
	Group(cols ...field.Expr) IPmsSkuAttrDo
	Having(conds ...gen.Condition) IPmsSkuAttrDo
	Limit(limit int) IPmsSkuAttrDo
	Offset(offset int) IPmsSkuAttrDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsSkuAttrDo
	Unscoped() IPmsSkuAttrDo
	Create(values ...*model.PmsSkuAttr) error
	CreateInBatches(values []*model.PmsSkuAttr, batchSize int) error
	Save(values ...*model.PmsSkuAttr) error
	First() (*model.PmsSkuAttr, error)
	Take() (*model.PmsSkuAttr, error)
	Last() (*model.PmsSkuAttr, error)
	Find() ([]*model.PmsSkuAttr, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsSkuAttr, err error)
	FindInBatches(result *[]*model.PmsSkuAttr, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PmsSkuAttr) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPmsSkuAttrDo
	Assign(attrs ...field.AssignExpr) IPmsSkuAttrDo
	Joins(fields ...field.RelationField) IPmsSkuAttrDo
	Preload(fields ...field.RelationField) IPmsSkuAttrDo
	FirstOrInit() (*model.PmsSkuAttr, error)
	FirstOrCreate() (*model.PmsSkuAttr, error)
	FindByPage(offset int, limit int) (result []*model.PmsSkuAttr, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPmsSkuAttrDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pmsSkuAttrDo) Debug() IPmsSkuAttrDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsSkuAttrDo) WithContext(ctx context.Context) IPmsSkuAttrDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsSkuAttrDo) ReadDB() IPmsSkuAttrDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsSkuAttrDo) WriteDB() IPmsSkuAttrDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsSkuAttrDo) Session(config *gorm.Session) IPmsSkuAttrDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsSkuAttrDo) Clauses(conds ...clause.Expression) IPmsSkuAttrDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsSkuAttrDo) Returning(value interface{}, columns ...string) IPmsSkuAttrDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsSkuAttrDo) Not(conds ...gen.Condition) IPmsSkuAttrDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsSkuAttrDo) Or(conds ...gen.Condition) IPmsSkuAttrDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsSkuAttrDo) Select(conds ...field.Expr) IPmsSkuAttrDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsSkuAttrDo) Where(conds ...gen.Condition) IPmsSkuAttrDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsSkuAttrDo) Order(conds ...field.Expr) IPmsSkuAttrDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsSkuAttrDo) Distinct(cols ...field.Expr) IPmsSkuAttrDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsSkuAttrDo) Omit(cols ...field.Expr) IPmsSkuAttrDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsSkuAttrDo) Join(table schema.Tabler, on ...field.Expr) IPmsSkuAttrDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsSkuAttrDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPmsSkuAttrDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsSkuAttrDo) RightJoin(table schema.Tabler, on ...field.Expr) IPmsSkuAttrDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsSkuAttrDo) Group(cols ...field.Expr) IPmsSkuAttrDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsSkuAttrDo) Having(conds ...gen.Condition) IPmsSkuAttrDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsSkuAttrDo) Limit(limit int) IPmsSkuAttrDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsSkuAttrDo) Offset(offset int) IPmsSkuAttrDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsSkuAttrDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsSkuAttrDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsSkuAttrDo) Unscoped() IPmsSkuAttrDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsSkuAttrDo) Create(values ...*model.PmsSkuAttr) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsSkuAttrDo) CreateInBatches(values []*model.PmsSkuAttr, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsSkuAttrDo) Save(values ...*model.PmsSkuAttr) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsSkuAttrDo) First() (*model.PmsSkuAttr, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuAttr), nil
	}
}

func (p pmsSkuAttrDo) Take() (*model.PmsSkuAttr, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuAttr), nil
	}
}

func (p pmsSkuAttrDo) Last() (*model.PmsSkuAttr, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuAttr), nil
	}
}

func (p pmsSkuAttrDo) Find() ([]*model.PmsSkuAttr, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsSkuAttr), err
}

func (p pmsSkuAttrDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsSkuAttr, err error) {
	buf := make([]*model.PmsSkuAttr, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsSkuAttrDo) FindInBatches(result *[]*model.PmsSkuAttr, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsSkuAttrDo) Attrs(attrs ...field.AssignExpr) IPmsSkuAttrDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsSkuAttrDo) Assign(attrs ...field.AssignExpr) IPmsSkuAttrDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsSkuAttrDo) Joins(fields ...field.RelationField) IPmsSkuAttrDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsSkuAttrDo) Preload(fields ...field.RelationField) IPmsSkuAttrDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsSkuAttrDo) FirstOrInit() (*model.PmsSkuAttr, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuAttr), nil
	}
}

func (p pmsSkuAttrDo) FirstOrCreate() (*model.PmsSkuAttr, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuAttr), nil
	}
}

func (p pmsSkuAttrDo) FindByPage(offset int, limit int) (result []*model.PmsSkuAttr, count int64, err error) {
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

func (p pmsSkuAttrDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsSkuAttrDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsSkuAttrDo) Delete(models ...*model.PmsSkuAttr) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsSkuAttrDo) withDO(do gen.Dao) *pmsSkuAttrDo {
	p.DO = *do.(*gen.DO)
	return p
}
