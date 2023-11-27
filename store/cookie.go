package store

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"
)

func CreateCookie(id string) (http.Cookie, error) {
	expires := time.Now().Add(24 * time.Hour)
	value, err := getHashValue(id, expires)
	if err != nil {
		return http.Cookie{}, err
	}
	return http.Cookie{
		Name:     "perseus",
		Value:    value,
		Expires:  expires,
		HttpOnly: true,
		// // HTTPS ONLY
		// Secure: true,
	}, nil
}

func CheckCookie(id string, cookie http.Cookie) (bool, error) {
	expectedValue, err := getHashValue(id, cookie.Expires)
	if err != nil {
		return false, err
	}
	return cookie.Value == expectedValue, nil
}

func getHashValue(name string, expires time.Time) (string, error) {
	input := fmt.Sprintf("%s-%v-Perseus_2023", name, expires.Unix())
	hasher := sha256.New()
	_, err := hasher.Write([]byte(input))
	if err != nil {
		return "", err
	}
	hash := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hash)[:32]
	return hashString, nil
}

func GetHashValue(name string, expires time.Time) (string, error) {
	input := fmt.Sprintf("%s-%v-Perseus_1110", name, expires.Unix())
	hasher := sha256.New()
	_, err := hasher.Write([]byte(input))
	if err != nil {
		return "", err
	}
	hash := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hash)[:32]
	return hashString, nil
}
