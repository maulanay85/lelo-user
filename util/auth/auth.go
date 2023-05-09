package auth

import (
	"errors"
	"lelo-user/config"
	entity "lelo-user/entity"
	"strings"
	"time"

	response "lelo-user/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (u *UtilAuthModule) HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 16)
	if err != nil {
		log.Infof("error hash password: %v", err)
		return "", err
	}
	return string(bytes), nil
}

func (u *UtilAuthModule) CheckHashPassword(pass string, hashPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
	return err == nil
}

func (u *UtilAuthModule) GenerateToken(userRole *entity.UserRoleEntityJoin) (*entity.TokenEntity, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	configData := config.ConfigData
	credentialData := config.CredentialData

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userRole.Email
	claims["name"] = userRole.Fullname
	claims["id"] = userRole.Id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(configData.Exp.Token)).Unix()
	claims["role"] = userRole.RoleName

	t, err := token.SignedString([]byte(credentialData.Jwt.SecretKey))

	if err != nil {
		log.Errorf("[util.auth]: Generate token err: %v", err)
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS512)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = userRole.Email
	rtClaims["id"] = userRole.Id
	rtClaims["exp"] = time.Now().Add(time.Minute * time.Duration(configData.Exp.RefreshToken)).Unix()

	rt, err := refreshToken.SignedString([]byte(credentialData.Jwt.SecretKey))
	if err != nil {
		log.Errorf("[util.auth]: Generate refresh token err: %v", err)
		return nil, err
	}

	return &entity.TokenEntity{
		AccessToken:  t,
		RefreshToken: rt,
	}, nil
}

func (u *UtilAuthModule) JwtTokenCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		res := response.ResponseEntity{Code: 401, Message: "unauthicated"}
		c.AbortWithStatusJSON(401, res)
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		res := response.ResponseEntity{Code: 500, Message: "internal server error"}
		c.AbortWithStatusJSON(500, res)
		return
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		res := response.ResponseEntity{Code: 500, Message: "internal server error"}
		c.AbortWithStatusJSON(500, res)
		return
	}
	id := claim["id"]
	role := claim["role"]

	c.Set("id", id)
	c.Set("role", role)
	c.Next()
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header for value")
	}
	jwt := strings.Split(header, " ")
	if len(jwt) != 2 {
		return "", errors.New("incorectly authorization header")
	}
	return jwt[1], nil
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	credentialData := config.CredentialData
	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad method signed received")
		}
		return []byte(credentialData.Jwt.SecretKey), nil
	})
	if err != nil {
		return nil, errors.New("bad jwt token")
	}
	return token, nil
}

func (u *UtilAuthModule) RefreshToken(rt string) (int64, error) {
	token, err := parseToken(rt)
	if err != nil {
		log.Errorf("error parse token: %v", err)
		return 0, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Errorf("error claims token: %v", err)
		return 0, err
	}
	id := claim["id"]

	return int64(id.(float64)), nil
}
