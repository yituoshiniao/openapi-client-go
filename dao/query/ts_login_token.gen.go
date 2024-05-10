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

func newTsLoginToken(db *gorm.DB, opts ...gen.DOOption) tsLoginToken {
	_tsLoginToken := tsLoginToken{}

	_tsLoginToken.tsLoginTokenDo.UseDB(db, opts...)
	_tsLoginToken.tsLoginTokenDo.UseModel(&model.TsLoginToken{})

	tableName := _tsLoginToken.tsLoginTokenDo.TableName()
	_tsLoginToken.ALL = field.NewAsterisk(tableName)
	_tsLoginToken.Token = field.NewString(tableName, "token")
	_tsLoginToken.UserID = field.NewInt64(tableName, "user_id")
	_tsLoginToken.TimeLogin = field.NewInt64(tableName, "time_login")
	_tsLoginToken.TimeExpire = field.NewInt64(tableName, "time_expire")
	_tsLoginToken.ClientIP = field.NewString(tableName, "client_ip")
	_tsLoginToken.ClientName = field.NewString(tableName, "client_name")
	_tsLoginToken.ClientID = field.NewString(tableName, "client_id")
	_tsLoginToken.NetID = field.NewInt32(tableName, "net_id")
	_tsLoginToken.ClientApp = field.NewString(tableName, "client_app")
	_tsLoginToken.Status = field.NewInt32(tableName, "status")
	_tsLoginToken.ID = field.NewInt64(tableName, "id")

	_tsLoginToken.fillFieldMap()

	return _tsLoginToken
}

type tsLoginToken struct {
	tsLoginTokenDo tsLoginTokenDo

	ALL        field.Asterisk
	Token      field.String
	UserID     field.Int64
	TimeLogin  field.Int64
	TimeExpire field.Int64
	ClientIP   field.String
	ClientName field.String
	ClientID   field.String
	NetID      field.Int32
	ClientApp  field.String
	Status     field.Int32
	ID         field.Int64

	fieldMap map[string]field.Expr
}

func (t tsLoginToken) Table(newTableName string) *tsLoginToken {
	t.tsLoginTokenDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tsLoginToken) As(alias string) *tsLoginToken {
	t.tsLoginTokenDo.DO = *(t.tsLoginTokenDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tsLoginToken) updateTableName(table string) *tsLoginToken {
	t.ALL = field.NewAsterisk(table)
	t.Token = field.NewString(table, "token")
	t.UserID = field.NewInt64(table, "user_id")
	t.TimeLogin = field.NewInt64(table, "time_login")
	t.TimeExpire = field.NewInt64(table, "time_expire")
	t.ClientIP = field.NewString(table, "client_ip")
	t.ClientName = field.NewString(table, "client_name")
	t.ClientID = field.NewString(table, "client_id")
	t.NetID = field.NewInt32(table, "net_id")
	t.ClientApp = field.NewString(table, "client_app")
	t.Status = field.NewInt32(table, "status")
	t.ID = field.NewInt64(table, "id")

	t.fillFieldMap()

	return t
}

func (t *tsLoginToken) WithContext(ctx context.Context) *tsLoginTokenDo {
	return t.tsLoginTokenDo.WithContext(ctx)
}

func (t tsLoginToken) TableName() string { return t.tsLoginTokenDo.TableName() }

func (t tsLoginToken) Alias() string { return t.tsLoginTokenDo.Alias() }

func (t tsLoginToken) Columns(cols ...field.Expr) gen.Columns {
	return t.tsLoginTokenDo.Columns(cols...)
}

func (t *tsLoginToken) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tsLoginToken) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 11)
	t.fieldMap["token"] = t.Token
	t.fieldMap["user_id"] = t.UserID
	t.fieldMap["time_login"] = t.TimeLogin
	t.fieldMap["time_expire"] = t.TimeExpire
	t.fieldMap["client_ip"] = t.ClientIP
	t.fieldMap["client_name"] = t.ClientName
	t.fieldMap["client_id"] = t.ClientID
	t.fieldMap["net_id"] = t.NetID
	t.fieldMap["client_app"] = t.ClientApp
	t.fieldMap["status"] = t.Status
	t.fieldMap["id"] = t.ID
}

func (t tsLoginToken) clone(db *gorm.DB) tsLoginToken {
	t.tsLoginTokenDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tsLoginToken) replaceDB(db *gorm.DB) tsLoginToken {
	t.tsLoginTokenDo.ReplaceDB(db)
	return t
}

type tsLoginTokenDo struct{ gen.DO }

func (t tsLoginTokenDo) Debug() *tsLoginTokenDo {
	return t.withDO(t.DO.Debug())
}

func (t tsLoginTokenDo) WithContext(ctx context.Context) *tsLoginTokenDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tsLoginTokenDo) ReadDB() *tsLoginTokenDo {
	return t.Clauses(dbresolver.Read)
}

func (t tsLoginTokenDo) WriteDB() *tsLoginTokenDo {
	return t.Clauses(dbresolver.Write)
}

func (t tsLoginTokenDo) Session(config *gorm.Session) *tsLoginTokenDo {
	return t.withDO(t.DO.Session(config))
}

func (t tsLoginTokenDo) Clauses(conds ...clause.Expression) *tsLoginTokenDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tsLoginTokenDo) Returning(value interface{}, columns ...string) *tsLoginTokenDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tsLoginTokenDo) Not(conds ...gen.Condition) *tsLoginTokenDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tsLoginTokenDo) Or(conds ...gen.Condition) *tsLoginTokenDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tsLoginTokenDo) Select(conds ...field.Expr) *tsLoginTokenDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tsLoginTokenDo) Where(conds ...gen.Condition) *tsLoginTokenDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tsLoginTokenDo) Order(conds ...field.Expr) *tsLoginTokenDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tsLoginTokenDo) Distinct(cols ...field.Expr) *tsLoginTokenDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tsLoginTokenDo) Omit(cols ...field.Expr) *tsLoginTokenDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tsLoginTokenDo) Join(table schema.Tabler, on ...field.Expr) *tsLoginTokenDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tsLoginTokenDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tsLoginTokenDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tsLoginTokenDo) RightJoin(table schema.Tabler, on ...field.Expr) *tsLoginTokenDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tsLoginTokenDo) Group(cols ...field.Expr) *tsLoginTokenDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tsLoginTokenDo) Having(conds ...gen.Condition) *tsLoginTokenDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tsLoginTokenDo) Limit(limit int) *tsLoginTokenDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tsLoginTokenDo) Offset(offset int) *tsLoginTokenDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tsLoginTokenDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tsLoginTokenDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tsLoginTokenDo) Unscoped() *tsLoginTokenDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tsLoginTokenDo) Create(values ...*model.TsLoginToken) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tsLoginTokenDo) CreateInBatches(values []*model.TsLoginToken, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tsLoginTokenDo) Save(values ...*model.TsLoginToken) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tsLoginTokenDo) First() (*model.TsLoginToken, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TsLoginToken), nil
	}
}

func (t tsLoginTokenDo) Take() (*model.TsLoginToken, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TsLoginToken), nil
	}
}

func (t tsLoginTokenDo) Last() (*model.TsLoginToken, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TsLoginToken), nil
	}
}

func (t tsLoginTokenDo) Find() ([]*model.TsLoginToken, error) {
	result, err := t.DO.Find()
	return result.([]*model.TsLoginToken), err
}

func (t tsLoginTokenDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TsLoginToken, err error) {
	buf := make([]*model.TsLoginToken, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tsLoginTokenDo) FindInBatches(result *[]*model.TsLoginToken, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tsLoginTokenDo) Attrs(attrs ...field.AssignExpr) *tsLoginTokenDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tsLoginTokenDo) Assign(attrs ...field.AssignExpr) *tsLoginTokenDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tsLoginTokenDo) Joins(fields ...field.RelationField) *tsLoginTokenDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tsLoginTokenDo) Preload(fields ...field.RelationField) *tsLoginTokenDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tsLoginTokenDo) FirstOrInit() (*model.TsLoginToken, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TsLoginToken), nil
	}
}

func (t tsLoginTokenDo) FirstOrCreate() (*model.TsLoginToken, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TsLoginToken), nil
	}
}

func (t tsLoginTokenDo) FindByPage(offset int, limit int) (result []*model.TsLoginToken, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tsLoginTokenDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tsLoginTokenDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tsLoginTokenDo) Delete(models ...*model.TsLoginToken) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tsLoginTokenDo) withDO(do gen.Dao) *tsLoginTokenDo {
	t.DO = *do.(*gen.DO)
	return t
}
