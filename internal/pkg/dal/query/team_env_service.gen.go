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

func newTeamEnvService(db *gorm.DB, opts ...gen.DOOption) teamEnvService {
	_teamEnvService := teamEnvService{}

	_teamEnvService.teamEnvServiceDo.UseDB(db, opts...)
	_teamEnvService.teamEnvServiceDo.UseModel(&model.TeamEnvService{})

	tableName := _teamEnvService.teamEnvServiceDo.TableName()
	_teamEnvService.ALL = field.NewAsterisk(tableName)
	_teamEnvService.ID = field.NewInt64(tableName, "id")
	_teamEnvService.TeamID = field.NewString(tableName, "team_id")
	_teamEnvService.TeamEnvID = field.NewInt64(tableName, "team_env_id")
	_teamEnvService.Name = field.NewString(tableName, "name")
	_teamEnvService.Content = field.NewString(tableName, "content")
	_teamEnvService.Sort = field.NewInt32(tableName, "sort")
	_teamEnvService.Status = field.NewInt32(tableName, "status")
	_teamEnvService.CreatedUserID = field.NewString(tableName, "created_user_id")
	_teamEnvService.RecentUserID = field.NewInt64(tableName, "recent_user_id")
	_teamEnvService.CreatedAt = field.NewTime(tableName, "created_at")
	_teamEnvService.UpdatedAt = field.NewTime(tableName, "updated_at")
	_teamEnvService.DeletedAt = field.NewField(tableName, "deleted_at")

	_teamEnvService.fillFieldMap()

	return _teamEnvService
}

type teamEnvService struct {
	teamEnvServiceDo teamEnvServiceDo

	ALL           field.Asterisk
	ID            field.Int64  // ID
	TeamID        field.String // 团队ID
	TeamEnvID     field.Int64  // 环境ID
	Name          field.String // 服务名称
	Content       field.String // 服务URL
	Sort          field.Int32  // 排序
	Status        field.Int32  // 状态：1-正常 2-删除
	CreatedUserID field.String // 创建人ID
	RecentUserID  field.Int64
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field

	fieldMap map[string]field.Expr
}

func (t teamEnvService) Table(newTableName string) *teamEnvService {
	t.teamEnvServiceDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t teamEnvService) As(alias string) *teamEnvService {
	t.teamEnvServiceDo.DO = *(t.teamEnvServiceDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *teamEnvService) updateTableName(table string) *teamEnvService {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.TeamID = field.NewString(table, "team_id")
	t.TeamEnvID = field.NewInt64(table, "team_env_id")
	t.Name = field.NewString(table, "name")
	t.Content = field.NewString(table, "content")
	t.Sort = field.NewInt32(table, "sort")
	t.Status = field.NewInt32(table, "status")
	t.CreatedUserID = field.NewString(table, "created_user_id")
	t.RecentUserID = field.NewInt64(table, "recent_user_id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *teamEnvService) WithContext(ctx context.Context) *teamEnvServiceDo {
	return t.teamEnvServiceDo.WithContext(ctx)
}

func (t teamEnvService) TableName() string { return t.teamEnvServiceDo.TableName() }

func (t teamEnvService) Alias() string { return t.teamEnvServiceDo.Alias() }

func (t *teamEnvService) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *teamEnvService) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 12)
	t.fieldMap["id"] = t.ID
	t.fieldMap["team_id"] = t.TeamID
	t.fieldMap["team_env_id"] = t.TeamEnvID
	t.fieldMap["name"] = t.Name
	t.fieldMap["content"] = t.Content
	t.fieldMap["sort"] = t.Sort
	t.fieldMap["status"] = t.Status
	t.fieldMap["created_user_id"] = t.CreatedUserID
	t.fieldMap["recent_user_id"] = t.RecentUserID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t teamEnvService) clone(db *gorm.DB) teamEnvService {
	t.teamEnvServiceDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t teamEnvService) replaceDB(db *gorm.DB) teamEnvService {
	t.teamEnvServiceDo.ReplaceDB(db)
	return t
}

type teamEnvServiceDo struct{ gen.DO }

func (t teamEnvServiceDo) Debug() *teamEnvServiceDo {
	return t.withDO(t.DO.Debug())
}

func (t teamEnvServiceDo) WithContext(ctx context.Context) *teamEnvServiceDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t teamEnvServiceDo) ReadDB() *teamEnvServiceDo {
	return t.Clauses(dbresolver.Read)
}

func (t teamEnvServiceDo) WriteDB() *teamEnvServiceDo {
	return t.Clauses(dbresolver.Write)
}

func (t teamEnvServiceDo) Session(config *gorm.Session) *teamEnvServiceDo {
	return t.withDO(t.DO.Session(config))
}

func (t teamEnvServiceDo) Clauses(conds ...clause.Expression) *teamEnvServiceDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t teamEnvServiceDo) Returning(value interface{}, columns ...string) *teamEnvServiceDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t teamEnvServiceDo) Not(conds ...gen.Condition) *teamEnvServiceDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t teamEnvServiceDo) Or(conds ...gen.Condition) *teamEnvServiceDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t teamEnvServiceDo) Select(conds ...field.Expr) *teamEnvServiceDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t teamEnvServiceDo) Where(conds ...gen.Condition) *teamEnvServiceDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t teamEnvServiceDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *teamEnvServiceDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t teamEnvServiceDo) Order(conds ...field.Expr) *teamEnvServiceDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t teamEnvServiceDo) Distinct(cols ...field.Expr) *teamEnvServiceDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t teamEnvServiceDo) Omit(cols ...field.Expr) *teamEnvServiceDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t teamEnvServiceDo) Join(table schema.Tabler, on ...field.Expr) *teamEnvServiceDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t teamEnvServiceDo) LeftJoin(table schema.Tabler, on ...field.Expr) *teamEnvServiceDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t teamEnvServiceDo) RightJoin(table schema.Tabler, on ...field.Expr) *teamEnvServiceDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t teamEnvServiceDo) Group(cols ...field.Expr) *teamEnvServiceDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t teamEnvServiceDo) Having(conds ...gen.Condition) *teamEnvServiceDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t teamEnvServiceDo) Limit(limit int) *teamEnvServiceDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t teamEnvServiceDo) Offset(offset int) *teamEnvServiceDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t teamEnvServiceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *teamEnvServiceDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t teamEnvServiceDo) Unscoped() *teamEnvServiceDo {
	return t.withDO(t.DO.Unscoped())
}

func (t teamEnvServiceDo) Create(values ...*model.TeamEnvService) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t teamEnvServiceDo) CreateInBatches(values []*model.TeamEnvService, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t teamEnvServiceDo) Save(values ...*model.TeamEnvService) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t teamEnvServiceDo) First() (*model.TeamEnvService, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamEnvService), nil
	}
}

func (t teamEnvServiceDo) Take() (*model.TeamEnvService, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamEnvService), nil
	}
}

func (t teamEnvServiceDo) Last() (*model.TeamEnvService, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamEnvService), nil
	}
}

func (t teamEnvServiceDo) Find() ([]*model.TeamEnvService, error) {
	result, err := t.DO.Find()
	return result.([]*model.TeamEnvService), err
}

func (t teamEnvServiceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TeamEnvService, err error) {
	buf := make([]*model.TeamEnvService, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t teamEnvServiceDo) FindInBatches(result *[]*model.TeamEnvService, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t teamEnvServiceDo) Attrs(attrs ...field.AssignExpr) *teamEnvServiceDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t teamEnvServiceDo) Assign(attrs ...field.AssignExpr) *teamEnvServiceDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t teamEnvServiceDo) Joins(fields ...field.RelationField) *teamEnvServiceDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t teamEnvServiceDo) Preload(fields ...field.RelationField) *teamEnvServiceDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t teamEnvServiceDo) FirstOrInit() (*model.TeamEnvService, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamEnvService), nil
	}
}

func (t teamEnvServiceDo) FirstOrCreate() (*model.TeamEnvService, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamEnvService), nil
	}
}

func (t teamEnvServiceDo) FindByPage(offset int, limit int) (result []*model.TeamEnvService, count int64, err error) {
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

func (t teamEnvServiceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t teamEnvServiceDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t teamEnvServiceDo) Delete(models ...*model.TeamEnvService) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *teamEnvServiceDo) withDO(do gen.Dao) *teamEnvServiceDo {
	t.DO = *do.(*gen.DO)
	return t
}
