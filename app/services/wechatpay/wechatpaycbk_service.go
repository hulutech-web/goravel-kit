package wechatpay

import (
	"context"
	"fmt"
	"github.com/goravel/framework/support/path"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type CallbackConfig struct {
	MchCertificateSerialNumber string
	MchID                      string
	MchAPIV3Key                string
}

type WechatpaycbkService struct {
	Handler *notify.Handler
	config  *CallbackConfig
}

func NewWechatpaycbkService(config *CallbackConfig) (*WechatpaycbkService, error) {
	// 1. 加载商户私钥
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(path.App("/keys/apiclient_key.pem"))
	if err != nil {
		return nil, fmt.Errorf("加载商户私钥失败: %w", err)
	}

	ctx := context.Background()

	// 2. 注册下载器并获取微信支付平台证书访问器
	if err1 := downloader.MgrInstance().RegisterDownloaderWithPrivateKey(
		ctx,
		mchPrivateKey,
		config.MchCertificateSerialNumber,
		config.MchID,
		config.MchAPIV3Key,
	); err1 != nil {
		return nil, fmt.Errorf("注册下载器失败: %w", err1)
	}

	// 3. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(config.MchID)

	// 4. 初始化 Notify Handler
	handler := notify.NewNotifyHandler(config.MchAPIV3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))

	svc := &WechatpaycbkService{
		Handler: handler,
		config:  config,
		// 如果后续有需要，也可以在这里初始化带有自动下载证书功能的 client
		// client: initializeWechatPayClient(ctx, config),
	}
	return svc, nil
}
