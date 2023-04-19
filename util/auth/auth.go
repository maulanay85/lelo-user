package auth

import (
	"lelo-user/config"
	entity "lelo-user/entity"
	"time"

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
