package utils

import (
	"cmp"
	"encoding/json"
	"strconv"
	"strings"
)

// 任意位置插入数字类型的元素
// @param slice []int 将指定元素插入的切片
// @param num int 指定元素
// @param index int 插入的指定位置
func arbitrarilyInsertElement(slice []string, num string, index int) []string {
	slice = append(slice[:index], append([]string{num}, slice[index:]...)...)
	return slice
}

//go版本大于1.21的情况，可以使用泛型参数

// 切片去重  泛型参数 利用map的key不能重复的特性 2次for循环
func Unique[T cmp.Ordered](ss []T) []T {
	size := len(ss)
	if size == 0 {
		return []T{}
	}
	// 这个地方利用了map数据的key不能重复的特性,将切片的值当做key放入map中,达到去重的目的
	m1 := make(map[T]T)
	var newSS []T
	for i := 0; i < size; i++ {
		_, ok := m1[ss[i]]
		if ok {
			continue
		} else {
			m1[ss[i]] = ss[i]
			newSS = append(newSS, ss[i])
		}
	}
	return newSS
}

func DeleteSlice(s []string, elem string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == elem {
			s = append(s[:i], s[i+1:]...)
			i-- // 因为切片缩短了，所以索引要回退一位
		}
	}
	return s
}

func LeftListSort(IdsStr string, FromId int, ToId int, Order string) string {
	IdsStr = strings.Trim(IdsStr, ",")
	Ids := strings.Split(IdsStr, ",")
	Ids = DeleteSlice(Ids, strconv.Itoa(FromId))
	var insertIndex = 0
	for index, Id := range Ids {
		IdInt, err := strconv.Atoi(Id)
		if err != nil {
			return ""
		}
		if IdInt == ToId {
			insertIndex = index
			break
		}
	}

	if Order == "before" {
		Ids = arbitrarilyInsertElement(Ids, strconv.Itoa(FromId), insertIndex)
	} else {
		if insertIndex+1 == len(Ids) {
			Ids = append(Ids, strconv.Itoa(FromId))
		} else {
			Ids = arbitrarilyInsertElement(Ids, strconv.Itoa(FromId), insertIndex+1)
		}
	}
	Ids = Unique(Ids)
	return "," + strings.Join(Ids, ",") + ","
}

func StructToMap(obj interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
