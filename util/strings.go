package util

import (
	"strconv"
	"strings"
)

func SplitFunc(s string, f func(rune) bool) []string {
	sa := make([]string, 0)
	if len(s) <= 0 {
		return sa
	} else {
		i := strings.IndexFunc(s, f)
		for i != -1 {
			sa = append(sa, s[0:i])
			s = s[i+1:]
			i = strings.IndexFunc(s, f)
		}
		return append(sa, s)
	}
}
func Split(s string) []string {
	return SplitFunc(s, func(r rune) bool {
		switch r {
		case ';', ',':
			return true
		}
		return false
	})
}
func String2Int(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}

	return res
}
func IsEmpty(s string) bool {
	if len(strings.TrimSpace(s)) <= 0 {
		return true
	} else {
		return false
	}
}
func IsEmptyStr(s interface{}) bool {
	return IsEmpty(String(s))
}

// 三目运算的函数
func Ternary(a bool, b, c interface{}) interface{} {
	if a {
		return b
	}
	return c
}

func Duplicate(list []string) (ret []string) {
	// 创建一个临时map用来存储数组元素
	temp := make(map[string]bool)
	var list2 []string
	index := 0
	for _, v := range list {
		// 遍历数组元素，判断此元素是否已经存在map中
		if len(v) > 0 {
			_, ok := temp[v]
			if !ok {
				list2 = append(list2, v)
				temp[v] = true
			}
			index++
		}
	}
	return list2
}
