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

func newEquipment(db *gorm.DB, opts ...gen.DOOption) equipment {
	_equipment := equipment{}

	_equipment.equipmentDo.UseDB(db, opts...)
	_equipment.equipmentDo.UseModel(&models.Equipment{})

	tableName := _equipment.equipmentDo.TableName()
	_equipment.ALL = field.NewAsterisk(tableName)
	_equipment.ID = field.NewInt64(tableName, "id")
	_equipment.CreatedTime = field.NewTime(tableName, "create_time")
	_equipment.UpdatedTime = field.NewTime(tableName, "update_time")
	_equipment.DeleteTime = field.NewField(tableName, "delete_time")
	_equipment.Vender = field.NewString(tableName, "vender")
	_equipment.Type = field.NewString(tableName, "type")
	_equipment.Sn = field.NewString(tableName, "sn")
	_equipment.Imei1 = field.NewString(tableName, "imei1")
	_equipment.Imei0 = field.NewString(tableName, "imei0")
	_equipment.Os = field.NewString(tableName, "os")
	_equipment.UserIDs = field.NewString(tableName, "user_ids")
	_equipment.IsPhysicalDevice = field.NewInt(tableName, "is_physical_device")
	_equipment.Fingerprint = field.NewString(tableName, "fingerprint")

	_equipment.fillFieldMap()

	return _equipment
}

type equipment struct {
	equipmentDo

	ALL              field.Asterisk
	ID               field.Int64
	CreatedTime      field.Time
	UpdatedTime      field.Time
	DeleteTime       field.Field
	Vender           field.String
	Type             field.String
	Sn               field.String
	Imei1            field.String
	Imei0            field.String
	Os               field.String
	UserIDs          field.String
	IsPhysicalDevice field.Int
	Fingerprint      field.String

	fieldMap map[string]field.Expr
}

func (e equipment) Table(newTableName string) *equipment {
	e.equipmentDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e equipment) As(alias string) *equipment {
	e.equipmentDo.DO = *(e.equipmentDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *equipment) updateTableName(table string) *equipment {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt64(table, "id")
	e.CreatedTime = field.NewTime(table, "create_time")
	e.UpdatedTime = field.NewTime(table, "update_time")
	e.DeleteTime = field.NewField(table, "delete_time")
	e.Vender = field.NewString(table, "vender")
	e.Type = field.NewString(table, "type")
	e.Sn = field.NewString(table, "sn")
	e.Imei1 = field.NewString(table, "imei1")
	e.Imei0 = field.NewString(table, "imei0")
	e.Os = field.NewString(table, "os")
	e.UserIDs = field.NewString(table, "user_ids")
	e.IsPhysicalDevice = field.NewInt(table, "is_physical_device")
	e.Fingerprint = field.NewString(table, "fingerprint")

	e.fillFieldMap()

	return e
}

func (e *equipment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *equipment) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 13)
	e.fieldMap["id"] = e.ID
	e.fieldMap["create_time"] = e.CreatedTime
	e.fieldMap["update_time"] = e.UpdatedTime
	e.fieldMap["delete_time"] = e.DeleteTime
	e.fieldMap["vender"] = e.Vender
	e.fieldMap["type"] = e.Type
	e.fieldMap["sn"] = e.Sn
	e.fieldMap["imei1"] = e.Imei1
	e.fieldMap["imei0"] = e.Imei0
	e.fieldMap["os"] = e.Os
	e.fieldMap["user_ids"] = e.UserIDs
	e.fieldMap["is_physical_device"] = e.IsPhysicalDevice
	e.fieldMap["fingerprint"] = e.Fingerprint
}

func (e equipment) clone(db *gorm.DB) equipment {
	e.equipmentDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e equipment) replaceDB(db *gorm.DB) equipment {
	e.equipmentDo.ReplaceDB(db)
	return e
}

type equipmentDo struct{ gen.DO }

type IEquipmentDo interface {
	gen.SubQuery
	Debug() IEquipmentDo
	WithContext(ctx context.Context) IEquipmentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEquipmentDo
	WriteDB() IEquipmentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEquipmentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEquipmentDo
	Not(conds ...gen.Condition) IEquipmentDo
	Or(conds ...gen.Condition) IEquipmentDo
	Select(conds ...field.Expr) IEquipmentDo
	Where(conds ...gen.Condition) IEquipmentDo
	Order(conds ...field.Expr) IEquipmentDo
	Distinct(cols ...field.Expr) IEquipmentDo
	Omit(cols ...field.Expr) IEquipmentDo
	Join(table schema.Tabler, on ...field.Expr) IEquipmentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEquipmentDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEquipmentDo
	Group(cols ...field.Expr) IEquipmentDo
	Having(conds ...gen.Condition) IEquipmentDo
	Limit(limit int) IEquipmentDo
	Offset(offset int) IEquipmentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEquipmentDo
	Unscoped() IEquipmentDo
	Create(values ...*models.Equipment) error
	CreateInBatches(values []*models.Equipment, batchSize int) error
	Save(values ...*models.Equipment) error
	First() (*models.Equipment, error)
	Take() (*models.Equipment, error)
	Last() (*models.Equipment, error)
	Find() ([]*models.Equipment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Equipment, err error)
	FindInBatches(result *[]*models.Equipment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Equipment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEquipmentDo
	Assign(attrs ...field.AssignExpr) IEquipmentDo
	Joins(fields ...field.RelationField) IEquipmentDo
	Preload(fields ...field.RelationField) IEquipmentDo
	FirstOrInit() (*models.Equipment, error)
	FirstOrCreate() (*models.Equipment, error)
	FindByPage(offset int, limit int) (result []*models.Equipment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEquipmentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e equipmentDo) Debug() IEquipmentDo {
	return e.withDO(e.DO.Debug())
}

func (e equipmentDo) WithContext(ctx context.Context) IEquipmentDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e equipmentDo) ReadDB() IEquipmentDo {
	return e.Clauses(dbresolver.Read)
}

func (e equipmentDo) WriteDB() IEquipmentDo {
	return e.Clauses(dbresolver.Write)
}

func (e equipmentDo) Session(config *gorm.Session) IEquipmentDo {
	return e.withDO(e.DO.Session(config))
}

func (e equipmentDo) Clauses(conds ...clause.Expression) IEquipmentDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e equipmentDo) Returning(value interface{}, columns ...string) IEquipmentDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e equipmentDo) Not(conds ...gen.Condition) IEquipmentDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e equipmentDo) Or(conds ...gen.Condition) IEquipmentDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e equipmentDo) Select(conds ...field.Expr) IEquipmentDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e equipmentDo) Where(conds ...gen.Condition) IEquipmentDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e equipmentDo) Order(conds ...field.Expr) IEquipmentDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e equipmentDo) Distinct(cols ...field.Expr) IEquipmentDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e equipmentDo) Omit(cols ...field.Expr) IEquipmentDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e equipmentDo) Join(table schema.Tabler, on ...field.Expr) IEquipmentDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e equipmentDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEquipmentDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e equipmentDo) RightJoin(table schema.Tabler, on ...field.Expr) IEquipmentDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e equipmentDo) Group(cols ...field.Expr) IEquipmentDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e equipmentDo) Having(conds ...gen.Condition) IEquipmentDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e equipmentDo) Limit(limit int) IEquipmentDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e equipmentDo) Offset(offset int) IEquipmentDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e equipmentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEquipmentDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e equipmentDo) Unscoped() IEquipmentDo {
	return e.withDO(e.DO.Unscoped())
}

func (e equipmentDo) Create(values ...*models.Equipment) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e equipmentDo) CreateInBatches(values []*models.Equipment, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e equipmentDo) Save(values ...*models.Equipment) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e equipmentDo) First() (*models.Equipment, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Equipment), nil
	}
}

func (e equipmentDo) Take() (*models.Equipment, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Equipment), nil
	}
}

func (e equipmentDo) Last() (*models.Equipment, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Equipment), nil
	}
}

func (e equipmentDo) Find() ([]*models.Equipment, error) {
	result, err := e.DO.Find()
	return result.([]*models.Equipment), err
}

func (e equipmentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Equipment, err error) {
	buf := make([]*models.Equipment, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e equipmentDo) FindInBatches(result *[]*models.Equipment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e equipmentDo) Attrs(attrs ...field.AssignExpr) IEquipmentDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e equipmentDo) Assign(attrs ...field.AssignExpr) IEquipmentDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e equipmentDo) Joins(fields ...field.RelationField) IEquipmentDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e equipmentDo) Preload(fields ...field.RelationField) IEquipmentDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e equipmentDo) FirstOrInit() (*models.Equipment, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Equipment), nil
	}
}

func (e equipmentDo) FirstOrCreate() (*models.Equipment, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Equipment), nil
	}
}

func (e equipmentDo) FindByPage(offset int, limit int) (result []*models.Equipment, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e equipmentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e equipmentDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e equipmentDo) Delete(models ...*models.Equipment) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *equipmentDo) withDO(do gen.Dao) *equipmentDo {
	e.DO = *do.(*gen.DO)
	return e
}
