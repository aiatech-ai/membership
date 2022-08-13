package model

type Organization struct {
	Id                 int32  `gorm:"column:id;type:INT;PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Name               string `gorm:"column:name;type:VARCHAR(255);NOT NULL"`
	Infor              string `gorm:"column:infor;type:VARCHAR(255);"`
	RootOrganizationId int32  `gorm:"column:root_organization_id;type:INT;"`
	Type               int    `gorm:"column:type;type:ENUM('SCHOOL', 'CENTER', 'COMPANY');NOT NULL"`
}
