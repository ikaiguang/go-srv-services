// Package repos
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package repos

import (
	context "context"
	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"

	entities "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/entity"
)

// UserRepo repo
type UserRepo interface {
	Create(ctx context.Context, dataModel *entities.User) error
	ExistCreate(ctx context.Context, dataModel *entities.User) (anotherModel *entities.User, isNotFound bool, err error)
	CreateInBatches(ctx context.Context, dataModels []*entities.User, batchSize int) error
	Insert(ctx context.Context, dataModels []*entities.User) error
	Update(ctx context.Context, dataModel *entities.User) error
	ExistUpdate(ctx context.Context, dataModel *entities.User) (anotherModel *entities.User, isNotFound bool, err error)
	QueryOneById(ctx context.Context, id interface{}) (dataModel *entities.User, isNotFound bool, err error)
	QueryOneByUserUuid(ctx context.Context, userUuid string) (dataModel *entities.User, isNotFound bool, err error)
	QueryOneByConditions(ctx context.Context, conditions map[string]interface{}) (dataModel *entities.User, isNotFound bool, err error)
	QueryAllByConditions(ctx context.Context, conditions map[string]interface{}) (dataModels []*entities.User, err error)
	List(ctx context.Context, conditions map[string]interface{}, paginatorArgs *gormutil.PaginatorArgs) (dataModels []*entities.User, totalNumber int64, err error)
	Delete(ctx context.Context, dataModel *entities.User) error
	DeleteByIds(ctx context.Context, ids interface{}) error
}