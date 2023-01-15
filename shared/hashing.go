package shared

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

func HashMd5(input any) string {
	json, _ := json.Marshal(input)
	hash := md5.Sum(json)
	return hex.EncodeToString(hash[:])
}
