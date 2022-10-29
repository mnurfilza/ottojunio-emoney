package authentication

import (
	"e-money-svc/domain/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var APPLICATION_NAME = "e-money-svc"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

func CreateToken(username string) (string, error) {
	claims := model.MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: &jwt.NumericDate{time.Now().Add(LOGIN_EXPIRATION_DURATION)},
		},
		Username: username,
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)

	if err != nil {
		return "", err
	}

	return signedToken, nil

}

func AutheticateToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header found"})
			return
		}

		tokens, err := jwt.ParseWithClaims(token, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		claims, ok := tokens.Claims.(*model.MyClaims)
		if !ok || !tokens.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not Valid Token or Token Expired"})
			return
		}
		c.Set("username", claims.Username)

	}

}
