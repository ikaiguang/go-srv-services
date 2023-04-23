// Package entities
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package entities

import (
	schemas "github.com/ikaiguang/go-srv-services/app/admin-service/internal/infra/schema"
	"time"
)

var _ = time.Time{}

// AdminRegUsername ENGINE InnoDB CHARSET utf8mb4 COMMENT '管理员-注册-用户名'
type AdminRegUsername struct {
	Id          uint64    `gorm:"column:id;primaryKey" json:"id"`          // ID
	UserName    string    `gorm:"column:user_name" json:"user_name"`       // 管理员名
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"` // 最后修改时间
	AdminUuid   string    `gorm:"column:admin_uuid" json:"admin_uuid"`     // uuid
}

// TableName table name
func (s *AdminRegUsername) TableName() string {
	return schemas.AdminRegUsernameSchema.TableName()
}
