package utils

import (
	"regexp"
	"unicode"
)

// VerifyPhone 电话校验
func VerifyPhone(mobileNum string) bool {
	regular := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// VerifyEmail 邮箱校验
func VerifyEmail(email string) bool {
	regular := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return regular.MatchString(email)
}

// VerifyCharPassword 密码校验 字母大小写与8-16位数字
func VerifyCharPassword(s string) bool {
	var hasNumber, hasLowercase, hasPunct bool
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsLower(c):
			hasLowercase = true
		case unicode.IsPunct(c):
			hasPunct = true
		default:
			return false
		}
	}
	return hasNumber && hasLowercase && hasPunct

}

// VerifyUserName 用户名校验 字母大小写与8-16位数字
func VerifyUserName(s string) bool {
	var hasNumber, hasLowercase bool
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsLower(c):
			hasLowercase = true
		default:
			return false
		}
	}
	return hasNumber && hasLowercase
}

func SpecialLetters(letter rune) (bool, []rune) {
	if unicode.IsPunct(letter) || unicode.IsSymbol(letter) {
		var chars []rune
		chars = append(chars, '\\', letter)
		return true, chars
	}
	return false, nil
}
