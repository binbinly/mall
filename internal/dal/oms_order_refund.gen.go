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

func newOmsOrderRefund(db *gorm.DB, opts ...gen.DOOption) omsOrderRefund {
	_omsOrderRefund := omsOrderRefund{}

	_omsOrderRefund.omsOrderRefundDo.UseDB(db, opts...)
	_omsOrderRefund.omsOrderRefundDo.UseModel(&model.OmsOrderRefund{})

	tableName := _omsOrderRefund.omsOrderRefundDo.TableName()
	_omsOrderRefund.ALL = field.NewAsterisk(tableName)
	_omsOrderRefund.ID = field.NewInt(tableName, "id")
	_omsOrderRefund.OrderID = field.NewInt(tableName, "order_id")
	_omsOrderRefund.Amount = field.NewInt(tableName, "amount")
	_omsOrderRefund.TradeNo = field.NewString(tableName, "trade_no")
	_omsOrderRefund.Status = field.NewInt8(tableName, "status")
	_omsOrderRefund.Channel = field.NewInt8(tableName, "channel")
	_omsOrderRefund.Content = field.NewString(tableName, "content")
	_omsOrderRefund.CreatedAt = field.NewTime(tableName, "created_at")
	_omsOrderRefund.UpdatedAt = field.NewTime(tableName, "updated_at")

	_omsOrderRefund.fillFieldMap()

	return _omsOrderRefund
}

type omsOrderRefund struct {
	omsOrderRefundDo

	ALL       field.Asterisk
	ID        field.Int    // ID
	OrderID   field.Int    // 订单id
	Amount    field.Int    // 退款金额
	TradeNo   field.String // 交易号
	Status    field.Int8   // 订单状态
	Channel   field.Int8   // 退款渠道[1-支付宝，2-微信，3-银联，4-汇款]
	Content   field.String // 退款内容
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (o omsOrderRefund) Table(newTableName string) *omsOrderRefund {
	o.omsOrderRefundDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o omsOrderRefund) As(alias string) *omsOrderRefund {
	o.omsOrderRefundDo.DO = *(o.omsOrderRefundDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *omsOrderRefund) updateTableName(table string) *omsOrderRefund {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt(table, "id")
	o.OrderID = field.NewInt(table, "order_id")
	o.Amount = field.NewInt(table, "amount")
	o.TradeNo = field.NewString(table, "trade_no")
	o.Status = field.NewInt8(table, "status")
	o.Channel = field.NewInt8(table, "channel")
	o.Content = field.NewString(table, "content")
	o.CreatedAt = field.NewTime(table, "created_at")
	o.UpdatedAt = field.NewTime(table, "updated_at")

	o.fillFieldMap()

	return o
}

func (o *omsOrderRefund) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *omsOrderRefund) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 9)
	o.fieldMap["id"] = o.ID
	o.fieldMap["order_id"] = o.OrderID
	o.fieldMap["amount"] = o.Amount
	o.fieldMap["trade_no"] = o.TradeNo
	o.fieldMap["status"] = o.Status
	o.fieldMap["channel"] = o.Channel
	o.fieldMap["content"] = o.Content
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
}

func (o omsOrderRefund) clone(db *gorm.DB) omsOrderRefund {
	o.omsOrderRefundDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o omsOrderRefund) replaceDB(db *gorm.DB) omsOrderRefund {
	o.omsOrderRefundDo.ReplaceDB(db)
	return o
}

type omsOrderRefundDo struct{ gen.DO }

type IOmsOrderRefundDo interface {
	gen.SubQuery
	Debug() IOmsOrderRefundDo
	WithContext(ctx context.Context) IOmsOrderRefundDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOmsOrderRefundDo
	WriteDB() IOmsOrderRefundDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOmsOrderRefundDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOmsOrderRefundDo
	Not(conds ...gen.Condition) IOmsOrderRefundDo
	Or(conds ...gen.Condition) IOmsOrderRefundDo
	Select(conds ...field.Expr) IOmsOrderRefundDo
	Where(conds ...gen.Condition) IOmsOrderRefundDo
	Order(conds ...field.Expr) IOmsOrderRefundDo
	Distinct(cols ...field.Expr) IOmsOrderRefundDo
	Omit(cols ...field.Expr) IOmsOrderRefundDo
	Join(table schema.Tabler, on ...field.Expr) IOmsOrderRefundDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderRefundDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderRefundDo
	Group(cols ...field.Expr) IOmsOrderRefundDo
	Having(conds ...gen.Condition) IOmsOrderRefundDo
	Limit(limit int) IOmsOrderRefundDo
	Offset(offset int) IOmsOrderRefundDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderRefundDo
	Unscoped() IOmsOrderRefundDo
	Create(values ...*model.OmsOrderRefund) error
	CreateInBatches(values []*model.OmsOrderRefund, batchSize int) error
	Save(values ...*model.OmsOrderRefund) error
	First() (*model.OmsOrderRefund, error)
	Take() (*model.OmsOrderRefund, error)
	Last() (*model.OmsOrderRefund, error)
	Find() ([]*model.OmsOrderRefund, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderRefund, err error)
	FindInBatches(result *[]*model.OmsOrderRefund, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OmsOrderRefund) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOmsOrderRefundDo
	Assign(attrs ...field.AssignExpr) IOmsOrderRefundDo
	Joins(fields ...field.RelationField) IOmsOrderRefundDo
	Preload(fields ...field.RelationField) IOmsOrderRefundDo
	FirstOrInit() (*model.OmsOrderRefund, error)
	FirstOrCreate() (*model.OmsOrderRefund, error)
	FindByPage(offset int, limit int) (result []*model.OmsOrderRefund, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOmsOrderRefundDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o omsOrderRefundDo) Debug() IOmsOrderRefundDo {
	return o.withDO(o.DO.Debug())
}

func (o omsOrderRefundDo) WithContext(ctx context.Context) IOmsOrderRefundDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o omsOrderRefundDo) ReadDB() IOmsOrderRefundDo {
	return o.Clauses(dbresolver.Read)
}

func (o omsOrderRefundDo) WriteDB() IOmsOrderRefundDo {
	return o.Clauses(dbresolver.Write)
}

func (o omsOrderRefundDo) Session(config *gorm.Session) IOmsOrderRefundDo {
	return o.withDO(o.DO.Session(config))
}

func (o omsOrderRefundDo) Clauses(conds ...clause.Expression) IOmsOrderRefundDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o omsOrderRefundDo) Returning(value interface{}, columns ...string) IOmsOrderRefundDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o omsOrderRefundDo) Not(conds ...gen.Condition) IOmsOrderRefundDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o omsOrderRefundDo) Or(conds ...gen.Condition) IOmsOrderRefundDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o omsOrderRefundDo) Select(conds ...field.Expr) IOmsOrderRefundDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o omsOrderRefundDo) Where(conds ...gen.Condition) IOmsOrderRefundDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o omsOrderRefundDo) Order(conds ...field.Expr) IOmsOrderRefundDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o omsOrderRefundDo) Distinct(cols ...field.Expr) IOmsOrderRefundDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o omsOrderRefundDo) Omit(cols ...field.Expr) IOmsOrderRefundDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o omsOrderRefundDo) Join(table schema.Tabler, on ...field.Expr) IOmsOrderRefundDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o omsOrderRefundDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderRefundDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o omsOrderRefundDo) RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderRefundDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o omsOrderRefundDo) Group(cols ...field.Expr) IOmsOrderRefundDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o omsOrderRefundDo) Having(conds ...gen.Condition) IOmsOrderRefundDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o omsOrderRefundDo) Limit(limit int) IOmsOrderRefundDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o omsOrderRefundDo) Offset(offset int) IOmsOrderRefundDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o omsOrderRefundDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderRefundDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o omsOrderRefundDo) Unscoped() IOmsOrderRefundDo {
	return o.withDO(o.DO.Unscoped())
}

func (o omsOrderRefundDo) Create(values ...*model.OmsOrderRefund) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o omsOrderRefundDo) CreateInBatches(values []*model.OmsOrderRefund, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o omsOrderRefundDo) Save(values ...*model.OmsOrderRefund) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o omsOrderRefundDo) First() (*model.OmsOrderRefund, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderRefund), nil
	}
}

func (o omsOrderRefundDo) Take() (*model.OmsOrderRefund, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderRefund), nil
	}
}

func (o omsOrderRefundDo) Last() (*model.OmsOrderRefund, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderRefund), nil
	}
}

func (o omsOrderRefundDo) Find() ([]*model.OmsOrderRefund, error) {
	result, err := o.DO.Find()
	return result.([]*model.OmsOrderRefund), err
}

func (o omsOrderRefundDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderRefund, err error) {
	buf := make([]*model.OmsOrderRefund, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o omsOrderRefundDo) FindInBatches(result *[]*model.OmsOrderRefund, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o omsOrderRefundDo) Attrs(attrs ...field.AssignExpr) IOmsOrderRefundDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o omsOrderRefundDo) Assign(attrs ...field.AssignExpr) IOmsOrderRefundDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o omsOrderRefundDo) Joins(fields ...field.RelationField) IOmsOrderRefundDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o omsOrderRefundDo) Preload(fields ...field.RelationField) IOmsOrderRefundDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o omsOrderRefundDo) FirstOrInit() (*model.OmsOrderRefund, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderRefund), nil
	}
}

func (o omsOrderRefundDo) FirstOrCreate() (*model.OmsOrderRefund, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderRefund), nil
	}
}

func (o omsOrderRefundDo) FindByPage(offset int, limit int) (result []*model.OmsOrderRefund, count int64, err error) {
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

func (o omsOrderRefundDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o omsOrderRefundDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o omsOrderRefundDo) Delete(models ...*model.OmsOrderRefund) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *omsOrderRefundDo) withDO(do gen.Dao) *omsOrderRefundDo {
	o.DO = *do.(*gen.DO)
	return o
}
