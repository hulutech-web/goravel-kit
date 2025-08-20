package wechatpay

import (
	"context"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// WechatPayService 微信支付服务
type WechatPayService struct {
	client         *core.Client
	merchantID     string
	merchantSerial string
	privateKeyPath string
	apiV3Key       string
}

// NewWechatPayService 创建微信支付服务实例
func NewWechatPayService(
	merchantID, merchantSerial, privateKeyPath, apiV3Key string,
) (*WechatPayService, error) {
	// 加载商户私钥
	privateKey, err := utils.LoadPrivateKeyWithPath(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("加载商户私钥失败: %w", err)
	}

	// 创建请求选项
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(merchantID, merchantSerial, privateKey, apiV3Key),
	}

	// 创建微信支付API客户端
	client, err := core.NewClient(context.Background(), opts...)
	if err != nil {
		return nil, fmt.Errorf("创建微信支付客户端失败: %w", err)
	}

	return &WechatPayService{
		client:         client,
		merchantID:     merchantID,
		merchantSerial: merchantSerial,
		privateKeyPath: privateKeyPath,
		apiV3Key:       apiV3Key,
	}, nil
}
