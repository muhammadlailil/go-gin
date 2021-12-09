package jwt_token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"
)

type JwtTokenData struct {
	Scope   string `json:"scope"`
	UsersId int8   `json:"users_id"`
	Email   string `json:"email"`
}

func GenerateToken(data *JwtTokenData) (string, error) {

	token_lifespan := 24 // (satuan jam)

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	claims["data"] = data
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractExpired(tokenString string) (float64, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims["exp"].(float64), nil
	}
	return 0, nil
}

func Data(tokenString string) (*JwtTokenData, error) {
	data := &JwtTokenData{}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return data, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		dataToken := claims["data"].(map[string]interface{})
		idfloat := dataToken["users_id"].(float64)
		data.Scope = dataToken["scope"].(string)
		if idfloat != 0 {
			data.UsersId = dataToken["users_id"].(int8)
		}
		data.Email = dataToken["email"].(string)
		return data, err
	}
	return data, err
}
