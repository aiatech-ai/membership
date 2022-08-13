package model

type RolePermission struct {
	Id           int32 `gorm:"column:id;type:INT;PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	RoleId       int32 `gorm:"column:role_id;type:INT;NOT NULL"`
	PermissionId int32 `gorm:"column:permission_id;type:INT;NOT NULL"`
}
