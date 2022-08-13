package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shinystarvn/membership/cmd/restapi/api"
	"github.com/shinystarvn/membership/internal/security/token"
	"github.com/shinystarvn/membership/module/user/dto"
	"github.com/shinystarvn/membership/module/user/service"
)

type UserController struct {
	engine              *gin.Engine
	authenticationToken token.AuthenticationToken
	userService         service.UserService
}

var _ api.Controller = (*UserController)(nil)

func NewUserController(engine *gin.Engine, authenticationToken token.AuthenticationToken, userService service.UserService) *UserController {
	return &UserController{
		engine:              engine,
		authenticationToken: authenticationToken,
		userService:         userService,
	}
}

func (s *UserController) Route() {
	router := s.engine.Group("/api/v1/user")
	router.Use(api.AuthorizationMiddleware(s.authenticationToken))
	router.GET("/profile", s.GetProfile)
	router.PUT("/:userID/profile", s.UpdateProfile)
	router.POST("/update-password", s.UpdatePassword)
}

func (s *UserController) GetProfile(c *gin.Context) {
	_, err := api.GetTokenClaims(c)
	if err != nil {
		api.Error(c, err)
		return
	}

	var req dto.ProfileRequest
	err = api.ParseRequest(c, &req)
	if err != nil {
		api.Error(c, err)
		return
	}

	user, err := s.userService.GetUserProfile(c.Request.Context(), req.Email)
	if err != nil {
		api.Error(c, err)
		return
	}

	api.OK(c, &dto.ProfileResponse{
		UserID:    user.Id,
		Email:     user.Email,
		FirstName: user.FullName,
		LastName:  user.FullName,
	})
}

func (s *UserController) UpdateProfile(c *gin.Context) {
	_, err := api.GetTokenClaims(c)
	if err != nil {
		api.Error(c, err)
		return
	}

	var req dto.UpdateProfileRequest
	err = api.ParseRequest(c, &req)
	if err != nil {
		api.Error(c, err)
		return
	}

	err = s.userService.UpdateUserProfile(c.Request.Context(), req.UserID, service.UpdateUserProfileInput{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		api.Error(c, err)
		return
	}

	api.OK(c, &dto.UpdateProfileResponse{
		UserID: req.UserID,
	})

}

func (s *UserController) UpdatePassword(c *gin.Context) {
	_, err := api.GetTokenClaims(c)
	if err != nil {
		api.Error(c, err)
		return
	}

	var req dto.UpdatePasswordRequest
	err = api.ParseRequest(c, &req)
	if err != nil {
		api.Error(c, err)
		return
	}

	err = s.userService.UpdatePassword(c.Request.Context(), req.UserID, service.UpdatePasswordInput{
		CurrentPassword: req.CurrentPassword,
		NewPassword:     req.NewPassword,
	})

	if err != nil {
		api.Error(c, err)
		return
	}

	api.OK(c, &dto.UpdateProfileResponse{
		UserID: req.UserID,
	})
}
