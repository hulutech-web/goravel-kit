package wechatpay

import (
	"context"
	"errors"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/path"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// MerchantConfig 商户配置信息
type MerchantConfig struct {
	MchID                string
	MchCertificateSerial string
	MchAPIv3Key          string
	AppID                string
	NotifyURL            string
	RefundURL            string
}

// WechatpayService 封装微信支付服务的结构体
type WechatpayService struct {
	client *core.Client
	config *MerchantConfig
}
type PrepayWithRequestPaymentResponse struct {
	// 预支付交易会话标识
	PrepayId *string `json:"prepay_id"` // revive:disable-line:var-naming
	// 应用ID
	Appid *string `json:"appId"`
	// 时间戳
	TimeStamp *string `json:"timeStamp"`
	// 随机字符串
	NonceStr *string `json:"nonceStr"`
	// 订单详情扩展字符串
	Package *string `json:"package"`
	// 签名方式
	SignType *string `json:"signType"`
	// 签名
	PaySign *string `json:"paySign"`
}

// 创建一个新的微信支付服务实例
func NewWechatpayService(config *MerchantConfig) (*WechatpayService, error) {
	// 加载商户私钥
	mchPrivateKey, err1 := utils.LoadPrivateKeyWithPath(path.App("/keys/apiclient_key.pem"))
	if err1 != nil {
		return nil, err1
	}
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(
			config.MchID,
			config.MchCertificateSerial,
			mchPrivateKey,
			config.MchAPIv3Key,
		),
	}

	ctx := context.Background()
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &WechatpayService{
		client: client,
		config: config,
	}, nil
}

// PrepareJSAPITransaction 准备JSAPI交易并获取预支付ID
func (s *WechatpayService) PrepareJSAPITransaction(ctx context.Context, description string, outTradeNo string, totalFee int64, openid string) (*PrepayWithRequestPaymentResponse, error) {
	prepayRequest := jsapi.PrepayRequest{
		Appid:       core.String(s.config.AppID),
		Mchid:       core.String(s.config.MchID),
		Description: core.String(description),
		OutTradeNo:  core.String(outTradeNo),
		Amount: &jsapi.Amount{
			Total: core.Int64(totalFee),
		},
		Payer: &jsapi.Payer{
			Openid: core.String(openid),
		},
		NotifyUrl: core.String(s.config.NotifyURL),
		//支持分账的订单
		SettleInfo: &jsapi.SettleInfo{
			ProfitSharing: core.Bool(true),
		},
	}
	// 调用预支付接口
	svc := jsapi.JsapiApiService{Client: s.client}
	prepayResp, _, err := svc.PrepayWithRequestPayment(ctx, prepayRequest)
	if err != nil {
		return nil, err
	}
	var response PrepayWithRequestPaymentResponse
	response.Appid = prepayResp.Appid
	response.TimeStamp = prepayResp.TimeStamp
	response.NonceStr = prepayResp.NonceStr
	response.Package = prepayResp.Package
	response.SignType = prepayResp.SignType
	response.PaySign = prepayResp.PaySign

	return &response, nil
}

// ===============以下是分账功能==================
// CreateProfitShare 发起分账请求
func (w *WechatpayService) CreateProfitShare(outTradeNo string, totalFee float64, transaction_id string) error {
	facades.Log().Info("============CreateProfitShare============\n")
	mch_id := w.config.MchID
	mchCertificateSerialNumber := w.config.MchCertificateSerial
	mch_secret_key := w.config.MchAPIv3Key
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(path.App("/keys/apiclient_key.pem"))
	if err != nil {
		facades.Log().Info("============load merchant private key error============\n")
	}

	haibo_mch_id := facades.Config().GetString("haibo.mch_id")
	if haibo_mch_id == "" {
		facades.Log().Info("============load haibo_mch_id merchant private key error============\n")
	}
	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mch_id, mchCertificateSerialNumber, mchPrivateKey, mch_secret_key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		facades.Log().Info("new wechat pay client err:%s", err.Error())
	}
	svc := profitsharing.OrdersApiService{Client: client}
	resp, result, err := svc.CreateOrder(ctx,
		profitsharing.CreateOrderRequest{
			Appid:      core.String(facades.Config().GetString("mini.app_id")),
			OutOrderNo: core.String(outTradeNo),
			Receivers: []profitsharing.CreateOrderReceiver{profitsharing.CreateOrderReceiver{
				Account:     core.String(haibo_mch_id),
				Amount:      core.Int64(1), //这里先写一个固定的数据1元，支付比例最大10%，应该支付至少10元，下单为15元测试
				Description: core.String("分给商户:海博科技"),
				Type:        core.String("MERCHANT_ID"),
			}},
			TransactionId:   core.String(transaction_id),
			UnfreezeUnsplit: core.Bool(true),
		},
	)
	if err != nil {
		// 处理错误
		facades.Log().Info("============CreateProfitShare============call CreateOrder err\n", err.Error())
	} else {
		// 处理返回结果
		facades.Log().Info("============CreateProfitShare============成功\n", result.Response.StatusCode, resp)
	}
	return nil
}
