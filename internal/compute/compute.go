package compute

import (
	"errors"
	"strconv"
	"strings"
)

// 算术结构体
type aggs struct {
	one      float64
	two      float64
	operator string
}

/*
 * 计算
 */
func Compute(s string) (float64, error) {
	var result float64

	box, err := cutApart(s)
	if err != nil {
		return result, err
	}

	if box.operator == "+" {
		result = box.one + box.two
	} else if box.operator == "-" {
		result = box.one - box.two
	} else if box.operator == "*" {
		result = box.one * box.two
	} else if box.operator == "/" {
		result = box.one / box.two
	} else {
		return result, errors.New("该运算符尚未定义")
	}

	return result, nil
}

/*
 * 字符串分割
 */
func cutApart(s string) (aggs, error) {
	var box aggs

	arr := strings.Fields(s)
	if len(arr) != 3 {
		return box, errors.New("表达式有误")
	}

	v1, err := strconv.ParseFloat(arr[0], 64)
	if err != nil {
		return box, err
	}

	v2, err := strconv.ParseFloat(arr[2], 64)
	if err != nil {
		return box, err
	}

	box = aggs{
		one:      v1,
		two:      v2,
		operator: arr[1],
	}

	return box, nil
}
