package auth

import (
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
