package md5

import (
	md "crypto/md5"
	"fmt"
)

// Get 获取字符串的 MD5
func Get(str string) string {
	return fmt.Sprintf("%x", md.Sum([]byte(str)))
}
