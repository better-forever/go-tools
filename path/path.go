package path

import (
	"regexp"
	"strings"
)

func Match(path string, whitelist []string) bool {
	for _, pattern := range whitelist {
		if pathMatchesPattern(path, pattern) {
			return true
		}
	}
	return false
}

// pathMatchesPattern 判断路径是否匹配单个模式
func pathMatchesPattern(path, pattern string) bool {
	// 转换模式为正则表达式
	reStr := "^"

	// 分割路径段
	segments := strings.Split(pattern, "/")
	for i, seg := range segments {
		if i > 0 {
			reStr += "/"
		}

		switch {
		case seg == "":
			// 空段，忽略
			continue
		case seg == "*":
			// 匹配任意字符串
			reStr += ".*"
		case strings.HasPrefix(seg, ":"):
			// 匹配路径参数（如 :id）
			reStr += "[^/]+"
		default:
			// 普通字符串，直接转义
			reStr += regexp.QuoteMeta(seg)
		}
	}

	reStr += "$"

	re, err := regexp.Compile(reStr)
	if err != nil {
		// 模式无效，返回 false
		return false
	}

	return re.MatchString(path)
}
