package utils

import (
	"math/rand"
	"regexp"
	"time"
)

// 用于检测字符串是否为有效的手机号格式
func IsPhone(phone string) bool {
	// 定义手机号的正则表达式模式
	pattern := `^1[3-9]\d{9}$`
	// 编译正则表达式
	reg, _ := regexp.Compile(pattern)
	// 使用正则表达式匹配字符串
	return reg.MatchString(phone)
}

// GenCode 生成安全的验证码,codeType 为 n,s,m 分别代表数字，字母，数字和字母混合；len 表示生成的验证码长度
func GenCode(codeType string, codeLen int) string {
	var charset string
	switch codeType {
	case "n":
		charset = "0123456789"
	case "s":
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "m":
		charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, codeLen)
	for i := range code {
		code[i] = charset[r.Intn(len(charset))]
	}
	return string(code)
}
