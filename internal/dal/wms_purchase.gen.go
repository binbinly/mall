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

func newWmsPurchase(db *gorm.DB, opts ...gen.DOOption) wmsPurchase {
	_wmsPurchase := wmsPurchase{}

	_wmsPurchase.wmsPurchaseDo.UseDB(db, opts...)
	_wmsPurchase.wmsPurchaseDo.UseModel(&model.WmsPurchase{})

	tableName := _wmsPurchase.wmsPurchaseDo.TableName()
	_wmsPurchase.ALL = field.NewAsterisk(tableName)
	_wmsPurchase.ID = field.NewInt(tableName, "id")
	_wmsPurchase.AssigneeID = field.NewInt(tableName, "assignee_id")
	_wmsPurchase.AssigneeName = field.NewString(tableName, "assignee_name")
	_wmsPurchase.Phone = field.NewString(tableName, "phone")
	_wmsPurchase.Priority = field.NewInt8(tableName, "priority")
	_wmsPurchase.Status = field.NewInt8(tableName, "status")
	_wmsPurchase.WareID = field.NewInt(tableName, "ware_id")
	_wmsPurchase.Amount = field.NewInt(tableName, "amount")
	_wmsPurchase.CreatedAt = field.NewTime(tableName, "created_at")
	_wmsPurchase.UpdatedAt = field.NewTime(tableName, "updated_at")

	_wmsPurchase.fillFieldMap()

	return _wmsPurchase
}

type wmsPurchase struct {
	wmsPurchaseDo

	ALL          field.Asterisk
	ID           field.Int    // ID
	AssigneeID   field.Int    // 采购人id
	AssigneeName field.String // 采购人
	Phone        field.String // 手机号
	Priority     field.Int8   // 优先级
	Status       field.Int8   // 状态
	WareID       field.Int    // 仓库id
	Amount       field.Int    // 总金额/分
	CreatedAt    field.Time   // 创建时间
	UpdatedAt    field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (w wmsPurchase) Table(newTableName string) *wmsPurchase {
	w.wmsPurchaseDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wmsPurchase) As(alias string) *wmsPurchase {
	w.wmsPurchaseDo.DO = *(w.wmsPurchaseDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wmsPurchase) updateTableName(table string) *wmsPurchase {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt(table, "id")
	w.AssigneeID = field.NewInt(table, "assignee_id")
	w.AssigneeName = field.NewString(table, "assignee_name")
	w.Phone = field.NewString(table, "phone")
	w.Priority = field.NewInt8(table, "priority")
	w.Status = field.NewInt8(table, "status")
	w.WareID = field.NewInt(table, "ware_id")
	w.Amount = field.NewInt(table, "amount")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *wmsPurchase) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wmsPurchase) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 10)
	w.fieldMap["id"] = w.ID
	w.fieldMap["assignee_id"] = w.AssigneeID
	w.fieldMap["assignee_name"] = w.AssigneeName
	w.fieldMap["phone"] = w.Phone
	w.fieldMap["priority"] = w.Priority
	w.fieldMap["status"] = w.Status
	w.fieldMap["ware_id"] = w.WareID
	w.fieldMap["amount"] = w.Amount
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w wmsPurchase) clone(db *gorm.DB) wmsPurchase {
	w.wmsPurchaseDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wmsPurchase) replaceDB(db *gorm.DB) wmsPurchase {
	w.wmsPurchaseDo.ReplaceDB(db)
	return w
}

type wmsPurchaseDo struct{ gen.DO }

type IWmsPurchaseDo interface {
	gen.SubQuery
	Debug() IWmsPurchaseDo
	WithContext(ctx context.Context) IWmsPurchaseDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWmsPurchaseDo
	WriteDB() IWmsPurchaseDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWmsPurchaseDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWmsPurchaseDo
	Not(conds ...gen.Condition) IWmsPurchaseDo
	Or(conds ...gen.Condition) IWmsPurchaseDo
	Select(conds ...field.Expr) IWmsPurchaseDo
	Where(conds ...gen.Condition) IWmsPurchaseDo
	Order(conds ...field.Expr) IWmsPurchaseDo
	Distinct(cols ...field.Expr) IWmsPurchaseDo
	Omit(cols ...field.Expr) IWmsPurchaseDo
	Join(table schema.Tabler, on ...field.Expr) IWmsPurchaseDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWmsPurchaseDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWmsPurchaseDo
	Group(cols ...field.Expr) IWmsPurchaseDo
	Having(conds ...gen.Condition) IWmsPurchaseDo
	Limit(limit int) IWmsPurchaseDo
	Offset(offset int) IWmsPurchaseDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWmsPurchaseDo
	Unscoped() IWmsPurchaseDo
	Create(values ...*model.WmsPurchase) error
	CreateInBatches(values []*model.WmsPurchase, batchSize int) error
	Save(values ...*model.WmsPurchase) error
	First() (*model.WmsPurchase, error)
	Take() (*model.WmsPurchase, error)
	Last() (*model.WmsPurchase, error)
	Find() ([]*model.WmsPurchase, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WmsPurchase, err error)
	FindInBatches(result *[]*model.WmsPurchase, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WmsPurchase) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWmsPurchaseDo
	Assign(attrs ...field.AssignExpr) IWmsPurchaseDo
	Joins(fields ...field.RelationField) IWmsPurchaseDo
	Preload(fields ...field.RelationField) IWmsPurchaseDo
	FirstOrInit() (*model.WmsPurchase, error)
	FirstOrCreate() (*model.WmsPurchase, error)
	FindByPage(offset int, limit int) (result []*model.WmsPurchase, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWmsPurchaseDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w wmsPurchaseDo) Debug() IWmsPurchaseDo {
	return w.withDO(w.DO.Debug())
}

func (w wmsPurchaseDo) WithContext(ctx context.Context) IWmsPurchaseDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wmsPurchaseDo) ReadDB() IWmsPurchaseDo {
	return w.Clauses(dbresolver.Read)
}

func (w wmsPurchaseDo) WriteDB() IWmsPurchaseDo {
	return w.Clauses(dbresolver.Write)
}

func (w wmsPurchaseDo) Session(config *gorm.Session) IWmsPurchaseDo {
	return w.withDO(w.DO.Session(config))
}

func (w wmsPurchaseDo) Clauses(conds ...clause.Expression) IWmsPurchaseDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wmsPurchaseDo) Returning(value interface{}, columns ...string) IWmsPurchaseDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wmsPurchaseDo) Not(conds ...gen.Condition) IWmsPurchaseDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wmsPurchaseDo) Or(conds ...gen.Condition) IWmsPurchaseDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wmsPurchaseDo) Select(conds ...field.Expr) IWmsPurchaseDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wmsPurchaseDo) Where(conds ...gen.Condition) IWmsPurchaseDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wmsPurchaseDo) Order(conds ...field.Expr) IWmsPurchaseDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wmsPurchaseDo) Distinct(cols ...field.Expr) IWmsPurchaseDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wmsPurchaseDo) Omit(cols ...field.Expr) IWmsPurchaseDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wmsPurchaseDo) Join(table schema.Tabler, on ...field.Expr) IWmsPurchaseDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wmsPurchaseDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWmsPurchaseDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wmsPurchaseDo) RightJoin(table schema.Tabler, on ...field.Expr) IWmsPurchaseDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wmsPurchaseDo) Group(cols ...field.Expr) IWmsPurchaseDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wmsPurchaseDo) Having(conds ...gen.Condition) IWmsPurchaseDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wmsPurchaseDo) Limit(limit int) IWmsPurchaseDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wmsPurchaseDo) Offset(offset int) IWmsPurchaseDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wmsPurchaseDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWmsPurchaseDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wmsPurchaseDo) Unscoped() IWmsPurchaseDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wmsPurchaseDo) Create(values ...*model.WmsPurchase) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wmsPurchaseDo) CreateInBatches(values []*model.WmsPurchase, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wmsPurchaseDo) Save(values ...*model.WmsPurchase) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wmsPurchaseDo) First() (*model.WmsPurchase, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WmsPurchase), nil
	}
}

func (w wmsPurchaseDo) Take() (*model.WmsPurchase, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WmsPurchase), nil
	}
}

func (w wmsPurchaseDo) Last() (*model.WmsPurchase, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WmsPurchase), nil
	}
}

func (w wmsPurchaseDo) Find() ([]*model.WmsPurchase, error) {
	result, err := w.DO.Find()
	return result.([]*model.WmsPurchase), err
}

func (w wmsPurchaseDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WmsPurchase, err error) {
	buf := make([]*model.WmsPurchase, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wmsPurchaseDo) FindInBatches(result *[]*model.WmsPurchase, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wmsPurchaseDo) Attrs(attrs ...field.AssignExpr) IWmsPurchaseDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wmsPurchaseDo) Assign(attrs ...field.AssignExpr) IWmsPurchaseDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wmsPurchaseDo) Joins(fields ...field.RelationField) IWmsPurchaseDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wmsPurchaseDo) Preload(fields ...field.RelationField) IWmsPurchaseDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wmsPurchaseDo) FirstOrInit() (*model.WmsPurchase, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WmsPurchase), nil
	}
}

func (w wmsPurchaseDo) FirstOrCreate() (*model.WmsPurchase, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WmsPurchase), nil
	}
}

func (w wmsPurchaseDo) FindByPage(offset int, limit int) (result []*model.WmsPurchase, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w wmsPurchaseDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wmsPurchaseDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wmsPurchaseDo) Delete(models ...*model.WmsPurchase) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wmsPurchaseDo) withDO(do gen.Dao) *wmsPurchaseDo {
	w.DO = *do.(*gen.DO)
	return w
}
