package util

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"goravel/packages/goravel-socket/setting"
	"goravel/packages/goravel-socket/tools/crypto"
	"strings"
)

// GenUUID 生成uuid
func GenUUID() string {
	uuidFunc := uuid.NewV4()
	uuidStr := uuidFunc.String()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	uuidByt := []rune(uuidStr)
	return string(uuidByt[8:24])
}

// 对称加密IP和端口，当做clientId
func GenClientId() string {

	cryptoKey := []byte(setting.CommonSetting.CryptoKey)
	if len(cryptoKey) != 16 && len(cryptoKey) != 24 && len(cryptoKey) != 32 {
		logrus.Fatalf("Invalid AES key length: %d, must be 16, 24 or 32 bytes", len(cryptoKey))
	}

	raw := []byte(setting.GlobalSetting.LocalHost + ":" + setting.CommonSetting.RPCPort)
	str, err := crypto.Encrypt(raw, cryptoKey)
	if err != nil {
		logrus.WithError(err).Error("crypto.Encrypt failed")
		panic(err)
	}
	return str
}

// 解析redis的地址格式
func ParseRedisAddrValue(redisValue string) (host string, port string, err error) {
	if redisValue == "" {
		err = errors.New("解析地址错误")
		return
	}
	addr := strings.Split(redisValue, ":")
	if len(addr) != 2 {
		err = errors.New("解析地址错误")
		return
	}
	host, port = addr[0], addr[1]

	return
}

// 判断地址是否为本机
func IsAddrLocal(host string, port string) bool {
	return host == setting.GlobalSetting.LocalHost && port == setting.CommonSetting.RPCPort
}

// 是否集群
func IsCluster() bool {
	return setting.CommonSetting.Cluster
}

// 获取client key地址信息
func GetAddrInfoAndIsLocal(clientId string) (addr string, host string, port string, isLocal bool, err error) {
	//解密ClientId
	addr, err = crypto.Decrypt(clientId, []byte(setting.CommonSetting.CryptoKey))
	if err != nil {
		return
	}

	host, port, err = ParseRedisAddrValue(addr)
	if err != nil {
		return
	}

	isLocal = IsAddrLocal(host, port)
	return
}

func GenGroupKey(systemId, groupName string) string {
	return systemId + ":" + groupName
}
