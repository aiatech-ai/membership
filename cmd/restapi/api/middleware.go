package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/shinystarvn/membership/internal/security/token"
)

const (
	AuthorizationHeader = "Authorization"
	TokenClaims         = "claims"
)

var ErrAuthorizationHeaderNotProvided = errors.New("authorization header is not provided")
var ErrGetTokenClaims = errors.New("could not retrieve token claims from context")

func AuthorizationMiddleware(authenticationToken token.AuthenticationToken) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader(AuthorizationHeader)
		if len(tokenHeader) <= 7 {
			Error(c, ErrAuthorizationHeaderNotProvided)
			c.Abort()
			return
		}

		accessToken := tokenHeader[7:]

		payload, err := authenticationToken.VerifyToken(accessToken)
		if err != nil {
			Error(c, err)
			c.Abort()
			return
		}

		c.Set(TokenClaims, payload)
		c.Next()
	}
}

func GetTokenClaims(c *gin.Context) (*token.Payload, error) {
	payload, exists := c.Get(TokenClaims)
	if !exists {
		return nil, ErrGetTokenClaims
	}
	return payload.(*token.Payload), nil
}
