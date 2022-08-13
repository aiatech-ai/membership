package dto

type ProfileRequest struct {
	Email string `json:"email"`
}

type ProfileResponse struct {
	UserID    int64  `json:"userId"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UpdateProfileRequest struct {
	UserID    int64  `url:"userID" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type UpdateProfileResponse struct {
	UserID int64 `json:"userId"`
}

type UpdatePasswordRequest struct {
	UserID          int64  `url:"userID" binding:"required"`
	CurrentPassword string `json:"currentPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
}

type UpdatePasswordResponse struct {
	UserID int64 `json:"userId"`
}
