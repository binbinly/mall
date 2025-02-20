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

func newSmsHomeAdv(db *gorm.DB, opts ...gen.DOOption) smsHomeAdv {
	_smsHomeAdv := smsHomeAdv{}

	_smsHomeAdv.smsHomeAdvDo.UseDB(db, opts...)
	_smsHomeAdv.smsHomeAdvDo.UseModel(&model.SmsHomeAdv{})

	tableName := _smsHomeAdv.smsHomeAdvDo.TableName()
	_smsHomeAdv.ALL = field.NewAsterisk(tableName)
	_smsHomeAdv.ID = field.NewInt(tableName, "id")
	_smsHomeAdv.Title = field.NewString(tableName, "title")
	_smsHomeAdv.Img = field.NewString(tableName, "img")
	_smsHomeAdv.StartAt = field.NewInt64(tableName, "start_at")
	_smsHomeAdv.EndAt = field.NewInt64(tableName, "end_at")
	_smsHomeAdv.ClickCount = field.NewInt64(tableName, "click_count")
	_smsHomeAdv.URL = field.NewString(tableName, "url")
	_smsHomeAdv.Note = field.NewString(tableName, "note")
	_smsHomeAdv.IsRelease = field.NewInt8(tableName, "is_release")
	_smsHomeAdv.Sort = field.NewInt(tableName, "sort")
	_smsHomeAdv.CreateBy = field.NewInt(tableName, "create_by")
	_smsHomeAdv.UpdateBy = field.NewInt(tableName, "update_by")
	_smsHomeAdv.CreatedAt = field.NewInt(tableName, "created_at")
	_smsHomeAdv.UpdatedAt = field.NewInt(tableName, "updated_at")

	_smsHomeAdv.fillFieldMap()

	return _smsHomeAdv
}

type smsHomeAdv struct {
	smsHomeAdvDo

	ALL        field.Asterisk
	ID         field.Int    // ID
	Title      field.String // 标题
	Img        field.String // 图片
	StartAt    field.Int64  // 开始时间
	EndAt      field.Int64  // 结束时间
	ClickCount field.Int64  // 点击数
	URL        field.String // 链接地址
	Note       field.String // 备注
	IsRelease  field.Int8   // 是否发布上线
	Sort       field.Int    // 排序
	CreateBy   field.Int    // 创建者
	UpdateBy   field.Int    // 更新者
	CreatedAt  field.Int    // 创建时间
	UpdatedAt  field.Int    // 更新时间

	fieldMap map[string]field.Expr
}

func (s smsHomeAdv) Table(newTableName string) *smsHomeAdv {
	s.smsHomeAdvDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s smsHomeAdv) As(alias string) *smsHomeAdv {
	s.smsHomeAdvDo.DO = *(s.smsHomeAdvDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *smsHomeAdv) updateTableName(table string) *smsHomeAdv {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt(table, "id")
	s.Title = field.NewString(table, "title")
	s.Img = field.NewString(table, "img")
	s.StartAt = field.NewInt64(table, "start_at")
	s.EndAt = field.NewInt64(table, "end_at")
	s.ClickCount = field.NewInt64(table, "click_count")
	s.URL = field.NewString(table, "url")
	s.Note = field.NewString(table, "note")
	s.IsRelease = field.NewInt8(table, "is_release")
	s.Sort = field.NewInt(table, "sort")
	s.CreateBy = field.NewInt(table, "create_by")
	s.UpdateBy = field.NewInt(table, "update_by")
	s.CreatedAt = field.NewInt(table, "created_at")
	s.UpdatedAt = field.NewInt(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *smsHomeAdv) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *smsHomeAdv) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["id"] = s.ID
	s.fieldMap["title"] = s.Title
	s.fieldMap["img"] = s.Img
	s.fieldMap["start_at"] = s.StartAt
	s.fieldMap["end_at"] = s.EndAt
	s.fieldMap["click_count"] = s.ClickCount
	s.fieldMap["url"] = s.URL
	s.fieldMap["note"] = s.Note
	s.fieldMap["is_release"] = s.IsRelease
	s.fieldMap["sort"] = s.Sort
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s smsHomeAdv) clone(db *gorm.DB) smsHomeAdv {
	s.smsHomeAdvDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s smsHomeAdv) replaceDB(db *gorm.DB) smsHomeAdv {
	s.smsHomeAdvDo.ReplaceDB(db)
	return s
}

type smsHomeAdvDo struct{ gen.DO }

type ISmsHomeAdvDo interface {
	gen.SubQuery
	Debug() ISmsHomeAdvDo
	WithContext(ctx context.Context) ISmsHomeAdvDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISmsHomeAdvDo
	WriteDB() ISmsHomeAdvDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISmsHomeAdvDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISmsHomeAdvDo
	Not(conds ...gen.Condition) ISmsHomeAdvDo
	Or(conds ...gen.Condition) ISmsHomeAdvDo
	Select(conds ...field.Expr) ISmsHomeAdvDo
	Where(conds ...gen.Condition) ISmsHomeAdvDo
	Order(conds ...field.Expr) ISmsHomeAdvDo
	Distinct(cols ...field.Expr) ISmsHomeAdvDo
	Omit(cols ...field.Expr) ISmsHomeAdvDo
	Join(table schema.Tabler, on ...field.Expr) ISmsHomeAdvDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISmsHomeAdvDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISmsHomeAdvDo
	Group(cols ...field.Expr) ISmsHomeAdvDo
	Having(conds ...gen.Condition) ISmsHomeAdvDo
	Limit(limit int) ISmsHomeAdvDo
	Offset(offset int) ISmsHomeAdvDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsHomeAdvDo
	Unscoped() ISmsHomeAdvDo
	Create(values ...*model.SmsHomeAdv) error
	CreateInBatches(values []*model.SmsHomeAdv, batchSize int) error
	Save(values ...*model.SmsHomeAdv) error
	First() (*model.SmsHomeAdv, error)
	Take() (*model.SmsHomeAdv, error)
	Last() (*model.SmsHomeAdv, error)
	Find() ([]*model.SmsHomeAdv, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsHomeAdv, err error)
	FindInBatches(result *[]*model.SmsHomeAdv, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SmsHomeAdv) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISmsHomeAdvDo
	Assign(attrs ...field.AssignExpr) ISmsHomeAdvDo
	Joins(fields ...field.RelationField) ISmsHomeAdvDo
	Preload(fields ...field.RelationField) ISmsHomeAdvDo
	FirstOrInit() (*model.SmsHomeAdv, error)
	FirstOrCreate() (*model.SmsHomeAdv, error)
	FindByPage(offset int, limit int) (result []*model.SmsHomeAdv, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISmsHomeAdvDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s smsHomeAdvDo) Debug() ISmsHomeAdvDo {
	return s.withDO(s.DO.Debug())
}

func (s smsHomeAdvDo) WithContext(ctx context.Context) ISmsHomeAdvDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s smsHomeAdvDo) ReadDB() ISmsHomeAdvDo {
	return s.Clauses(dbresolver.Read)
}

func (s smsHomeAdvDo) WriteDB() ISmsHomeAdvDo {
	return s.Clauses(dbresolver.Write)
}

func (s smsHomeAdvDo) Session(config *gorm.Session) ISmsHomeAdvDo {
	return s.withDO(s.DO.Session(config))
}

func (s smsHomeAdvDo) Clauses(conds ...clause.Expression) ISmsHomeAdvDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s smsHomeAdvDo) Returning(value interface{}, columns ...string) ISmsHomeAdvDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s smsHomeAdvDo) Not(conds ...gen.Condition) ISmsHomeAdvDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s smsHomeAdvDo) Or(conds ...gen.Condition) ISmsHomeAdvDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s smsHomeAdvDo) Select(conds ...field.Expr) ISmsHomeAdvDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s smsHomeAdvDo) Where(conds ...gen.Condition) ISmsHomeAdvDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s smsHomeAdvDo) Order(conds ...field.Expr) ISmsHomeAdvDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s smsHomeAdvDo) Distinct(cols ...field.Expr) ISmsHomeAdvDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s smsHomeAdvDo) Omit(cols ...field.Expr) ISmsHomeAdvDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s smsHomeAdvDo) Join(table schema.Tabler, on ...field.Expr) ISmsHomeAdvDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s smsHomeAdvDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISmsHomeAdvDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s smsHomeAdvDo) RightJoin(table schema.Tabler, on ...field.Expr) ISmsHomeAdvDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s smsHomeAdvDo) Group(cols ...field.Expr) ISmsHomeAdvDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s smsHomeAdvDo) Having(conds ...gen.Condition) ISmsHomeAdvDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s smsHomeAdvDo) Limit(limit int) ISmsHomeAdvDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s smsHomeAdvDo) Offset(offset int) ISmsHomeAdvDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s smsHomeAdvDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsHomeAdvDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s smsHomeAdvDo) Unscoped() ISmsHomeAdvDo {
	return s.withDO(s.DO.Unscoped())
}

func (s smsHomeAdvDo) Create(values ...*model.SmsHomeAdv) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s smsHomeAdvDo) CreateInBatches(values []*model.SmsHomeAdv, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s smsHomeAdvDo) Save(values ...*model.SmsHomeAdv) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s smsHomeAdvDo) First() (*model.SmsHomeAdv, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeAdv), nil
	}
}

func (s smsHomeAdvDo) Take() (*model.SmsHomeAdv, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeAdv), nil
	}
}

func (s smsHomeAdvDo) Last() (*model.SmsHomeAdv, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeAdv), nil
	}
}

func (s smsHomeAdvDo) Find() ([]*model.SmsHomeAdv, error) {
	result, err := s.DO.Find()
	return result.([]*model.SmsHomeAdv), err
}

func (s smsHomeAdvDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsHomeAdv, err error) {
	buf := make([]*model.SmsHomeAdv, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s smsHomeAdvDo) FindInBatches(result *[]*model.SmsHomeAdv, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s smsHomeAdvDo) Attrs(attrs ...field.AssignExpr) ISmsHomeAdvDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s smsHomeAdvDo) Assign(attrs ...field.AssignExpr) ISmsHomeAdvDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s smsHomeAdvDo) Joins(fields ...field.RelationField) ISmsHomeAdvDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s smsHomeAdvDo) Preload(fields ...field.RelationField) ISmsHomeAdvDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s smsHomeAdvDo) FirstOrInit() (*model.SmsHomeAdv, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeAdv), nil
	}
}

func (s smsHomeAdvDo) FirstOrCreate() (*model.SmsHomeAdv, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsHomeAdv), nil
	}
}

func (s smsHomeAdvDo) FindByPage(offset int, limit int) (result []*model.SmsHomeAdv, count int64, err error) {
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

func (s smsHomeAdvDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s smsHomeAdvDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s smsHomeAdvDo) Delete(models ...*model.SmsHomeAdv) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *smsHomeAdvDo) withDO(do gen.Dao) *smsHomeAdvDo {
	s.DO = *do.(*gen.DO)
	return s
}
