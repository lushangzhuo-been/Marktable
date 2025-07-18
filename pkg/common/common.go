package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"mark3/types/common"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GetUUID() string {
	timestamp := time.Now().UnixNano()
	str := strconv.FormatInt(timestamp, 10)
	return Md5(str)
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func FormatConfig(config map[string]string) []common.BaseConfig {
	var format []common.BaseConfig
	// 键进行升序
	keys := make([]string, 0, len(config))
	for k := range config {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		format = append(format, common.BaseConfig{Value: k, Label: config[k]})
	}
	return format
}

func RandCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
