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

	"github.com/coze-dev/coze-studio/backend/domain/memory/database/internal/dal/model"
)

func newDraftDatabaseInfo(db *gorm.DB, opts ...gen.DOOption) draftDatabaseInfo {
	_draftDatabaseInfo := draftDatabaseInfo{}

	_draftDatabaseInfo.draftDatabaseInfoDo.UseDB(db, opts...)
	_draftDatabaseInfo.draftDatabaseInfoDo.UseModel(&model.DraftDatabaseInfo{})

	tableName := _draftDatabaseInfo.draftDatabaseInfoDo.TableName()
	_draftDatabaseInfo.ALL = field.NewAsterisk(tableName)
	_draftDatabaseInfo.ID = field.NewInt64(tableName, "id")
	_draftDatabaseInfo.AppID = field.NewInt64(tableName, "app_id")
	_draftDatabaseInfo.SpaceID = field.NewInt64(tableName, "space_id")
	_draftDatabaseInfo.RelatedOnlineID = field.NewInt64(tableName, "related_online_id")
	_draftDatabaseInfo.IsVisible = field.NewInt32(tableName, "is_visible")
	_draftDatabaseInfo.PromptDisabled = field.NewInt32(tableName, "prompt_disabled")
	_draftDatabaseInfo.TableName_ = field.NewString(tableName, "table_name")
	_draftDatabaseInfo.TableDesc = field.NewString(tableName, "table_desc")
	_draftDatabaseInfo.TableField = field.NewField(tableName, "table_field")
	_draftDatabaseInfo.CreatorID = field.NewInt64(tableName, "creator_id")
	_draftDatabaseInfo.IconURI = field.NewString(tableName, "icon_uri")
	_draftDatabaseInfo.PhysicalTableName = field.NewString(tableName, "physical_table_name")
	_draftDatabaseInfo.RwMode = field.NewInt64(tableName, "rw_mode")
	_draftDatabaseInfo.CreatedAt = field.NewInt64(tableName, "created_at")
	_draftDatabaseInfo.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_draftDatabaseInfo.DeletedAt = field.NewField(tableName, "deleted_at")

	_draftDatabaseInfo.fillFieldMap()

	return _draftDatabaseInfo
}

// draftDatabaseInfo draft database info
type draftDatabaseInfo struct {
	draftDatabaseInfoDo

	ALL               field.Asterisk
	ID                field.Int64  // ID
	AppID             field.Int64  // App ID
	SpaceID           field.Int64  // Space ID
	RelatedOnlineID   field.Int64  // The primary key ID of online_database_info table
	IsVisible         field.Int32  // Visibility: 0 invisible, 1 visible
	PromptDisabled    field.Int32  // Support prompt calls: 1 not supported, 0 supported
	TableName_        field.String // Table name
	TableDesc         field.String // Table description
	TableField        field.Field  // Table field info
	CreatorID         field.Int64  // Creator ID
	IconURI           field.String // Icon Uri
	PhysicalTableName field.String // The name of the real physical table
	RwMode            field.Int64  // Read and write permission modes: 1. Limited read and write mode 2. Read-only mode 3. Full read and write mode
	CreatedAt         field.Int64  // Create Time in Milliseconds
	UpdatedAt         field.Int64  // Update Time in Milliseconds
	DeletedAt         field.Field  // Delete Time

	fieldMap map[string]field.Expr
}

func (d draftDatabaseInfo) Table(newTableName string) *draftDatabaseInfo {
	d.draftDatabaseInfoDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d draftDatabaseInfo) As(alias string) *draftDatabaseInfo {
	d.draftDatabaseInfoDo.DO = *(d.draftDatabaseInfoDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *draftDatabaseInfo) updateTableName(table string) *draftDatabaseInfo {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt64(table, "id")
	d.AppID = field.NewInt64(table, "app_id")
	d.SpaceID = field.NewInt64(table, "space_id")
	d.RelatedOnlineID = field.NewInt64(table, "related_online_id")
	d.IsVisible = field.NewInt32(table, "is_visible")
	d.PromptDisabled = field.NewInt32(table, "prompt_disabled")
	d.TableName_ = field.NewString(table, "table_name")
	d.TableDesc = field.NewString(table, "table_desc")
	d.TableField = field.NewField(table, "table_field")
	d.CreatorID = field.NewInt64(table, "creator_id")
	d.IconURI = field.NewString(table, "icon_uri")
	d.PhysicalTableName = field.NewString(table, "physical_table_name")
	d.RwMode = field.NewInt64(table, "rw_mode")
	d.CreatedAt = field.NewInt64(table, "created_at")
	d.UpdatedAt = field.NewInt64(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")

	d.fillFieldMap()

	return d
}

func (d *draftDatabaseInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *draftDatabaseInfo) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 16)
	d.fieldMap["id"] = d.ID
	d.fieldMap["app_id"] = d.AppID
	d.fieldMap["space_id"] = d.SpaceID
	d.fieldMap["related_online_id"] = d.RelatedOnlineID
	d.fieldMap["is_visible"] = d.IsVisible
	d.fieldMap["prompt_disabled"] = d.PromptDisabled
	d.fieldMap["table_name"] = d.TableName_
	d.fieldMap["table_desc"] = d.TableDesc
	d.fieldMap["table_field"] = d.TableField
	d.fieldMap["creator_id"] = d.CreatorID
	d.fieldMap["icon_uri"] = d.IconURI
	d.fieldMap["physical_table_name"] = d.PhysicalTableName
	d.fieldMap["rw_mode"] = d.RwMode
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
}

func (d draftDatabaseInfo) clone(db *gorm.DB) draftDatabaseInfo {
	d.draftDatabaseInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d draftDatabaseInfo) replaceDB(db *gorm.DB) draftDatabaseInfo {
	d.draftDatabaseInfoDo.ReplaceDB(db)
	return d
}

type draftDatabaseInfoDo struct{ gen.DO }

type IDraftDatabaseInfoDo interface {
	gen.SubQuery
	Debug() IDraftDatabaseInfoDo
	WithContext(ctx context.Context) IDraftDatabaseInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDraftDatabaseInfoDo
	WriteDB() IDraftDatabaseInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDraftDatabaseInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDraftDatabaseInfoDo
	Not(conds ...gen.Condition) IDraftDatabaseInfoDo
	Or(conds ...gen.Condition) IDraftDatabaseInfoDo
	Select(conds ...field.Expr) IDraftDatabaseInfoDo
	Where(conds ...gen.Condition) IDraftDatabaseInfoDo
	Order(conds ...field.Expr) IDraftDatabaseInfoDo
	Distinct(cols ...field.Expr) IDraftDatabaseInfoDo
	Omit(cols ...field.Expr) IDraftDatabaseInfoDo
	Join(table schema.Tabler, on ...field.Expr) IDraftDatabaseInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDraftDatabaseInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDraftDatabaseInfoDo
	Group(cols ...field.Expr) IDraftDatabaseInfoDo
	Having(conds ...gen.Condition) IDraftDatabaseInfoDo
	Limit(limit int) IDraftDatabaseInfoDo
	Offset(offset int) IDraftDatabaseInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDraftDatabaseInfoDo
	Unscoped() IDraftDatabaseInfoDo
	Create(values ...*model.DraftDatabaseInfo) error
	CreateInBatches(values []*model.DraftDatabaseInfo, batchSize int) error
	Save(values ...*model.DraftDatabaseInfo) error
	First() (*model.DraftDatabaseInfo, error)
	Take() (*model.DraftDatabaseInfo, error)
	Last() (*model.DraftDatabaseInfo, error)
	Find() ([]*model.DraftDatabaseInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DraftDatabaseInfo, err error)
	FindInBatches(result *[]*model.DraftDatabaseInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DraftDatabaseInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDraftDatabaseInfoDo
	Assign(attrs ...field.AssignExpr) IDraftDatabaseInfoDo
	Joins(fields ...field.RelationField) IDraftDatabaseInfoDo
	Preload(fields ...field.RelationField) IDraftDatabaseInfoDo
	FirstOrInit() (*model.DraftDatabaseInfo, error)
	FirstOrCreate() (*model.DraftDatabaseInfo, error)
	FindByPage(offset int, limit int) (result []*model.DraftDatabaseInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDraftDatabaseInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d draftDatabaseInfoDo) Debug() IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Debug())
}

func (d draftDatabaseInfoDo) WithContext(ctx context.Context) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d draftDatabaseInfoDo) ReadDB() IDraftDatabaseInfoDo {
	return d.Clauses(dbresolver.Read)
}

func (d draftDatabaseInfoDo) WriteDB() IDraftDatabaseInfoDo {
	return d.Clauses(dbresolver.Write)
}

func (d draftDatabaseInfoDo) Session(config *gorm.Session) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Session(config))
}

func (d draftDatabaseInfoDo) Clauses(conds ...clause.Expression) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d draftDatabaseInfoDo) Returning(value interface{}, columns ...string) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d draftDatabaseInfoDo) Not(conds ...gen.Condition) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d draftDatabaseInfoDo) Or(conds ...gen.Condition) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d draftDatabaseInfoDo) Select(conds ...field.Expr) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d draftDatabaseInfoDo) Where(conds ...gen.Condition) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d draftDatabaseInfoDo) Order(conds ...field.Expr) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d draftDatabaseInfoDo) Distinct(cols ...field.Expr) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d draftDatabaseInfoDo) Omit(cols ...field.Expr) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d draftDatabaseInfoDo) Join(table schema.Tabler, on ...field.Expr) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d draftDatabaseInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d draftDatabaseInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d draftDatabaseInfoDo) Group(cols ...field.Expr) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d draftDatabaseInfoDo) Having(conds ...gen.Condition) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d draftDatabaseInfoDo) Limit(limit int) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d draftDatabaseInfoDo) Offset(offset int) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d draftDatabaseInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d draftDatabaseInfoDo) Unscoped() IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Unscoped())
}

func (d draftDatabaseInfoDo) Create(values ...*model.DraftDatabaseInfo) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d draftDatabaseInfoDo) CreateInBatches(values []*model.DraftDatabaseInfo, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d draftDatabaseInfoDo) Save(values ...*model.DraftDatabaseInfo) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d draftDatabaseInfoDo) First() (*model.DraftDatabaseInfo, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DraftDatabaseInfo), nil
	}
}

func (d draftDatabaseInfoDo) Take() (*model.DraftDatabaseInfo, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DraftDatabaseInfo), nil
	}
}

func (d draftDatabaseInfoDo) Last() (*model.DraftDatabaseInfo, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DraftDatabaseInfo), nil
	}
}

func (d draftDatabaseInfoDo) Find() ([]*model.DraftDatabaseInfo, error) {
	result, err := d.DO.Find()
	return result.([]*model.DraftDatabaseInfo), err
}

func (d draftDatabaseInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DraftDatabaseInfo, err error) {
	buf := make([]*model.DraftDatabaseInfo, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d draftDatabaseInfoDo) FindInBatches(result *[]*model.DraftDatabaseInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d draftDatabaseInfoDo) Attrs(attrs ...field.AssignExpr) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d draftDatabaseInfoDo) Assign(attrs ...field.AssignExpr) IDraftDatabaseInfoDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d draftDatabaseInfoDo) Joins(fields ...field.RelationField) IDraftDatabaseInfoDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d draftDatabaseInfoDo) Preload(fields ...field.RelationField) IDraftDatabaseInfoDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d draftDatabaseInfoDo) FirstOrInit() (*model.DraftDatabaseInfo, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DraftDatabaseInfo), nil
	}
}

func (d draftDatabaseInfoDo) FirstOrCreate() (*model.DraftDatabaseInfo, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DraftDatabaseInfo), nil
	}
}

func (d draftDatabaseInfoDo) FindByPage(offset int, limit int) (result []*model.DraftDatabaseInfo, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d draftDatabaseInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d draftDatabaseInfoDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d draftDatabaseInfoDo) Delete(models ...*model.DraftDatabaseInfo) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *draftDatabaseInfoDo) withDO(do gen.Dao) *draftDatabaseInfoDo {
	d.DO = *do.(*gen.DO)
	return d
}
