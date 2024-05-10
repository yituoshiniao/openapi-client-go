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

	"github.com/yituoshiniao/gin-api-http/dao/model"
)

func newAppleProductPrice(db *gorm.DB, opts ...gen.DOOption) appleProductPrice {
	_appleProductPrice := appleProductPrice{}

	_appleProductPrice.appleProductPriceDo.UseDB(db, opts...)
	_appleProductPrice.appleProductPriceDo.UseModel(&model.AppleProductPrice{})

	tableName := _appleProductPrice.appleProductPriceDo.TableName()
	_appleProductPrice.ALL = field.NewAsterisk(tableName)
	_appleProductPrice.ID = field.NewUint64(tableName, "id")
	_appleProductPrice.Number = field.NewInt64(tableName, "number")
	_appleProductPrice.Channel = field.NewString(tableName, "channel")
	_appleProductPrice.CustomerPrice = field.NewFloat64(tableName, "customer_price")
	_appleProductPrice.Proceeds = field.NewFloat64(tableName, "proceeds")
	_appleProductPrice.ProceedsYear2 = field.NewFloat64(tableName, "proceeds_year2")
	_appleProductPrice.Territory = field.NewString(tableName, "territory")
	_appleProductPrice.Currency = field.NewString(tableName, "currency")
	_appleProductPrice.Name = field.NewString(tableName, "name")
	_appleProductPrice.FamilySharable = field.NewInt32(tableName, "family_sharable")
	_appleProductPrice.AvailableInAllTerritories = field.NewInt32(tableName, "available_in_all_territories")
	_appleProductPrice.InAppPurchaseType = field.NewString(tableName, "in_app_purchase_type")
	_appleProductPrice.ProductID = field.NewString(tableName, "product_id")
	_appleProductPrice.FirstPrice = field.NewFloat64(tableName, "first_price")
	_appleProductPrice.ManualPricesID = field.NewString(tableName, "manual_prices_id")
	_appleProductPrice.SubscriptionPricesID = field.NewString(tableName, "subscription_prices_id")
	_appleProductPrice.SubscriptionPricePointDataID = field.NewString(tableName, "subscription_price_point_data_id")
	_appleProductPrice.SubscriptionPeriod = field.NewString(tableName, "subscription_period")
	_appleProductPrice.GroupLevel = field.NewInt32(tableName, "group_level")
	_appleProductPrice.State = field.NewString(tableName, "state")
	_appleProductPrice.Type = field.NewInt32(tableName, "type")
	_appleProductPrice.IsDel = field.NewBool(tableName, "is_del")
	_appleProductPrice.FirstPriceDetail = field.NewString(tableName, "first_price_detail")
	_appleProductPrice.CreateTime = field.NewTime(tableName, "create_time")
	_appleProductPrice.UpdateTime = field.NewTime(tableName, "update_time")
	_appleProductPrice.OfferKeyDetail = field.NewString(tableName, "offer_key_detail")

	_appleProductPrice.fillFieldMap()

	return _appleProductPrice
}

// appleProductPrice 苹果购买项价格信息
type appleProductPrice struct {
	appleProductPriceDo appleProductPriceDo

	ALL                          field.Asterisk
	ID                           field.Uint64  // 主键id
	Number                       field.Int64   // 全局唯一number
	Channel                      field.String  // 渠道来源: cs (全能扫描王) | camexam (蜜蜂) 等
	CustomerPrice                field.Float64 // 购买项价格
	Proceeds                     field.Float64 // 首次收益
	ProceedsYear2                field.Float64 // 次年收益(订阅项)
	Territory                    field.String  // 国家地区码大写 三位(如: CHN  中国)
	Currency                     field.String  // 货币单位(如: CNY 人民币)
	Name                         field.String  // 购买项 显示名
	FamilySharable               field.Int32   // 是否 家庭共享: 0 否 ;1 是;
	AvailableInAllTerritories    field.Int32   // 是否 适用于所有地区: 0 否 ;1 是;
	InAppPurchaseType            field.String  // 非订阅项特有; 应用内购买类型 [NON_RENEWING_SUBSCRIPTION]; 消耗品: CONSUMABLE (如 传真、点卷 )
	ProductID                    field.String  // 购买项id
	FirstPrice                   field.Float64 // 首单优惠价格 (订阅项)
	ManualPricesID               field.String  //   非订阅类型 manualPricesId
	SubscriptionPricesID         field.String  //   订阅类型: SubscriptionPricesID [接口: v1/subscriptions/1641205786/prices 中的 data.id]
	SubscriptionPricePointDataID field.String  // 订阅类型 subscriptionsPricePointsId [接口: v1/subscriptions/1641205786/prices 中的 data.relationships.subscriptionPricePoint.data.id]
	SubscriptionPeriod           field.String  // 订阅周期 [年-ONE_YEAR]
	GroupLevel                   field.Int32   // 订阅项特有 组级别 [1]
	State                        field.String  // 售卖状态 APPROVED(唯一可以用), 其他:MISSING_METADATA, WAITING_FOR_UPLOAD, PROCESSING_CONTENT, READY_TO_SUBMIT, WAITING_FOR_REVIEW, IN_REVIEW
	Type                         field.Int32   // 订阅类型: 1 订阅型； 2 非订阅型, 3 消耗品
	IsDel                        field.Bool    // 是否删除: 1 是； 0 否; 删除表示 历史信息
	FirstPriceDetail             field.String  // 首单优惠详情(订阅项)
	CreateTime                   field.Time    // 创建时间
	UpdateTime                   field.Time    // 更新时间
	OfferKeyDetail               field.String  // promotion优惠详情

	fieldMap map[string]field.Expr
}

func (a appleProductPrice) Table(newTableName string) *appleProductPrice {
	a.appleProductPriceDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a appleProductPrice) As(alias string) *appleProductPrice {
	a.appleProductPriceDo.DO = *(a.appleProductPriceDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *appleProductPrice) updateTableName(table string) *appleProductPrice {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint64(table, "id")
	a.Number = field.NewInt64(table, "number")
	a.Channel = field.NewString(table, "channel")
	a.CustomerPrice = field.NewFloat64(table, "customer_price")
	a.Proceeds = field.NewFloat64(table, "proceeds")
	a.ProceedsYear2 = field.NewFloat64(table, "proceeds_year2")
	a.Territory = field.NewString(table, "territory")
	a.Currency = field.NewString(table, "currency")
	a.Name = field.NewString(table, "name")
	a.FamilySharable = field.NewInt32(table, "family_sharable")
	a.AvailableInAllTerritories = field.NewInt32(table, "available_in_all_territories")
	a.InAppPurchaseType = field.NewString(table, "in_app_purchase_type")
	a.ProductID = field.NewString(table, "product_id")
	a.FirstPrice = field.NewFloat64(table, "first_price")
	a.ManualPricesID = field.NewString(table, "manual_prices_id")
	a.SubscriptionPricesID = field.NewString(table, "subscription_prices_id")
	a.SubscriptionPricePointDataID = field.NewString(table, "subscription_price_point_data_id")
	a.SubscriptionPeriod = field.NewString(table, "subscription_period")
	a.GroupLevel = field.NewInt32(table, "group_level")
	a.State = field.NewString(table, "state")
	a.Type = field.NewInt32(table, "type")
	a.IsDel = field.NewBool(table, "is_del")
	a.FirstPriceDetail = field.NewString(table, "first_price_detail")
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")
	a.OfferKeyDetail = field.NewString(table, "offer_key_detail")

	a.fillFieldMap()

	return a
}

func (a *appleProductPrice) WithContext(ctx context.Context) *appleProductPriceDo {
	return a.appleProductPriceDo.WithContext(ctx)
}

func (a appleProductPrice) TableName() string { return a.appleProductPriceDo.TableName() }

func (a appleProductPrice) Alias() string { return a.appleProductPriceDo.Alias() }

func (a appleProductPrice) Columns(cols ...field.Expr) gen.Columns {
	return a.appleProductPriceDo.Columns(cols...)
}

func (a *appleProductPrice) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *appleProductPrice) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 26)
	a.fieldMap["id"] = a.ID
	a.fieldMap["number"] = a.Number
	a.fieldMap["channel"] = a.Channel
	a.fieldMap["customer_price"] = a.CustomerPrice
	a.fieldMap["proceeds"] = a.Proceeds
	a.fieldMap["proceeds_year2"] = a.ProceedsYear2
	a.fieldMap["territory"] = a.Territory
	a.fieldMap["currency"] = a.Currency
	a.fieldMap["name"] = a.Name
	a.fieldMap["family_sharable"] = a.FamilySharable
	a.fieldMap["available_in_all_territories"] = a.AvailableInAllTerritories
	a.fieldMap["in_app_purchase_type"] = a.InAppPurchaseType
	a.fieldMap["product_id"] = a.ProductID
	a.fieldMap["first_price"] = a.FirstPrice
	a.fieldMap["manual_prices_id"] = a.ManualPricesID
	a.fieldMap["subscription_prices_id"] = a.SubscriptionPricesID
	a.fieldMap["subscription_price_point_data_id"] = a.SubscriptionPricePointDataID
	a.fieldMap["subscription_period"] = a.SubscriptionPeriod
	a.fieldMap["group_level"] = a.GroupLevel
	a.fieldMap["state"] = a.State
	a.fieldMap["type"] = a.Type
	a.fieldMap["is_del"] = a.IsDel
	a.fieldMap["first_price_detail"] = a.FirstPriceDetail
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
	a.fieldMap["offer_key_detail"] = a.OfferKeyDetail
}

func (a appleProductPrice) clone(db *gorm.DB) appleProductPrice {
	a.appleProductPriceDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a appleProductPrice) replaceDB(db *gorm.DB) appleProductPrice {
	a.appleProductPriceDo.ReplaceDB(db)
	return a
}

type appleProductPriceDo struct{ gen.DO }

func (a appleProductPriceDo) Debug() *appleProductPriceDo {
	return a.withDO(a.DO.Debug())
}

func (a appleProductPriceDo) WithContext(ctx context.Context) *appleProductPriceDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a appleProductPriceDo) ReadDB() *appleProductPriceDo {
	return a.Clauses(dbresolver.Read)
}

func (a appleProductPriceDo) WriteDB() *appleProductPriceDo {
	return a.Clauses(dbresolver.Write)
}

func (a appleProductPriceDo) Session(config *gorm.Session) *appleProductPriceDo {
	return a.withDO(a.DO.Session(config))
}

func (a appleProductPriceDo) Clauses(conds ...clause.Expression) *appleProductPriceDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a appleProductPriceDo) Returning(value interface{}, columns ...string) *appleProductPriceDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a appleProductPriceDo) Not(conds ...gen.Condition) *appleProductPriceDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a appleProductPriceDo) Or(conds ...gen.Condition) *appleProductPriceDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a appleProductPriceDo) Select(conds ...field.Expr) *appleProductPriceDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a appleProductPriceDo) Where(conds ...gen.Condition) *appleProductPriceDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a appleProductPriceDo) Order(conds ...field.Expr) *appleProductPriceDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a appleProductPriceDo) Distinct(cols ...field.Expr) *appleProductPriceDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a appleProductPriceDo) Omit(cols ...field.Expr) *appleProductPriceDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a appleProductPriceDo) Join(table schema.Tabler, on ...field.Expr) *appleProductPriceDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a appleProductPriceDo) LeftJoin(table schema.Tabler, on ...field.Expr) *appleProductPriceDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a appleProductPriceDo) RightJoin(table schema.Tabler, on ...field.Expr) *appleProductPriceDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a appleProductPriceDo) Group(cols ...field.Expr) *appleProductPriceDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a appleProductPriceDo) Having(conds ...gen.Condition) *appleProductPriceDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a appleProductPriceDo) Limit(limit int) *appleProductPriceDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a appleProductPriceDo) Offset(offset int) *appleProductPriceDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a appleProductPriceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *appleProductPriceDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a appleProductPriceDo) Unscoped() *appleProductPriceDo {
	return a.withDO(a.DO.Unscoped())
}

func (a appleProductPriceDo) Create(values ...*model.AppleProductPrice) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a appleProductPriceDo) CreateInBatches(values []*model.AppleProductPrice, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a appleProductPriceDo) Save(values ...*model.AppleProductPrice) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a appleProductPriceDo) First() (*model.AppleProductPrice, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppleProductPrice), nil
	}
}

func (a appleProductPriceDo) Take() (*model.AppleProductPrice, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppleProductPrice), nil
	}
}

func (a appleProductPriceDo) Last() (*model.AppleProductPrice, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppleProductPrice), nil
	}
}

func (a appleProductPriceDo) Find() ([]*model.AppleProductPrice, error) {
	result, err := a.DO.Find()
	return result.([]*model.AppleProductPrice), err
}

func (a appleProductPriceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AppleProductPrice, err error) {
	buf := make([]*model.AppleProductPrice, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a appleProductPriceDo) FindInBatches(result *[]*model.AppleProductPrice, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a appleProductPriceDo) Attrs(attrs ...field.AssignExpr) *appleProductPriceDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a appleProductPriceDo) Assign(attrs ...field.AssignExpr) *appleProductPriceDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a appleProductPriceDo) Joins(fields ...field.RelationField) *appleProductPriceDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a appleProductPriceDo) Preload(fields ...field.RelationField) *appleProductPriceDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a appleProductPriceDo) FirstOrInit() (*model.AppleProductPrice, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppleProductPrice), nil
	}
}

func (a appleProductPriceDo) FirstOrCreate() (*model.AppleProductPrice, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppleProductPrice), nil
	}
}

func (a appleProductPriceDo) FindByPage(offset int, limit int) (result []*model.AppleProductPrice, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a appleProductPriceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a appleProductPriceDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a appleProductPriceDo) Delete(models ...*model.AppleProductPrice) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *appleProductPriceDo) withDO(do gen.Dao) *appleProductPriceDo {
	a.DO = *do.(*gen.DO)
	return a
}
