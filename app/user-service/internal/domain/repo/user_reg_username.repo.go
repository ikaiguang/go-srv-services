// Package repos
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package repos

import (
	context "context"
	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"

	entities "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/entity"
)

// UserRegUsernameRepo repo
type UserRegUsernameRepo interface {
	Create(ctx context.Context, dataModel *entities.UserRegUsername) error
	ExistCreate(ctx context.Context, dataModel *entities.UserRegUsername) (anotherModel *entities.UserRegUsername, isNotFound bool, err error)
	CreateInBatches(ctx context.Context, dataModels []*entities.UserRegUsername, batchSize int) error
	Insert(ctx context.Context, dataModels []*entities.UserRegUsername) error
	Update(ctx context.Context, dataModel *entities.UserRegUsername) error
	ExistUpdate(ctx context.Context, dataModel *entities.UserRegUsername) (anotherModel *entities.UserRegUsername, isNotFound bool, err error)
	QueryOneById(ctx context.Context, id interface{}) (dataModel *entities.UserRegUsername, isNotFound bool, err error)
	QueryOneByConditions(ctx context.Context, conditions map[string]interface{}) (dataModel *entities.UserRegUsername, isNotFound bool, err error)
	QueryAllByConditions(ctx context.Context, conditions map[string]interface{}) (dataModels []*entities.UserRegUsername, err error)
	List(ctx context.Context, conditions map[string]interface{}, paginatorArgs *gormutil.PaginatorArgs) (dataModels []*entities.UserRegUsername, totalNumber int64, err error)
	Delete(ctx context.Context, dataModel *entities.UserRegUsername) error
	DeleteByIds(ctx context.Context, ids interface{}) error
}
