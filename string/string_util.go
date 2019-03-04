package string

/**
 * Author: Yauri Attamimi (yauritux@gmail.com)
 * Version: 2.0.0-RC1
 * Description: Common utility for String.
 *
 */

import (
	"regexp"
)

const (
	BASE64 string = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
)

var (
	rxBase64 = regexp.MustCompile(BASE64)
)

func IsBase64(str string) bool {
	return rxBase64.MatchString(str)
}
