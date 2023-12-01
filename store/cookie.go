package store

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CreateCookie(id string) (http.Cookie, error) {
	expires := time.Now().Add(24 * time.Hour)
	hashValue, err := getHashValue(id, fmt.Sprintf("%v", expires.Unix()))
	if err != nil {
		return http.Cookie{}, err
	}
	return http.Cookie{
		Name:    "perseus",
		Value:   fmt.Sprintf("%s|%v|%s", id, expires.Unix(), hashValue),
		Expires: expires,
		// HttpOnly: true,
		// Secure: true, // Uncomment this when using HTTPS
	}, nil
}

func CheckCookie(cookie http.Cookie) (bool, error) {
	parts := strings.Split(cookie.Value, "|")
	if len(parts) != 3 {
		return false, fmt.Errorf("invalid cookie format")
	}

	id, expiresPart, hashValue := parts[0], parts[1], parts[2]
	if id == "" || expiresPart == "" || hashValue == "" {
		return false, fmt.Errorf("missing cookie components")
	}

	expiresUnix, err := strconv.ParseInt(expiresPart, 10, 64)
	if err != nil {
		return false, fmt.Errorf("invalid expiration time")
	}

	if time.Now().After(time.Unix(expiresUnix, 0)) {
		return false, fmt.Errorf("cookie expired")
	}

	expectedValue, err := getHashValue(id, expiresPart)
	if err != nil {
		return false, err
	}

	return hashValue == expectedValue, nil
}

func getHashValue(name, expires string) (string, error) {
	input := name + expires + "Perseus_1110"
	hasher := sha256.New()
	_, err := hasher.Write([]byte(input))
	if err != nil {
		return "", err
	}
	hash := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hash)[:32]
	return hashString, nil
}

func GetHashValue(name, expires string) (string, error) {
	input := name + expires + "Perseus_2023"
	hasher := sha256.New()
	_, err := hasher.Write([]byte(input))
	if err != nil {
		return "", err
	}
	hash := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hash)[:32]
	return hashString, nil
}
