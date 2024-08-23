package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Captcha(n int) string {
	code := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		code = fmt.Sprintf("%s%d", code, rand.Intn(10))
	}
	return code
}
