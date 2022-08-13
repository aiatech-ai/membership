package model

type UserRole struct {
	Id     int64 `gorm:"column:id;type:BIGINT;PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	UserId int32 `gorm:"column:user_id;type:INT;NOT NULL"`
	RoleId int32 `gorm:"column:role_id;type:INT;NOT NULL"`
}
