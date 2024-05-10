// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAppleProductPrice = "apple_product_price"

// AppleProductPrice 苹果购买项价格信息
type AppleProductPrice struct {
	ID                           uint64    `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true;comment:主键id" json:"id"`                                                                                                                                                         // 主键id
	Number                       *int64    `gorm:"column:number;type:bigint;comment:全局唯一number" json:"number"`                                                                                                                                                                                  // 全局唯一number
	Channel                      string    `gorm:"column:channel;type:varchar(255);not null;uniqueIndex:uk_product_id_channel_territory_create_time,priority:2;index:idx_product_id_channel_territory,priority:2;comment:渠道来源: cs (全能扫描王) | camexam (蜜蜂) 等" json:"channel"`                     // 渠道来源: cs (全能扫描王) | camexam (蜜蜂) 等
	CustomerPrice                float64   `gorm:"column:customer_price;type:decimal(10,2);not null;default:0.00;comment:购买项价格" json:"customerPrice"`                                                                                                                                           // 购买项价格
	Proceeds                     float64   `gorm:"column:proceeds;type:decimal(10,2);not null;default:0.00;comment:首次收益" json:"proceeds"`                                                                                                                                                       // 首次收益
	ProceedsYear2                float64   `gorm:"column:proceeds_year2;type:decimal(10,2);not null;default:0.00;comment:次年收益(订阅项)" json:"proceedsYear2"`                                                                                                                                       // 次年收益(订阅项)
	Territory                    string    `gorm:"column:territory;type:varchar(10);not null;uniqueIndex:uk_product_id_channel_territory_create_time,priority:3;index:idx_product_id_channel_territory,priority:3;comment:国家地区码大写 三位(如: CHN  中国)" json:"territory"`                             // 国家地区码大写 三位(如: CHN  中国)
	Currency                     string    `gorm:"column:currency;type:varchar(10);not null;comment:货币单位(如: CNY 人民币)" json:"currency"`                                                                                                                                                          // 货币单位(如: CNY 人民币)
	Name                         string    `gorm:"column:name;type:varchar(255);not null;comment:购买项 显示名" json:"name"`                                                                                                                                                                          // 购买项 显示名
	FamilySharable               int32     `gorm:"column:family_sharable;type:tinyint;not null;comment:是否 家庭共享: 0 否 ;1 是;" json:"familySharable"`                                                                                                                                               // 是否 家庭共享: 0 否 ;1 是;
	AvailableInAllTerritories    int32     `gorm:"column:available_in_all_territories;type:tinyint;not null;comment:是否 适用于所有地区: 0 否 ;1 是;" json:"availableInAllTerritories"`                                                                                                                    // 是否 适用于所有地区: 0 否 ;1 是;
	InAppPurchaseType            string    `gorm:"column:in_app_purchase_type;type:varchar(50);not null;comment:非订阅项特有; 应用内购买类型 [NON_RENEWING_SUBSCRIPTION]; 消耗品: CONSUMABLE (如 传真、点卷 )" json:"inAppPurchaseType"`                                                                              // 非订阅项特有; 应用内购买类型 [NON_RENEWING_SUBSCRIPTION]; 消耗品: CONSUMABLE (如 传真、点卷 )
	ProductID                    string    `gorm:"column:product_id;type:varchar(200);not null;uniqueIndex:uk_product_id_channel_territory_create_time,priority:1;index:idx_product_id_channel_territory,priority:1;comment:购买项id" json:"productId"`                                            // 购买项id
	FirstPrice                   float64   `gorm:"column:first_price;type:decimal(10,2);not null;default:0.00;comment:首单优惠价格 (订阅项)" json:"firstPrice"`                                                                                                                                          // 首单优惠价格 (订阅项)
	ManualPricesID               string    `gorm:"column:manual_prices_id;type:varchar(80);not null;comment:  非订阅类型 manualPricesId" json:"manualPricesId"`                                                                                                                                      //   非订阅类型 manualPricesId
	SubscriptionPricesID         string    `gorm:"column:subscription_prices_id;type:varchar(80);not null;comment:  订阅类型: SubscriptionPricesID [接口: v1/subscriptions/1641205786/prices 中的 data.id]" json:"subscriptionPricesId"`                                                                //   订阅类型: SubscriptionPricesID [接口: v1/subscriptions/1641205786/prices 中的 data.id]
	SubscriptionPricePointDataID string    `gorm:"column:subscription_price_point_data_id;type:varchar(80);not null;comment:订阅类型 subscriptionsPricePointsId [接口: v1/subscriptions/1641205786/prices 中的 data.relationships.subscriptionPricePoint.data.id]" json:"subscriptionPricePointDataId"` // 订阅类型 subscriptionsPricePointsId [接口: v1/subscriptions/1641205786/prices 中的 data.relationships.subscriptionPricePoint.data.id]
	SubscriptionPeriod           string    `gorm:"column:subscription_period;type:varchar(80);not null;comment:订阅周期 [年-ONE_YEAR]" json:"subscriptionPeriod"`                                                                                                                                    // 订阅周期 [年-ONE_YEAR]
	GroupLevel                   int32     `gorm:"column:group_level;type:tinyint;not null;comment:订阅项特有 组级别 [1]" json:"groupLevel"`                                                                                                                                                            // 订阅项特有 组级别 [1]
	State                        string    `gorm:"column:state;type:varchar(80);not null;comment:售卖状态 APPROVED(唯一可以用), 其他:MISSING_METADATA, WAITING_FOR_UPLOAD, PROCESSING_CONTENT, READY_TO_SUBMIT, WAITING_FOR_REVIEW, IN_REVIEW" json:"state"`                                               // 售卖状态 APPROVED(唯一可以用), 其他:MISSING_METADATA, WAITING_FOR_UPLOAD, PROCESSING_CONTENT, READY_TO_SUBMIT, WAITING_FOR_REVIEW, IN_REVIEW
	Type                         int32     `gorm:"column:type;type:tinyint;not null;comment:订阅类型: 1 订阅型； 2 非订阅型, 3 消耗品" json:"type"`                                                                                                                                                            // 订阅类型: 1 订阅型； 2 非订阅型, 3 消耗品
	IsDel                        bool      `gorm:"column:is_del;type:tinyint(1);not null;comment:是否删除: 1 是； 0 否; 删除表示 历史信息" json:"isDel"`                                                                                                                                                       // 是否删除: 1 是； 0 否; 删除表示 历史信息
	FirstPriceDetail             string    `gorm:"column:first_price_detail;type:varchar(500);not null;comment:首单优惠详情(订阅项)" json:"firstPriceDetail"`                                                                                                                                            // 首单优惠详情(订阅项)
	CreateTime                   time.Time `gorm:"column:create_time;type:datetime;not null;uniqueIndex:uk_product_id_channel_territory_create_time,priority:4;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"`                                                                       // 创建时间
	UpdateTime                   time.Time `gorm:"column:update_time;type:datetime;not null;comment:更新时间" json:"updateTime"`                                                                                                                                                           // 更新时间
	OfferKeyDetail               string    `gorm:"column:offer_key_detail;type:varchar(1000);not null;comment:promotion优惠详情" json:"offer_key_detail"`                                                                                                                                               // promotion优惠详情
}

// TableName AppleProductPrice's table name
func (*AppleProductPrice) TableName() string {
	return TableNameAppleProductPrice
}
