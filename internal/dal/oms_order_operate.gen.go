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

func newOmsOrderOperate(db *gorm.DB, opts ...gen.DOOption) omsOrderOperate {
	_omsOrderOperate := omsOrderOperate{}

	_omsOrderOperate.omsOrderOperateDo.UseDB(db, opts...)
	_omsOrderOperate.omsOrderOperateDo.UseModel(&model.OmsOrderOperate{})

	tableName := _omsOrderOperate.omsOrderOperateDo.TableName()
	_omsOrderOperate.ALL = field.NewAsterisk(tableName)
	_omsOrderOperate.ID = field.NewInt(tableName, "id")
	_omsOrderOperate.OrderID = field.NewInt(tableName, "order_id")
	_omsOrderOperate.OrderStatus = field.NewInt8(tableName, "order_status")
	_omsOrderOperate.Note = field.NewString(tableName, "note")
	_omsOrderOperate.OperateName = field.NewString(tableName, "operate_name")
	_omsOrderOperate.CreatedAt = field.NewTime(tableName, "created_at")

	_omsOrderOperate.fillFieldMap()

	return _omsOrderOperate
}

type omsOrderOperate struct {
	omsOrderOperateDo

	ALL         field.Asterisk
	ID          field.Int    // ID
	OrderID     field.Int    // 订单id
	OrderStatus field.Int8   // 订单状态
	Note        field.String // 备注
	OperateName field.String // 操作人
	CreatedAt   field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (o omsOrderOperate) Table(newTableName string) *omsOrderOperate {
	o.omsOrderOperateDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o omsOrderOperate) As(alias string) *omsOrderOperate {
	o.omsOrderOperateDo.DO = *(o.omsOrderOperateDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *omsOrderOperate) updateTableName(table string) *omsOrderOperate {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt(table, "id")
	o.OrderID = field.NewInt(table, "order_id")
	o.OrderStatus = field.NewInt8(table, "order_status")
	o.Note = field.NewString(table, "note")
	o.OperateName = field.NewString(table, "operate_name")
	o.CreatedAt = field.NewTime(table, "created_at")

	o.fillFieldMap()

	return o
}

func (o *omsOrderOperate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *omsOrderOperate) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 6)
	o.fieldMap["id"] = o.ID
	o.fieldMap["order_id"] = o.OrderID
	o.fieldMap["order_status"] = o.OrderStatus
	o.fieldMap["note"] = o.Note
	o.fieldMap["operate_name"] = o.OperateName
	o.fieldMap["created_at"] = o.CreatedAt
}

func (o omsOrderOperate) clone(db *gorm.DB) omsOrderOperate {
	o.omsOrderOperateDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o omsOrderOperate) replaceDB(db *gorm.DB) omsOrderOperate {
	o.omsOrderOperateDo.ReplaceDB(db)
	return o
}

type omsOrderOperateDo struct{ gen.DO }

type IOmsOrderOperateDo interface {
	gen.SubQuery
	Debug() IOmsOrderOperateDo
	WithContext(ctx context.Context) IOmsOrderOperateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOmsOrderOperateDo
	WriteDB() IOmsOrderOperateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOmsOrderOperateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOmsOrderOperateDo
	Not(conds ...gen.Condition) IOmsOrderOperateDo
	Or(conds ...gen.Condition) IOmsOrderOperateDo
	Select(conds ...field.Expr) IOmsOrderOperateDo
	Where(conds ...gen.Condition) IOmsOrderOperateDo
	Order(conds ...field.Expr) IOmsOrderOperateDo
	Distinct(cols ...field.Expr) IOmsOrderOperateDo
	Omit(cols ...field.Expr) IOmsOrderOperateDo
	Join(table schema.Tabler, on ...field.Expr) IOmsOrderOperateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderOperateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderOperateDo
	Group(cols ...field.Expr) IOmsOrderOperateDo
	Having(conds ...gen.Condition) IOmsOrderOperateDo
	Limit(limit int) IOmsOrderOperateDo
	Offset(offset int) IOmsOrderOperateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderOperateDo
	Unscoped() IOmsOrderOperateDo
	Create(values ...*model.OmsOrderOperate) error
	CreateInBatches(values []*model.OmsOrderOperate, batchSize int) error
	Save(values ...*model.OmsOrderOperate) error
	First() (*model.OmsOrderOperate, error)
	Take() (*model.OmsOrderOperate, error)
	Last() (*model.OmsOrderOperate, error)
	Find() ([]*model.OmsOrderOperate, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderOperate, err error)
	FindInBatches(result *[]*model.OmsOrderOperate, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OmsOrderOperate) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOmsOrderOperateDo
	Assign(attrs ...field.AssignExpr) IOmsOrderOperateDo
	Joins(fields ...field.RelationField) IOmsOrderOperateDo
	Preload(fields ...field.RelationField) IOmsOrderOperateDo
	FirstOrInit() (*model.OmsOrderOperate, error)
	FirstOrCreate() (*model.OmsOrderOperate, error)
	FindByPage(offset int, limit int) (result []*model.OmsOrderOperate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOmsOrderOperateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o omsOrderOperateDo) Debug() IOmsOrderOperateDo {
	return o.withDO(o.DO.Debug())
}

func (o omsOrderOperateDo) WithContext(ctx context.Context) IOmsOrderOperateDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o omsOrderOperateDo) ReadDB() IOmsOrderOperateDo {
	return o.Clauses(dbresolver.Read)
}

func (o omsOrderOperateDo) WriteDB() IOmsOrderOperateDo {
	return o.Clauses(dbresolver.Write)
}

func (o omsOrderOperateDo) Session(config *gorm.Session) IOmsOrderOperateDo {
	return o.withDO(o.DO.Session(config))
}

func (o omsOrderOperateDo) Clauses(conds ...clause.Expression) IOmsOrderOperateDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o omsOrderOperateDo) Returning(value interface{}, columns ...string) IOmsOrderOperateDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o omsOrderOperateDo) Not(conds ...gen.Condition) IOmsOrderOperateDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o omsOrderOperateDo) Or(conds ...gen.Condition) IOmsOrderOperateDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o omsOrderOperateDo) Select(conds ...field.Expr) IOmsOrderOperateDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o omsOrderOperateDo) Where(conds ...gen.Condition) IOmsOrderOperateDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o omsOrderOperateDo) Order(conds ...field.Expr) IOmsOrderOperateDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o omsOrderOperateDo) Distinct(cols ...field.Expr) IOmsOrderOperateDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o omsOrderOperateDo) Omit(cols ...field.Expr) IOmsOrderOperateDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o omsOrderOperateDo) Join(table schema.Tabler, on ...field.Expr) IOmsOrderOperateDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o omsOrderOperateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderOperateDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o omsOrderOperateDo) RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderOperateDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o omsOrderOperateDo) Group(cols ...field.Expr) IOmsOrderOperateDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o omsOrderOperateDo) Having(conds ...gen.Condition) IOmsOrderOperateDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o omsOrderOperateDo) Limit(limit int) IOmsOrderOperateDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o omsOrderOperateDo) Offset(offset int) IOmsOrderOperateDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o omsOrderOperateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderOperateDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o omsOrderOperateDo) Unscoped() IOmsOrderOperateDo {
	return o.withDO(o.DO.Unscoped())
}

func (o omsOrderOperateDo) Create(values ...*model.OmsOrderOperate) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o omsOrderOperateDo) CreateInBatches(values []*model.OmsOrderOperate, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o omsOrderOperateDo) Save(values ...*model.OmsOrderOperate) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o omsOrderOperateDo) First() (*model.OmsOrderOperate, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderOperate), nil
	}
}

func (o omsOrderOperateDo) Take() (*model.OmsOrderOperate, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderOperate), nil
	}
}

func (o omsOrderOperateDo) Last() (*model.OmsOrderOperate, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderOperate), nil
	}
}

func (o omsOrderOperateDo) Find() ([]*model.OmsOrderOperate, error) {
	result, err := o.DO.Find()
	return result.([]*model.OmsOrderOperate), err
}

func (o omsOrderOperateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderOperate, err error) {
	buf := make([]*model.OmsOrderOperate, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o omsOrderOperateDo) FindInBatches(result *[]*model.OmsOrderOperate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o omsOrderOperateDo) Attrs(attrs ...field.AssignExpr) IOmsOrderOperateDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o omsOrderOperateDo) Assign(attrs ...field.AssignExpr) IOmsOrderOperateDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o omsOrderOperateDo) Joins(fields ...field.RelationField) IOmsOrderOperateDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o omsOrderOperateDo) Preload(fields ...field.RelationField) IOmsOrderOperateDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o omsOrderOperateDo) FirstOrInit() (*model.OmsOrderOperate, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderOperate), nil
	}
}

func (o omsOrderOperateDo) FirstOrCreate() (*model.OmsOrderOperate, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderOperate), nil
	}
}

func (o omsOrderOperateDo) FindByPage(offset int, limit int) (result []*model.OmsOrderOperate, count int64, err error) {
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

func (o omsOrderOperateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o omsOrderOperateDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o omsOrderOperateDo) Delete(models ...*model.OmsOrderOperate) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *omsOrderOperateDo) withDO(do gen.Dao) *omsOrderOperateDo {
	o.DO = *do.(*gen.DO)
	return o
}
