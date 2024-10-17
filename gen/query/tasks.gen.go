// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"gin_study/gen/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newTask(db *gorm.DB, opts ...gen.DOOption) task {
	_task := task{}

	_task.taskDo.UseDB(db, opts...)
	_task.taskDo.UseModel(&models.Task{})

	tableName := _task.taskDo.TableName()
	_task.ALL = field.NewAsterisk(tableName)
	_task.ID = field.NewInt64(tableName, "id")
	_task.CreatedTime = field.NewTime(tableName, "create_time")
	_task.UpdatedTime = field.NewTime(tableName, "update_time")
	_task.DeleteTime = field.NewField(tableName, "delete_time")
	_task.UserID = field.NewInt64(tableName, "user_id")
	_task.Description = field.NewString(tableName, "description")
	_task.SpentTime = field.NewInt64(tableName, "spent_time")
	_task.CategoryPath = field.NewString(tableName, "category_path")
	_task.StartTime = field.NewInt64(tableName, "start_time")
	_task.EndTime = field.NewInt64(tableName, "end_time")
	_task.CategoryID = field.NewInt64(tableName, "categor_id")

	_task.fillFieldMap()

	return _task
}

type task struct {
	taskDo

	ALL          field.Asterisk
	ID           field.Int64
	CreatedTime  field.Time
	UpdatedTime  field.Time
	DeleteTime   field.Field
	UserID       field.Int64
	Description  field.String
	SpentTime    field.Int64
	CategoryPath field.String
	StartTime    field.Int64
	EndTime      field.Int64
	CategoryID   field.Int64

	fieldMap map[string]field.Expr
}

func (t task) Table(newTableName string) *task {
	t.taskDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t task) As(alias string) *task {
	t.taskDo.DO = *(t.taskDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *task) updateTableName(table string) *task {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.CreatedTime = field.NewTime(table, "create_time")
	t.UpdatedTime = field.NewTime(table, "update_time")
	t.DeleteTime = field.NewField(table, "delete_time")
	t.UserID = field.NewInt64(table, "user_id")
	t.Description = field.NewString(table, "description")
	t.SpentTime = field.NewInt64(table, "spent_time")
	t.CategoryPath = field.NewString(table, "category_path")
	t.StartTime = field.NewInt64(table, "start_time")
	t.EndTime = field.NewInt64(table, "end_time")
	t.CategoryID = field.NewInt64(table, "categor_id")

	t.fillFieldMap()

	return t
}

func (t *task) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *task) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 11)
	t.fieldMap["id"] = t.ID
	t.fieldMap["create_time"] = t.CreatedTime
	t.fieldMap["update_time"] = t.UpdatedTime
	t.fieldMap["delete_time"] = t.DeleteTime
	t.fieldMap["user_id"] = t.UserID
	t.fieldMap["description"] = t.Description
	t.fieldMap["spent_time"] = t.SpentTime
	t.fieldMap["category_path"] = t.CategoryPath
	t.fieldMap["start_time"] = t.StartTime
	t.fieldMap["end_time"] = t.EndTime
	t.fieldMap["categor_id"] = t.CategoryID
}

func (t task) clone(db *gorm.DB) task {
	t.taskDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t task) replaceDB(db *gorm.DB) task {
	t.taskDo.ReplaceDB(db)
	return t
}

type taskDo struct{ gen.DO }

type ITaskDo interface {
	gen.SubQuery
	Debug() ITaskDo
	WithContext(ctx context.Context) ITaskDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITaskDo
	WriteDB() ITaskDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITaskDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITaskDo
	Not(conds ...gen.Condition) ITaskDo
	Or(conds ...gen.Condition) ITaskDo
	Select(conds ...field.Expr) ITaskDo
	Where(conds ...gen.Condition) ITaskDo
	Order(conds ...field.Expr) ITaskDo
	Distinct(cols ...field.Expr) ITaskDo
	Omit(cols ...field.Expr) ITaskDo
	Join(table schema.Tabler, on ...field.Expr) ITaskDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITaskDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITaskDo
	Group(cols ...field.Expr) ITaskDo
	Having(conds ...gen.Condition) ITaskDo
	Limit(limit int) ITaskDo
	Offset(offset int) ITaskDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITaskDo
	Unscoped() ITaskDo
	Create(values ...*models.Task) error
	CreateInBatches(values []*models.Task, batchSize int) error
	Save(values ...*models.Task) error
	First() (*models.Task, error)
	Take() (*models.Task, error)
	Last() (*models.Task, error)
	Find() ([]*models.Task, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Task, err error)
	FindInBatches(result *[]*models.Task, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Task) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITaskDo
	Assign(attrs ...field.AssignExpr) ITaskDo
	Joins(fields ...field.RelationField) ITaskDo
	Preload(fields ...field.RelationField) ITaskDo
	FirstOrInit() (*models.Task, error)
	FirstOrCreate() (*models.Task, error)
	FindByPage(offset int, limit int) (result []*models.Task, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITaskDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t taskDo) Debug() ITaskDo {
	return t.withDO(t.DO.Debug())
}

func (t taskDo) WithContext(ctx context.Context) ITaskDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t taskDo) ReadDB() ITaskDo {
	return t.Clauses(dbresolver.Read)
}

func (t taskDo) WriteDB() ITaskDo {
	return t.Clauses(dbresolver.Write)
}

func (t taskDo) Session(config *gorm.Session) ITaskDo {
	return t.withDO(t.DO.Session(config))
}

func (t taskDo) Clauses(conds ...clause.Expression) ITaskDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t taskDo) Returning(value interface{}, columns ...string) ITaskDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t taskDo) Not(conds ...gen.Condition) ITaskDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t taskDo) Or(conds ...gen.Condition) ITaskDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t taskDo) Select(conds ...field.Expr) ITaskDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t taskDo) Where(conds ...gen.Condition) ITaskDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t taskDo) Order(conds ...field.Expr) ITaskDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t taskDo) Distinct(cols ...field.Expr) ITaskDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t taskDo) Omit(cols ...field.Expr) ITaskDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t taskDo) Join(table schema.Tabler, on ...field.Expr) ITaskDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t taskDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITaskDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t taskDo) RightJoin(table schema.Tabler, on ...field.Expr) ITaskDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t taskDo) Group(cols ...field.Expr) ITaskDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t taskDo) Having(conds ...gen.Condition) ITaskDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t taskDo) Limit(limit int) ITaskDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t taskDo) Offset(offset int) ITaskDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t taskDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITaskDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t taskDo) Unscoped() ITaskDo {
	return t.withDO(t.DO.Unscoped())
}

func (t taskDo) Create(values ...*models.Task) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t taskDo) CreateInBatches(values []*models.Task, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t taskDo) Save(values ...*models.Task) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t taskDo) First() (*models.Task, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Task), nil
	}
}

func (t taskDo) Take() (*models.Task, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Task), nil
	}
}

func (t taskDo) Last() (*models.Task, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Task), nil
	}
}

func (t taskDo) Find() ([]*models.Task, error) {
	result, err := t.DO.Find()
	return result.([]*models.Task), err
}

func (t taskDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Task, err error) {
	buf := make([]*models.Task, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t taskDo) FindInBatches(result *[]*models.Task, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t taskDo) Attrs(attrs ...field.AssignExpr) ITaskDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t taskDo) Assign(attrs ...field.AssignExpr) ITaskDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t taskDo) Joins(fields ...field.RelationField) ITaskDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t taskDo) Preload(fields ...field.RelationField) ITaskDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t taskDo) FirstOrInit() (*models.Task, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Task), nil
	}
}

func (t taskDo) FirstOrCreate() (*models.Task, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Task), nil
	}
}

func (t taskDo) FindByPage(offset int, limit int) (result []*models.Task, count int64, err error) {
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

func (t taskDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t taskDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t taskDo) Delete(models ...*models.Task) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *taskDo) withDO(do gen.Dao) *taskDo {
	t.DO = *do.(*gen.DO)
	return t
}
