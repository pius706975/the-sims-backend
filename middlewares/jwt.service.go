package middlewares

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	envConfig "github.com/pius706975/the-sims-backend/config"
)


// ===========================================
// Payload & Claims
// ===========================================
type TokenPayload struct {
	UserId      string
	RoleId      *string
	Email       string
	Username    string
	Name        string
	IsActivated bool
	IsSuperUser bool
}

type Claims struct {
	UserId      string  `json:"user_id"`
	RoleId      *string `json:"role_id,omitempty"`
	Email       string  `json:"email"`
	Username    string  `json:"username"`
	Name        string  `json:"name"`
	IsActivated bool    `json:"is_activated"`
	IsSuperUser bool    `json:"is_superuser"`
	jwt.RegisteredClaims
}


// ===========================================
// Claim Builder
// ===========================================
func NewAccessToken(payload TokenPayload, expiresIn time.Duration) *Claims {
	return &Claims{
		UserId:      payload.UserId,
		RoleId:      payload.RoleId,
		Email:       payload.Email,
		Username:    payload.Username,
		Name:        payload.Name,
		IsActivated: payload.IsActivated,
		IsSuperUser: payload.IsSuperUser,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		},
	}
}

// ===========================================
// Token Signer
// ===========================================
func CreateTokenWithSecret(claim *Claims, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(secret)
}

// ===========================================
// Token Verifier (Access Token)
// ===========================================
func VerifyAccessToken(tokenString string) (*Claims, error) {
	cfg := envConfig.LoadConfig()
	secret := []byte(cfg.JwtSecret)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// ===========================================
// Decode Refresh Token
// ===========================================
func DecodeRefreshToken(tokenString string) (*TokenPayload, error) {
	cfg := envConfig.LoadConfig()
	secret := []byte(cfg.JwtRefreshTokenSecret)

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to parse claims")
	}

	var roleId *string
	if val, ok := claims["role_id"]; ok && val != nil {
		str := val.(string)
		roleId = &str
	}

	return &TokenPayload{
		UserId:      claims["user_id"].(string),
		RoleId:      roleId,
		Email:       claims["email"].(string),
		Username:    claims["username"].(string),
		Name:        claims["name"].(string),
		IsActivated: claims["is_activated"].(bool),
		IsSuperUser: claims["is_superuser"].(bool),
	}, nil
}
