package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"math/rand/v2"
	"strconv"
	"time"
)

func PasswordHash(password string) (string, error) {
	b, e := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(b), e
}

func HashCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(base string) string {
	timeStamp := strconv.Itoa(int(time.Now().Unix()))
	magicNumber := strconv.Itoa(rand.Int())

	var ogtxt []byte = []byte(timeStamp + base + magicNumber)
	h := sha256.New()
	h.Write(ogtxt)

	final_token := hex.EncodeToString(h.Sum(nil))
	return final_token
}
