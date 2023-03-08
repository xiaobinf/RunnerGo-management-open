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

func newTargetDebugLog(db *gorm.DB, opts ...gen.DOOption) targetDebugLog {
	_targetDebugLog := targetDebugLog{}

	_targetDebugLog.targetDebugLogDo.UseDB(db, opts...)
	_targetDebugLog.targetDebugLogDo.UseModel(&model.TargetDebugLog{})

	tableName := _targetDebugLog.targetDebugLogDo.TableName()
	_targetDebugLog.ALL = field.NewAsterisk(tableName)
	_targetDebugLog.ID = field.NewInt64(tableName, "id")
	_targetDebugLog.TargetID = field.NewString(tableName, "target_id")
	_targetDebugLog.TargetType = field.NewInt32(tableName, "target_type")
	_targetDebugLog.TeamID = field.NewString(tableName, "team_id")
	_targetDebugLog.CreatedAt = field.NewTime(tableName, "created_at")
	_targetDebugLog.UpdatedAt = field.NewTime(tableName, "updated_at")
	_targetDebugLog.DeletedAt = field.NewField(tableName, "deleted_at")

	_targetDebugLog.fillFieldMap()

	return _targetDebugLog
}

type targetDebugLog struct {
	targetDebugLogDo targetDebugLogDo

	ALL        field.Asterisk
	ID         field.Int64  // 主键ID
	TargetID   field.String // 目标唯一ID
	TargetType field.Int32  // 目标类型：1-api，2-scene
	TeamID     field.String // 团队ID
	CreatedAt  field.Time   // 创建时间
	UpdatedAt  field.Time   // 更新时间
	DeletedAt  field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (t targetDebugLog) Table(newTableName string) *targetDebugLog {
	t.targetDebugLogDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t targetDebugLog) As(alias string) *targetDebugLog {
	t.targetDebugLogDo.DO = *(t.targetDebugLogDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *targetDebugLog) updateTableName(table string) *targetDebugLog {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.TargetID = field.NewString(table, "target_id")
	t.TargetType = field.NewInt32(table, "target_type")
	t.TeamID = field.NewString(table, "team_id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *targetDebugLog) WithContext(ctx context.Context) *targetDebugLogDo {
	return t.targetDebugLogDo.WithContext(ctx)
}

func (t targetDebugLog) TableName() string { return t.targetDebugLogDo.TableName() }

func (t targetDebugLog) Alias() string { return t.targetDebugLogDo.Alias() }

func (t *targetDebugLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *targetDebugLog) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["id"] = t.ID
	t.fieldMap["target_id"] = t.TargetID
	t.fieldMap["target_type"] = t.TargetType
	t.fieldMap["team_id"] = t.TeamID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t targetDebugLog) clone(db *gorm.DB) targetDebugLog {
	t.targetDebugLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t targetDebugLog) replaceDB(db *gorm.DB) targetDebugLog {
	t.targetDebugLogDo.ReplaceDB(db)
	return t
}

type targetDebugLogDo struct{ gen.DO }

func (t targetDebugLogDo) Debug() *targetDebugLogDo {
	return t.withDO(t.DO.Debug())
}

func (t targetDebugLogDo) WithContext(ctx context.Context) *targetDebugLogDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t targetDebugLogDo) ReadDB() *targetDebugLogDo {
	return t.Clauses(dbresolver.Read)
}

func (t targetDebugLogDo) WriteDB() *targetDebugLogDo {
	return t.Clauses(dbresolver.Write)
}

func (t targetDebugLogDo) Session(config *gorm.Session) *targetDebugLogDo {
	return t.withDO(t.DO.Session(config))
}

func (t targetDebugLogDo) Clauses(conds ...clause.Expression) *targetDebugLogDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t targetDebugLogDo) Returning(value interface{}, columns ...string) *targetDebugLogDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t targetDebugLogDo) Not(conds ...gen.Condition) *targetDebugLogDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t targetDebugLogDo) Or(conds ...gen.Condition) *targetDebugLogDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t targetDebugLogDo) Select(conds ...field.Expr) *targetDebugLogDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t targetDebugLogDo) Where(conds ...gen.Condition) *targetDebugLogDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t targetDebugLogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *targetDebugLogDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t targetDebugLogDo) Order(conds ...field.Expr) *targetDebugLogDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t targetDebugLogDo) Distinct(cols ...field.Expr) *targetDebugLogDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t targetDebugLogDo) Omit(cols ...field.Expr) *targetDebugLogDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t targetDebugLogDo) Join(table schema.Tabler, on ...field.Expr) *targetDebugLogDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t targetDebugLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *targetDebugLogDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t targetDebugLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *targetDebugLogDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t targetDebugLogDo) Group(cols ...field.Expr) *targetDebugLogDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t targetDebugLogDo) Having(conds ...gen.Condition) *targetDebugLogDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t targetDebugLogDo) Limit(limit int) *targetDebugLogDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t targetDebugLogDo) Offset(offset int) *targetDebugLogDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t targetDebugLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *targetDebugLogDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t targetDebugLogDo) Unscoped() *targetDebugLogDo {
	return t.withDO(t.DO.Unscoped())
}

func (t targetDebugLogDo) Create(values ...*model.TargetDebugLog) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t targetDebugLogDo) CreateInBatches(values []*model.TargetDebugLog, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t targetDebugLogDo) Save(values ...*model.TargetDebugLog) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t targetDebugLogDo) First() (*model.TargetDebugLog, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TargetDebugLog), nil
	}
}

func (t targetDebugLogDo) Take() (*model.TargetDebugLog, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TargetDebugLog), nil
	}
}

func (t targetDebugLogDo) Last() (*model.TargetDebugLog, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TargetDebugLog), nil
	}
}

func (t targetDebugLogDo) Find() ([]*model.TargetDebugLog, error) {
	result, err := t.DO.Find()
	return result.([]*model.TargetDebugLog), err
}

func (t targetDebugLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TargetDebugLog, err error) {
	buf := make([]*model.TargetDebugLog, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t targetDebugLogDo) FindInBatches(result *[]*model.TargetDebugLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t targetDebugLogDo) Attrs(attrs ...field.AssignExpr) *targetDebugLogDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t targetDebugLogDo) Assign(attrs ...field.AssignExpr) *targetDebugLogDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t targetDebugLogDo) Joins(fields ...field.RelationField) *targetDebugLogDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t targetDebugLogDo) Preload(fields ...field.RelationField) *targetDebugLogDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t targetDebugLogDo) FirstOrInit() (*model.TargetDebugLog, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TargetDebugLog), nil
	}
}

func (t targetDebugLogDo) FirstOrCreate() (*model.TargetDebugLog, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TargetDebugLog), nil
	}
}

func (t targetDebugLogDo) FindByPage(offset int, limit int) (result []*model.TargetDebugLog, count int64, err error) {
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

func (t targetDebugLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t targetDebugLogDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t targetDebugLogDo) Delete(models ...*model.TargetDebugLog) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *targetDebugLogDo) withDO(do gen.Dao) *targetDebugLogDo {
	t.DO = *do.(*gen.DO)
	return t
}
