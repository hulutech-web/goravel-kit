package wechatpay

import (
	"context"
	"errors"
	"fmt"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/path"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cast"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"goravel/app/models"
	"strings"
)

type RefundConfig struct {
	MchID                string
	MchCertificateSerial string
	MchAPIv3Key          string
	AppID                string
	NotifyURL            string
	RefundURL            string
}

type WechatRefundService struct {
	Handler *notify.Handler
	config  *RefundConfig
}

// NewWechatpayService 创建一个新的微信支付服务实例
func NewWechatRefundService(config *RefundConfig) (*WechatRefundService, error) {
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
		config.MchCertificateSerial,
		config.MchID,
		config.MchAPIv3Key,
	); err1 != nil {
		return nil, fmt.Errorf("注册下载器失败: %w", err1)
	}

	// 3. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(config.MchID)

	// 4. 初始化 Notify Handler
	handler := notify.NewNotifyHandler(config.MchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))

	svc := &WechatRefundService{
		Handler: handler,
		config:  config,
		// 如果后续有需要，也可以在这里初始化带有自动下载证书功能的 client
		// client: initializeWechatPayClient(ctx, config),
	}
	return svc, nil
}

func (s *WechatRefundService) CreateRefund(order models.CouponOrder) error {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(path.App("/keys/apiclient_key.pem"))

	if err != nil {
		facades.Log().Info("load private key error:" + err.Error())
		return err
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(s.config.MchID, s.config.MchCertificateSerial, mchPrivateKey, s.config.MchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		facades.Log().Info("NewClient error:" + err.Error())
		return errors.New(err.Error())
	}
	//必须转换一下
	refundAmount := cast.ToFloat64(order.PayAmount)
	refundAmount = refundAmount * 100
	refundInt64 := cast.ToInt64(refundAmount)
	svc := refunddomestic.RefundsApiService{Client: client}
	_, _, err1 := svc.Create(ctx,
		refunddomestic.CreateRequest{
			OutTradeNo:  core.String(order.OutTradeNo),
			OutRefundNo: core.String(s.MakeRefundNo()),
			Reason:      core.String("客户要求退货"),
			NotifyUrl:   core.String(s.config.RefundURL),
			Amount: &refunddomestic.AmountReq{
				Currency: core.String("CNY"),
				Refund:   core.Int64(refundInt64),
				Total:    core.Int64(refundInt64),
			},
		},
	)

	if err1 != nil {
		// 处理错误
		facades.Log().Info("call Create err:" + err1.Error())
		return errors.New("call Create err:" + err1.Error())
	} else {
		// 处理返回结果
		return nil
	}
}

func (s *WechatRefundService) MakeRefundNo() string {
	//	32位，不重复，uuid
	uuidV4 := uuid.NewV4()
	// 将 UUID 转换为字符串
	uuidStr := uuidV4.String()
	// 如果需要去掉 UUID 中的 "-"
	uuidWithoutHyphens := uuidStr[:8] + uuidStr[9:13] + uuidStr[14:18] + uuidStr[19:23] + uuidStr[24:]
	uuidWithoutHyphens = strings.ToUpper(uuidWithoutHyphens)
	return uuidWithoutHyphens
}
