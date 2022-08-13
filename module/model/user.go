package model

import "time"

type User struct {
	Id           int64     `gorm:"column:id;type:INT;PRIMARY_KEY;NOT NULL"`
	FullName     string    `gorm:"column:full_name;type:VARCHAR(255);"`
	Email        string    `gorm:"column:email;type:VARCHAR(255);NOT NULL"`
	Tel          string    `gorm:"column:tel;type:VARCHAR(255);NOT NULL"`
	PasswordHash string    `gorm:"column:password_hash;type:VARCHAR(255);NOT NULL"`
	Otp          string    `gorm:"column:otp;type:VARCHAR(255);"`
	OtpTimeout   time.Time `gorm:"column:otp_timeout;type:DATETIME;"`
	IsActive     bool      `gorm:"column:is_active;type:BOOL;"`
	AvatarPath   string    `gorm:"column:avatar_path;type:VARCHAR(255);"`
	UpdateReason int       `gorm:"column:update_reason;type:ENUM('CREATE_NEW', 'UPDATE_FULL_NAME', 'UPDATE_PASSWORD', 'UPDATE_OTP', 'UPDATE_TO_ACTIVE', 'UPDATE_TO_DEACTIVE');NOT NULL"`
	CreatedAt    time.Time `gorm:"column:created_at;type:DATETIME;NOT NULL"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:DATETIME;NOT NULL"`
}
