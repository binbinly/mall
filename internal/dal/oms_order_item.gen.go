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

func newOmsOrderItem(db *gorm.DB, opts ...gen.DOOption) omsOrderItem {
	_omsOrderItem := omsOrderItem{}

	_omsOrderItem.omsOrderItemDo.UseDB(db, opts...)
	_omsOrderItem.omsOrderItemDo.UseModel(&model.OmsOrderItem{})

	tableName := _omsOrderItem.omsOrderItemDo.TableName()
	_omsOrderItem.ALL = field.NewAsterisk(tableName)
	_omsOrderItem.ID = field.NewInt(tableName, "id")
	_omsOrderItem.OrderID = field.NewInt(tableName, "order_id")
	_omsOrderItem.OrderNo = field.NewString(tableName, "order_no")
	_omsOrderItem.SkuID = field.NewInt(tableName, "sku_id")
	_omsOrderItem.SkuName = field.NewString(tableName, "sku_name")
	_omsOrderItem.SkuImg = field.NewString(tableName, "sku_img")
	_omsOrderItem.SkuPrice = field.NewInt(tableName, "sku_price")
	_omsOrderItem.SkuAttrs = field.NewString(tableName, "sku_attrs")
	_omsOrderItem.SkuNum = field.NewInt(tableName, "sku_num")
	_omsOrderItem.PromotionAmount = field.NewInt(tableName, "promotion_amount")
	_omsOrderItem.IntegrationAmount = field.NewInt(tableName, "integration_amount")
	_omsOrderItem.CouponAmount = field.NewInt(tableName, "coupon_amount")
	_omsOrderItem.RealAmount = field.NewInt(tableName, "real_amount")
	_omsOrderItem.GiveIntegration = field.NewInt(tableName, "give_integration")
	_omsOrderItem.GiveGrowth = field.NewInt(tableName, "give_growth")

	_omsOrderItem.fillFieldMap()

	return _omsOrderItem
}

type omsOrderItem struct {
	omsOrderItemDo

	ALL               field.Asterisk
	ID                field.Int    // ID
	OrderID           field.Int    // 订单id
	OrderNo           field.String // 订单号
	SkuID             field.Int    // sku_id
	SkuName           field.String // 商品sku名
	SkuImg            field.String // 商品sku图片
	SkuPrice          field.Int    // 商品单价/分
	SkuAttrs          field.String // 商品销售属性组合（JSON）
	SkuNum            field.Int    // 购买数量
	PromotionAmount   field.Int    // 促销优惠金额（促销价、满减、阶梯价）
	IntegrationAmount field.Int    // 积分抵扣金额
	CouponAmount      field.Int    // 优惠券折扣金额
	RealAmount        field.Int    // 优惠后的真实金额
	GiveIntegration   field.Int    // 赠送积分
	GiveGrowth        field.Int    // 赠送成长值

	fieldMap map[string]field.Expr
}

func (o omsOrderItem) Table(newTableName string) *omsOrderItem {
	o.omsOrderItemDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o omsOrderItem) As(alias string) *omsOrderItem {
	o.omsOrderItemDo.DO = *(o.omsOrderItemDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *omsOrderItem) updateTableName(table string) *omsOrderItem {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt(table, "id")
	o.OrderID = field.NewInt(table, "order_id")
	o.OrderNo = field.NewString(table, "order_no")
	o.SkuID = field.NewInt(table, "sku_id")
	o.SkuName = field.NewString(table, "sku_name")
	o.SkuImg = field.NewString(table, "sku_img")
	o.SkuPrice = field.NewInt(table, "sku_price")
	o.SkuAttrs = field.NewString(table, "sku_attrs")
	o.SkuNum = field.NewInt(table, "sku_num")
	o.PromotionAmount = field.NewInt(table, "promotion_amount")
	o.IntegrationAmount = field.NewInt(table, "integration_amount")
	o.CouponAmount = field.NewInt(table, "coupon_amount")
	o.RealAmount = field.NewInt(table, "real_amount")
	o.GiveIntegration = field.NewInt(table, "give_integration")
	o.GiveGrowth = field.NewInt(table, "give_growth")

	o.fillFieldMap()

	return o
}

func (o *omsOrderItem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *omsOrderItem) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 15)
	o.fieldMap["id"] = o.ID
	o.fieldMap["order_id"] = o.OrderID
	o.fieldMap["order_no"] = o.OrderNo
	o.fieldMap["sku_id"] = o.SkuID
	o.fieldMap["sku_name"] = o.SkuName
	o.fieldMap["sku_img"] = o.SkuImg
	o.fieldMap["sku_price"] = o.SkuPrice
	o.fieldMap["sku_attrs"] = o.SkuAttrs
	o.fieldMap["sku_num"] = o.SkuNum
	o.fieldMap["promotion_amount"] = o.PromotionAmount
	o.fieldMap["integration_amount"] = o.IntegrationAmount
	o.fieldMap["coupon_amount"] = o.CouponAmount
	o.fieldMap["real_amount"] = o.RealAmount
	o.fieldMap["give_integration"] = o.GiveIntegration
	o.fieldMap["give_growth"] = o.GiveGrowth
}

func (o omsOrderItem) clone(db *gorm.DB) omsOrderItem {
	o.omsOrderItemDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o omsOrderItem) replaceDB(db *gorm.DB) omsOrderItem {
	o.omsOrderItemDo.ReplaceDB(db)
	return o
}

type omsOrderItemDo struct{ gen.DO }

type IOmsOrderItemDo interface {
	gen.SubQuery
	Debug() IOmsOrderItemDo
	WithContext(ctx context.Context) IOmsOrderItemDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOmsOrderItemDo
	WriteDB() IOmsOrderItemDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOmsOrderItemDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOmsOrderItemDo
	Not(conds ...gen.Condition) IOmsOrderItemDo
	Or(conds ...gen.Condition) IOmsOrderItemDo
	Select(conds ...field.Expr) IOmsOrderItemDo
	Where(conds ...gen.Condition) IOmsOrderItemDo
	Order(conds ...field.Expr) IOmsOrderItemDo
	Distinct(cols ...field.Expr) IOmsOrderItemDo
	Omit(cols ...field.Expr) IOmsOrderItemDo
	Join(table schema.Tabler, on ...field.Expr) IOmsOrderItemDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderItemDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderItemDo
	Group(cols ...field.Expr) IOmsOrderItemDo
	Having(conds ...gen.Condition) IOmsOrderItemDo
	Limit(limit int) IOmsOrderItemDo
	Offset(offset int) IOmsOrderItemDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderItemDo
	Unscoped() IOmsOrderItemDo
	Create(values ...*model.OmsOrderItem) error
	CreateInBatches(values []*model.OmsOrderItem, batchSize int) error
	Save(values ...*model.OmsOrderItem) error
	First() (*model.OmsOrderItem, error)
	Take() (*model.OmsOrderItem, error)
	Last() (*model.OmsOrderItem, error)
	Find() ([]*model.OmsOrderItem, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderItem, err error)
	FindInBatches(result *[]*model.OmsOrderItem, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OmsOrderItem) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOmsOrderItemDo
	Assign(attrs ...field.AssignExpr) IOmsOrderItemDo
	Joins(fields ...field.RelationField) IOmsOrderItemDo
	Preload(fields ...field.RelationField) IOmsOrderItemDo
	FirstOrInit() (*model.OmsOrderItem, error)
	FirstOrCreate() (*model.OmsOrderItem, error)
	FindByPage(offset int, limit int) (result []*model.OmsOrderItem, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOmsOrderItemDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o omsOrderItemDo) Debug() IOmsOrderItemDo {
	return o.withDO(o.DO.Debug())
}

func (o omsOrderItemDo) WithContext(ctx context.Context) IOmsOrderItemDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o omsOrderItemDo) ReadDB() IOmsOrderItemDo {
	return o.Clauses(dbresolver.Read)
}

func (o omsOrderItemDo) WriteDB() IOmsOrderItemDo {
	return o.Clauses(dbresolver.Write)
}

func (o omsOrderItemDo) Session(config *gorm.Session) IOmsOrderItemDo {
	return o.withDO(o.DO.Session(config))
}

func (o omsOrderItemDo) Clauses(conds ...clause.Expression) IOmsOrderItemDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o omsOrderItemDo) Returning(value interface{}, columns ...string) IOmsOrderItemDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o omsOrderItemDo) Not(conds ...gen.Condition) IOmsOrderItemDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o omsOrderItemDo) Or(conds ...gen.Condition) IOmsOrderItemDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o omsOrderItemDo) Select(conds ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o omsOrderItemDo) Where(conds ...gen.Condition) IOmsOrderItemDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o omsOrderItemDo) Order(conds ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o omsOrderItemDo) Distinct(cols ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o omsOrderItemDo) Omit(cols ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o omsOrderItemDo) Join(table schema.Tabler, on ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o omsOrderItemDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o omsOrderItemDo) RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o omsOrderItemDo) Group(cols ...field.Expr) IOmsOrderItemDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o omsOrderItemDo) Having(conds ...gen.Condition) IOmsOrderItemDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o omsOrderItemDo) Limit(limit int) IOmsOrderItemDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o omsOrderItemDo) Offset(offset int) IOmsOrderItemDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o omsOrderItemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderItemDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o omsOrderItemDo) Unscoped() IOmsOrderItemDo {
	return o.withDO(o.DO.Unscoped())
}

func (o omsOrderItemDo) Create(values ...*model.OmsOrderItem) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o omsOrderItemDo) CreateInBatches(values []*model.OmsOrderItem, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o omsOrderItemDo) Save(values ...*model.OmsOrderItem) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o omsOrderItemDo) First() (*model.OmsOrderItem, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderItem), nil
	}
}

func (o omsOrderItemDo) Take() (*model.OmsOrderItem, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderItem), nil
	}
}

func (o omsOrderItemDo) Last() (*model.OmsOrderItem, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderItem), nil
	}
}

func (o omsOrderItemDo) Find() ([]*model.OmsOrderItem, error) {
	result, err := o.DO.Find()
	return result.([]*model.OmsOrderItem), err
}

func (o omsOrderItemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderItem, err error) {
	buf := make([]*model.OmsOrderItem, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o omsOrderItemDo) FindInBatches(result *[]*model.OmsOrderItem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o omsOrderItemDo) Attrs(attrs ...field.AssignExpr) IOmsOrderItemDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o omsOrderItemDo) Assign(attrs ...field.AssignExpr) IOmsOrderItemDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o omsOrderItemDo) Joins(fields ...field.RelationField) IOmsOrderItemDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o omsOrderItemDo) Preload(fields ...field.RelationField) IOmsOrderItemDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o omsOrderItemDo) FirstOrInit() (*model.OmsOrderItem, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderItem), nil
	}
}

func (o omsOrderItemDo) FirstOrCreate() (*model.OmsOrderItem, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderItem), nil
	}
}

func (o omsOrderItemDo) FindByPage(offset int, limit int) (result []*model.OmsOrderItem, count int64, err error) {
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

func (o omsOrderItemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o omsOrderItemDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o omsOrderItemDo) Delete(models ...*model.OmsOrderItem) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *omsOrderItemDo) withDO(do gen.Dao) *omsOrderItemDo {
	o.DO = *do.(*gen.DO)
	return o
}
