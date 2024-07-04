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

func newSmsSeckillActivity(db *gorm.DB, opts ...gen.DOOption) smsSeckillActivity {
	_smsSeckillActivity := smsSeckillActivity{}

	_smsSeckillActivity.smsSeckillActivityDo.UseDB(db, opts...)
	_smsSeckillActivity.smsSeckillActivityDo.UseModel(&model.SmsSeckillActivity{})

	tableName := _smsSeckillActivity.smsSeckillActivityDo.TableName()
	_smsSeckillActivity.ALL = field.NewAsterisk(tableName)
	_smsSeckillActivity.ID = field.NewInt(tableName, "id")
	_smsSeckillActivity.Title = field.NewString(tableName, "title")
	_smsSeckillActivity.Cover = field.NewString(tableName, "cover")
	_smsSeckillActivity.StartAt = field.NewInt64(tableName, "start_at")
	_smsSeckillActivity.EndAt = field.NewInt64(tableName, "end_at")
	_smsSeckillActivity.IsRelease = field.NewInt8(tableName, "is_release")
	_smsSeckillActivity.CreateBy = field.NewInt(tableName, "create_by")
	_smsSeckillActivity.UpdateBy = field.NewInt(tableName, "update_by")
	_smsSeckillActivity.CreatedAt = field.NewInt(tableName, "created_at")
	_smsSeckillActivity.UpdatedAt = field.NewInt(tableName, "updated_at")

	_smsSeckillActivity.fillFieldMap()

	return _smsSeckillActivity
}

type smsSeckillActivity struct {
	smsSeckillActivityDo

	ALL       field.Asterisk
	ID        field.Int    // ID
	Title     field.String // 标题
	Cover     field.String // 活动封面
	StartAt   field.Int64  // 开始时间
	EndAt     field.Int64  // 结束时间
	IsRelease field.Int8   // 是否发布上线
	CreateBy  field.Int    // 创建者
	UpdateBy  field.Int    // 更新者
	CreatedAt field.Int    // 创建时间
	UpdatedAt field.Int    // 更新时间

	fieldMap map[string]field.Expr
}

func (s smsSeckillActivity) Table(newTableName string) *smsSeckillActivity {
	s.smsSeckillActivityDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s smsSeckillActivity) As(alias string) *smsSeckillActivity {
	s.smsSeckillActivityDo.DO = *(s.smsSeckillActivityDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *smsSeckillActivity) updateTableName(table string) *smsSeckillActivity {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt(table, "id")
	s.Title = field.NewString(table, "title")
	s.Cover = field.NewString(table, "cover")
	s.StartAt = field.NewInt64(table, "start_at")
	s.EndAt = field.NewInt64(table, "end_at")
	s.IsRelease = field.NewInt8(table, "is_release")
	s.CreateBy = field.NewInt(table, "create_by")
	s.UpdateBy = field.NewInt(table, "update_by")
	s.CreatedAt = field.NewInt(table, "created_at")
	s.UpdatedAt = field.NewInt(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *smsSeckillActivity) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *smsSeckillActivity) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 10)
	s.fieldMap["id"] = s.ID
	s.fieldMap["title"] = s.Title
	s.fieldMap["cover"] = s.Cover
	s.fieldMap["start_at"] = s.StartAt
	s.fieldMap["end_at"] = s.EndAt
	s.fieldMap["is_release"] = s.IsRelease
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s smsSeckillActivity) clone(db *gorm.DB) smsSeckillActivity {
	s.smsSeckillActivityDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s smsSeckillActivity) replaceDB(db *gorm.DB) smsSeckillActivity {
	s.smsSeckillActivityDo.ReplaceDB(db)
	return s
}

type smsSeckillActivityDo struct{ gen.DO }

type ISmsSeckillActivityDo interface {
	gen.SubQuery
	Debug() ISmsSeckillActivityDo
	WithContext(ctx context.Context) ISmsSeckillActivityDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISmsSeckillActivityDo
	WriteDB() ISmsSeckillActivityDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISmsSeckillActivityDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISmsSeckillActivityDo
	Not(conds ...gen.Condition) ISmsSeckillActivityDo
	Or(conds ...gen.Condition) ISmsSeckillActivityDo
	Select(conds ...field.Expr) ISmsSeckillActivityDo
	Where(conds ...gen.Condition) ISmsSeckillActivityDo
	Order(conds ...field.Expr) ISmsSeckillActivityDo
	Distinct(cols ...field.Expr) ISmsSeckillActivityDo
	Omit(cols ...field.Expr) ISmsSeckillActivityDo
	Join(table schema.Tabler, on ...field.Expr) ISmsSeckillActivityDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISmsSeckillActivityDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISmsSeckillActivityDo
	Group(cols ...field.Expr) ISmsSeckillActivityDo
	Having(conds ...gen.Condition) ISmsSeckillActivityDo
	Limit(limit int) ISmsSeckillActivityDo
	Offset(offset int) ISmsSeckillActivityDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsSeckillActivityDo
	Unscoped() ISmsSeckillActivityDo
	Create(values ...*model.SmsSeckillActivity) error
	CreateInBatches(values []*model.SmsSeckillActivity, batchSize int) error
	Save(values ...*model.SmsSeckillActivity) error
	First() (*model.SmsSeckillActivity, error)
	Take() (*model.SmsSeckillActivity, error)
	Last() (*model.SmsSeckillActivity, error)
	Find() ([]*model.SmsSeckillActivity, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsSeckillActivity, err error)
	FindInBatches(result *[]*model.SmsSeckillActivity, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SmsSeckillActivity) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISmsSeckillActivityDo
	Assign(attrs ...field.AssignExpr) ISmsSeckillActivityDo
	Joins(fields ...field.RelationField) ISmsSeckillActivityDo
	Preload(fields ...field.RelationField) ISmsSeckillActivityDo
	FirstOrInit() (*model.SmsSeckillActivity, error)
	FirstOrCreate() (*model.SmsSeckillActivity, error)
	FindByPage(offset int, limit int) (result []*model.SmsSeckillActivity, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISmsSeckillActivityDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s smsSeckillActivityDo) Debug() ISmsSeckillActivityDo {
	return s.withDO(s.DO.Debug())
}

func (s smsSeckillActivityDo) WithContext(ctx context.Context) ISmsSeckillActivityDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s smsSeckillActivityDo) ReadDB() ISmsSeckillActivityDo {
	return s.Clauses(dbresolver.Read)
}

func (s smsSeckillActivityDo) WriteDB() ISmsSeckillActivityDo {
	return s.Clauses(dbresolver.Write)
}

func (s smsSeckillActivityDo) Session(config *gorm.Session) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Session(config))
}

func (s smsSeckillActivityDo) Clauses(conds ...clause.Expression) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s smsSeckillActivityDo) Returning(value interface{}, columns ...string) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s smsSeckillActivityDo) Not(conds ...gen.Condition) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s smsSeckillActivityDo) Or(conds ...gen.Condition) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s smsSeckillActivityDo) Select(conds ...field.Expr) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s smsSeckillActivityDo) Where(conds ...gen.Condition) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s smsSeckillActivityDo) Order(conds ...field.Expr) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s smsSeckillActivityDo) Distinct(cols ...field.Expr) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s smsSeckillActivityDo) Omit(cols ...field.Expr) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s smsSeckillActivityDo) Join(table schema.Tabler, on ...field.Expr) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s smsSeckillActivityDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISmsSeckillActivityDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s smsSeckillActivityDo) RightJoin(table schema.Tabler, on ...field.Expr) ISmsSeckillActivityDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s smsSeckillActivityDo) Group(cols ...field.Expr) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s smsSeckillActivityDo) Having(conds ...gen.Condition) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s smsSeckillActivityDo) Limit(limit int) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s smsSeckillActivityDo) Offset(offset int) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s smsSeckillActivityDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s smsSeckillActivityDo) Unscoped() ISmsSeckillActivityDo {
	return s.withDO(s.DO.Unscoped())
}

func (s smsSeckillActivityDo) Create(values ...*model.SmsSeckillActivity) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s smsSeckillActivityDo) CreateInBatches(values []*model.SmsSeckillActivity, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s smsSeckillActivityDo) Save(values ...*model.SmsSeckillActivity) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s smsSeckillActivityDo) First() (*model.SmsSeckillActivity, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsSeckillActivity), nil
	}
}

func (s smsSeckillActivityDo) Take() (*model.SmsSeckillActivity, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsSeckillActivity), nil
	}
}

func (s smsSeckillActivityDo) Last() (*model.SmsSeckillActivity, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsSeckillActivity), nil
	}
}

func (s smsSeckillActivityDo) Find() ([]*model.SmsSeckillActivity, error) {
	result, err := s.DO.Find()
	return result.([]*model.SmsSeckillActivity), err
}

func (s smsSeckillActivityDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsSeckillActivity, err error) {
	buf := make([]*model.SmsSeckillActivity, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s smsSeckillActivityDo) FindInBatches(result *[]*model.SmsSeckillActivity, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s smsSeckillActivityDo) Attrs(attrs ...field.AssignExpr) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s smsSeckillActivityDo) Assign(attrs ...field.AssignExpr) ISmsSeckillActivityDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s smsSeckillActivityDo) Joins(fields ...field.RelationField) ISmsSeckillActivityDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s smsSeckillActivityDo) Preload(fields ...field.RelationField) ISmsSeckillActivityDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s smsSeckillActivityDo) FirstOrInit() (*model.SmsSeckillActivity, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsSeckillActivity), nil
	}
}

func (s smsSeckillActivityDo) FirstOrCreate() (*model.SmsSeckillActivity, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsSeckillActivity), nil
	}
}

func (s smsSeckillActivityDo) FindByPage(offset int, limit int) (result []*model.SmsSeckillActivity, count int64, err error) {
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

func (s smsSeckillActivityDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s smsSeckillActivityDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s smsSeckillActivityDo) Delete(models ...*model.SmsSeckillActivity) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *smsSeckillActivityDo) withDO(do gen.Dao) *smsSeckillActivityDo {
	s.DO = *do.(*gen.DO)
	return s
}
