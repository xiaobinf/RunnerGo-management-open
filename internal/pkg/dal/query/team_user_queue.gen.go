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

func newTeamUserQueue(db *gorm.DB, opts ...gen.DOOption) teamUserQueue {
	_teamUserQueue := teamUserQueue{}

	_teamUserQueue.teamUserQueueDo.UseDB(db, opts...)
	_teamUserQueue.teamUserQueueDo.UseModel(&model.TeamUserQueue{})

	tableName := _teamUserQueue.teamUserQueueDo.TableName()
	_teamUserQueue.ALL = field.NewAsterisk(tableName)
	_teamUserQueue.ID = field.NewInt64(tableName, "id")
	_teamUserQueue.Email = field.NewString(tableName, "email")
	_teamUserQueue.TeamID = field.NewString(tableName, "team_id")
	_teamUserQueue.CreatedAt = field.NewTime(tableName, "created_at")
	_teamUserQueue.UpdatedAt = field.NewTime(tableName, "updated_at")
	_teamUserQueue.DeletedAt = field.NewField(tableName, "deleted_at")

	_teamUserQueue.fillFieldMap()

	return _teamUserQueue
}

type teamUserQueue struct {
	teamUserQueueDo teamUserQueueDo

	ALL       field.Asterisk
	ID        field.Int64
	Email     field.String // 邮箱
	TeamID    field.String // 团队id
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (t teamUserQueue) Table(newTableName string) *teamUserQueue {
	t.teamUserQueueDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t teamUserQueue) As(alias string) *teamUserQueue {
	t.teamUserQueueDo.DO = *(t.teamUserQueueDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *teamUserQueue) updateTableName(table string) *teamUserQueue {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.Email = field.NewString(table, "email")
	t.TeamID = field.NewString(table, "team_id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *teamUserQueue) WithContext(ctx context.Context) *teamUserQueueDo {
	return t.teamUserQueueDo.WithContext(ctx)
}

func (t teamUserQueue) TableName() string { return t.teamUserQueueDo.TableName() }

func (t teamUserQueue) Alias() string { return t.teamUserQueueDo.Alias() }

func (t *teamUserQueue) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *teamUserQueue) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["id"] = t.ID
	t.fieldMap["email"] = t.Email
	t.fieldMap["team_id"] = t.TeamID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t teamUserQueue) clone(db *gorm.DB) teamUserQueue {
	t.teamUserQueueDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t teamUserQueue) replaceDB(db *gorm.DB) teamUserQueue {
	t.teamUserQueueDo.ReplaceDB(db)
	return t
}

type teamUserQueueDo struct{ gen.DO }

func (t teamUserQueueDo) Debug() *teamUserQueueDo {
	return t.withDO(t.DO.Debug())
}

func (t teamUserQueueDo) WithContext(ctx context.Context) *teamUserQueueDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t teamUserQueueDo) ReadDB() *teamUserQueueDo {
	return t.Clauses(dbresolver.Read)
}

func (t teamUserQueueDo) WriteDB() *teamUserQueueDo {
	return t.Clauses(dbresolver.Write)
}

func (t teamUserQueueDo) Session(config *gorm.Session) *teamUserQueueDo {
	return t.withDO(t.DO.Session(config))
}

func (t teamUserQueueDo) Clauses(conds ...clause.Expression) *teamUserQueueDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t teamUserQueueDo) Returning(value interface{}, columns ...string) *teamUserQueueDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t teamUserQueueDo) Not(conds ...gen.Condition) *teamUserQueueDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t teamUserQueueDo) Or(conds ...gen.Condition) *teamUserQueueDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t teamUserQueueDo) Select(conds ...field.Expr) *teamUserQueueDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t teamUserQueueDo) Where(conds ...gen.Condition) *teamUserQueueDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t teamUserQueueDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *teamUserQueueDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t teamUserQueueDo) Order(conds ...field.Expr) *teamUserQueueDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t teamUserQueueDo) Distinct(cols ...field.Expr) *teamUserQueueDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t teamUserQueueDo) Omit(cols ...field.Expr) *teamUserQueueDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t teamUserQueueDo) Join(table schema.Tabler, on ...field.Expr) *teamUserQueueDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t teamUserQueueDo) LeftJoin(table schema.Tabler, on ...field.Expr) *teamUserQueueDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t teamUserQueueDo) RightJoin(table schema.Tabler, on ...field.Expr) *teamUserQueueDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t teamUserQueueDo) Group(cols ...field.Expr) *teamUserQueueDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t teamUserQueueDo) Having(conds ...gen.Condition) *teamUserQueueDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t teamUserQueueDo) Limit(limit int) *teamUserQueueDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t teamUserQueueDo) Offset(offset int) *teamUserQueueDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t teamUserQueueDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *teamUserQueueDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t teamUserQueueDo) Unscoped() *teamUserQueueDo {
	return t.withDO(t.DO.Unscoped())
}

func (t teamUserQueueDo) Create(values ...*model.TeamUserQueue) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t teamUserQueueDo) CreateInBatches(values []*model.TeamUserQueue, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t teamUserQueueDo) Save(values ...*model.TeamUserQueue) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t teamUserQueueDo) First() (*model.TeamUserQueue, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamUserQueue), nil
	}
}

func (t teamUserQueueDo) Take() (*model.TeamUserQueue, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamUserQueue), nil
	}
}

func (t teamUserQueueDo) Last() (*model.TeamUserQueue, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamUserQueue), nil
	}
}

func (t teamUserQueueDo) Find() ([]*model.TeamUserQueue, error) {
	result, err := t.DO.Find()
	return result.([]*model.TeamUserQueue), err
}

func (t teamUserQueueDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TeamUserQueue, err error) {
	buf := make([]*model.TeamUserQueue, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t teamUserQueueDo) FindInBatches(result *[]*model.TeamUserQueue, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t teamUserQueueDo) Attrs(attrs ...field.AssignExpr) *teamUserQueueDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t teamUserQueueDo) Assign(attrs ...field.AssignExpr) *teamUserQueueDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t teamUserQueueDo) Joins(fields ...field.RelationField) *teamUserQueueDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t teamUserQueueDo) Preload(fields ...field.RelationField) *teamUserQueueDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t teamUserQueueDo) FirstOrInit() (*model.TeamUserQueue, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamUserQueue), nil
	}
}

func (t teamUserQueueDo) FirstOrCreate() (*model.TeamUserQueue, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamUserQueue), nil
	}
}

func (t teamUserQueueDo) FindByPage(offset int, limit int) (result []*model.TeamUserQueue, count int64, err error) {
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

func (t teamUserQueueDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t teamUserQueueDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t teamUserQueueDo) Delete(models ...*model.TeamUserQueue) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *teamUserQueueDo) withDO(do gen.Dao) *teamUserQueueDo {
	t.DO = *do.(*gen.DO)
	return t
}
