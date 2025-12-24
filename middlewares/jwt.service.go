package middlewares

import (
	"errors"
	"fmt"
	"time"

	envConfig "github.com/pius706975/the-sims-backend/config"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(envConfig.LoadConfig().JwtSecret)

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

func NewToken(payload TokenPayload, expiresIn time.Duration) *Claims {
	return &Claims{
		UserId:           payload.UserId,
		RoleId:           payload.RoleId,
		Email:            payload.Email,
		Username:         payload.Username,
		Name:             payload.Name,
		IsActivated:      payload.IsActivated,
		IsSuperUser:      payload.IsSuperUser,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn))},
	}
}

func (claim *Claims) CreateToken() (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return accessToken.SignedString(jwtSecret)
}

func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
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

func DecodeRefreshToken(tokenString string) (*TokenPayload, error) {
	envCfg := envConfig.LoadConfig()
	secretKey := []byte(envCfg.JwtRefreshTokenSecret)

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to parse claims")
	}

	var roleId *string
	if val, exists := claims["role_id"]; exists && val != nil {
		str := val.(string)
		roleId = &str
	}

	payload := &TokenPayload{
		UserId:      claims["user_id"].(string),
		RoleId:      roleId,
		Email:       claims["email"].(string),
		Username:    claims["username"].(string),
		Name:        claims["name"].(string),
		IsActivated: claims["is_activated"].(bool),
		IsSuperUser: claims["is_superuser"].(bool),
	}

	return payload, nil
}
