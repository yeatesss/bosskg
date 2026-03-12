package bosskg

// ---------------------------------------------------------------------------
// 5.13 充值查询（6018）
// ---------------------------------------------------------------------------

// QueryRechargeReq 充值查询请求参数。
type QueryRechargeReq struct {
	ProviderID uint64 `json:"providerId"` // 服务商 ID
	StartDate  string `json:"startDate"`  // 开始日期（yyyy-MM-dd）
	EndDate    string `json:"endDate"`    // 结束日期（yyyy-MM-dd）
}

// QueryRechargeResp 充值查询响应参数（数组）。
type QueryRechargeResp []RechargeRecord

// RechargeRecord 充值记录明细。
type RechargeRecord struct {
	ProviderID      uint64 `json:"providerId,omitempty"`        // 服务商 ID
	ProviderName    string `json:"providerName,omitempty"`      // 服务商名称
	EnterpriseOrder string `json:"enterpriseOrderNo,omitempty"` // 企业订单号
	OrderNo         string `json:"orderNo,omitempty"`           // 平台订单号
	RechargeAmt     int64  `json:"rechargeAmt,omitempty"`       // 充值金额（单位：分）
	FeeAmt          int64  `json:"feeAmt,omitempty"`            // 手续费（单位：分）
	AccountingAmt   int64  `json:"accountingAmt,omitempty"`     // 入账金额（单位：分）
	RechargeState   string `json:"rechargeState,omitempty"`     // 充值状态：PROCESSING / SUCCESS / FAIL
	BankRemark      string `json:"bankRemark,omitempty"`        // 银行备注
	ReceiveBankNo   string `json:"receiveBankNo,omitempty"`     // 收款银行账号
	ReceiveBankName string `json:"receiveBankName,omitempty"`   // 收款银行名称
	PayBankName     string `json:"payBankName,omitempty"`       // 付款银行名称
	PayBankNo       string `json:"payBankNo,omitempty"`         // 付款银行账号
	CreateTime      string `json:"createTime,omitempty"`        // 创建时间
	UpdateTime      string `json:"updateTime,omitempty"`        // 更新时间
	ErrMsg          string `json:"errMsg,omitempty"`            // 错误信息
}

// ---------------------------------------------------------------------------
// 5.14 查询可充值金额（6019）
// ---------------------------------------------------------------------------

// QueryAvailableRechargeReq 查询可充值金额请求参数。
type QueryAvailableRechargeReq struct {
	SubAccNo string `json:"subAccNo"` // 银行电子账户号
}

// QueryAvailableRechargeResp 查询可充值金额响应参数。
type QueryAvailableRechargeResp struct {
	AvailableRechargeAmt int64  `json:"availableRechargeAmt,omitempty"` // 可充值金额（单位：分）
	SubAccAmt            string `json:"subAccAmt,omitempty"`            // 子账户金额（单位：分，文档为 string 类型）
	SubAccNo             string `json:"subAccNo,omitempty"`             // 银行电子账户号
}

// ---------------------------------------------------------------------------
// 5.16 申请充值（6020）
// ---------------------------------------------------------------------------

// ApplyRechargeReq 申请充值请求参数。
type ApplyRechargeReq struct {
	ProviderID      uint64 `json:"providerId"`          // 服务商 ID
	SubAccNo        string `json:"subAccNo"`            // 银行电子账户号
	RechargeAmt     uint64 `json:"rechargeAmt"`         // 充值金额（单位：分）
	EnterpriseOrder string `json:"enterpriseOrderNo"`   // 企业订单号
	NotifyURL       string `json:"notifyUrl,omitempty"` // 异步通知地址
}

// ApplyRechargeResp 申请充值响应参数。
type ApplyRechargeResp struct {
	ProviderID      uint64 `json:"providerId,omitempty"`        // 服务商 ID
	EnterpriseOrder string `json:"enterpriseOrderNo,omitempty"` // 企业订单号
	OrderNo         string `json:"orderNo,omitempty"`           // 平台订单号
	RechargeAmt     string `json:"rechargeAmt,omitempty"`       // 充值金额（单位：分，文档为 string 类型）
	FeeAmt          string `json:"feeAmt,omitempty"`            // 手续费（单位：分）
	AccountingAmt   int64  `json:"accountingAmt,omitempty"`     // 入账金额（单位：分）
	RechargeState   string `json:"rechargeState,omitempty"`     // 充值状态
	CreateTime      string `json:"createTime,omitempty"`        // 创建时间
	UpdateTime      string `json:"updateTime,omitempty"`        // 更新时间
	ErrMsg          string `json:"errMsg,omitempty"`            // 错误信息
}

// ---------------------------------------------------------------------------
// 5.17 充值结果查询（6021）
// ---------------------------------------------------------------------------

// QueryRechargeResultReq 充值结果查询请求参数。
type QueryRechargeResultReq struct {
	EnterpriseOrder string `json:"enterpriseOrderNo,omitempty"` // 企业订单号（可选）
	OrderNo         string `json:"orderNo,omitempty"`           // 平台订单号（可选）
}

// QueryRechargeResultResp 充值结果查询响应参数。
type QueryRechargeResultResp struct {
	RechargeRecordList []RechargeResultRecord `json:"rechargeRecordList,omitempty"` // 充值记录列表
}

// RechargeResultRecord 充值结果明细。
type RechargeResultRecord struct {
	ProviderID      uint64 `json:"providerId,omitempty"`        // 服务商 ID
	ProviderName    string `json:"providerName,omitempty"`      // 服务商名称
	RechargeState   string `json:"rechargeState,omitempty"`     // 充值状态：PROCESSING / SUCCESS / FAIL
	AccountingAmt   int64  `json:"accountingAmt,omitempty"`     // 入账金额（单位：分）
	RechargeAmt     int64  `json:"rechargeAmt,omitempty"`       // 充值金额（单位：分）
	FeeAmt          int64  `json:"feeAmt,omitempty"`            // 手续费（单位：分）
	EnterpriseOrder string `json:"enterpriseOrderNo,omitempty"` // 企业订单号
	OrderNo         string `json:"orderNo,omitempty"`           // 平台订单号
	CreateTime      string `json:"createTime,omitempty"`        // 创建时间
	UpdateTime      string `json:"updateTime,omitempty"`        // 更新时间
	ErrorMsg        string `json:"errorMsg,omitempty"`          // 错误信息
}
