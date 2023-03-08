// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"kp-management/internal/pkg/dal/model"
)

func newOrder(db *gorm.DB, opts ...gen.DOOption) order {
	_order := order{}

	_order.orderDo.UseDB(db, opts...)
	_order.orderDo.UseModel(&model.Order{})

	tableName := _order.orderDo.TableName()
	_order.ALL = field.NewAsterisk(tableName)
	_order.ID = field.NewInt32(tableName, "id")
	_order.OrderID = field.NewString(tableName, "order_id")
	_order.PayTradeNo = field.NewString(tableName, "pay_trade_no")
	_order.TradeNo = field.NewString(tableName, "trade_no")
	_order.TeamID = field.NewString(tableName, "team_id")
	_order.GoogsName = field.NewString(tableName, "googs_name")
	_order.OrderType = field.NewInt32(tableName, "order_type")
	_order.VumBuyVersionType = field.NewInt32(tableName, "vum_buy_version_type")
	_order.TeamBuyVersionType = field.NewInt32(tableName, "team_buy_version_type")
	_order.MaxConcurrence = field.NewInt32(tableName, "max_concurrence")
	_order.BuyNumber = field.NewInt64(tableName, "buy_number")
	_order.OrderAmount = field.NewFloat64(tableName, "order_amount")
	_order.Discounts = field.NewFloat64(tableName, "discounts")
	_order.RealAmount = field.NewFloat64(tableName, "real_amount")
	_order.PayType = field.NewInt32(tableName, "pay_type")
	_order.PayStatus = field.NewInt32(tableName, "pay_status")
	_order.GoodsValidDate = field.NewInt32(tableName, "goods_valid_date")
	_order.FinishPayTime = field.NewTime(tableName, "finish_pay_time")
	_order.OpenInvoiceState = field.NewInt32(tableName, "open_invoice_state")
	_order.PayUserID = field.NewString(tableName, "pay_user_id")
	_order.CreatedAt = field.NewTime(tableName, "created_at")
	_order.UpdatedAt = field.NewTime(tableName, "updated_at")
	_order.DeletedAt = field.NewField(tableName, "deleted_at")

	_order.fillFieldMap()

	return _order
}

type order struct {
	orderDo orderDo

	ALL                field.Asterisk
	ID                 field.Int32   // 主键id
	OrderID            field.String  // 订单号id
	PayTradeNo         field.String  // 支付中心订单号
	TradeNo            field.String  // 第三方订单号
	TeamID             field.String  // 团队id
	GoogsName          field.String  // 商品名称
	OrderType          field.Int32   // 订单类型：1-新建团队，2-VUM资源包，3-升级团队，4-增加席位，5-套餐续期
	VumBuyVersionType  field.Int32   // VUM套餐类型：1-A,2-B,3-C,4-D
	TeamBuyVersionType field.Int32   // 团队套餐类型：1-个人版，2-团队版，3-企业版
	MaxConcurrence     field.Int32   // 最大并发数
	BuyNumber          field.Int64   // 购买数量（席位数量/VUM资源包数量）
	OrderAmount        field.Float64 // 订单金额
	Discounts          field.Float64 // 优惠金额
	RealAmount         field.Float64 // 实际付款金额
	PayType            field.Int32   // 支付方式：0-未知，1-微信，2-支付宝，3-银联，4-PayPal
	PayStatus          field.Int32   // 支付状态：0-待支付，1-支付成功，2-支付失败，3-支付中，4-已过期，5-已关闭
	GoodsValidDate     field.Int32   // 商品的有效期（单位：月）
	FinishPayTime      field.Time    // 到账时间
	OpenInvoiceState   field.Int32   // 发票申请状态：0-待开票，1-已开票，2-已作废
	PayUserID          field.String  // 支付人id
	CreatedAt          field.Time    // 创建时间
	UpdatedAt          field.Time    // 更新时间
	DeletedAt          field.Field   // 删除时间

	fieldMap map[string]field.Expr
}

func (o order) Table(newTableName string) *order {
	o.orderDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o order) As(alias string) *order {
	o.orderDo.DO = *(o.orderDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *order) updateTableName(table string) *order {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt32(table, "id")
	o.OrderID = field.NewString(table, "order_id")
	o.PayTradeNo = field.NewString(table, "pay_trade_no")
	o.TradeNo = field.NewString(table, "trade_no")
	o.TeamID = field.NewString(table, "team_id")
	o.GoogsName = field.NewString(table, "googs_name")
	o.OrderType = field.NewInt32(table, "order_type")
	o.VumBuyVersionType = field.NewInt32(table, "vum_buy_version_type")
	o.TeamBuyVersionType = field.NewInt32(table, "team_buy_version_type")
	o.MaxConcurrence = field.NewInt32(table, "max_concurrence")
	o.BuyNumber = field.NewInt64(table, "buy_number")
	o.OrderAmount = field.NewFloat64(table, "order_amount")
	o.Discounts = field.NewFloat64(table, "discounts")
	o.RealAmount = field.NewFloat64(table, "real_amount")
	o.PayType = field.NewInt32(table, "pay_type")
	o.PayStatus = field.NewInt32(table, "pay_status")
	o.GoodsValidDate = field.NewInt32(table, "goods_valid_date")
	o.FinishPayTime = field.NewTime(table, "finish_pay_time")
	o.OpenInvoiceState = field.NewInt32(table, "open_invoice_state")
	o.PayUserID = field.NewString(table, "pay_user_id")
	o.CreatedAt = field.NewTime(table, "created_at")
	o.UpdatedAt = field.NewTime(table, "updated_at")
	o.DeletedAt = field.NewField(table, "deleted_at")

	o.fillFieldMap()

	return o
}

func (o *order) WithContext(ctx context.Context) *orderDo { return o.orderDo.WithContext(ctx) }

func (o order) TableName() string { return o.orderDo.TableName() }

func (o order) Alias() string { return o.orderDo.Alias() }

func (o *order) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *order) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 23)
	o.fieldMap["id"] = o.ID
	o.fieldMap["order_id"] = o.OrderID
	o.fieldMap["pay_trade_no"] = o.PayTradeNo
	o.fieldMap["trade_no"] = o.TradeNo
	o.fieldMap["team_id"] = o.TeamID
	o.fieldMap["googs_name"] = o.GoogsName
	o.fieldMap["order_type"] = o.OrderType
	o.fieldMap["vum_buy_version_type"] = o.VumBuyVersionType
	o.fieldMap["team_buy_version_type"] = o.TeamBuyVersionType
	o.fieldMap["max_concurrence"] = o.MaxConcurrence
	o.fieldMap["buy_number"] = o.BuyNumber
	o.fieldMap["order_amount"] = o.OrderAmount
	o.fieldMap["discounts"] = o.Discounts
	o.fieldMap["real_amount"] = o.RealAmount
	o.fieldMap["pay_type"] = o.PayType
	o.fieldMap["pay_status"] = o.PayStatus
	o.fieldMap["goods_valid_date"] = o.GoodsValidDate
	o.fieldMap["finish_pay_time"] = o.FinishPayTime
	o.fieldMap["open_invoice_state"] = o.OpenInvoiceState
	o.fieldMap["pay_user_id"] = o.PayUserID
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
	o.fieldMap["deleted_at"] = o.DeletedAt
}

func (o order) clone(db *gorm.DB) order {
	o.orderDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o order) replaceDB(db *gorm.DB) order {
	o.orderDo.ReplaceDB(db)
	return o
}

type orderDo struct{ gen.DO }

func (o orderDo) Debug() *orderDo {
	return o.withDO(o.DO.Debug())
}

func (o orderDo) WithContext(ctx context.Context) *orderDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o orderDo) ReadDB() *orderDo {
	return o.Clauses(dbresolver.Read)
}

func (o orderDo) WriteDB() *orderDo {
	return o.Clauses(dbresolver.Write)
}

func (o orderDo) Session(config *gorm.Session) *orderDo {
	return o.withDO(o.DO.Session(config))
}

func (o orderDo) Clauses(conds ...clause.Expression) *orderDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o orderDo) Returning(value interface{}, columns ...string) *orderDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o orderDo) Not(conds ...gen.Condition) *orderDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o orderDo) Or(conds ...gen.Condition) *orderDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o orderDo) Select(conds ...field.Expr) *orderDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o orderDo) Where(conds ...gen.Condition) *orderDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o orderDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *orderDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o orderDo) Order(conds ...field.Expr) *orderDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o orderDo) Distinct(cols ...field.Expr) *orderDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o orderDo) Omit(cols ...field.Expr) *orderDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o orderDo) Join(table schema.Tabler, on ...field.Expr) *orderDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o orderDo) LeftJoin(table schema.Tabler, on ...field.Expr) *orderDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o orderDo) RightJoin(table schema.Tabler, on ...field.Expr) *orderDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o orderDo) Group(cols ...field.Expr) *orderDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o orderDo) Having(conds ...gen.Condition) *orderDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o orderDo) Limit(limit int) *orderDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o orderDo) Offset(offset int) *orderDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o orderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *orderDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o orderDo) Unscoped() *orderDo {
	return o.withDO(o.DO.Unscoped())
}

func (o orderDo) Create(values ...*model.Order) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o orderDo) CreateInBatches(values []*model.Order, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o orderDo) Save(values ...*model.Order) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o orderDo) First() (*model.Order, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) Take() (*model.Order, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) Last() (*model.Order, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) Find() ([]*model.Order, error) {
	result, err := o.DO.Find()
	return result.([]*model.Order), err
}

func (o orderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Order, err error) {
	buf := make([]*model.Order, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o orderDo) FindInBatches(result *[]*model.Order, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o orderDo) Attrs(attrs ...field.AssignExpr) *orderDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o orderDo) Assign(attrs ...field.AssignExpr) *orderDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o orderDo) Joins(fields ...field.RelationField) *orderDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o orderDo) Preload(fields ...field.RelationField) *orderDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o orderDo) FirstOrInit() (*model.Order, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) FirstOrCreate() (*model.Order, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Order), nil
	}
}

func (o orderDo) FindByPage(offset int, limit int) (result []*model.Order, count int64, err error) {
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

func (o orderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o orderDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o orderDo) Delete(models ...*model.Order) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *orderDo) withDO(do gen.Dao) *orderDo {
	o.DO = *do.(*gen.DO)
	return o
}
