package ds

type JWTClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
	Role   role.Role
}
