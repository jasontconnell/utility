package utility

import (
	"fmt"
	"crypto/md5"
)

func MD5(content []byte) string {
	sum := md5.Sum(content)
	//final := base64.URLEncoding.EncodeToString(sum)
	return fmt.Sprintf("%x", sum)
}