package model

type Role struct {
	Id          int32  `gorm:"column:id;type:INT;PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Name        string `gorm:"column:name;type:VARCHAR(255);"`
	Description string `gorm:"column:description;type:VARCHAR(255);"`
}
