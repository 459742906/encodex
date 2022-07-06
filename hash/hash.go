package hash

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	h :=md5.New()
	h.Write([]byte(str))
	re := h.Sum(nil)
	md5Str := fmt.Sprintf("%#v",re)

	return md5Str+"v1.0.2"
}