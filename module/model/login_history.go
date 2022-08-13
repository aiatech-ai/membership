package model

import "time"

type LoginHistory struct {
	Id        int64     `gorm:"column:id;type:BIGINT;PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	UseId     int32     `gorm:"column:use_id;type:INT;NOT NULL"`
	Device    string    `gorm:"column:device;type:VARCHAR(255);NOT NULL"`
	IP        string    `gorm:"column:IP;type:VARCHAR(255);NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;type:DATETIME;NOT NULL"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:DATETIME;NOT NULL"`
}
