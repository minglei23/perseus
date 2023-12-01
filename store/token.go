package store

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CreateToken(id string) string {
	expires := time.Now().Add(24 * time.Hour)
	hashValue := getHashValue(id, fmt.Sprintf("%v", expires.Unix()))
	return fmt.Sprintf("%s|%v|%s", id, expires.Unix(), hashValue)
}

func CheckToken(token string) bool {
	parts := strings.Split(token, "|")
	if len(parts) != 3 {
		return false
	}
	id, expiresPart, hashValue := parts[0], parts[1], parts[2]
	if id == "" || expiresPart == "" || hashValue == "" {
		return false
	}
	expiresUnix, err := strconv.ParseInt(expiresPart, 10, 64)
	if err != nil {
		return false
	}
	if time.Now().After(time.Unix(expiresUnix, 0)) {
		return false
	}
	return hashValue == getHashValue(id, expiresPart)
}

func getHashValue(name, expires string) string {
	input := name + expires + "Perseus_1110"
	hasher := sha256.New()
	_, err := hasher.Write([]byte(input))
	if err != nil {
		return ""
	}
	hash := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hash)[:32]
	return hashString
}

func GetHashValue(name, expires string) string {
	input := name + expires + "Perseus_2023"
	hasher := sha256.New()
	_, err := hasher.Write([]byte(input))
	if err != nil {
		return ""
	}
	hash := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hash)[:32]
	return hashString
}
