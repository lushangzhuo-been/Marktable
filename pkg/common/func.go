package common

import (
	"fmt"
	"reflect"
)

func InArray(needle string, haystack []string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func ByteToOther(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func SliceRemoveDuplicate(slice interface{}) (interface{}, error) {
	sValue := reflect.ValueOf(slice)
	if sValue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("func may only be called on slices")
	}

	// Allocate a new slice of the same type and size as the original
	rValue := reflect.MakeSlice(sValue.Type(), 0, sValue.Len())
	rMap := make(map[interface{}]bool)

	for i := 0; i < sValue.Len(); i++ {
		item := sValue.Index(i).Interface()
		if _, found := rMap[item]; !found {
			rMap[item] = true
			rValue = reflect.Append(rValue, reflect.ValueOf(item))
		}
	}

	return rValue.Interface(), nil
}
