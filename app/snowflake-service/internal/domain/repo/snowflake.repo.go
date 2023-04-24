// Package repos
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package repos

import (
	context "context"
	entities "github.com/ikaiguang/go-srv-services/app/snowflake-service/internal/domain/entity"
)

// SnowflakeNodeIDRepo repo
type SnowflakeNodeIDRepo interface {
	Create(ctx context.Context, dataModel *entities.SnowflakeNodeID) error
	Update(ctx context.Context, dataModel *entities.SnowflakeNodeID) error
	ExtendNodeID(ctx context.Context, dataModel *entities.SnowflakeNodeID) (err error)
	QueryOneByNodeUUID(ctx context.Context, instanceID string, snowflakeNodeID int64) (dataModel *entities.SnowflakeNodeID, isNotFound bool, err error)
	QueryOneByIDAndNodeUUID(ctx context.Context, req *entities.SnowflakeNodeID) (dataModel *entities.SnowflakeNodeID, isNotFound bool, err error)
	QueryMaxNodeIDByInstanceID(ctx context.Context, instanceID string) (dataModels []*entities.InstanceMaxNodeID, err error)
	QueryIdleNodeIDByInstanceID(ctx context.Context, req *entities.InstanceIdleNodeIDReq) (dataModel *entities.InstanceMaxNodeID, isNotFound bool, err error)
}