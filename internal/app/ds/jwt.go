package ds

import (
	"github.com/golang-jwt/jwt"
	"rip2023/internal/app/role"
)

type JWTClaims struct {
	jwt.StandardClaims

	UserID uint `json:"user_id"`

	Role role.Role
}
