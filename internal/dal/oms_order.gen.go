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

func newOmsOrder(db *gorm.DB, opts ...gen.DOOption) omsOrder {
	_omsOrder := omsOrder{}

	_omsOrder.omsOrderDo.UseDB(db, opts...)
	_omsOrder.omsOrderDo.UseModel(&model.OmsOrder{})

	tableName := _omsOrder.omsOrderDo.TableName()
	_omsOrder.ALL = field.NewAsterisk(tableName)
	_omsOrder.ID = field.NewInt(tableName, "id")
	_omsOrder.MemberID = field.NewInt(tableName, "member_id")
	_omsOrder.OrderNo = field.NewString(tableName, "order_no")
	_omsOrder.CouponID = field.NewInt(tableName, "coupon_id")
	_omsOrder.Username = field.NewString(tableName, "username")
	_omsOrder.TotalAmount = field.NewInt(tableName, "total_amount")
	_omsOrder.FreightAmount = field.NewInt(tableName, "freight_amount")
	_omsOrder.PromotionAmount = field.NewInt(tableName, "promotion_amount")
	_omsOrder.IntegrationAmount = field.NewInt(tableName, "integration_amount")
	_omsOrder.CouponAmount = field.NewInt(tableName, "coupon_amount")
	_omsOrder.DiscountAmount = field.NewInt(tableName, "discount_amount")
	_omsOrder.Amount = field.NewInt(tableName, "amount")
	_omsOrder.PayAmount = field.NewInt(tableName, "pay_amount")
	_omsOrder.PayType = field.NewInt8(tableName, "pay_type")
	_omsOrder.SourceType = field.NewInt8(tableName, "source_type")
	_omsOrder.Status = field.NewInt8(tableName, "status")
	_omsOrder.DeliveryCompany = field.NewString(tableName, "delivery_company")
	_omsOrder.DeliveryNo = field.NewString(tableName, "delivery_no")
	_omsOrder.Integration = field.NewInt(tableName, "integration")
	_omsOrder.Growth = field.NewInt(tableName, "growth")
	_omsOrder.AddressName = field.NewString(tableName, "address_name")
	_omsOrder.AddressPhone = field.NewString(tableName, "address_phone")
	_omsOrder.AddressProvince = field.NewString(tableName, "address_province")
	_omsOrder.AddressCity = field.NewString(tableName, "address_city")
	_omsOrder.AddressCounty = field.NewString(tableName, "address_county")
	_omsOrder.AddressDetail = field.NewString(tableName, "address_detail")
	_omsOrder.Note = field.NewString(tableName, "note")
	_omsOrder.IsConfirm = field.NewInt8(tableName, "is_confirm")
	_omsOrder.UseIntegration = field.NewInt64(tableName, "use_integration")
	_omsOrder.TradeNo = field.NewString(tableName, "trade_no")
	_omsOrder.TransHash = field.NewString(tableName, "trans_hash")
	_omsOrder.PayAt = field.NewTime(tableName, "pay_at")
	_omsOrder.DeliveryAt = field.NewTime(tableName, "delivery_at")
	_omsOrder.ReceiveAt = field.NewTime(tableName, "receive_at")
	_omsOrder.CommentAt = field.NewTime(tableName, "comment_at")
	_omsOrder.CreatedAt = field.NewTime(tableName, "created_at")
	_omsOrder.UpdatedAt = field.NewTime(tableName, "updated_at")
	_omsOrder.DeletedAt = field.NewField(tableName, "deleted_at")

	_omsOrder.fillFieldMap()

	return _omsOrder
}

type omsOrder struct {
	omsOrderDo

	ALL               field.Asterisk
	ID                field.Int    // ID
	MemberID          field.Int    // 用户id
	OrderNo           field.String // 订单号
	CouponID          field.Int    // 优惠券id
	Username          field.String // 用户名
	TotalAmount       field.Int    // 订单总额/分
	FreightAmount     field.Int    // 运费/分
	PromotionAmount   field.Int    // 促销优惠金额（促销价、满减、阶梯价）
	IntegrationAmount field.Int    // 积分抵扣金额
	CouponAmount      field.Int    // 优惠券折扣金额
	DiscountAmount    field.Int    // 后台调整订单使用的折扣金额
	Amount            field.Int    // 优惠后的真实金额
	PayAmount         field.Int    // 支付金额/分
	PayType           field.Int8   // 支付方式
	SourceType        field.Int8   // 订单来源[0->PC订单；1->app订单]
	Status            field.Int8   // 订单状态
	DeliveryCompany   field.String // 物流公司(配送方式)
	DeliveryNo        field.String // 物流单号
	Integration       field.Int    // 可以获得的积分
	Growth            field.Int    // 可以获得的成长值
	AddressName       field.String // 收货人姓名
	AddressPhone      field.String // 收货人手机
	AddressProvince   field.String // 省
	AddressCity       field.String // 市
	AddressCounty     field.String // 县/区
	AddressDetail     field.String // 详情
	Note              field.String // 备注
	IsConfirm         field.Int8   // 是否确认收货
	UseIntegration    field.Int64  // 下单时使用的积分
	TradeNo           field.String // 交易号
	TransHash         field.String // 交易hash
	PayAt             field.Time   // 支付时间
	DeliveryAt        field.Time   // 发货时间
	ReceiveAt         field.Time   // 确认收货时间
	CommentAt         field.Time   // 评价时间
	CreatedAt         field.Time   // 创建时间
	UpdatedAt         field.Time   // 更新时间
	DeletedAt         field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (o omsOrder) Table(newTableName string) *omsOrder {
	o.omsOrderDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o omsOrder) As(alias string) *omsOrder {
	o.omsOrderDo.DO = *(o.omsOrderDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *omsOrder) updateTableName(table string) *omsOrder {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt(table, "id")
	o.MemberID = field.NewInt(table, "member_id")
	o.OrderNo = field.NewString(table, "order_no")
	o.CouponID = field.NewInt(table, "coupon_id")
	o.Username = field.NewString(table, "username")
	o.TotalAmount = field.NewInt(table, "total_amount")
	o.FreightAmount = field.NewInt(table, "freight_amount")
	o.PromotionAmount = field.NewInt(table, "promotion_amount")
	o.IntegrationAmount = field.NewInt(table, "integration_amount")
	o.CouponAmount = field.NewInt(table, "coupon_amount")
	o.DiscountAmount = field.NewInt(table, "discount_amount")
	o.Amount = field.NewInt(table, "amount")
	o.PayAmount = field.NewInt(table, "pay_amount")
	o.PayType = field.NewInt8(table, "pay_type")
	o.SourceType = field.NewInt8(table, "source_type")
	o.Status = field.NewInt8(table, "status")
	o.DeliveryCompany = field.NewString(table, "delivery_company")
	o.DeliveryNo = field.NewString(table, "delivery_no")
	o.Integration = field.NewInt(table, "integration")
	o.Growth = field.NewInt(table, "growth")
	o.AddressName = field.NewString(table, "address_name")
	o.AddressPhone = field.NewString(table, "address_phone")
	o.AddressProvince = field.NewString(table, "address_province")
	o.AddressCity = field.NewString(table, "address_city")
	o.AddressCounty = field.NewString(table, "address_county")
	o.AddressDetail = field.NewString(table, "address_detail")
	o.Note = field.NewString(table, "note")
	o.IsConfirm = field.NewInt8(table, "is_confirm")
	o.UseIntegration = field.NewInt64(table, "use_integration")
	o.TradeNo = field.NewString(table, "trade_no")
	o.TransHash = field.NewString(table, "trans_hash")
	o.PayAt = field.NewTime(table, "pay_at")
	o.DeliveryAt = field.NewTime(table, "delivery_at")
	o.ReceiveAt = field.NewTime(table, "receive_at")
	o.CommentAt = field.NewTime(table, "comment_at")
	o.CreatedAt = field.NewTime(table, "created_at")
	o.UpdatedAt = field.NewTime(table, "updated_at")
	o.DeletedAt = field.NewField(table, "deleted_at")

	o.fillFieldMap()

	return o
}

func (o *omsOrder) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *omsOrder) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 38)
	o.fieldMap["id"] = o.ID
	o.fieldMap["member_id"] = o.MemberID
	o.fieldMap["order_no"] = o.OrderNo
	o.fieldMap["coupon_id"] = o.CouponID
	o.fieldMap["username"] = o.Username
	o.fieldMap["total_amount"] = o.TotalAmount
	o.fieldMap["freight_amount"] = o.FreightAmount
	o.fieldMap["promotion_amount"] = o.PromotionAmount
	o.fieldMap["integration_amount"] = o.IntegrationAmount
	o.fieldMap["coupon_amount"] = o.CouponAmount
	o.fieldMap["discount_amount"] = o.DiscountAmount
	o.fieldMap["amount"] = o.Amount
	o.fieldMap["pay_amount"] = o.PayAmount
	o.fieldMap["pay_type"] = o.PayType
	o.fieldMap["source_type"] = o.SourceType
	o.fieldMap["status"] = o.Status
	o.fieldMap["delivery_company"] = o.DeliveryCompany
	o.fieldMap["delivery_no"] = o.DeliveryNo
	o.fieldMap["integration"] = o.Integration
	o.fieldMap["growth"] = o.Growth
	o.fieldMap["address_name"] = o.AddressName
	o.fieldMap["address_phone"] = o.AddressPhone
	o.fieldMap["address_province"] = o.AddressProvince
	o.fieldMap["address_city"] = o.AddressCity
	o.fieldMap["address_county"] = o.AddressCounty
	o.fieldMap["address_detail"] = o.AddressDetail
	o.fieldMap["note"] = o.Note
	o.fieldMap["is_confirm"] = o.IsConfirm
	o.fieldMap["use_integration"] = o.UseIntegration
	o.fieldMap["trade_no"] = o.TradeNo
	o.fieldMap["trans_hash"] = o.TransHash
	o.fieldMap["pay_at"] = o.PayAt
	o.fieldMap["delivery_at"] = o.DeliveryAt
	o.fieldMap["receive_at"] = o.ReceiveAt
	o.fieldMap["comment_at"] = o.CommentAt
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
	o.fieldMap["deleted_at"] = o.DeletedAt
}

func (o omsOrder) clone(db *gorm.DB) omsOrder {
	o.omsOrderDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o omsOrder) replaceDB(db *gorm.DB) omsOrder {
	o.omsOrderDo.ReplaceDB(db)
	return o
}

type omsOrderDo struct{ gen.DO }

type IOmsOrderDo interface {
	gen.SubQuery
	Debug() IOmsOrderDo
	WithContext(ctx context.Context) IOmsOrderDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOmsOrderDo
	WriteDB() IOmsOrderDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOmsOrderDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOmsOrderDo
	Not(conds ...gen.Condition) IOmsOrderDo
	Or(conds ...gen.Condition) IOmsOrderDo
	Select(conds ...field.Expr) IOmsOrderDo
	Where(conds ...gen.Condition) IOmsOrderDo
	Order(conds ...field.Expr) IOmsOrderDo
	Distinct(cols ...field.Expr) IOmsOrderDo
	Omit(cols ...field.Expr) IOmsOrderDo
	Join(table schema.Tabler, on ...field.Expr) IOmsOrderDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderDo
	Group(cols ...field.Expr) IOmsOrderDo
	Having(conds ...gen.Condition) IOmsOrderDo
	Limit(limit int) IOmsOrderDo
	Offset(offset int) IOmsOrderDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderDo
	Unscoped() IOmsOrderDo
	Create(values ...*model.OmsOrder) error
	CreateInBatches(values []*model.OmsOrder, batchSize int) error
	Save(values ...*model.OmsOrder) error
	First() (*model.OmsOrder, error)
	Take() (*model.OmsOrder, error)
	Last() (*model.OmsOrder, error)
	Find() ([]*model.OmsOrder, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrder, err error)
	FindInBatches(result *[]*model.OmsOrder, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OmsOrder) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOmsOrderDo
	Assign(attrs ...field.AssignExpr) IOmsOrderDo
	Joins(fields ...field.RelationField) IOmsOrderDo
	Preload(fields ...field.RelationField) IOmsOrderDo
	FirstOrInit() (*model.OmsOrder, error)
	FirstOrCreate() (*model.OmsOrder, error)
	FindByPage(offset int, limit int) (result []*model.OmsOrder, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOmsOrderDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o omsOrderDo) Debug() IOmsOrderDo {
	return o.withDO(o.DO.Debug())
}

func (o omsOrderDo) WithContext(ctx context.Context) IOmsOrderDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o omsOrderDo) ReadDB() IOmsOrderDo {
	return o.Clauses(dbresolver.Read)
}

func (o omsOrderDo) WriteDB() IOmsOrderDo {
	return o.Clauses(dbresolver.Write)
}

func (o omsOrderDo) Session(config *gorm.Session) IOmsOrderDo {
	return o.withDO(o.DO.Session(config))
}

func (o omsOrderDo) Clauses(conds ...clause.Expression) IOmsOrderDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o omsOrderDo) Returning(value interface{}, columns ...string) IOmsOrderDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o omsOrderDo) Not(conds ...gen.Condition) IOmsOrderDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o omsOrderDo) Or(conds ...gen.Condition) IOmsOrderDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o omsOrderDo) Select(conds ...field.Expr) IOmsOrderDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o omsOrderDo) Where(conds ...gen.Condition) IOmsOrderDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o omsOrderDo) Order(conds ...field.Expr) IOmsOrderDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o omsOrderDo) Distinct(cols ...field.Expr) IOmsOrderDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o omsOrderDo) Omit(cols ...field.Expr) IOmsOrderDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o omsOrderDo) Join(table schema.Tabler, on ...field.Expr) IOmsOrderDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o omsOrderDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o omsOrderDo) RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o omsOrderDo) Group(cols ...field.Expr) IOmsOrderDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o omsOrderDo) Having(conds ...gen.Condition) IOmsOrderDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o omsOrderDo) Limit(limit int) IOmsOrderDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o omsOrderDo) Offset(offset int) IOmsOrderDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o omsOrderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o omsOrderDo) Unscoped() IOmsOrderDo {
	return o.withDO(o.DO.Unscoped())
}

func (o omsOrderDo) Create(values ...*model.OmsOrder) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o omsOrderDo) CreateInBatches(values []*model.OmsOrder, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o omsOrderDo) Save(values ...*model.OmsOrder) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o omsOrderDo) First() (*model.OmsOrder, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrder), nil
	}
}

func (o omsOrderDo) Take() (*model.OmsOrder, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrder), nil
	}
}

func (o omsOrderDo) Last() (*model.OmsOrder, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrder), nil
	}
}

func (o omsOrderDo) Find() ([]*model.OmsOrder, error) {
	result, err := o.DO.Find()
	return result.([]*model.OmsOrder), err
}

func (o omsOrderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrder, err error) {
	buf := make([]*model.OmsOrder, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o omsOrderDo) FindInBatches(result *[]*model.OmsOrder, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o omsOrderDo) Attrs(attrs ...field.AssignExpr) IOmsOrderDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o omsOrderDo) Assign(attrs ...field.AssignExpr) IOmsOrderDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o omsOrderDo) Joins(fields ...field.RelationField) IOmsOrderDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o omsOrderDo) Preload(fields ...field.RelationField) IOmsOrderDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o omsOrderDo) FirstOrInit() (*model.OmsOrder, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrder), nil
	}
}

func (o omsOrderDo) FirstOrCreate() (*model.OmsOrder, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrder), nil
	}
}

func (o omsOrderDo) FindByPage(offset int, limit int) (result []*model.OmsOrder, count int64, err error) {
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

func (o omsOrderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o omsOrderDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o omsOrderDo) Delete(models ...*model.OmsOrder) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *omsOrderDo) withDO(do gen.Dao) *omsOrderDo {
	o.DO = *do.(*gen.DO)
	return o
}
