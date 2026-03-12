package bosskg

// ---------------------------------------------------------------------------
// 5.5 商户账户余额查询（6003）
// ---------------------------------------------------------------------------

// QueryMerIDBalanceReq 商户账户余额查询请求参数。
type QueryMerIDBalanceReq struct {
	ProviderID  uint64  `json:"providerId"`            // 服务商 ID
	PaymentType *uint64 `json:"paymentType,omitempty"` // 账户类型：0=银行卡 1=支付宝 2=微信（可选）
}

// QueryMerIDBalanceResp 商户账户余额查询响应参数。
type QueryMerIDBalanceResp struct {
	Balance    int64  `json:"balance,omitempty"`    // 账户余额（单位：分）
	ProviderID uint64 `json:"providerId,omitempty"` // 服务商 ID
}

// ---------------------------------------------------------------------------
// 5.6 对账文件下载（6004）
// ---------------------------------------------------------------------------

// DownloadReconciliationReq 对账文件下载请求参数。
type DownloadReconciliationReq struct {
	BillDate string `json:"billDate"` // 对账日期（yyyy-MM-dd）
}

// DownloadReconciliationResp 对账文件下载响应参数。
type DownloadReconciliationResp struct {
	FilePath string `json:"filePath,omitempty"` // 文件下载链接
	BillDate string `json:"billDate,omitempty"` // 对账日期
}

// ---------------------------------------------------------------------------
// 5.7 自由职业者余额查询（6005）
// ---------------------------------------------------------------------------

// QueryFreelancerBalanceReq 自由职业者余额查询请求参数。
type QueryFreelancerBalanceReq struct {
	ProviderID uint64 `json:"providerId"` // 服务商 ID
	Name       string `json:"name"`       // 姓名
	IDCard     string `json:"idCard"`     // 身份证号（18 位）
}

// QueryFreelancerBalanceResp 自由职业者余额查询响应参数。
type QueryFreelancerBalanceResp struct {
	ProviderID uint64 `json:"providerId,omitempty"` // 服务商 ID
	Name       string `json:"name,omitempty"`       // 姓名
	IDCard     string `json:"idCard,omitempty"`     // 身份证号
	Balance    int64  `json:"balance,omitempty"`    // 剩余额度（单位：分）
}

// ---------------------------------------------------------------------------
// 5.8 试算（6006）
// ---------------------------------------------------------------------------

// TrialCalcReq 试算请求参数。
type TrialCalcReq struct {
	ProviderID uint64          `json:"providerId"`          // 服务商 ID
	IfReverse  *bool           `json:"ifReverse,omitempty"` // true=反算（由到手金额反推） false=正算（默认）
	UserList   []TrialCalcUser `json:"userList"`            // 试算用户列表（最多 50 条）
}

// TrialCalcUser 试算用户明细。
type TrialCalcUser struct {
	MerOrderID string `json:"merOrderId,omitempty"` // 商户订单号（最长 32 位，可选）
	Name       string `json:"name"`                 // 姓名（最长 32 位）
	IDCardNo   string `json:"idCardNo"`             // 身份证号（18 位）
	Amt        uint64 `json:"amt"`                  // 金额（单位：分）
}

// TrialCalcResp 试算响应参数（数组）。
type TrialCalcResp []TrialCalcResult

// TrialCalcResult 试算结果明细。
type TrialCalcResult struct {
	MerOrderID   string `json:"merOrderId,omitempty"`   // 商户订单号
	Name         string `json:"name,omitempty"`         // 姓名
	IDCardNo     string `json:"idCardNo,omitempty"`     // 身份证号
	ProviderID   uint64 `json:"providerId,omitempty"`   // 服务商 ID
	Amt          uint64 `json:"amt,omitempty"`          // 金额（单位：分）
	OrderAmt     uint64 `json:"orderAmt,omitempty"`     // 订单金额（单位：分）
	UserFeeRatio string `json:"userFeeRatio,omitempty"` // 个人服务费率/税率
	UserFee      uint64 `json:"userFee,omitempty"`      // 个税金额（单位：分）
	VATax        uint64 `json:"vaTax,omitempty"`        // 增值税（单位：分）
	VAAddTax     uint64 `json:"vaAddTax,omitempty"`     // 增值税附加（单位：分）
	UserDueAmt   uint64 `json:"userDueAmt,omitempty"`   // 实际到账金额（单位：分）
	Status       bool   `json:"status,omitempty"`       // 是否成功
	ErrMsg       string `json:"errMsg,omitempty"`       // 错误信息
}

// ---------------------------------------------------------------------------
// 5.26 任务列表查询（6031）
// ---------------------------------------------------------------------------

// QueryTaskListReq 任务列表查询请求参数（无业务参数）。
type QueryTaskListReq struct{}

// QueryTaskListResp 任务列表查询响应参数（数组）。
type QueryTaskListResp []TaskInfo

// TaskInfo 任务信息明细。
type TaskInfo struct {
	TaskID     uint64 `json:"taskId,omitempty"`     // 任务 ID
	TaskName   string `json:"taskName,omitempty"`   // 任务名称
	TaskStatus string `json:"taskStatus,omitempty"` // 任务状态
	StartTime  string `json:"startTime,omitempty"`  // 开始时间（yyyy-MM-dd）
	EndTime    string `json:"endTime,omitempty"`    // 结束时间（yyyy-MM-dd）
}

// ---------------------------------------------------------------------------
// 5.29 报税查询（6047）
// ---------------------------------------------------------------------------

// QueryTaxReportReq 报税查询请求参数。
type QueryTaxReportReq struct {
	ProviderID    uint64 `json:"providerId"`    // 服务商 ID
	DeclarePeriod string `json:"declarePeriod"` // 申报期（yyyyMM 格式，最长 6 位）
}

// QueryTaxReportResp 报税查询响应参数。
type QueryTaxReportResp struct {
	ResData string `json:"resData,omitempty"` // 报税文件下载地址（最长 256 位）
}
