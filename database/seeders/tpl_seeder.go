package seeders

import (
	goravelfacades "github.com/goravel/framework/facades"
	"goravel/packages/goravel_pdf_gen/models"
)

type TplSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *TplSeeder) Signature() string {
	return "TplSeeder"
}

// Run executes the seeder logic.
func (s *TplSeeder) Run() error {
	htmlTpl := "<html>\n<head>\n  <title>[[.contractTitle]]</title>\n  <style>\n    .contract-container { \n      font-family: 'Times New Roman', Times, serif; \n      line-height: 1.8;\n      color: #333;\n      max-width: 800px;\n      margin: 0 auto;\n      padding: 40px;\n      background: #fff;\n    }\n    .header {\n      text-align: center;\n      margin-bottom: 40px;\n    }\n    h1 {\n      font-size: 28px;\n      color: #1a3a6c;\n      margin-bottom: 10px;\n      border-bottom: 2px solid #1a3a6c;\n      padding-bottom: 15px;\n    }\n    .contract-info {\n      display: flex;\n      justify-content: space-between;\n      margin-bottom: 30px;\n      font-size: 16px;\n    }\n    .parties {\n      display: flex;\n      justify-content: space-between;\n      margin: 30px 0;\n    }\n    .party {\n      width: 45%;\n    }\n    .party h2 {\n      font-size: 20px;\n      border-bottom: 1px solid #ccc;\n      padding-bottom: 8px;\n      margin-bottom: 15px;\n    }\n    .terms-section {\n      margin: 30px 0;\n    }\n    .terms-section h2 {\n      font-size: 20px;\n      color: #1a3a6c;\n      border-left: 4px solid #1a3a6c;\n      padding-left: 10px;\n      margin: 25px 0 15px;\n    }\n    .signature-area {\n      display: flex;\n      justify-content: space-between;\n      margin-top: 80px;\n    }\n    .signature-box {\n      width: 45%;\n      text-align: center;\n    }\n    .signature-line {\n      height: 1px;\n      background: #333;\n      margin: 40px 0 10px;\n    }\n    .variable {\n      background: #fff8e1;\n      padding: 2px 5px;\n      border-radius: 4px;\n      border: 1px dashed #ffc107;\n    }\n  </style>\n</head>\n<body class=\"contract-container\">\n  <div class=\"header\">\n    <h1>[[.contractTitle]]</h1>\n    <p>合同编号：<span class=\"variable\">[[.contractNo]]</span></p>\n  </div>\n  \n  <div class=\"contract-info\">\n    <div>签订日期：<span class=\"variable\">[[.signDate]]</span></div>\n    <div>生效日期：<span class=\"variable\">[[.signDate]]</span></div>\n  </div>\n  \n  <div class=\"parties\">\n    <div class=\"party\">\n      <h2>甲方（卖方）</h2>\n      <p>名称：<span class=\"variable\">[[.sellerName]]</span></p>\n    </div>\n    \n    <div class=\"party\">\n      <h2>乙方（买方）</h2>\n      <p>名称：<span class=\"variable\">[[.buyerName]]</span></p>\n    </div>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第一条 产品信息</h2>\n    <p>1.1 产品名称：<span class=\"variable\">[[.productName]]</span></p>\n    <p>1.2 产品数量：<span class=\"variable\">[[.quantity]]</span></p>\n    <p>1.3 产品单价：人民币 <span class=\"variable\">[[.unitPrice]]</span> 元</p>\n    <p>1.4 总金额：人民币 <span class=\"variable\">[[.totalAmount]]</span> 元</p>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第二条 交货条款</h2>\n    <p>2.1 交货日期：<span class=\"variable\">[[.deliveryDate]]</span></p>\n    <p>2.2 交货地点：买方指定地点</p>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第三条 付款条款</h2>\n    <p>3.1 付款方式：<span class=\"variable\">[[.paymentTerms]]</span></p>\n    <p>3.2 付款期限：自合同签订之日起30日内</p>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第四条 保修条款</h2>\n    <p>4.1 保修期限：<span class=\"variable\">[[.warrantyPeriod]]</span></p>\n    <p>4.2 保修范围：产品制造缺陷</p>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第五条 特别条款</h2>\n    <p><span class=\"variable\">[[.specialTerms]]</span></p>\n  </div>\n  \n  <div class=\"signature-area\">\n    <div class=\"signature-box\">\n      <p>甲方（卖方）签字：</p>\n      <div class=\"signature-line\"></div>\n      <p>日期：<span class=\"variable\">[[.signDate]]</span></p>\n    </div>\n    \n    <div class=\"signature-box\">\n      <p>乙方（买方）签字：</p>\n      <div class=\"signature-line\"></div>\n      <p>日期：<span class=\"variable\">[[.signDate]]</span></p>\n    </div>\n  </div>\n</body>\n</html>"
	json := `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "销售合同",
  "description": "销售合同表单数据",
  "type": "object",
  "properties": {
    "contractTitle": {
      "type": "string",
      "title": "合同标题",
      "default": "产品销售合同"
    },
    "contractNo": {
      "type": "string",
      "title": "合同编号",
      "pattern": "^CT-[0-9]{8}$"
    },
    "signDate": {
      "type": "string",
      "title": "签署日期",
      "format": "date"
    },
    "sellerName": {
      "type": "string",
      "title": "卖方名称"
    },
    "buyerName": {
      "type": "string",
      "title": "买方名称"
    },
    "productName": {
      "type": "string",
      "title": "产品名称"
    },
    "quantity": {
      "type": "number",
      "title": "数量",
      "minimum": 1
    },
    "unitPrice": {
      "type": "number",
      "title": "单价",
      "minimum": 0
    },
    "totalAmount": {
      "type": "number",
      "title": "总金额",
      "readOnly": true
    },
    "deliveryDate": {
      "type": "string",
      "title": "交货日期",
      "format": "date"
    },
    "paymentTerms": {
      "type": "string",
      "title": "付款条款",
      "enum": [
        "预付全款",
        "货到付款",
        "30天账期",
        "60天账期"
      ]
    },
    "warrantyPeriod": {
      "type": "string",
      "title": "保修期",
      "enum": [
        "6个月",
        "1年",
        "2年",
        "3年",
        "5年"
      ]
    },
    "specialTerms": {
      "type": "string",
      "title": "特别条款",
      "format": "textarea",
      "rows": 4
    }
  },
  "required": [
    "contractTitle",
    "contractNo",
    "sellerName",
    "buyerName",
    "productName",
    "quantity",
    "unitPrice"
  ]
}`

	pdfGen2 := models.PdfGen{
		Name:   "pdf_gen2",
		Html:   htmlTpl,
		Params: models.JsonMap(json),
	}
	goravelfacades.Orm().Query().Create(&pdfGen2)
	return nil
}
