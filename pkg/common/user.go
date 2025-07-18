package common

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) ([]byte, error) {
	// 将密码转换为字节数组
	passwordBytes := []byte(password)

	// 生成哈希值
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// 返回哈希值
	return hashedPassword, nil
}

// ValidatePassword 验证密码函数
func ValidatePassword(hashedPassword []byte, password string) error {
	// 将密码转换为字节数组
	passwordBytes := []byte(password)

	// 验证密码
	err := bcrypt.CompareHashAndPassword(hashedPassword, passwordBytes)
	if err != nil {
		return err
	}

	// 密码匹配，返回nil错误
	return nil
}
