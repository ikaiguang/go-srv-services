// Package schemas
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package schemas

import (
	gorm "gorm.io/gorm"
	"time"

	migrationuitl "github.com/ikaiguang/go-srv-kit/data/migration"
)

var _ = time.Time{}

// ExampleSchema example
var ExampleSchema Example

// NewExample new schema
func NewExample() *Example {
	return &Example{}
}

// Example ENGINE InnoDB CHARSET utf8mb4 COMMENT '例子'
type Example struct {
	Id          uint64     `gorm:"column:id;primaryKey;type:uint;autoIncrement;comment:ID" json:"id"`
	ExampleUuid string     `gorm:"column:example_uuid;unique;type:string;size:255;not null;default:'';comment:uuid" json:"example_uuid"`
	CreatedTime time.Time  `gorm:"column:created_time;type:time;not null;comment:创建时间" json:"created_time"`
	UpdatedTime time.Time  `gorm:"column:updated_time;type:time;not null;comment:最后修改时间" json:"updated_time"`
	DeletedTime *time.Time `gorm:"column:deleted_time;type:time;comment:删除时间" json:"deleted_time"`
	IsDeleted   bool       `gorm:"column:is_deleted;index;type:bool;not null;default:0;comment:是否已删除" json:"is_deleted"`
}

// TableName table name
func (s *Example) TableName() string {
	return "srv_example"
}

// CreateTableMigrator create table migrator
func (s *Example) CreateTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationRepo {
	return migrationuitl.NewCreateTable(migrator, s)
}

// DropTableMigrator create table migrator
func (s *Example) DropTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationRepo {
	return migrationuitl.NewDropTable(migrator, s)
}