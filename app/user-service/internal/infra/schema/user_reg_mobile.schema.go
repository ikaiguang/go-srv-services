// Package schemas
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package schemas

import (
	migrationuitl "github.com/ikaiguang/go-srv-kit/data/migration"
	gorm "gorm.io/gorm"
	"time"
)

var _ = time.Time{}

// UserRegMobileSchema UserRegMobile
var UserRegMobileSchema UserRegMobile

// NewUserRegMobile new schema
func NewUserRegMobile() *UserRegMobile {
	return &UserRegMobile{}
}

// UserRegMobile ENGINE InnoDB CHARSET utf8mb4 COMMENT '用户-注册-手机号'
type UserRegMobile struct {
	Id          uint64    `gorm:"column:id;primaryKey;type:uint;autoIncrement;comment:ID" json:"id"`
	UserMobile  string    `gorm:"column:user_mobile;unique;type:string;size:255;not null;default:'';comment:用户手机号码" json:"user_mobile"`
	CreatedTime time.Time `gorm:"column:created_time;type:time;not null;comment:创建时间" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time;type:time;not null;comment:最后修改时间" json:"updated_time"`
	UserUuid    string    `gorm:"column:user_uuid;unique;type:string;size:255;not null;default:'';comment:uuid" json:"user_uuid"`
}

// TableName table name
func (s *UserRegMobile) TableName() string {
	return "srv_user_reg_mobile"
}

// CreateTableMigrator create table migrator
func (s *UserRegMobile) CreateTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationRepo {
	return migrationuitl.NewCreateTable(migrator, s)
}

// DropTableMigrator create table migrator
func (s *UserRegMobile) DropTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationRepo {
	return migrationuitl.NewDropTable(migrator, s)
}
