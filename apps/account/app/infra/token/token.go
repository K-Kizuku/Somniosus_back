package token

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/K-Kizuku/Somniosus_back/apps/account/app/domain/model/account"
	"github.com/K-Kizuku/Somniosus_back/apps/account/config"
	"github.com/golang-jwt/jwt/v5"
)

type (
	tokenManager struct {
		secretKey     string
		tokenLifeTime int
	}
	TokenManager interface {
		GenerateToken(ctx context.Context, id string, expiresAt time.Time, status account.AccountStatus) (string, error)
		GenerateTempToken(ctx context.Context, id, name, telNum string, birthDay time.Time, expiresAt time.Time) (string, error)
		VerifyToken(ctx context.Context, tokenString string) (*Token, error)
		VerifyTempToken(ctx context.Context, tokenString string) (*CustomClaims, error)
	}
	Token struct {
		ID          string
		ExpiresAt   time.Time
		AccessToken string
	}
	TempToken struct {
		ID          string
		Name        string
		TelNum      string
		BirthDay    time.Time
		ExpiresAt   time.Time
		AccessToken string
	}
	CustomClaims struct {
		Name     string
		TelNum   string
		BirthDay time.Time
		jwt.RegisteredClaims
	}
)

const (
	ExpiresTime int = 36000
)

func NewTokenManager() TokenManager {
	secretKey := config.SigningKey
	tokenLifeTime, err := strconv.Atoi(config.TokenLifeTime)
	if err != nil {
		return nil
	}
	return &tokenManager{
		secretKey:     string(secretKey),
		tokenLifeTime: tokenLifeTime,
	}
}

func (t *tokenManager) GenerateToken(ctx context.Context, id string, expiresAt time.Time, status account.AccountStatus) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
		Issuer:    "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(t.secretKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func (t *tokenManager) GenerateTempToken(ctx context.Context, id, name, telNum string, birthDay time.Time, expiresAt time.Time) (string, error) {
	claims := CustomClaims{
		name,
		telNum,
		birthDay,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(t.secretKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func (t *tokenManager) VerifyToken(ctx context.Context, tokenString string) (*Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return t.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		id, err := token.Claims.GetIssuer()
		if err != nil {
			return nil, err
		}
		expiredAt, err := token.Claims.GetExpirationTime()
		return &Token{
			ID:          id,
			ExpiresAt:   expiredAt.Time,
			AccessToken: tokenString,
		}, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, jwt.ErrTokenMalformed
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return nil, jwt.ErrTokenSignatureInvalid
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return nil, jwt.ErrTokenExpired
	} else {
		return nil, err
	}
}

func (t *tokenManager) VerifyTempToken(ctx context.Context, tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return t.secretKey, nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
