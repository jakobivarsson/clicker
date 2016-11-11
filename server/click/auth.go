package click

import (
	"crypto/sha256"
	"golang.org/x/crypto/bcrypt"
)

func hash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func CreateUser(user string, pass string) {
	db := GetDB()
	sh := sha256.New()
	sh.Write([]byte(user))
	db.StorePassword(string(sh.Sum(nil)), hash(pass))
}

func AuthenticateUser(user string, pass string) bool {
	db := GetDB()
	sh := sha256.New()
	sh.Write([]byte(user))
	if bcrypt.CompareHashAndPassword([]byte(db.GetPassword(string(sh.Sum(nil)))), []byte(pass)) == nil {
		return true
	}
	return false
}