package sn

import (
	md "crypto/md5"
	"fmt"
)

// Get 获取序列号
func Get() (string, error) {
	str, err := get()
	if err == nil {
		return fmt.Sprintf("%x", md.Sum([]byte(str))), nil
	}
	return "", err
}
