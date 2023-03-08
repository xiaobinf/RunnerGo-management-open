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

func newVariable(db *gorm.DB, opts ...gen.DOOption) variable {
	_variable := variable{}

	_variable.variableDo.UseDB(db, opts...)
	_variable.variableDo.UseModel(&model.Variable{})

	tableName := _variable.variableDo.TableName()
	_variable.ALL = field.NewAsterisk(tableName)
	_variable.ID = field.NewInt64(tableName, "id")
	_variable.TeamID = field.NewString(tableName, "team_id")
	_variable.SceneID = field.NewString(tableName, "scene_id")
	_variable.Type = field.NewInt32(tableName, "type")
	_variable.Var = field.NewString(tableName, "var")
	_variable.Val = field.NewString(tableName, "val")
	_variable.Description = field.NewString(tableName, "description")
	_variable.Status = field.NewInt32(tableName, "status")
	_variable.CreatedAt = field.NewTime(tableName, "created_at")
	_variable.UpdatedAt = field.NewTime(tableName, "updated_at")
	_variable.DeletedAt = field.NewField(tableName, "deleted_at")

	_variable.fillFieldMap()

	return _variable
}

type variable struct {
	variableDo variableDo

	ALL         field.Asterisk
	ID          field.Int64
	TeamID      field.String // 团队id
	SceneID     field.String // 场景ID
	Type        field.Int32  // 使用范围：1-全局变量，2-场景变量
	Var         field.String // 变量名
	Val         field.String // 变量值
	Description field.String // 描述
	Status      field.Int32  // 开关状态：1-开启，2-关闭
	CreatedAt   field.Time   // 创建时间
	UpdatedAt   field.Time   // 修改时间
	DeletedAt   field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (v variable) Table(newTableName string) *variable {
	v.variableDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v variable) As(alias string) *variable {
	v.variableDo.DO = *(v.variableDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *variable) updateTableName(table string) *variable {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewInt64(table, "id")
	v.TeamID = field.NewString(table, "team_id")
	v.SceneID = field.NewString(table, "scene_id")
	v.Type = field.NewInt32(table, "type")
	v.Var = field.NewString(table, "var")
	v.Val = field.NewString(table, "val")
	v.Description = field.NewString(table, "description")
	v.Status = field.NewInt32(table, "status")
	v.CreatedAt = field.NewTime(table, "created_at")
	v.UpdatedAt = field.NewTime(table, "updated_at")
	v.DeletedAt = field.NewField(table, "deleted_at")

	v.fillFieldMap()

	return v
}

func (v *variable) WithContext(ctx context.Context) *variableDo { return v.variableDo.WithContext(ctx) }

func (v variable) TableName() string { return v.variableDo.TableName() }

func (v variable) Alias() string { return v.variableDo.Alias() }

func (v *variable) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *variable) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 11)
	v.fieldMap["id"] = v.ID
	v.fieldMap["team_id"] = v.TeamID
	v.fieldMap["scene_id"] = v.SceneID
	v.fieldMap["type"] = v.Type
	v.fieldMap["var"] = v.Var
	v.fieldMap["val"] = v.Val
	v.fieldMap["description"] = v.Description
	v.fieldMap["status"] = v.Status
	v.fieldMap["created_at"] = v.CreatedAt
	v.fieldMap["updated_at"] = v.UpdatedAt
	v.fieldMap["deleted_at"] = v.DeletedAt
}

func (v variable) clone(db *gorm.DB) variable {
	v.variableDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v variable) replaceDB(db *gorm.DB) variable {
	v.variableDo.ReplaceDB(db)
	return v
}

type variableDo struct{ gen.DO }

func (v variableDo) Debug() *variableDo {
	return v.withDO(v.DO.Debug())
}

func (v variableDo) WithContext(ctx context.Context) *variableDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v variableDo) ReadDB() *variableDo {
	return v.Clauses(dbresolver.Read)
}

func (v variableDo) WriteDB() *variableDo {
	return v.Clauses(dbresolver.Write)
}

func (v variableDo) Session(config *gorm.Session) *variableDo {
	return v.withDO(v.DO.Session(config))
}

func (v variableDo) Clauses(conds ...clause.Expression) *variableDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v variableDo) Returning(value interface{}, columns ...string) *variableDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v variableDo) Not(conds ...gen.Condition) *variableDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v variableDo) Or(conds ...gen.Condition) *variableDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v variableDo) Select(conds ...field.Expr) *variableDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v variableDo) Where(conds ...gen.Condition) *variableDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v variableDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *variableDo {
	return v.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (v variableDo) Order(conds ...field.Expr) *variableDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v variableDo) Distinct(cols ...field.Expr) *variableDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v variableDo) Omit(cols ...field.Expr) *variableDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v variableDo) Join(table schema.Tabler, on ...field.Expr) *variableDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v variableDo) LeftJoin(table schema.Tabler, on ...field.Expr) *variableDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v variableDo) RightJoin(table schema.Tabler, on ...field.Expr) *variableDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v variableDo) Group(cols ...field.Expr) *variableDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v variableDo) Having(conds ...gen.Condition) *variableDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v variableDo) Limit(limit int) *variableDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v variableDo) Offset(offset int) *variableDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v variableDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *variableDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v variableDo) Unscoped() *variableDo {
	return v.withDO(v.DO.Unscoped())
}

func (v variableDo) Create(values ...*model.Variable) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v variableDo) CreateInBatches(values []*model.Variable, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v variableDo) Save(values ...*model.Variable) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v variableDo) First() (*model.Variable, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Variable), nil
	}
}

func (v variableDo) Take() (*model.Variable, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Variable), nil
	}
}

func (v variableDo) Last() (*model.Variable, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Variable), nil
	}
}

func (v variableDo) Find() ([]*model.Variable, error) {
	result, err := v.DO.Find()
	return result.([]*model.Variable), err
}

func (v variableDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Variable, err error) {
	buf := make([]*model.Variable, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v variableDo) FindInBatches(result *[]*model.Variable, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v variableDo) Attrs(attrs ...field.AssignExpr) *variableDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v variableDo) Assign(attrs ...field.AssignExpr) *variableDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v variableDo) Joins(fields ...field.RelationField) *variableDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v variableDo) Preload(fields ...field.RelationField) *variableDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v variableDo) FirstOrInit() (*model.Variable, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Variable), nil
	}
}

func (v variableDo) FirstOrCreate() (*model.Variable, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Variable), nil
	}
}

func (v variableDo) FindByPage(offset int, limit int) (result []*model.Variable, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v variableDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v variableDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v variableDo) Delete(models ...*model.Variable) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *variableDo) withDO(do gen.Dao) *variableDo {
	v.DO = *do.(*gen.DO)
	return v
}
