package helper

import (
	err "errors"
	"log"
	"os"
	"strconv"

	"time"

	"github.com/CRobinDev/Gemastik/model"
	"github.com/CRobinDev/Gemastik/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID
	jwt.RegisteredClaims
}

func SignJWTToken(userID uuid.UUID) (string, error) {
	expiredTime, err := strconv.Atoi(os.Getenv("JWT_EXP_TIME"))
	if err != nil {
		log.Fatalf("failed set expired time for jwt : %v", err.Error())
	}
	ExpiredAt := time.Duration(expiredTime) * 24
	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ExpiredAt) * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTSecretKey))

	if err != nil {
		return "", errors.ErrSigningJWT
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (uuid.UUID, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	var (
		claims Claims
		userID uuid.UUID
	)

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return userID, errors.ErrClaimsJWT
	}

	if !token.Valid {
		return userID, err
	}

	userID = claims.UserID

	return userID, nil
}

func GetLoginUser(ctx *gin.Context) (model.ServiceResponse, error) {
	user, ok := ctx.Get("user")
	if !ok {
		return model.ServiceResponse{}, err.New("failed to get user")
	}

	if u, ok := user.(model.ServiceResponse); ok {
		return u, nil
	} else {
		return model.ServiceResponse{}, err.New("user is not of type entity.User")
	}
}
