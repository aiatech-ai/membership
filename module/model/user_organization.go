package model

import "time"

type UserOrganization struct {
	Id             int64     `gorm:"column:id;type:BIGINT;PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	UserId         int32     `gorm:"column:user_id;type:INT;NOT NULL"`
	OrganizationId int32     `gorm:"column:organization_id;type:INT;NOT NULL"`
	IsActive       bool      `gorm:"column:is_active;type:BOOL;NOT NULL"`
	CreatedAt      time.Time `gorm:"column:created_at;type:DATETIME;NOT NULL"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:DATETIME;"`
}
