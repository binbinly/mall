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

func newOmsOrderReturnReason(db *gorm.DB, opts ...gen.DOOption) omsOrderReturnReason {
	_omsOrderReturnReason := omsOrderReturnReason{}

	_omsOrderReturnReason.omsOrderReturnReasonDo.UseDB(db, opts...)
	_omsOrderReturnReason.omsOrderReturnReasonDo.UseModel(&model.OmsOrderReturnReason{})

	tableName := _omsOrderReturnReason.omsOrderReturnReasonDo.TableName()
	_omsOrderReturnReason.ALL = field.NewAsterisk(tableName)
	_omsOrderReturnReason.ID = field.NewInt(tableName, "id")
	_omsOrderReturnReason.Name = field.NewString(tableName, "name")
	_omsOrderReturnReason.Status = field.NewInt8(tableName, "status")
	_omsOrderReturnReason.Sort = field.NewInt(tableName, "sort")
	_omsOrderReturnReason.CreatedAt = field.NewInt(tableName, "created_at")
	_omsOrderReturnReason.UpdatedAt = field.NewInt(tableName, "updated_at")

	_omsOrderReturnReason.fillFieldMap()

	return _omsOrderReturnReason
}

type omsOrderReturnReason struct {
	omsOrderReturnReasonDo

	ALL       field.Asterisk
	ID        field.Int    // ID
	Name      field.String // 退货原因名
	Status    field.Int8   // 订单状态【0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单】
	Sort      field.Int    // 排序
	CreatedAt field.Int    // 创建时间
	UpdatedAt field.Int    // 更新时间

	fieldMap map[string]field.Expr
}

func (o omsOrderReturnReason) Table(newTableName string) *omsOrderReturnReason {
	o.omsOrderReturnReasonDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o omsOrderReturnReason) As(alias string) *omsOrderReturnReason {
	o.omsOrderReturnReasonDo.DO = *(o.omsOrderReturnReasonDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *omsOrderReturnReason) updateTableName(table string) *omsOrderReturnReason {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt(table, "id")
	o.Name = field.NewString(table, "name")
	o.Status = field.NewInt8(table, "status")
	o.Sort = field.NewInt(table, "sort")
	o.CreatedAt = field.NewInt(table, "created_at")
	o.UpdatedAt = field.NewInt(table, "updated_at")

	o.fillFieldMap()

	return o
}

func (o *omsOrderReturnReason) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *omsOrderReturnReason) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 6)
	o.fieldMap["id"] = o.ID
	o.fieldMap["name"] = o.Name
	o.fieldMap["status"] = o.Status
	o.fieldMap["sort"] = o.Sort
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
}

func (o omsOrderReturnReason) clone(db *gorm.DB) omsOrderReturnReason {
	o.omsOrderReturnReasonDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o omsOrderReturnReason) replaceDB(db *gorm.DB) omsOrderReturnReason {
	o.omsOrderReturnReasonDo.ReplaceDB(db)
	return o
}

type omsOrderReturnReasonDo struct{ gen.DO }

type IOmsOrderReturnReasonDo interface {
	gen.SubQuery
	Debug() IOmsOrderReturnReasonDo
	WithContext(ctx context.Context) IOmsOrderReturnReasonDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOmsOrderReturnReasonDo
	WriteDB() IOmsOrderReturnReasonDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOmsOrderReturnReasonDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOmsOrderReturnReasonDo
	Not(conds ...gen.Condition) IOmsOrderReturnReasonDo
	Or(conds ...gen.Condition) IOmsOrderReturnReasonDo
	Select(conds ...field.Expr) IOmsOrderReturnReasonDo
	Where(conds ...gen.Condition) IOmsOrderReturnReasonDo
	Order(conds ...field.Expr) IOmsOrderReturnReasonDo
	Distinct(cols ...field.Expr) IOmsOrderReturnReasonDo
	Omit(cols ...field.Expr) IOmsOrderReturnReasonDo
	Join(table schema.Tabler, on ...field.Expr) IOmsOrderReturnReasonDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderReturnReasonDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderReturnReasonDo
	Group(cols ...field.Expr) IOmsOrderReturnReasonDo
	Having(conds ...gen.Condition) IOmsOrderReturnReasonDo
	Limit(limit int) IOmsOrderReturnReasonDo
	Offset(offset int) IOmsOrderReturnReasonDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderReturnReasonDo
	Unscoped() IOmsOrderReturnReasonDo
	Create(values ...*model.OmsOrderReturnReason) error
	CreateInBatches(values []*model.OmsOrderReturnReason, batchSize int) error
	Save(values ...*model.OmsOrderReturnReason) error
	First() (*model.OmsOrderReturnReason, error)
	Take() (*model.OmsOrderReturnReason, error)
	Last() (*model.OmsOrderReturnReason, error)
	Find() ([]*model.OmsOrderReturnReason, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderReturnReason, err error)
	FindInBatches(result *[]*model.OmsOrderReturnReason, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OmsOrderReturnReason) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOmsOrderReturnReasonDo
	Assign(attrs ...field.AssignExpr) IOmsOrderReturnReasonDo
	Joins(fields ...field.RelationField) IOmsOrderReturnReasonDo
	Preload(fields ...field.RelationField) IOmsOrderReturnReasonDo
	FirstOrInit() (*model.OmsOrderReturnReason, error)
	FirstOrCreate() (*model.OmsOrderReturnReason, error)
	FindByPage(offset int, limit int) (result []*model.OmsOrderReturnReason, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOmsOrderReturnReasonDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o omsOrderReturnReasonDo) Debug() IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Debug())
}

func (o omsOrderReturnReasonDo) WithContext(ctx context.Context) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o omsOrderReturnReasonDo) ReadDB() IOmsOrderReturnReasonDo {
	return o.Clauses(dbresolver.Read)
}

func (o omsOrderReturnReasonDo) WriteDB() IOmsOrderReturnReasonDo {
	return o.Clauses(dbresolver.Write)
}

func (o omsOrderReturnReasonDo) Session(config *gorm.Session) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Session(config))
}

func (o omsOrderReturnReasonDo) Clauses(conds ...clause.Expression) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o omsOrderReturnReasonDo) Returning(value interface{}, columns ...string) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o omsOrderReturnReasonDo) Not(conds ...gen.Condition) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o omsOrderReturnReasonDo) Or(conds ...gen.Condition) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o omsOrderReturnReasonDo) Select(conds ...field.Expr) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o omsOrderReturnReasonDo) Where(conds ...gen.Condition) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o omsOrderReturnReasonDo) Order(conds ...field.Expr) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o omsOrderReturnReasonDo) Distinct(cols ...field.Expr) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o omsOrderReturnReasonDo) Omit(cols ...field.Expr) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o omsOrderReturnReasonDo) Join(table schema.Tabler, on ...field.Expr) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o omsOrderReturnReasonDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o omsOrderReturnReasonDo) RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o omsOrderReturnReasonDo) Group(cols ...field.Expr) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o omsOrderReturnReasonDo) Having(conds ...gen.Condition) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o omsOrderReturnReasonDo) Limit(limit int) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o omsOrderReturnReasonDo) Offset(offset int) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o omsOrderReturnReasonDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o omsOrderReturnReasonDo) Unscoped() IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Unscoped())
}

func (o omsOrderReturnReasonDo) Create(values ...*model.OmsOrderReturnReason) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o omsOrderReturnReasonDo) CreateInBatches(values []*model.OmsOrderReturnReason, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o omsOrderReturnReasonDo) Save(values ...*model.OmsOrderReturnReason) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o omsOrderReturnReasonDo) First() (*model.OmsOrderReturnReason, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturnReason), nil
	}
}

func (o omsOrderReturnReasonDo) Take() (*model.OmsOrderReturnReason, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturnReason), nil
	}
}

func (o omsOrderReturnReasonDo) Last() (*model.OmsOrderReturnReason, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturnReason), nil
	}
}

func (o omsOrderReturnReasonDo) Find() ([]*model.OmsOrderReturnReason, error) {
	result, err := o.DO.Find()
	return result.([]*model.OmsOrderReturnReason), err
}

func (o omsOrderReturnReasonDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderReturnReason, err error) {
	buf := make([]*model.OmsOrderReturnReason, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o omsOrderReturnReasonDo) FindInBatches(result *[]*model.OmsOrderReturnReason, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o omsOrderReturnReasonDo) Attrs(attrs ...field.AssignExpr) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o omsOrderReturnReasonDo) Assign(attrs ...field.AssignExpr) IOmsOrderReturnReasonDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o omsOrderReturnReasonDo) Joins(fields ...field.RelationField) IOmsOrderReturnReasonDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o omsOrderReturnReasonDo) Preload(fields ...field.RelationField) IOmsOrderReturnReasonDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o omsOrderReturnReasonDo) FirstOrInit() (*model.OmsOrderReturnReason, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturnReason), nil
	}
}

func (o omsOrderReturnReasonDo) FirstOrCreate() (*model.OmsOrderReturnReason, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturnReason), nil
	}
}

func (o omsOrderReturnReasonDo) FindByPage(offset int, limit int) (result []*model.OmsOrderReturnReason, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o omsOrderReturnReasonDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o omsOrderReturnReasonDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o omsOrderReturnReasonDo) Delete(models ...*model.OmsOrderReturnReason) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *omsOrderReturnReasonDo) withDO(do gen.Dao) *omsOrderReturnReasonDo {
	o.DO = *do.(*gen.DO)
	return o
}
