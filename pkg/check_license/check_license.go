package check_license

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"mark3/pkg/common"
	"net"
	"os"
	"time"
)

type License struct {
	Username   string    `json:"username"`            // 用户名
	Secret     string    `json:"secret"`              // 密钥
	Expiration time.Time `json:"expiration"`          // 到期时间
	Signature  []byte    `json:"signature,omitempty"` // 签名
}

func GetMac() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v\n", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}

var (
	Username = ""
	Secret   = "secret-key-123"
)

// 验证许可证签名
func verifyLicense(license *License, publicKey *rsa.PublicKey) error {
	licenseCopy := *license
	licenseCopy.Signature = nil

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(licenseCopy)
	if err != nil {
		return err
	}
	licenseData := buffer.Bytes()

	hash := sha256.Sum256(licenseData)
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], license.Signature)
}

func CheckLicense() {
	// 读取公钥
	publicKeyData, err := os.ReadFile("./pkg/check_license/public.pem")
	if err != nil {
		panic(err)
	}
	publicBlock, _ := pem.Decode(publicKeyData)
	publicKey, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		panic(err)
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		panic("不是 RSA 公钥")
	}

	// 读取二进制许可证
	licenseData, err := os.ReadFile("./pkg/check_license/license.lic")
	if err != nil {
		fmt.Println("未找到许可证文件")
		panic(err)
	}

	var license License
	buffer := bytes.NewBuffer(licenseData)
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(&license)
	if err != nil {
		fmt.Println("解析许可证文件失败")
		panic(err)
	}

	// 验证签名
	err = verifyLicense(&license, rsaPublicKey)
	if err != nil {
		fmt.Println("许可证验证失败:", err)
		panic(err)
	}

	// 检查过期时间 这里是使用本地时间校验的，也可以通过网络时间校验
	if time.Now().After(license.Expiration) {
		fmt.Println("许可证已过期")
		panic(err)
	}

	// 检查用户名
	Usernames := GetMac()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if !(common.InArray(license.Username, Usernames) && license.Secret == Secret) {
		fmt.Println("许可证无效")
		panic(err)
	}

	fmt.Printf("许可证有效，程序可以运行，有效期至%s\n", license.Expiration.Format("2006-01-02 15:04:05"))
	// 你的应用程序逻辑在这里
}
