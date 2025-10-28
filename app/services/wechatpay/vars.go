package wechatpay

// Amount结构体用于解析amount字段
type Amount struct {
	Total            int    `json:"total"`
	Refund           int    `json:"refund"`
	From             []From `json:"from"`
	PayerTotal       int    `json:"payer_total"`
	PayerRefund      int    `json:"payer_refund"`
	SettlementRefund int    `json:"settlement_refund"`
	SettlementTotal  int    `json:"settlement_total"`
	DiscountRefund   int    `json:"discount_refund"`
	Currency         string `json:"currency"`
	RefundFee        int    `json:"refund_fee"`
}

// From结构体用于解析amount.from字段中的元素
type From struct {
	Account string `json:"account"`
	Amount  int    `json:"amount"`
}

// GoodsDetail结构体用于解析promotion_detail.goods_detail字段中的元素
type GoodsDetail struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`
	WechatpayGoodsId string `json:"wechatpay_goods_id"`
	GoodsName        string `json:"goods_name"`
	UnitPrice        int    `json:"unit_price"`
	RefundAmount     int    `json:"refund_amount"`
	RefundQuantity   int    `json:"refund_quantity"`
}

// PromotionDetail结构体用于解析promotion_detail字段
type PromotionDetail struct {
	PromotionId  string        `json:"promotion_id"`
	Scope        string        `json:"scope"`
	Type         string        `json:"type"`
	Amount       int           `json:"amount"`
	RefundAmount int           `json:"refund_amount"`
	GoodsDetail  []GoodsDetail `json:"goods_detail"`
}

// Refund结构体用于解析整个退款数据结构
type Refund struct {
	RefundId            string            `json:"refund_id"`
	OutRefundNo         string            `json:"out_refund_no"`
	TransactionId       string            `json:"transaction_id"`
	OutTradeNo          string            `json:"out_trade_no"`
	Channel             string            `json:"channel"`
	UserReceivedAccount string            `json:"user_received_account"`
	SuccessTime         string            `json:"success_time"`
	CreateTime          string            `json:"create_time"`
	Status              string            `json:"status"`
	FundsAccount        string            `json:"funds_account"`
	Amount              Amount            `json:"amount"`
	PromotionDetail     []PromotionDetail `json:"promotion_detail"`
}
