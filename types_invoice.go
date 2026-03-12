package bosskg

// ---------------------------------------------------------------------------
// 5.9 开票类目查询（6015）
// ---------------------------------------------------------------------------

// QueryInvoiceTypeReq 开票类目查询请求参数。
type QueryInvoiceTypeReq struct {
	ProviderID uint64 `json:"providerId"` // 服务商 ID
}

// QueryInvoiceTypeResp 开票类目查询响应参数（数组）。
type QueryInvoiceTypeResp []InvoiceType

// InvoiceType 开票类目明细。
type InvoiceType struct {
	ProviderID    uint64 `json:"providerId,omitempty"`    // 服务商 ID
	MerID         string `json:"merId,omitempty"`         // 商户号
	InvoiceTypeID uint64 `json:"invoiceTypeId,omitempty"` // 开票类目 ID
	InvoiceName   string `json:"invoiceName,omitempty"`   // 开票类目名称
}

// ---------------------------------------------------------------------------
// 5.10 可开票金额查询（6012）
// ---------------------------------------------------------------------------

// QueryInvoiceAmountReq 可开票金额查询请求参数。
type QueryInvoiceAmountReq struct {
	ProviderID uint64 `json:"providerId"` // 服务商 ID
}

// QueryInvoiceAmountResp 可开票金额查询响应参数。
type QueryInvoiceAmountResp struct {
	ProviderID   uint64 `json:"providerId,omitempty"`   // 服务商 ID
	AvailableAmt int64  `json:"availableAmt,omitempty"` // 可开票金额（单位：分）
}

// ---------------------------------------------------------------------------
// 5.11 申请开票（6013）
// ---------------------------------------------------------------------------

// ApplyInvoiceReq 申请开票请求参数。
type ApplyInvoiceReq struct {
	ProviderID    uint64 `json:"providerId"`            // 服务商 ID
	InvoiceTypeID uint64 `json:"invoiceTypeId"`         // 开票类目 ID
	Amt           uint64 `json:"amt"`                   // 开票金额（单位：分）
	InvoiceType   string `json:"invoiceType"`           // 发票类型：SPECIAL=专用发票 PLAIN=普通发票
	InvoiceMemo   string `json:"invoiceMemo,omitempty"` // 开票备注
	Contact       string `json:"contact,omitempty"`     // 联系人
	Mobile        string `json:"mobile,omitempty"`      // 联系电话
	PostAddress   string `json:"postAddress,omitempty"` // 邮寄地址
	TicketType    string `json:"ticketType"`            // 票据类型：PAPER=纸质 ELECTRONIC=电子
}

// ApplyInvoiceResp 申请开票响应参数。
type ApplyInvoiceResp struct {
	ProviderID     uint64 `json:"providerId,omitempty"`     // 服务商 ID
	InvoiceApplyNo string `json:"invoiceApplyNo,omitempty"` // 开票申请编号
	InvoiceTypeID  uint64 `json:"invoiceTypeId,omitempty"`  // 开票类目 ID
	Amt            uint64 `json:"amt,omitempty"`            // 开票金额（单位：分）
	InvoiceMemo    string `json:"invoiceMemo,omitempty"`    // 开票备注
	Contact        string `json:"contact,omitempty"`        // 联系人
	Mobile         string `json:"mobile,omitempty"`         // 联系电话
	PostAddress    string `json:"postAddress,omitempty"`    // 邮寄地址
}

// ---------------------------------------------------------------------------
// 5.12 开票结果查询（6014）
// ---------------------------------------------------------------------------

// QueryInvoiceResultReq 开票结果查询请求参数。
type QueryInvoiceResultReq struct {
	InvoiceApplyNo string `json:"invoiceApplyNo,omitempty"` // 开票申请编号（可选）
	ProviderID     uint64 `json:"providerId,omitempty"`     // 服务商 ID（可选）
	StartDate      string `json:"startDate,omitempty"`      // 开始日期（yyyy-MM-dd，可选）
	EndDate        string `json:"endDate,omitempty"`        // 结束日期（yyyy-MM-dd，可选）
}

// QueryInvoiceResultResp 开票结果查询响应参数（数组）。
type QueryInvoiceResultResp []InvoiceResult

// InvoiceResult 开票结果明细。
type InvoiceResult struct {
	MerID           string `json:"merId,omitempty"`           // 商户号
	ProviderID      uint64 `json:"providerId,omitempty"`      // 服务商 ID
	InvoiceApplyNo  string `json:"invoiceApplyNo,omitempty"`  // 开票申请编号
	CreateTime      string `json:"createTime,omitempty"`      // 创建时间（yyyy-MM-dd HH:mm:ss）
	InvoiceTypeID   uint64 `json:"invoiceTypeId,omitempty"`   // 开票类目 ID
	Amt             string `json:"amt,omitempty"`             // 开票金额（单位：分，文档为 string 类型）
	InvoiceMemo     string `json:"invoiceMemo,omitempty"`     // 开票备注
	State           uint64 `json:"state,omitempty"`           // 开票状态：0=处理中 1=已开票 2=已驳回 3=已作废 4=已邮寄 5=已冲红
	Contact         string `json:"contact,omitempty"`         // 联系人
	Mobile          string `json:"mobile,omitempty"`          // 联系电话
	PostAddress     string `json:"postAddress,omitempty"`     // 邮寄地址
	InvoiceNum      string `json:"invoiceNum,omitempty"`      // 发票号码
	InvoiceCode     string `json:"invoiceCode,omitempty"`     // 发票代码
	ExpressID       string `json:"expressId,omitempty"`       // 快递公司
	TrackNo         string `json:"trackNo,omitempty"`         // 快递单号
	InvoiceFileList string `json:"invoiceFileList,omitempty"` // 发票文件列表
}
