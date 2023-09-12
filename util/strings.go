package util

import (
	"encoding/binary"
	json2 "encoding/json"
	"math"
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
func Int(t1 interface{}) int {
	var t2 int
	switch t1.(type) {
	case uint:
		t2 = int(t1.(uint))
		break
	case int8:
		t2 = int(t1.(int8))
		break
	case uint8:
		t2 = int(t1.(uint8))
		break
	case int16:
		t2 = int(t1.(int16))
		break
	case uint16:
		t2 = int(t1.(uint16))
		break
	case int32:
		t2 = int(t1.(int32))
		break
	case uint32:
		t2 = int(t1.(uint32))
		break
	case int64:
		t2 = int(t1.(int64))
		break
	case uint64:
		t2 = int(t1.(uint64))
		break
	case float32:
		t2 = int(t1.(float32))
		break
	case float64:
		t2 = int(t1.(float64))
		break
	case string:
		t2, _ = strconv.Atoi(t1.(string))
		break
	default:
		t2 = t1.(int)
		break
	}
	return t2
}
func String(value interface{}) string {
	var key string
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json2.Marshal(value)
		key = string(newValue)
	}
	return key
}
func Int64(t1 interface{}) int64 {
	var t2 int64
	switch t1.(type) {
	case uint:
		t2 = int64(t1.(uint))
		break
	case int8:
		t2 = int64(t1.(int8))
		break
	case uint8:
		t2 = int64(t1.(uint8))
		break
	case int16:
		t2 = int64(t1.(int16))
		break
	case uint16:
		t2 = int64(t1.(uint16))
		break
	case int32:
		t2 = int64(t1.(int32))
		break
	case uint32:
		t2 = int64(t1.(uint32))
		break
	case int64:
		t2 = int64(t1.(int64))
		break
	case uint64:
		t2 = int64(t1.(uint64))
		break
	case float32:
		t2 = int64(t1.(float32))
		break
	case float64:
		t2 = int64(t1.(float64))
		break
	case string:
		_t2, _ := strconv.Atoi(t1.(string))
		t2 = int64(_t2)
		break
	default:
		t2 = t1.(int64)
		break
	}
	return t2
}

// Float32 converts `any` to float32.
func Float32(any interface{}) float32 {
	if any == nil {
		return 0
	}
	switch value := any.(type) {
	case float32:
		return value
	case float64:
		return float32(value)
	case []byte:
		return math.Float32frombits(binary.LittleEndian.Uint32(LeFillUpSize(value, 4)))
	default:
		v, _ := strconv.ParseFloat(String(any), 64)
		return float32(v)
	}
}
func LeFillUpSize(b []byte, l int) []byte {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]byte, l)
	copy(c, b)
	return c
}

// Float64 converts `any` to float64.
func Float64(any interface{}) float64 {
	if any == nil {
		return 0
	}
	switch value := any.(type) {
	case float32:
		return float64(value)
	case float64:
		return value
	case []byte:
		return math.Float64frombits(binary.LittleEndian.Uint64(LeFillUpSize(value, 8)))
	default:
		v, _ := strconv.ParseFloat(String(any), 64)
		return v
	}
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
