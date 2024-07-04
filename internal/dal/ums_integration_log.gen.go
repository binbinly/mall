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

func newUmsIntegrationLog(db *gorm.DB, opts ...gen.DOOption) umsIntegrationLog {
	_umsIntegrationLog := umsIntegrationLog{}

	_umsIntegrationLog.umsIntegrationLogDo.UseDB(db, opts...)
	_umsIntegrationLog.umsIntegrationLogDo.UseModel(&model.UmsIntegrationLog{})

	tableName := _umsIntegrationLog.umsIntegrationLogDo.TableName()
	_umsIntegrationLog.ALL = field.NewAsterisk(tableName)
	_umsIntegrationLog.ID = field.NewInt(tableName, "id")
	_umsIntegrationLog.MemberID = field.NewInt(tableName, "member_id")
	_umsIntegrationLog.ChangeCount = field.NewInt64(tableName, "change_count")
	_umsIntegrationLog.Note = field.NewString(tableName, "note")
	_umsIntegrationLog.SourceType = field.NewInt8(tableName, "source_type")
	_umsIntegrationLog.CreatedAt = field.NewInt(tableName, "created_at")

	_umsIntegrationLog.fillFieldMap()

	return _umsIntegrationLog
}

type umsIntegrationLog struct {
	umsIntegrationLogDo

	ALL         field.Asterisk
	ID          field.Int    // ID
	MemberID    field.Int    // 用户id
	ChangeCount field.Int64  // 改变的值（正负计数）
	Note        field.String // 备注
	SourceType  field.Int8   // 积分来源[0-购物，1-管理员修改]
	CreatedAt   field.Int    // 创建时间

	fieldMap map[string]field.Expr
}

func (u umsIntegrationLog) Table(newTableName string) *umsIntegrationLog {
	u.umsIntegrationLogDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u umsIntegrationLog) As(alias string) *umsIntegrationLog {
	u.umsIntegrationLogDo.DO = *(u.umsIntegrationLogDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *umsIntegrationLog) updateTableName(table string) *umsIntegrationLog {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt(table, "id")
	u.MemberID = field.NewInt(table, "member_id")
	u.ChangeCount = field.NewInt64(table, "change_count")
	u.Note = field.NewString(table, "note")
	u.SourceType = field.NewInt8(table, "source_type")
	u.CreatedAt = field.NewInt(table, "created_at")

	u.fillFieldMap()

	return u
}

func (u *umsIntegrationLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *umsIntegrationLog) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["id"] = u.ID
	u.fieldMap["member_id"] = u.MemberID
	u.fieldMap["change_count"] = u.ChangeCount
	u.fieldMap["note"] = u.Note
	u.fieldMap["source_type"] = u.SourceType
	u.fieldMap["created_at"] = u.CreatedAt
}

func (u umsIntegrationLog) clone(db *gorm.DB) umsIntegrationLog {
	u.umsIntegrationLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u umsIntegrationLog) replaceDB(db *gorm.DB) umsIntegrationLog {
	u.umsIntegrationLogDo.ReplaceDB(db)
	return u
}

type umsIntegrationLogDo struct{ gen.DO }

type IUmsIntegrationLogDo interface {
	gen.SubQuery
	Debug() IUmsIntegrationLogDo
	WithContext(ctx context.Context) IUmsIntegrationLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUmsIntegrationLogDo
	WriteDB() IUmsIntegrationLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUmsIntegrationLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUmsIntegrationLogDo
	Not(conds ...gen.Condition) IUmsIntegrationLogDo
	Or(conds ...gen.Condition) IUmsIntegrationLogDo
	Select(conds ...field.Expr) IUmsIntegrationLogDo
	Where(conds ...gen.Condition) IUmsIntegrationLogDo
	Order(conds ...field.Expr) IUmsIntegrationLogDo
	Distinct(cols ...field.Expr) IUmsIntegrationLogDo
	Omit(cols ...field.Expr) IUmsIntegrationLogDo
	Join(table schema.Tabler, on ...field.Expr) IUmsIntegrationLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUmsIntegrationLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUmsIntegrationLogDo
	Group(cols ...field.Expr) IUmsIntegrationLogDo
	Having(conds ...gen.Condition) IUmsIntegrationLogDo
	Limit(limit int) IUmsIntegrationLogDo
	Offset(offset int) IUmsIntegrationLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsIntegrationLogDo
	Unscoped() IUmsIntegrationLogDo
	Create(values ...*model.UmsIntegrationLog) error
	CreateInBatches(values []*model.UmsIntegrationLog, batchSize int) error
	Save(values ...*model.UmsIntegrationLog) error
	First() (*model.UmsIntegrationLog, error)
	Take() (*model.UmsIntegrationLog, error)
	Last() (*model.UmsIntegrationLog, error)
	Find() ([]*model.UmsIntegrationLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsIntegrationLog, err error)
	FindInBatches(result *[]*model.UmsIntegrationLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UmsIntegrationLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUmsIntegrationLogDo
	Assign(attrs ...field.AssignExpr) IUmsIntegrationLogDo
	Joins(fields ...field.RelationField) IUmsIntegrationLogDo
	Preload(fields ...field.RelationField) IUmsIntegrationLogDo
	FirstOrInit() (*model.UmsIntegrationLog, error)
	FirstOrCreate() (*model.UmsIntegrationLog, error)
	FindByPage(offset int, limit int) (result []*model.UmsIntegrationLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUmsIntegrationLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u umsIntegrationLogDo) Debug() IUmsIntegrationLogDo {
	return u.withDO(u.DO.Debug())
}

func (u umsIntegrationLogDo) WithContext(ctx context.Context) IUmsIntegrationLogDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u umsIntegrationLogDo) ReadDB() IUmsIntegrationLogDo {
	return u.Clauses(dbresolver.Read)
}

func (u umsIntegrationLogDo) WriteDB() IUmsIntegrationLogDo {
	return u.Clauses(dbresolver.Write)
}

func (u umsIntegrationLogDo) Session(config *gorm.Session) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Session(config))
}

func (u umsIntegrationLogDo) Clauses(conds ...clause.Expression) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u umsIntegrationLogDo) Returning(value interface{}, columns ...string) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u umsIntegrationLogDo) Not(conds ...gen.Condition) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u umsIntegrationLogDo) Or(conds ...gen.Condition) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u umsIntegrationLogDo) Select(conds ...field.Expr) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u umsIntegrationLogDo) Where(conds ...gen.Condition) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u umsIntegrationLogDo) Order(conds ...field.Expr) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u umsIntegrationLogDo) Distinct(cols ...field.Expr) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u umsIntegrationLogDo) Omit(cols ...field.Expr) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u umsIntegrationLogDo) Join(table schema.Tabler, on ...field.Expr) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u umsIntegrationLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUmsIntegrationLogDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u umsIntegrationLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IUmsIntegrationLogDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u umsIntegrationLogDo) Group(cols ...field.Expr) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u umsIntegrationLogDo) Having(conds ...gen.Condition) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u umsIntegrationLogDo) Limit(limit int) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u umsIntegrationLogDo) Offset(offset int) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u umsIntegrationLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u umsIntegrationLogDo) Unscoped() IUmsIntegrationLogDo {
	return u.withDO(u.DO.Unscoped())
}

func (u umsIntegrationLogDo) Create(values ...*model.UmsIntegrationLog) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u umsIntegrationLogDo) CreateInBatches(values []*model.UmsIntegrationLog, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u umsIntegrationLogDo) Save(values ...*model.UmsIntegrationLog) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u umsIntegrationLogDo) First() (*model.UmsIntegrationLog, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationLog), nil
	}
}

func (u umsIntegrationLogDo) Take() (*model.UmsIntegrationLog, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationLog), nil
	}
}

func (u umsIntegrationLogDo) Last() (*model.UmsIntegrationLog, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationLog), nil
	}
}

func (u umsIntegrationLogDo) Find() ([]*model.UmsIntegrationLog, error) {
	result, err := u.DO.Find()
	return result.([]*model.UmsIntegrationLog), err
}

func (u umsIntegrationLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsIntegrationLog, err error) {
	buf := make([]*model.UmsIntegrationLog, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u umsIntegrationLogDo) FindInBatches(result *[]*model.UmsIntegrationLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u umsIntegrationLogDo) Attrs(attrs ...field.AssignExpr) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u umsIntegrationLogDo) Assign(attrs ...field.AssignExpr) IUmsIntegrationLogDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u umsIntegrationLogDo) Joins(fields ...field.RelationField) IUmsIntegrationLogDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u umsIntegrationLogDo) Preload(fields ...field.RelationField) IUmsIntegrationLogDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u umsIntegrationLogDo) FirstOrInit() (*model.UmsIntegrationLog, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationLog), nil
	}
}

func (u umsIntegrationLogDo) FirstOrCreate() (*model.UmsIntegrationLog, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationLog), nil
	}
}

func (u umsIntegrationLogDo) FindByPage(offset int, limit int) (result []*model.UmsIntegrationLog, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u umsIntegrationLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u umsIntegrationLogDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u umsIntegrationLogDo) Delete(models ...*model.UmsIntegrationLog) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *umsIntegrationLogDo) withDO(do gen.Dao) *umsIntegrationLogDo {
	u.DO = *do.(*gen.DO)
	return u
}
