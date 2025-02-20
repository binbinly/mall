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

func newPasswordResetToken(db *gorm.DB, opts ...gen.DOOption) passwordResetToken {
	_passwordResetToken := passwordResetToken{}

	_passwordResetToken.passwordResetTokenDo.UseDB(db, opts...)
	_passwordResetToken.passwordResetTokenDo.UseModel(&model.PasswordResetToken{})

	tableName := _passwordResetToken.passwordResetTokenDo.TableName()
	_passwordResetToken.ALL = field.NewAsterisk(tableName)
	_passwordResetToken.Email = field.NewString(tableName, "email")
	_passwordResetToken.Token = field.NewString(tableName, "token")
	_passwordResetToken.CreatedAt = field.NewTime(tableName, "created_at")

	_passwordResetToken.fillFieldMap()

	return _passwordResetToken
}

type passwordResetToken struct {
	passwordResetTokenDo

	ALL       field.Asterisk
	Email     field.String
	Token     field.String
	CreatedAt field.Time

	fieldMap map[string]field.Expr
}

func (p passwordResetToken) Table(newTableName string) *passwordResetToken {
	p.passwordResetTokenDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p passwordResetToken) As(alias string) *passwordResetToken {
	p.passwordResetTokenDo.DO = *(p.passwordResetTokenDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *passwordResetToken) updateTableName(table string) *passwordResetToken {
	p.ALL = field.NewAsterisk(table)
	p.Email = field.NewString(table, "email")
	p.Token = field.NewString(table, "token")
	p.CreatedAt = field.NewTime(table, "created_at")

	p.fillFieldMap()

	return p
}

func (p *passwordResetToken) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *passwordResetToken) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 3)
	p.fieldMap["email"] = p.Email
	p.fieldMap["token"] = p.Token
	p.fieldMap["created_at"] = p.CreatedAt
}

func (p passwordResetToken) clone(db *gorm.DB) passwordResetToken {
	p.passwordResetTokenDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p passwordResetToken) replaceDB(db *gorm.DB) passwordResetToken {
	p.passwordResetTokenDo.ReplaceDB(db)
	return p
}

type passwordResetTokenDo struct{ gen.DO }

type IPasswordResetTokenDo interface {
	gen.SubQuery
	Debug() IPasswordResetTokenDo
	WithContext(ctx context.Context) IPasswordResetTokenDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPasswordResetTokenDo
	WriteDB() IPasswordResetTokenDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPasswordResetTokenDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPasswordResetTokenDo
	Not(conds ...gen.Condition) IPasswordResetTokenDo
	Or(conds ...gen.Condition) IPasswordResetTokenDo
	Select(conds ...field.Expr) IPasswordResetTokenDo
	Where(conds ...gen.Condition) IPasswordResetTokenDo
	Order(conds ...field.Expr) IPasswordResetTokenDo
	Distinct(cols ...field.Expr) IPasswordResetTokenDo
	Omit(cols ...field.Expr) IPasswordResetTokenDo
	Join(table schema.Tabler, on ...field.Expr) IPasswordResetTokenDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPasswordResetTokenDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPasswordResetTokenDo
	Group(cols ...field.Expr) IPasswordResetTokenDo
	Having(conds ...gen.Condition) IPasswordResetTokenDo
	Limit(limit int) IPasswordResetTokenDo
	Offset(offset int) IPasswordResetTokenDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPasswordResetTokenDo
	Unscoped() IPasswordResetTokenDo
	Create(values ...*model.PasswordResetToken) error
	CreateInBatches(values []*model.PasswordResetToken, batchSize int) error
	Save(values ...*model.PasswordResetToken) error
	First() (*model.PasswordResetToken, error)
	Take() (*model.PasswordResetToken, error)
	Last() (*model.PasswordResetToken, error)
	Find() ([]*model.PasswordResetToken, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PasswordResetToken, err error)
	FindInBatches(result *[]*model.PasswordResetToken, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PasswordResetToken) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPasswordResetTokenDo
	Assign(attrs ...field.AssignExpr) IPasswordResetTokenDo
	Joins(fields ...field.RelationField) IPasswordResetTokenDo
	Preload(fields ...field.RelationField) IPasswordResetTokenDo
	FirstOrInit() (*model.PasswordResetToken, error)
	FirstOrCreate() (*model.PasswordResetToken, error)
	FindByPage(offset int, limit int) (result []*model.PasswordResetToken, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPasswordResetTokenDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p passwordResetTokenDo) Debug() IPasswordResetTokenDo {
	return p.withDO(p.DO.Debug())
}

func (p passwordResetTokenDo) WithContext(ctx context.Context) IPasswordResetTokenDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p passwordResetTokenDo) ReadDB() IPasswordResetTokenDo {
	return p.Clauses(dbresolver.Read)
}

func (p passwordResetTokenDo) WriteDB() IPasswordResetTokenDo {
	return p.Clauses(dbresolver.Write)
}

func (p passwordResetTokenDo) Session(config *gorm.Session) IPasswordResetTokenDo {
	return p.withDO(p.DO.Session(config))
}

func (p passwordResetTokenDo) Clauses(conds ...clause.Expression) IPasswordResetTokenDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p passwordResetTokenDo) Returning(value interface{}, columns ...string) IPasswordResetTokenDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p passwordResetTokenDo) Not(conds ...gen.Condition) IPasswordResetTokenDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p passwordResetTokenDo) Or(conds ...gen.Condition) IPasswordResetTokenDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p passwordResetTokenDo) Select(conds ...field.Expr) IPasswordResetTokenDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p passwordResetTokenDo) Where(conds ...gen.Condition) IPasswordResetTokenDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p passwordResetTokenDo) Order(conds ...field.Expr) IPasswordResetTokenDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p passwordResetTokenDo) Distinct(cols ...field.Expr) IPasswordResetTokenDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p passwordResetTokenDo) Omit(cols ...field.Expr) IPasswordResetTokenDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p passwordResetTokenDo) Join(table schema.Tabler, on ...field.Expr) IPasswordResetTokenDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p passwordResetTokenDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPasswordResetTokenDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p passwordResetTokenDo) RightJoin(table schema.Tabler, on ...field.Expr) IPasswordResetTokenDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p passwordResetTokenDo) Group(cols ...field.Expr) IPasswordResetTokenDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p passwordResetTokenDo) Having(conds ...gen.Condition) IPasswordResetTokenDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p passwordResetTokenDo) Limit(limit int) IPasswordResetTokenDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p passwordResetTokenDo) Offset(offset int) IPasswordResetTokenDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p passwordResetTokenDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPasswordResetTokenDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p passwordResetTokenDo) Unscoped() IPasswordResetTokenDo {
	return p.withDO(p.DO.Unscoped())
}

func (p passwordResetTokenDo) Create(values ...*model.PasswordResetToken) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p passwordResetTokenDo) CreateInBatches(values []*model.PasswordResetToken, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p passwordResetTokenDo) Save(values ...*model.PasswordResetToken) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p passwordResetTokenDo) First() (*model.PasswordResetToken, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PasswordResetToken), nil
	}
}

func (p passwordResetTokenDo) Take() (*model.PasswordResetToken, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PasswordResetToken), nil
	}
}

func (p passwordResetTokenDo) Last() (*model.PasswordResetToken, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PasswordResetToken), nil
	}
}

func (p passwordResetTokenDo) Find() ([]*model.PasswordResetToken, error) {
	result, err := p.DO.Find()
	return result.([]*model.PasswordResetToken), err
}

func (p passwordResetTokenDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PasswordResetToken, err error) {
	buf := make([]*model.PasswordResetToken, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p passwordResetTokenDo) FindInBatches(result *[]*model.PasswordResetToken, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p passwordResetTokenDo) Attrs(attrs ...field.AssignExpr) IPasswordResetTokenDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p passwordResetTokenDo) Assign(attrs ...field.AssignExpr) IPasswordResetTokenDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p passwordResetTokenDo) Joins(fields ...field.RelationField) IPasswordResetTokenDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p passwordResetTokenDo) Preload(fields ...field.RelationField) IPasswordResetTokenDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p passwordResetTokenDo) FirstOrInit() (*model.PasswordResetToken, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PasswordResetToken), nil
	}
}

func (p passwordResetTokenDo) FirstOrCreate() (*model.PasswordResetToken, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PasswordResetToken), nil
	}
}

func (p passwordResetTokenDo) FindByPage(offset int, limit int) (result []*model.PasswordResetToken, count int64, err error) {
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

func (p passwordResetTokenDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p passwordResetTokenDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p passwordResetTokenDo) Delete(models ...*model.PasswordResetToken) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *passwordResetTokenDo) withDO(do gen.Dao) *passwordResetTokenDo {
	p.DO = *do.(*gen.DO)
	return p
}
