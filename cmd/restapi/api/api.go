package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinystarvn/membership/internal/security/token"
)

const (
	StatusFailed    = 0
	StatusSucceeded = 1
)

var ErrInternalServerError = errors.New("internal server error")
var ErrBadRequest = errors.New("bad request")

type Controller interface {
	Route()
}

func SetupRouter(controllers ...Controller) {
	for _, c := range controllers {
		c.Route()
	}
}

func ParseRequest(c *gin.Context, obj interface{}) error {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		return fmt.Errorf("could not parse request: %v %w", err, ErrBadRequest)
	}
	return nil
}

func WriteResponse(c *gin.Context, code, status int, message string, data interface{}) {
	m := gin.H{
		"status":  status,
		"message": code,
		"debug":   message,
	}
	if data != nil {
		m["data"] = data
	}
	c.JSON(code, m)
}

func BadRequest(c *gin.Context, err error) {
	WriteResponse(c, http.StatusBadRequest, StatusFailed, err.Error(), nil)
}

func Unauthorized(c *gin.Context, err error) {
	WriteResponse(c, http.StatusUnauthorized, StatusFailed, err.Error(), nil)
}

func NotFound(c *gin.Context, err error) {
	WriteResponse(c, http.StatusNotFound, StatusFailed, err.Error(), nil)
}

func Conflict(c *gin.Context, err error) {
	WriteResponse(c, http.StatusConflict, StatusFailed, err.Error(), nil)
}

func InternalServerError(c *gin.Context, err error) {
	WriteResponse(c, http.StatusInternalServerError, StatusFailed, err.Error(), nil)
}

func OK(c *gin.Context, data interface{}) {
	WriteResponse(c, http.StatusOK, StatusSucceeded, "Succeed", data)
}

func Error(c *gin.Context, err error) {
	log.Println(err)
	e := errors.Unwrap(err)
	if e != nil {
		err = e
	}
	switch err {
	case ErrBadRequest:
		BadRequest(c, err)
	case ErrAuthorizationHeaderNotProvided, token.ErrInvalidToken, token.ErrExpiredToken:
		Unauthorized(c, err)
	default:
		InternalServerError(c, ErrInternalServerError)
	}
}
