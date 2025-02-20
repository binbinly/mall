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

func newSmsAppNotice(db *gorm.DB, opts ...gen.DOOption) smsAppNotice {
	_smsAppNotice := smsAppNotice{}

	_smsAppNotice.smsAppNoticeDo.UseDB(db, opts...)
	_smsAppNotice.smsAppNoticeDo.UseModel(&model.SmsAppNotice{})

	tableName := _smsAppNotice.smsAppNoticeDo.TableName()
	_smsAppNotice.ALL = field.NewAsterisk(tableName)
	_smsAppNotice.ID = field.NewInt(tableName, "id")
	_smsAppNotice.Title = field.NewString(tableName, "title")
	_smsAppNotice.Content = field.NewString(tableName, "content")
	_smsAppNotice.CreateBy = field.NewInt(tableName, "create_by")
	_smsAppNotice.UpdateBy = field.NewInt(tableName, "update_by")
	_smsAppNotice.CreatedAt = field.NewInt(tableName, "created_at")
	_smsAppNotice.UpdatedAt = field.NewInt(tableName, "updated_at")

	_smsAppNotice.fillFieldMap()

	return _smsAppNotice
}

type smsAppNotice struct {
	smsAppNoticeDo

	ALL       field.Asterisk
	ID        field.Int    // ID
	Title     field.String // 标题
	Content   field.String // 内容
	CreateBy  field.Int    // 创建者
	UpdateBy  field.Int    // 更新者
	CreatedAt field.Int    // 创建时间
	UpdatedAt field.Int    // 更新时间

	fieldMap map[string]field.Expr
}

func (s smsAppNotice) Table(newTableName string) *smsAppNotice {
	s.smsAppNoticeDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s smsAppNotice) As(alias string) *smsAppNotice {
	s.smsAppNoticeDo.DO = *(s.smsAppNoticeDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *smsAppNotice) updateTableName(table string) *smsAppNotice {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt(table, "id")
	s.Title = field.NewString(table, "title")
	s.Content = field.NewString(table, "content")
	s.CreateBy = field.NewInt(table, "create_by")
	s.UpdateBy = field.NewInt(table, "update_by")
	s.CreatedAt = field.NewInt(table, "created_at")
	s.UpdatedAt = field.NewInt(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *smsAppNotice) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *smsAppNotice) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["id"] = s.ID
	s.fieldMap["title"] = s.Title
	s.fieldMap["content"] = s.Content
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s smsAppNotice) clone(db *gorm.DB) smsAppNotice {
	s.smsAppNoticeDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s smsAppNotice) replaceDB(db *gorm.DB) smsAppNotice {
	s.smsAppNoticeDo.ReplaceDB(db)
	return s
}

type smsAppNoticeDo struct{ gen.DO }

type ISmsAppNoticeDo interface {
	gen.SubQuery
	Debug() ISmsAppNoticeDo
	WithContext(ctx context.Context) ISmsAppNoticeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISmsAppNoticeDo
	WriteDB() ISmsAppNoticeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISmsAppNoticeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISmsAppNoticeDo
	Not(conds ...gen.Condition) ISmsAppNoticeDo
	Or(conds ...gen.Condition) ISmsAppNoticeDo
	Select(conds ...field.Expr) ISmsAppNoticeDo
	Where(conds ...gen.Condition) ISmsAppNoticeDo
	Order(conds ...field.Expr) ISmsAppNoticeDo
	Distinct(cols ...field.Expr) ISmsAppNoticeDo
	Omit(cols ...field.Expr) ISmsAppNoticeDo
	Join(table schema.Tabler, on ...field.Expr) ISmsAppNoticeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISmsAppNoticeDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISmsAppNoticeDo
	Group(cols ...field.Expr) ISmsAppNoticeDo
	Having(conds ...gen.Condition) ISmsAppNoticeDo
	Limit(limit int) ISmsAppNoticeDo
	Offset(offset int) ISmsAppNoticeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsAppNoticeDo
	Unscoped() ISmsAppNoticeDo
	Create(values ...*model.SmsAppNotice) error
	CreateInBatches(values []*model.SmsAppNotice, batchSize int) error
	Save(values ...*model.SmsAppNotice) error
	First() (*model.SmsAppNotice, error)
	Take() (*model.SmsAppNotice, error)
	Last() (*model.SmsAppNotice, error)
	Find() ([]*model.SmsAppNotice, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsAppNotice, err error)
	FindInBatches(result *[]*model.SmsAppNotice, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SmsAppNotice) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISmsAppNoticeDo
	Assign(attrs ...field.AssignExpr) ISmsAppNoticeDo
	Joins(fields ...field.RelationField) ISmsAppNoticeDo
	Preload(fields ...field.RelationField) ISmsAppNoticeDo
	FirstOrInit() (*model.SmsAppNotice, error)
	FirstOrCreate() (*model.SmsAppNotice, error)
	FindByPage(offset int, limit int) (result []*model.SmsAppNotice, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISmsAppNoticeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s smsAppNoticeDo) Debug() ISmsAppNoticeDo {
	return s.withDO(s.DO.Debug())
}

func (s smsAppNoticeDo) WithContext(ctx context.Context) ISmsAppNoticeDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s smsAppNoticeDo) ReadDB() ISmsAppNoticeDo {
	return s.Clauses(dbresolver.Read)
}

func (s smsAppNoticeDo) WriteDB() ISmsAppNoticeDo {
	return s.Clauses(dbresolver.Write)
}

func (s smsAppNoticeDo) Session(config *gorm.Session) ISmsAppNoticeDo {
	return s.withDO(s.DO.Session(config))
}

func (s smsAppNoticeDo) Clauses(conds ...clause.Expression) ISmsAppNoticeDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s smsAppNoticeDo) Returning(value interface{}, columns ...string) ISmsAppNoticeDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s smsAppNoticeDo) Not(conds ...gen.Condition) ISmsAppNoticeDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s smsAppNoticeDo) Or(conds ...gen.Condition) ISmsAppNoticeDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s smsAppNoticeDo) Select(conds ...field.Expr) ISmsAppNoticeDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s smsAppNoticeDo) Where(conds ...gen.Condition) ISmsAppNoticeDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s smsAppNoticeDo) Order(conds ...field.Expr) ISmsAppNoticeDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s smsAppNoticeDo) Distinct(cols ...field.Expr) ISmsAppNoticeDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s smsAppNoticeDo) Omit(cols ...field.Expr) ISmsAppNoticeDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s smsAppNoticeDo) Join(table schema.Tabler, on ...field.Expr) ISmsAppNoticeDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s smsAppNoticeDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISmsAppNoticeDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s smsAppNoticeDo) RightJoin(table schema.Tabler, on ...field.Expr) ISmsAppNoticeDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s smsAppNoticeDo) Group(cols ...field.Expr) ISmsAppNoticeDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s smsAppNoticeDo) Having(conds ...gen.Condition) ISmsAppNoticeDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s smsAppNoticeDo) Limit(limit int) ISmsAppNoticeDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s smsAppNoticeDo) Offset(offset int) ISmsAppNoticeDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s smsAppNoticeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsAppNoticeDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s smsAppNoticeDo) Unscoped() ISmsAppNoticeDo {
	return s.withDO(s.DO.Unscoped())
}

func (s smsAppNoticeDo) Create(values ...*model.SmsAppNotice) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s smsAppNoticeDo) CreateInBatches(values []*model.SmsAppNotice, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s smsAppNoticeDo) Save(values ...*model.SmsAppNotice) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s smsAppNoticeDo) First() (*model.SmsAppNotice, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsAppNotice), nil
	}
}

func (s smsAppNoticeDo) Take() (*model.SmsAppNotice, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsAppNotice), nil
	}
}

func (s smsAppNoticeDo) Last() (*model.SmsAppNotice, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsAppNotice), nil
	}
}

func (s smsAppNoticeDo) Find() ([]*model.SmsAppNotice, error) {
	result, err := s.DO.Find()
	return result.([]*model.SmsAppNotice), err
}

func (s smsAppNoticeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsAppNotice, err error) {
	buf := make([]*model.SmsAppNotice, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s smsAppNoticeDo) FindInBatches(result *[]*model.SmsAppNotice, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s smsAppNoticeDo) Attrs(attrs ...field.AssignExpr) ISmsAppNoticeDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s smsAppNoticeDo) Assign(attrs ...field.AssignExpr) ISmsAppNoticeDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s smsAppNoticeDo) Joins(fields ...field.RelationField) ISmsAppNoticeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s smsAppNoticeDo) Preload(fields ...field.RelationField) ISmsAppNoticeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s smsAppNoticeDo) FirstOrInit() (*model.SmsAppNotice, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsAppNotice), nil
	}
}

func (s smsAppNoticeDo) FirstOrCreate() (*model.SmsAppNotice, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsAppNotice), nil
	}
}

func (s smsAppNoticeDo) FindByPage(offset int, limit int) (result []*model.SmsAppNotice, count int64, err error) {
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

func (s smsAppNoticeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s smsAppNoticeDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s smsAppNoticeDo) Delete(models ...*model.SmsAppNotice) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *smsAppNoticeDo) withDO(do gen.Dao) *smsAppNoticeDo {
	s.DO = *do.(*gen.DO)
	return s
}
