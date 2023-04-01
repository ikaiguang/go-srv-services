// Package datas
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package datas

import (
	"bytes"
	context "context"
	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"
	gorm "gorm.io/gorm"
	"strings"

	entities "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/entity"
	repos "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/repo"
	schemas "github.com/ikaiguang/go-srv-services/app/user-service/internal/infra/schema"
)

// userRegMobileRepo repo
type userRegMobileRepo struct {
	dbConn              *gorm.DB              // *gorm.DB
	UserRegMobileSchema schemas.UserRegMobile // UserRegMobile
}

// NewUserRegMobileRepo new data repo
func NewUserRegMobileRepo(dbConn *gorm.DB) repos.UserRegMobileRepo {
	return &userRegMobileRepo{
		dbConn: dbConn,
	}
}

// =============== 创建 ===============

// create insert one
func (s *userRegMobileRepo) create(ctx context.Context, dbConn *gorm.DB, dataModel *entities.UserRegMobile) (err error) {
	err = dbConn.WithContext(ctx).Create(dataModel).Error
	if err != nil {
		return err
	}
	return
}

// Create insert one
func (s *userRegMobileRepo) Create(ctx context.Context, dataModel *entities.UserRegMobile) error {
	return s.create(ctx, s.dbConn, dataModel)
}

// CreateWithDBConn create
func (s *userRegMobileRepo) CreateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *entities.UserRegMobile) error {
	return s.create(ctx, dbConn, dataModel)
}

// existCreate exist create
func (s *userRegMobileRepo) existCreate(ctx context.Context, dbConn *gorm.DB, dataModel *entities.UserRegMobile) (anotherModel *entities.UserRegMobile, isNotFound bool, err error) {
	anotherModel = new(entities.UserRegMobile)
	err = dbConn.WithContext(ctx).
		Table(s.UserRegMobileSchema.TableName()).
		Where("id = ?", dataModel.Id).
		First(anotherModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			isNotFound = true
			err = nil
		}
		return
	}
	return
}

// ExistCreate exist create
func (s *userRegMobileRepo) ExistCreate(ctx context.Context, dataModel *entities.UserRegMobile) (anotherModel *entities.UserRegMobile, isNotFound bool, err error) {
	return s.existCreate(ctx, s.dbConn, dataModel)
}

// ExistCreateWithDBConn exist create
func (s *userRegMobileRepo) ExistCreateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *entities.UserRegMobile) (anotherModel *entities.UserRegMobile, isNotFound bool, err error) {
	return s.existCreate(ctx, dbConn, dataModel)
}

// createInBatches create many
func (s *userRegMobileRepo) createInBatches(ctx context.Context, dbConn *gorm.DB, dataModels []*entities.UserRegMobile, batchSize int) (err error) {
	err = dbConn.WithContext(ctx).CreateInBatches(dataModels, batchSize).Error
	if err != nil {
		return err
	}
	return
}

// CreateInBatches create many
func (s *userRegMobileRepo) CreateInBatches(ctx context.Context, dataModels []*entities.UserRegMobile, batchSize int) error {
	return s.createInBatches(ctx, s.dbConn, dataModels, batchSize)
}

// CreateInBatchesWithDBConn create many
func (s *userRegMobileRepo) CreateInBatchesWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModels []*entities.UserRegMobile, batchSize int) error {
	return s.createInBatches(ctx, dbConn, dataModels, batchSize)
}

// =============== 更新 ===============

// update update
func (s *userRegMobileRepo) update(ctx context.Context, dbConn *gorm.DB, dataModel *entities.UserRegMobile) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserRegMobileSchema.TableName()).
		Where("id = ?", dataModel.Id).
		Save(dataModel).Error
	if err != nil {
		return err
	}
	return
}

// Update update
func (s *userRegMobileRepo) Update(ctx context.Context, dataModel *entities.UserRegMobile) error {
	return s.update(ctx, s.dbConn, dataModel)
}

// UpdateWithDBConn update
func (s *userRegMobileRepo) UpdateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *entities.UserRegMobile) error {
	return s.update(ctx, dbConn, dataModel)
}

// existUpdate exist update
func (s *userRegMobileRepo) existUpdate(ctx context.Context, dbConn *gorm.DB, dataModel *entities.UserRegMobile) (anotherModel *entities.UserRegMobile, isNotFound bool, err error) {
	anotherModel = new(entities.UserRegMobile)
	err = dbConn.WithContext(ctx).
		Table(s.UserRegMobileSchema.TableName()).
		Where("id = ?", dataModel.Id).
		Where("id != ?", dataModel.Id).
		First(anotherModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			isNotFound = true
			err = nil
		}
		return
	}
	return
}

// ExistUpdate exist update
func (s *userRegMobileRepo) ExistUpdate(ctx context.Context, dataModel *entities.UserRegMobile) (anotherModel *entities.UserRegMobile, isNotFound bool, err error) {
	return s.existUpdate(ctx, s.dbConn, dataModel)
}

// ExistUpdateWithDBConn exist update
func (s *userRegMobileRepo) ExistUpdateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *entities.UserRegMobile) (anotherModel *entities.UserRegMobile, isNotFound bool, err error) {
	return s.existUpdate(ctx, dbConn, dataModel)
}

// =============== query one : 查一个 ===============

// queryOneById query one by id
func (s *userRegMobileRepo) queryOneById(ctx context.Context, dbConn *gorm.DB, id interface{}) (dataModel *entities.UserRegMobile, isNotFound bool, err error) {
	dataModel = new(entities.UserRegMobile)
	err = dbConn.WithContext(ctx).
		Table(s.UserRegMobileSchema.TableName()).
		Where("id = ?", id).
		First(dataModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			isNotFound = true
		}
		return
	}
	return
}

// QueryOneById query one by id
func (s *userRegMobileRepo) QueryOneById(ctx context.Context, id interface{}) (dataModel *entities.UserRegMobile, isNotFound bool, err error) {
	return s.queryOneById(ctx, s.dbConn, id)
}

// QueryOneByIdWithDBConn query one by id
func (s *userRegMobileRepo) QueryOneByIdWithDBConn(ctx context.Context, dbConn *gorm.DB, id interface{}) (dataModel *entities.UserRegMobile, isNotFound bool, err error) {
	return s.queryOneById(ctx, dbConn, id)
}

// queryOneByConditions query one by conditions
func (s *userRegMobileRepo) queryOneByConditions(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModel *entities.UserRegMobile, isNotFound bool, err error) {
	dataModel = new(entities.UserRegMobile)
	dbConn = dbConn.WithContext(ctx).Table(s.UserRegMobileSchema.TableName())
	err = s.WhereConditions(dbConn, conditions).
		First(dataModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			isNotFound = true
		}
		return
	}
	return
}

// QueryOneByConditions query one by conditions
func (s *userRegMobileRepo) QueryOneByConditions(ctx context.Context, conditions map[string]interface{}) (dataModel *entities.UserRegMobile, isNotFound bool, err error) {
	return s.queryOneByConditions(ctx, s.dbConn, conditions)
}

// QueryOneByConditionsWithDBConn query one by conditions
func (s *userRegMobileRepo) QueryOneByConditionsWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModel *entities.UserRegMobile, isNotFound bool, err error) {
	return s.queryOneByConditions(ctx, dbConn, conditions)
}

// =============== query all : 查全部 ===============

// queryAllByConditions query all by conditions
func (s *userRegMobileRepo) queryAllByConditions(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModels []*entities.UserRegMobile, err error) {
	dbConn = dbConn.WithContext(ctx).Table(s.UserRegMobileSchema.TableName())
	err = s.WhereConditions(dbConn, conditions).
		Find(&dataModels).Error
	if err != nil {
		return dataModels, err
	}
	return
}

// QueryAllByConditions query all by conditions
func (s *userRegMobileRepo) QueryAllByConditions(ctx context.Context, conditions map[string]interface{}) ([]*entities.UserRegMobile, error) {
	return s.queryAllByConditions(ctx, s.dbConn, conditions)
}

// QueryAllByConditionsWithDBConn query all by conditions
func (s *userRegMobileRepo) QueryAllByConditionsWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) ([]*entities.UserRegMobile, error) {
	return s.queryAllByConditions(ctx, dbConn, conditions)
}

// =============== list : 列表 ===============

// list 列表
func (s *userRegMobileRepo) list(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}, paginatorArgs *gormutil.PaginatorArgs) (dataModels []*entities.UserRegMobile, recordCount int64, err error) {
	// query where
	dbConn = dbConn.WithContext(ctx).Table(s.UserRegMobileSchema.TableName())
	dbConn = s.WhereConditions(dbConn, conditions)
	dbConn = gormutil.AssembleWheres(dbConn, paginatorArgs.PageWheres)

	err = dbConn.Count(&recordCount).Error
	if err != nil {
		return
	} else if recordCount == 0 {
		return // empty
	}

	// pagination
	dbConn = gormutil.AssembleOrders(dbConn, paginatorArgs.PageOrders)
	err = gormutil.Paginator(dbConn, paginatorArgs.PageOption).
		Find(&dataModels).Error
	if err != nil {
		return
	}
	return
}

// List 列表
func (s *userRegMobileRepo) List(ctx context.Context, conditions map[string]interface{}, paginatorArgs *gormutil.PaginatorArgs) ([]*entities.UserRegMobile, int64, error) {
	return s.list(ctx, s.dbConn, conditions, paginatorArgs)
}

// ListWithDBConn 列表
func (s *userRegMobileRepo) ListWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}, paginatorArgs *gormutil.PaginatorArgs) ([]*entities.UserRegMobile, int64, error) {
	return s.list(ctx, dbConn, conditions, paginatorArgs)
}

// =============== delete : 删除 ===============

// delete delete one
func (s *userRegMobileRepo) delete(ctx context.Context, dbConn *gorm.DB, dataModel *entities.UserRegMobile) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserRegMobileSchema.TableName()).
		Where("id = ?", dataModel.Id).
		Delete(dataModel).Error
	if err != nil {
		return err
	}
	return
}

// Delete delete one
func (s *userRegMobileRepo) Delete(ctx context.Context, dataModel *entities.UserRegMobile) error {
	return s.delete(ctx, s.dbConn, dataModel)
}

// DeleteWithDBConn delete one
func (s *userRegMobileRepo) DeleteWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *entities.UserRegMobile) error {
	return s.delete(ctx, dbConn, dataModel)
}

// deleteByIds delete by ids
func (s *userRegMobileRepo) deleteByIds(ctx context.Context, dbConn *gorm.DB, ids interface{}) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserRegMobileSchema.TableName()).
		Where("id in (?)", ids).
		Delete(entities.UserRegMobile{}).Error
	if err != nil {
		return err
	}
	return
}

// DeleteByIds delete by ids
func (s *userRegMobileRepo) DeleteByIds(ctx context.Context, ids interface{}) error {
	return s.deleteByIds(ctx, s.dbConn, ids)
}

// DeleteByIdsWithDBConn delete by ids
func (s *userRegMobileRepo) DeleteByIdsWithDBConn(ctx context.Context, dbConn *gorm.DB, ids interface{}) error {
	return s.deleteByIds(ctx, dbConn, ids)
}

// =============== insert : 批量入库 ===============

var _ gormutil.BatchInsertRepo = new(UserRegMobileSlice)

// UserRegMobileSlice 表切片
type UserRegMobileSlice []*entities.UserRegMobile

// TableName 表名
func (s *UserRegMobileSlice) TableName() string {
	if len(*s) > 0 {
		return (*s)[0].TableName()
	}
	return (&entities.UserRegMobile{}).TableName()
}

// Len 长度
func (s *UserRegMobileSlice) Len() int {
	return len(*s)
}

// InsertColumns 批量入库的列
func (s *UserRegMobileSlice) InsertColumns() (columnList []string, placeholder string) {
	// columns
	columnList = []string{
		"user_mobile", "created_time", "updated_time", "user_uuid",
	}

	// placeholders
	insertPlaceholderBytes := bytes.Repeat([]byte("?, "), len(columnList))
	insertPlaceholderBytes = bytes.TrimSuffix(insertPlaceholderBytes, []byte(", "))

	return columnList, string(insertPlaceholderBytes)
}

// InsertValues 批量入库的值
func (s *UserRegMobileSlice) InsertValues(args *gormutil.BatchInsertValueArgs) (prepareData []interface{}, placeholderSlice []string) {
	dataModels := (*s)[args.StepStart:args.StepEnd]
	for index := range dataModels {
		// placeholder
		placeholderSlice = append(placeholderSlice, "("+args.InsertPlaceholder+")")

		// prepare data
		prepareData = append(prepareData, dataModels[index].UserMobile)
		prepareData = append(prepareData, dataModels[index].CreatedTime)
		prepareData = append(prepareData, dataModels[index].UpdatedTime)
		prepareData = append(prepareData, dataModels[index].UserUuid)
	}
	return prepareData, placeholderSlice
}

// UpdateColumns 批量入库的列
func (s *UserRegMobileSlice) UpdateColumns() (columnList []string) {
	// columns
	columnList = []string{
		"user_mobile=excluded.user_mobile", "created_time=excluded.created_time", "updated_time=excluded.updated_time",
		"user_uuid=excluded.user_uuid",
	}
	return columnList
}

// ConflictActionForMySQL 更新冲突时的操作
func (s *UserRegMobileSlice) ConflictActionForMySQL() (req *gormutil.BatchInsertConflictActionReq) {
	req = &gormutil.BatchInsertConflictActionReq{
		OnConflictValueAlias:  "AS excluded",
		OnConflictTarget:      "ON DUPLICATE KEY",
		OnConflictAction:      "UPDATE " + strings.Join(s.UpdateColumns(), ", "),
		OnConflictPrepareData: nil,
	}
	return req
}

// ConflictActionForPostgres 更新冲突时的操作
func (s *UserRegMobileSlice) ConflictActionForPostgres() (req *gormutil.BatchInsertConflictActionReq) {
	req = &gormutil.BatchInsertConflictActionReq{
		OnConflictValueAlias:  "",
		OnConflictTarget:      "ON CONFLICT(id)",
		OnConflictAction:      "DO UPDATE SET " + strings.Join(s.UpdateColumns(), ", "),
		OnConflictPrepareData: nil,
	}
	return req
}

// insert 批量插入
func (s *userRegMobileRepo) insert(ctx context.Context, dbConn *gorm.DB, dataModels UserRegMobileSlice) error {
	return gormutil.BatchInsertWithContext(ctx, dbConn, &dataModels)
}

// Insert 批量插入
func (s *userRegMobileRepo) Insert(ctx context.Context, dataModels []*entities.UserRegMobile) error {
	return s.insert(ctx, s.dbConn, dataModels)
}

// InsertWithDBConn 批量插入
func (s *userRegMobileRepo) InsertWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModels []*entities.UserRegMobile) error {
	return s.insert(ctx, dbConn, dataModels)
}

// =============== conditions : 条件 ===============

// WhereConditions orm where
func (s *userRegMobileRepo) WhereConditions(dbConn *gorm.DB, conditions map[string]interface{}) *gorm.DB {

	// table name
	//tableName := s.UserRegMobileSchema.TableName()

	// On-demand loading

	// id
	//if data, ok := conditions["id"]; ok {
	//	dbConn = dbConn.Where(tableName+".id = ?", data)
	//}

	// user_mobile
	//if data, ok := conditions["user_mobile"]; ok {
	//	dbConn = dbConn.Where(tableName+".user_mobile = ?", data)
	//}

	// created_time
	//if data, ok := conditions["created_time"]; ok {
	//	dbConn = dbConn.Where(tableName+".created_time = ?", data)
	//}

	// updated_time
	//if data, ok := conditions["updated_time"]; ok {
	//	dbConn = dbConn.Where(tableName+".updated_time = ?", data)
	//}

	// user_uuid
	//if data, ok := conditions["user_uuid"]; ok {
	//	dbConn = dbConn.Where(tableName+".user_uuid = ?", data)
	//}

	return dbConn
}

// UpdateColumns update columns
func (s *userRegMobileRepo) UpdateColumns(conditions map[string]interface{}) map[string]interface{} {

	// update columns
	updateColumns := make(map[string]interface{})

	// On-demand loading

	// id
	//if data, ok := conditions["id"]; ok {
	//	updateColumns["id"] = data
	//}

	// user_mobile
	//if data, ok := conditions["user_mobile"]; ok {
	//	updateColumns["user_mobile"] = data
	//}

	// created_time
	//if data, ok := conditions["created_time"]; ok {
	//	updateColumns["created_time"] = data
	//}

	// updated_time
	//if data, ok := conditions["updated_time"]; ok {
	//	updateColumns["updated_time"] = data
	//}

	// user_uuid
	//if data, ok := conditions["user_uuid"]; ok {
	//	updateColumns["user_uuid"] = data
	//}

	return updateColumns
}
