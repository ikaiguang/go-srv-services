// Package repos
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package repos

import (
	context "context"
	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"

	entities "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/entity"
)

// UserRegEmailRepo repo
type UserRegEmailRepo interface {
	Create(ctx context.Context, dataModel *entities.UserRegEmail) error
	ExistCreate(ctx context.Context, dataModel *entities.UserRegEmail) (anotherModel *entities.UserRegEmail, isNotFound bool, err error)
	CreateInBatches(ctx context.Context, dataModels []*entities.UserRegEmail, batchSize int) error
	Insert(ctx context.Context, dataModels []*entities.UserRegEmail) error
	Update(ctx context.Context, dataModel *entities.UserRegEmail) error
	ExistUpdate(ctx context.Context, dataModel *entities.UserRegEmail) (anotherModel *entities.UserRegEmail, isNotFound bool, err error)
	QueryOneById(ctx context.Context, id interface{}) (dataModel *entities.UserRegEmail, isNotFound bool, err error)
	QueryOneByUserEmail(ctx context.Context, adminEmail string) (dataModel *entities.UserRegEmail, isNotFound bool, err error)
	QueryOneByConditions(ctx context.Context, conditions map[string]interface{}) (dataModel *entities.UserRegEmail, isNotFound bool, err error)
	QueryAllByConditions(ctx context.Context, conditions map[string]interface{}) (dataModels []*entities.UserRegEmail, err error)
	List(ctx context.Context, conditions map[string]interface{}, paginatorArgs *gormutil.PaginatorArgs) (dataModels []*entities.UserRegEmail, totalNumber int64, err error)
	Delete(ctx context.Context, dataModel *entities.UserRegEmail) error
	DeleteByIds(ctx context.Context, ids interface{}) error
}
