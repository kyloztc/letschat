package encrypt

import (
	"crypto/md5"
	"fmt"
)

func Md5Encrypt(str string) string {
	_hash := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", _hash)
}
