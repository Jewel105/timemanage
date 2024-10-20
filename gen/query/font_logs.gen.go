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

func newFontLogs(db *gorm.DB, opts ...gen.DOOption) fontLogs {
	_fontLogs := fontLogs{}

	_fontLogs.fontLogsDo.UseDB(db, opts...)
	_fontLogs.fontLogsDo.UseModel(&models.FontLogs{})

	tableName := _fontLogs.fontLogsDo.TableName()
	_fontLogs.ALL = field.NewAsterisk(tableName)
	_fontLogs.ID = field.NewInt64(tableName, "id")
	_fontLogs.CreatedTime = field.NewTime(tableName, "create_time")
	_fontLogs.UpdatedTime = field.NewTime(tableName, "update_time")
	_fontLogs.DeleteTime = field.NewField(tableName, "delete_time")
	_fontLogs.EquipmentID = field.NewInt64(tableName, "equipment_id")
	_fontLogs.UserID = field.NewInt64(tableName, "user_id")
	_fontLogs.Version = field.NewString(tableName, "version")
	_fontLogs.Stack = field.NewString(tableName, "stack")
	_fontLogs.Error = field.NewString(tableName, "error")

	_fontLogs.fillFieldMap()

	return _fontLogs
}

type fontLogs struct {
	fontLogsDo

	ALL         field.Asterisk
	ID          field.Int64
	CreatedTime field.Time
	UpdatedTime field.Time
	DeleteTime  field.Field
	EquipmentID field.Int64
	UserID      field.Int64
	Version     field.String
	Stack       field.String
	Error       field.String

	fieldMap map[string]field.Expr
}

func (f fontLogs) Table(newTableName string) *fontLogs {
	f.fontLogsDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fontLogs) As(alias string) *fontLogs {
	f.fontLogsDo.DO = *(f.fontLogsDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fontLogs) updateTableName(table string) *fontLogs {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.CreatedTime = field.NewTime(table, "create_time")
	f.UpdatedTime = field.NewTime(table, "update_time")
	f.DeleteTime = field.NewField(table, "delete_time")
	f.EquipmentID = field.NewInt64(table, "equipment_id")
	f.UserID = field.NewInt64(table, "user_id")
	f.Version = field.NewString(table, "version")
	f.Stack = field.NewString(table, "stack")
	f.Error = field.NewString(table, "error")

	f.fillFieldMap()

	return f
}

func (f *fontLogs) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fontLogs) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 9)
	f.fieldMap["id"] = f.ID
	f.fieldMap["create_time"] = f.CreatedTime
	f.fieldMap["update_time"] = f.UpdatedTime
	f.fieldMap["delete_time"] = f.DeleteTime
	f.fieldMap["equipment_id"] = f.EquipmentID
	f.fieldMap["user_id"] = f.UserID
	f.fieldMap["version"] = f.Version
	f.fieldMap["stack"] = f.Stack
	f.fieldMap["error"] = f.Error
}

func (f fontLogs) clone(db *gorm.DB) fontLogs {
	f.fontLogsDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fontLogs) replaceDB(db *gorm.DB) fontLogs {
	f.fontLogsDo.ReplaceDB(db)
	return f
}

type fontLogsDo struct{ gen.DO }

type IFontLogsDo interface {
	gen.SubQuery
	Debug() IFontLogsDo
	WithContext(ctx context.Context) IFontLogsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFontLogsDo
	WriteDB() IFontLogsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFontLogsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFontLogsDo
	Not(conds ...gen.Condition) IFontLogsDo
	Or(conds ...gen.Condition) IFontLogsDo
	Select(conds ...field.Expr) IFontLogsDo
	Where(conds ...gen.Condition) IFontLogsDo
	Order(conds ...field.Expr) IFontLogsDo
	Distinct(cols ...field.Expr) IFontLogsDo
	Omit(cols ...field.Expr) IFontLogsDo
	Join(table schema.Tabler, on ...field.Expr) IFontLogsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFontLogsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFontLogsDo
	Group(cols ...field.Expr) IFontLogsDo
	Having(conds ...gen.Condition) IFontLogsDo
	Limit(limit int) IFontLogsDo
	Offset(offset int) IFontLogsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFontLogsDo
	Unscoped() IFontLogsDo
	Create(values ...*models.FontLogs) error
	CreateInBatches(values []*models.FontLogs, batchSize int) error
	Save(values ...*models.FontLogs) error
	First() (*models.FontLogs, error)
	Take() (*models.FontLogs, error)
	Last() (*models.FontLogs, error)
	Find() ([]*models.FontLogs, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.FontLogs, err error)
	FindInBatches(result *[]*models.FontLogs, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.FontLogs) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFontLogsDo
	Assign(attrs ...field.AssignExpr) IFontLogsDo
	Joins(fields ...field.RelationField) IFontLogsDo
	Preload(fields ...field.RelationField) IFontLogsDo
	FirstOrInit() (*models.FontLogs, error)
	FirstOrCreate() (*models.FontLogs, error)
	FindByPage(offset int, limit int) (result []*models.FontLogs, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFontLogsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fontLogsDo) Debug() IFontLogsDo {
	return f.withDO(f.DO.Debug())
}

func (f fontLogsDo) WithContext(ctx context.Context) IFontLogsDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fontLogsDo) ReadDB() IFontLogsDo {
	return f.Clauses(dbresolver.Read)
}

func (f fontLogsDo) WriteDB() IFontLogsDo {
	return f.Clauses(dbresolver.Write)
}

func (f fontLogsDo) Session(config *gorm.Session) IFontLogsDo {
	return f.withDO(f.DO.Session(config))
}

func (f fontLogsDo) Clauses(conds ...clause.Expression) IFontLogsDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fontLogsDo) Returning(value interface{}, columns ...string) IFontLogsDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fontLogsDo) Not(conds ...gen.Condition) IFontLogsDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fontLogsDo) Or(conds ...gen.Condition) IFontLogsDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fontLogsDo) Select(conds ...field.Expr) IFontLogsDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fontLogsDo) Where(conds ...gen.Condition) IFontLogsDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fontLogsDo) Order(conds ...field.Expr) IFontLogsDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fontLogsDo) Distinct(cols ...field.Expr) IFontLogsDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fontLogsDo) Omit(cols ...field.Expr) IFontLogsDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fontLogsDo) Join(table schema.Tabler, on ...field.Expr) IFontLogsDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fontLogsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFontLogsDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fontLogsDo) RightJoin(table schema.Tabler, on ...field.Expr) IFontLogsDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fontLogsDo) Group(cols ...field.Expr) IFontLogsDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fontLogsDo) Having(conds ...gen.Condition) IFontLogsDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fontLogsDo) Limit(limit int) IFontLogsDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fontLogsDo) Offset(offset int) IFontLogsDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fontLogsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFontLogsDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fontLogsDo) Unscoped() IFontLogsDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fontLogsDo) Create(values ...*models.FontLogs) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fontLogsDo) CreateInBatches(values []*models.FontLogs, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fontLogsDo) Save(values ...*models.FontLogs) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fontLogsDo) First() (*models.FontLogs, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.FontLogs), nil
	}
}

func (f fontLogsDo) Take() (*models.FontLogs, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.FontLogs), nil
	}
}

func (f fontLogsDo) Last() (*models.FontLogs, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.FontLogs), nil
	}
}

func (f fontLogsDo) Find() ([]*models.FontLogs, error) {
	result, err := f.DO.Find()
	return result.([]*models.FontLogs), err
}

func (f fontLogsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.FontLogs, err error) {
	buf := make([]*models.FontLogs, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fontLogsDo) FindInBatches(result *[]*models.FontLogs, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fontLogsDo) Attrs(attrs ...field.AssignExpr) IFontLogsDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fontLogsDo) Assign(attrs ...field.AssignExpr) IFontLogsDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fontLogsDo) Joins(fields ...field.RelationField) IFontLogsDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fontLogsDo) Preload(fields ...field.RelationField) IFontLogsDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fontLogsDo) FirstOrInit() (*models.FontLogs, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.FontLogs), nil
	}
}

func (f fontLogsDo) FirstOrCreate() (*models.FontLogs, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.FontLogs), nil
	}
}

func (f fontLogsDo) FindByPage(offset int, limit int) (result []*models.FontLogs, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fontLogsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fontLogsDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fontLogsDo) Delete(models ...*models.FontLogs) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fontLogsDo) withDO(do gen.Dao) *fontLogsDo {
	f.DO = *do.(*gen.DO)
	return f
}
