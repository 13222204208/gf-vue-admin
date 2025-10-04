package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT配置
const (
	// JWT密钥，实际项目中应该从配置文件读取
	JWTSecret = "gf-vue-admin-jwt-secret-key-2024"
	// Access Token过期时间（2小时）
	AccessTokenExpire = 2 * time.Hour
	// Refresh Token过期时间（7天）
	RefreshTokenExpire = 7 * 24 * time.Hour
)

// Claims JWT载荷结构
type Claims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateTokens 生成访问令牌和刷新令牌
func GenerateTokens(userID uint64, username, email string) (accessToken, refreshToken string, err error) {
	// 生成访问令牌
	accessToken, err = GenerateAccessToken(userID, username, email)
	if err != nil {
		return "", "", err
	}

	// 生成刷新令牌
	refreshToken, err = GenerateRefreshToken(userID, username, email)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// GenerateAccessToken 生成访问令牌
func GenerateAccessToken(userID uint64, username, email string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "gf-vue-admin",
			Subject:   "access_token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecret))
}

// GenerateRefreshToken 生成刷新令牌
func GenerateRefreshToken(userID uint64, username, email string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "gf-vue-admin",
			Subject:   "refresh_token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecret))
}

// ParseToken 解析JWT令牌
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ValidateToken 验证JWT令牌
func ValidateToken(tokenString string) (*Claims, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	// 检查令牌是否过期
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

// RefreshAccessToken 使用刷新令牌生成新的访问令牌
func RefreshAccessToken(refreshTokenString string) (string, error) {
	claims, err := ParseToken(refreshTokenString)
	if err != nil {
		return "", err
	}

	// 验证是否为刷新令牌
	if claims.Subject != "refresh_token" {
		return "", errors.New("invalid refresh token")
	}

	// 检查刷新令牌是否过期
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return "", errors.New("refresh token expired")
	}

	// 生成新的访问令牌
	return GenerateAccessToken(claims.UserID, claims.Username, claims.Email)
}