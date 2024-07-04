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

func newPersonalAccessToken(db *gorm.DB, opts ...gen.DOOption) personalAccessToken {
	_personalAccessToken := personalAccessToken{}

	_personalAccessToken.personalAccessTokenDo.UseDB(db, opts...)
	_personalAccessToken.personalAccessTokenDo.UseModel(&model.PersonalAccessToken{})

	tableName := _personalAccessToken.personalAccessTokenDo.TableName()
	_personalAccessToken.ALL = field.NewAsterisk(tableName)
	_personalAccessToken.ID = field.NewInt64(tableName, "id")
	_personalAccessToken.TokenableType = field.NewString(tableName, "tokenable_type")
	_personalAccessToken.TokenableID = field.NewInt64(tableName, "tokenable_id")
	_personalAccessToken.Name = field.NewString(tableName, "name")
	_personalAccessToken.Token = field.NewString(tableName, "token")
	_personalAccessToken.Abilities = field.NewString(tableName, "abilities")
	_personalAccessToken.LastUsedAt = field.NewTime(tableName, "last_used_at")
	_personalAccessToken.ExpiresAt = field.NewTime(tableName, "expires_at")
	_personalAccessToken.CreatedAt = field.NewTime(tableName, "created_at")
	_personalAccessToken.UpdatedAt = field.NewTime(tableName, "updated_at")

	_personalAccessToken.fillFieldMap()

	return _personalAccessToken
}

type personalAccessToken struct {
	personalAccessTokenDo

	ALL           field.Asterisk
	ID            field.Int64
	TokenableType field.String
	TokenableID   field.Int64
	Name          field.String
	Token         field.String
	Abilities     field.String
	LastUsedAt    field.Time
	ExpiresAt     field.Time
	CreatedAt     field.Time
	UpdatedAt     field.Time

	fieldMap map[string]field.Expr
}

func (p personalAccessToken) Table(newTableName string) *personalAccessToken {
	p.personalAccessTokenDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p personalAccessToken) As(alias string) *personalAccessToken {
	p.personalAccessTokenDo.DO = *(p.personalAccessTokenDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *personalAccessToken) updateTableName(table string) *personalAccessToken {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.TokenableType = field.NewString(table, "tokenable_type")
	p.TokenableID = field.NewInt64(table, "tokenable_id")
	p.Name = field.NewString(table, "name")
	p.Token = field.NewString(table, "token")
	p.Abilities = field.NewString(table, "abilities")
	p.LastUsedAt = field.NewTime(table, "last_used_at")
	p.ExpiresAt = field.NewTime(table, "expires_at")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")

	p.fillFieldMap()

	return p
}

func (p *personalAccessToken) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *personalAccessToken) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 10)
	p.fieldMap["id"] = p.ID
	p.fieldMap["tokenable_type"] = p.TokenableType
	p.fieldMap["tokenable_id"] = p.TokenableID
	p.fieldMap["name"] = p.Name
	p.fieldMap["token"] = p.Token
	p.fieldMap["abilities"] = p.Abilities
	p.fieldMap["last_used_at"] = p.LastUsedAt
	p.fieldMap["expires_at"] = p.ExpiresAt
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
}

func (p personalAccessToken) clone(db *gorm.DB) personalAccessToken {
	p.personalAccessTokenDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p personalAccessToken) replaceDB(db *gorm.DB) personalAccessToken {
	p.personalAccessTokenDo.ReplaceDB(db)
	return p
}

type personalAccessTokenDo struct{ gen.DO }

type IPersonalAccessTokenDo interface {
	gen.SubQuery
	Debug() IPersonalAccessTokenDo
	WithContext(ctx context.Context) IPersonalAccessTokenDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPersonalAccessTokenDo
	WriteDB() IPersonalAccessTokenDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPersonalAccessTokenDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPersonalAccessTokenDo
	Not(conds ...gen.Condition) IPersonalAccessTokenDo
	Or(conds ...gen.Condition) IPersonalAccessTokenDo
	Select(conds ...field.Expr) IPersonalAccessTokenDo
	Where(conds ...gen.Condition) IPersonalAccessTokenDo
	Order(conds ...field.Expr) IPersonalAccessTokenDo
	Distinct(cols ...field.Expr) IPersonalAccessTokenDo
	Omit(cols ...field.Expr) IPersonalAccessTokenDo
	Join(table schema.Tabler, on ...field.Expr) IPersonalAccessTokenDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPersonalAccessTokenDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPersonalAccessTokenDo
	Group(cols ...field.Expr) IPersonalAccessTokenDo
	Having(conds ...gen.Condition) IPersonalAccessTokenDo
	Limit(limit int) IPersonalAccessTokenDo
	Offset(offset int) IPersonalAccessTokenDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPersonalAccessTokenDo
	Unscoped() IPersonalAccessTokenDo
	Create(values ...*model.PersonalAccessToken) error
	CreateInBatches(values []*model.PersonalAccessToken, batchSize int) error
	Save(values ...*model.PersonalAccessToken) error
	First() (*model.PersonalAccessToken, error)
	Take() (*model.PersonalAccessToken, error)
	Last() (*model.PersonalAccessToken, error)
	Find() ([]*model.PersonalAccessToken, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PersonalAccessToken, err error)
	FindInBatches(result *[]*model.PersonalAccessToken, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PersonalAccessToken) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPersonalAccessTokenDo
	Assign(attrs ...field.AssignExpr) IPersonalAccessTokenDo
	Joins(fields ...field.RelationField) IPersonalAccessTokenDo
	Preload(fields ...field.RelationField) IPersonalAccessTokenDo
	FirstOrInit() (*model.PersonalAccessToken, error)
	FirstOrCreate() (*model.PersonalAccessToken, error)
	FindByPage(offset int, limit int) (result []*model.PersonalAccessToken, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPersonalAccessTokenDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p personalAccessTokenDo) Debug() IPersonalAccessTokenDo {
	return p.withDO(p.DO.Debug())
}

func (p personalAccessTokenDo) WithContext(ctx context.Context) IPersonalAccessTokenDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p personalAccessTokenDo) ReadDB() IPersonalAccessTokenDo {
	return p.Clauses(dbresolver.Read)
}

func (p personalAccessTokenDo) WriteDB() IPersonalAccessTokenDo {
	return p.Clauses(dbresolver.Write)
}

func (p personalAccessTokenDo) Session(config *gorm.Session) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Session(config))
}

func (p personalAccessTokenDo) Clauses(conds ...clause.Expression) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p personalAccessTokenDo) Returning(value interface{}, columns ...string) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p personalAccessTokenDo) Not(conds ...gen.Condition) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p personalAccessTokenDo) Or(conds ...gen.Condition) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p personalAccessTokenDo) Select(conds ...field.Expr) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p personalAccessTokenDo) Where(conds ...gen.Condition) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p personalAccessTokenDo) Order(conds ...field.Expr) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p personalAccessTokenDo) Distinct(cols ...field.Expr) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p personalAccessTokenDo) Omit(cols ...field.Expr) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p personalAccessTokenDo) Join(table schema.Tabler, on ...field.Expr) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p personalAccessTokenDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPersonalAccessTokenDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p personalAccessTokenDo) RightJoin(table schema.Tabler, on ...field.Expr) IPersonalAccessTokenDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p personalAccessTokenDo) Group(cols ...field.Expr) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p personalAccessTokenDo) Having(conds ...gen.Condition) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p personalAccessTokenDo) Limit(limit int) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p personalAccessTokenDo) Offset(offset int) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p personalAccessTokenDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p personalAccessTokenDo) Unscoped() IPersonalAccessTokenDo {
	return p.withDO(p.DO.Unscoped())
}

func (p personalAccessTokenDo) Create(values ...*model.PersonalAccessToken) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p personalAccessTokenDo) CreateInBatches(values []*model.PersonalAccessToken, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p personalAccessTokenDo) Save(values ...*model.PersonalAccessToken) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p personalAccessTokenDo) First() (*model.PersonalAccessToken, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PersonalAccessToken), nil
	}
}

func (p personalAccessTokenDo) Take() (*model.PersonalAccessToken, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PersonalAccessToken), nil
	}
}

func (p personalAccessTokenDo) Last() (*model.PersonalAccessToken, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PersonalAccessToken), nil
	}
}

func (p personalAccessTokenDo) Find() ([]*model.PersonalAccessToken, error) {
	result, err := p.DO.Find()
	return result.([]*model.PersonalAccessToken), err
}

func (p personalAccessTokenDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PersonalAccessToken, err error) {
	buf := make([]*model.PersonalAccessToken, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p personalAccessTokenDo) FindInBatches(result *[]*model.PersonalAccessToken, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p personalAccessTokenDo) Attrs(attrs ...field.AssignExpr) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p personalAccessTokenDo) Assign(attrs ...field.AssignExpr) IPersonalAccessTokenDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p personalAccessTokenDo) Joins(fields ...field.RelationField) IPersonalAccessTokenDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p personalAccessTokenDo) Preload(fields ...field.RelationField) IPersonalAccessTokenDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p personalAccessTokenDo) FirstOrInit() (*model.PersonalAccessToken, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PersonalAccessToken), nil
	}
}

func (p personalAccessTokenDo) FirstOrCreate() (*model.PersonalAccessToken, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PersonalAccessToken), nil
	}
}

func (p personalAccessTokenDo) FindByPage(offset int, limit int) (result []*model.PersonalAccessToken, count int64, err error) {
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

func (p personalAccessTokenDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p personalAccessTokenDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p personalAccessTokenDo) Delete(models ...*model.PersonalAccessToken) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *personalAccessTokenDo) withDO(do gen.Dao) *personalAccessTokenDo {
	p.DO = *do.(*gen.DO)
	return p
}
