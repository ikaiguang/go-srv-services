// Package schemas
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package schemas

import (
	migrationuitl "github.com/ikaiguang/go-srv-kit/data/migration"
	gorm "gorm.io/gorm"
	"time"
)

var _ = time.Time{}

// SnowflakeNodeIDSchema SnowflakeNodeID
var SnowflakeNodeIDSchema SnowflakeNodeID

// NewSnowflakeNodeID new schema
func NewSnowflakeNodeID() *SnowflakeNodeID {
	return &SnowflakeNodeID{}
}

// SnowflakeNodeID ENGINE InnoDB CHARSET utf8mb4 COMMENT '雪花算法节点ID'
type SnowflakeNodeID struct {
	Id                   uint64    `gorm:"column:id;primaryKey;type:uint;autoIncrement;comment:id" json:"id"`
	InstanceId           string    `gorm:"column:instance_id;uniqueIndex:idx_srv_snowflake_node_id_instance_id_nid;type:string;size:255;not null;default:'';comment:实例ID" json:"instance_id"`
	SnowflakeNodeId      int32     `gorm:"column:snowflake_node_id;uniqueIndex:idx_srv_snowflake_node_id_instance_id_nid;type:int;not null;default:0;comment:雪花算法节点id" json:"snowflake_node_id"`
	InstanceLaunchTime   time.Time `gorm:"column:instance_launch_time;type:time;not null;comment:实例启动时间" json:"instance_launch_time"`
	InstanceExtendTime   time.Time `gorm:"column:instance_extend_time;index;type:time;not null;comment:实例续期时间" json:"instance_extend_time"`
	InstanceName         string    `gorm:"column:instance_name;type:string;size:255;not null;default:'';comment:实例名称" json:"instance_name"`
	InstanceEndpointList string    `gorm:"column:instance_endpoint_list;type:json;not null;comment:实例端点数组" json:"instance_endpoint_list"`
	InstanceMetadata     *string   `gorm:"column:instance_metadata;type:json;comment:实例元数据" json:"instance_metadata"`
	CreatedTime          time.Time `gorm:"column:created_time;type:time;not null;comment:创建时间" json:"created_time"`
}

// TableName table name
func (s *SnowflakeNodeID) TableName() string {
	return "srv_snowflake_node_id"
}

// CreateTableMigrator create table migrator
func (s *SnowflakeNodeID) CreateTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationRepo {
	return migrationuitl.NewCreateTable(migrator, s)
}

// DropTableMigrator create table migrator
func (s *SnowflakeNodeID) DropTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationRepo {
	return migrationuitl.NewDropTable(migrator, s)
}

// CreateUniqueIndexForInstanceIDAndNodeID 创建唯一索引
func (s *SnowflakeNodeID) CreateUniqueIndexForInstanceIDAndNodeID(migrator gorm.Migrator) {
	//migrator.CreateIndex()
}
