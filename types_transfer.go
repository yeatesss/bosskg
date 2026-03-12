package bosskg

// ---------------------------------------------------------------------------
// 付款共用类型
// ---------------------------------------------------------------------------

// PayItem 批量付款明细（6001/6022/6029 共用）。
type PayItem struct {
	MerOrderID  string `json:"merOrderId"`          // 商户订单号（商户维度唯一，最长 32 位）
	Amt         uint64 `json:"amt"`                 // 付款金额（单位：分，最低 1000 即 10 元，最高 9800000 即 98000 元）
	PayeeName   string `json:"payeeName"`           // 收款人姓名（最长 50 位）
	PayeeAcc    string `json:"payeeAcc"`            // 收款账号：银行卡号 / 支付宝账号 / 微信 OpenID（最长 28 位）
	IDCard      string `json:"idCard"`              // 身份证号（18 位）
	Mobile      string `json:"mobile"`              // 收款人手机号（11 位）
	Memo        string `json:"memo,omitempty"`      // 备注（最长 20 位，敏感词限制）
	PaymentType uint64 `json:"paymentType"`         // 付款方式：0=银行卡 1=支付宝 2=微信
	NotifyURL   string `json:"notifyUrl,omitempty"` // 异步通知地址（不填则不通知，最长 100 位）
}

// PayResult 付款返回明细（6001/6029 共用）。
type PayResult struct {
	MerOrderID  string `json:"merOrderId,omitempty"`  // 商户订单号
	OrderNo     string `json:"orderNo,omitempty"`     // 平台订单号（全局唯一，最长 25 位）
	Amt         uint64 `json:"amt,omitempty"`         // 付款金额（单位：分）
	Fee         uint64 `json:"fee,omitempty"`         // 服务费（单位：分）
	PackageInfo string `json:"packageInfo,omitempty"` // 微信零钱新模式：跳转收款页 package 信息
	MchID       string `json:"mchId,omitempty"`       // 微信零钱新模式：跳转收款页 mchId 信息
	ResCode     string `json:"resCode,omitempty"`     // 受理响应码（非交易终态）
	ResMsg      string `json:"resMsg,omitempty"`      // 受理响应信息
}

// TransferQueryItem 付款查询条件（6002/6023/6030 共用）。
type TransferQueryItem struct {
	MerOrderID string `json:"merOrderId,omitempty"` // 商户订单号（最长 32 位）
	OrderNo    string `json:"orderNo,omitempty"`    // 平台订单号（最长 25 位）
}

// TransferResultItem 商户批量付款查询结果明细（6002/6030 共用）。
type TransferResultItem struct {
	MerOrderID   string `json:"merOrderId,omitempty"`   // 商户订单号
	OrderNo      string `json:"orderNo,omitempty"`      // 平台订单号（最长 25 位）
	State        uint64 `json:"state,omitempty"`        // 交易状态：1=付款中 3=成功 4=失败 6=待用户确认 7=已取消
	Amt          uint64 `json:"amt,omitempty"`          // 付款金额（单位：分）
	Fee          uint64 `json:"fee,omitempty"`          // 平台管理费（单位：分）
	UserFee      uint64 `json:"userFee,omitempty"`      // 个人服务费/个税（单位：分）
	VATax        uint64 `json:"vaTax,omitempty"`        // 个人增值税（单位：分）
	VAAddTax     uint64 `json:"vaAddTax,omitempty"`     // 个人增值税附加（单位：分）
	UserDueAmt   uint64 `json:"userDueAmt,omitempty"`   // 个人实际到账金额（单位：分）
	UserFeeRatio string `json:"userFeeRatio,omitempty"` // 个人服务费率/个税税率
	ResMsg       string `json:"resMsg,omitempty"`       // 响应信息
	CreateTime   string `json:"createTime,omitempty"`   // 创建时间（yyyy-MM-dd HH:mm:ss）
	EndTime      string `json:"endTime,omitempty"`      // 完成时间（yyyy-MM-dd HH:mm:ss）
}

// ---------------------------------------------------------------------------
// 5.3 商户批量付款（6001）
// ---------------------------------------------------------------------------

// BatchTransferReq 商户批量付款请求参数。
type BatchTransferReq struct {
	MerBatchID string    `json:"merBatchId"` // 商户批次号（建议 yyyymmddHHSS + 8 位随机数，商户维度唯一）
	PayItems   []PayItem `json:"payItems"`   // 付款明细列表
	TaskID     uint64    `json:"taskId"`     // 任务 ID（商户平台获取）
	ProviderID uint64    `json:"providerId"` // 服务商 ID
}

// BatchTransferResp 商户批量付款响应参数。
type BatchTransferResp struct {
	SuccessNum    uint64      `json:"successNum,omitempty"`    // 该批次订单受理成功笔数
	FailureNum    uint64      `json:"failureNum,omitempty"`    // 该批次订单受理失败笔数
	MerBatchID    string      `json:"merBatchId,omitempty"`    // 商户批次号
	PayResultList []PayResult `json:"payResultList,omitempty"` // 付款返回数据列表
}

// ---------------------------------------------------------------------------
// 5.4 商户批量付款查询（6002）
// ---------------------------------------------------------------------------

// QueryBatchTransferResultReq 商户批量付款查询请求参数。
type QueryBatchTransferResultReq struct {
	MerBatchID string              `json:"merBatchId"`           // 商户批次号（最长 32 位）
	QueryItems []TransferQueryItem `json:"queryItems,omitempty"` // 查询条件（为空则返回该批次全部订单）
}

// QueryBatchTransferResultResp 商户批量付款查询响应参数。
type QueryBatchTransferResultResp struct {
	MerID      string               `json:"merId,omitempty"`      // 商户号
	MerBatchID string               `json:"merBatchId,omitempty"` // 商户批次号
	QueryItems []TransferResultItem `json:"queryItems,omitempty"` // 查询结果列表
}

// ---------------------------------------------------------------------------
// 5.20 批量上传（6022）
// ---------------------------------------------------------------------------

// BatchUploadReq 批量上传请求参数。
type BatchUploadReq struct {
	MerBatchID string    `json:"merBatchId"` // 商户批次号（最长 32 位，商户维度唯一）
	PayItems   []PayItem `json:"payItems"`   // 付款明细列表（复用 PayItem）
	TaskID     uint64    `json:"taskId"`     // 任务 ID
	ProviderID uint64    `json:"providerId"` // 服务商 ID
}

// BatchUploadResp 批量上传响应参数。
type BatchUploadResp struct {
	BatchNo uint64 `json:"batchNo,omitempty"` // 批次号
	ResCode string `json:"resCode,omitempty"` // 响应码
	ResMsg  string `json:"resMsg,omitempty"`  // 响应信息
}

// ---------------------------------------------------------------------------
// 5.21 批量订单查询（6023）
// ---------------------------------------------------------------------------

// QueryBatchOrderReq 批量订单查询请求参数。
type QueryBatchOrderReq struct {
	MerID      string              `json:"merId"`                // 商户号
	MerBatchID string              `json:"merBatchId"`           // 商户批次号（最长 32 位）
	QueryItems []TransferQueryItem `json:"queryItems,omitempty"` // 查询条件（复用 TransferQueryItem）
}

// QueryBatchOrderResp 批量订单查询响应参数。
type QueryBatchOrderResp struct {
	MerID      string             `json:"merId,omitempty"`      // 商户号
	MerBatchID string             `json:"merBatchId,omitempty"` // 商户批次号
	QueryItems []BatchOrderResult `json:"queryItems,omitempty"` // 查询结果列表
}

// BatchOrderResult 批量订单查询结果明细。
type BatchOrderResult struct {
	MerOrderID   string `json:"merOrderId,omitempty"`   // 商户订单号
	OrderNo      string `json:"orderNo,omitempty"`      // 平台订单号（最长 25 位）
	State        uint64 `json:"state,omitempty"`        // 交易状态：0=待处理 1=付款中 3=成功 4=失败 7=已取消
	Amt          uint64 `json:"amt,omitempty"`          // 付款金额（单位：分）
	Fee          uint64 `json:"fee,omitempty"`          // 平台管理费（单位：分）
	UserFee      uint64 `json:"userFee,omitempty"`      // 个人服务费/个税（单位：分）
	VATax        uint64 `json:"vaTax,omitempty"`        // 个人增值税（单位：分）
	VAAddTax     uint64 `json:"vaAddTax,omitempty"`     // 个人增值税附加（单位：分）
	UserDueAmt   uint64 `json:"userDueAmt,omitempty"`   // 个人实际到账金额（单位：分）
	UserFeeRatio string `json:"userFeeRatio,omitempty"` // 个人服务费率/个税税率
	ResMsg       string `json:"resMsg,omitempty"`       // 响应信息
	CreateTime   string `json:"createTime,omitempty"`   // 创建时间（yyyy-MM-dd HH:mm:ss）
	EndTime      string `json:"endTime,omitempty"`      // 完成时间（yyyy-MM-dd HH:mm:ss）
}

// ---------------------------------------------------------------------------
// 5.22 回单查询（6024）
// ---------------------------------------------------------------------------

// QueryReceiptReq 回单查询请求参数。
type QueryReceiptReq struct {
	MerBatchID string `json:"merBatchId,omitempty"` // 商户批次号（最长 32 位，可选）
	MerOrderID string `json:"merOrderId"`           // 商户订单号（最长 32 位）
}

// QueryReceiptResp 回单查询响应参数。
type QueryReceiptResp struct {
	MerID      string `json:"merId,omitempty"`      // 商户号
	MerBatchID string `json:"merBatchId,omitempty"` // 商户批次号
	MerOrderID string `json:"merOrderId,omitempty"` // 商户订单号
	ReceiptURL string `json:"receiptUrl,omitempty"` // 回单下载地址（最长 100 位）
}

// ---------------------------------------------------------------------------
// 5.24 一键发放（6029）
// ---------------------------------------------------------------------------

// AllInOneReleaseReq 一键发放请求参数。
type AllInOneReleaseReq struct {
	MerBatchID string    `json:"merBatchId"` // 商户批次号（最长 32 位，商户维度唯一）
	PayItems   []PayItem `json:"payItems"`   // 付款明细列表（复用 PayItem）
	TaskID     uint64    `json:"taskId"`     // 任务 ID
	ProviderID uint64    `json:"providerId"` // 服务商 ID
}

// AllInOneReleaseResp 一键发放响应参数。
type AllInOneReleaseResp struct {
	SuccessNum    uint64      `json:"successNum,omitempty"`    // 受理成功笔数
	FailureNum    uint64      `json:"failureNum,omitempty"`    // 受理失败笔数
	MerBatchID    string      `json:"merBatchId,omitempty"`    // 商户批次号
	PayResultList []PayResult `json:"payResultList,omitempty"` // 付款返回数据列表（复用 PayResult）
}

// ---------------------------------------------------------------------------
// 5.25 一键发放查询（6030）
// ---------------------------------------------------------------------------

// QueryAllInOneResultReq 一键发放查询请求参数。
type QueryAllInOneResultReq struct {
	MerBatchID string              `json:"merBatchId"`           // 商户批次号（最长 32 位）
	QueryItems []TransferQueryItem `json:"queryItems,omitempty"` // 查询条件（复用 TransferQueryItem）
}

// QueryAllInOneResultResp 一键发放查询响应参数。
type QueryAllInOneResultResp struct {
	MerID      string               `json:"merId,omitempty"`      // 商户号
	MerBatchID string               `json:"merBatchId,omitempty"` // 商户批次号
	QueryItems []TransferResultItem `json:"queryItems,omitempty"` // 查询结果列表（复用 TransferResultItem）
}

// ---------------------------------------------------------------------------
// 5.27 微信取消转账（6043）
// ---------------------------------------------------------------------------

// WeChatCancelTransferReq 微信取消转账请求参数（仅限微信零钱新模式，处理状态 6=待用户确认的订单）。
type WeChatCancelTransferReq struct {
	MerBatchID string `json:"merBatchId,omitempty"` // 商户批次号（最长 32 位，可选）
	MerOrderID string `json:"merOrderId,omitempty"` // 商户订单号（最长 32 位，可选）
	OrderNo    string `json:"orderNo,omitempty"`    // 平台订单号（最长 25 位，可选）
}

// WeChatCancelTransferResp 微信取消转账响应参数。
type WeChatCancelTransferResp struct {
	MerBatchID string `json:"merBatchId,omitempty"` // 商户批次号
	MerOrderID string `json:"merOrderId,omitempty"` // 商户订单号
	OrderNo    string `json:"orderNo,omitempty"`    // 平台订单号
}
