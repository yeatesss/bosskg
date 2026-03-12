package bosskg

// FunCode 接口功能码。
type FunCode string

const (
	// FunCodeBatchTransfer 商户批量付款（5.3）。
	FunCodeBatchTransfer FunCode = "6001"
	// FunCodeQueryBatchTransferResult 商户批量付款查询（5.4）。
	FunCodeQueryBatchTransferResult FunCode = "6002"
	// FunCodeQueryMerIDBalance 商户账户余额查询（5.5）。
	FunCodeQueryMerIDBalance FunCode = "6003"
	// FunCodeDownloadReconciliation 对账文件下载（5.6）。
	FunCodeDownloadReconciliation FunCode = "6004"
	// FunCodeQueryFreelancerBalance 自由职业者余额查询（5.7）。
	FunCodeQueryFreelancerBalance FunCode = "6005"
	// FunCodeTrialCalc 试算（5.8）。
	FunCodeTrialCalc FunCode = "6006"
	// FunCodeSyncFaceAuth 同步人脸识别信息（5.19）。
	FunCodeSyncFaceAuth FunCode = "6008"
	// FunCodeFaceRecognition 人脸识别 H5（5.18）。
	FunCodeFaceRecognition FunCode = "6009"
	// FunCodeProfessionalSign 自由职业者无感签约（5.1）。
	FunCodeProfessionalSign FunCode = "6010"
	// FunCodeQueryProfessionalSignResult 自由职业者签约查询（5.2）。
	FunCodeQueryProfessionalSignResult FunCode = "6011"
	// FunCodeQueryInvoiceAmount 可开票金额查询（5.10）。
	FunCodeQueryInvoiceAmount FunCode = "6012"
	// FunCodeApplyInvoice 申请开票（5.11）。
	FunCodeApplyInvoice FunCode = "6013"
	// FunCodeQueryInvoiceResult 开票结果查询（5.12）。
	FunCodeQueryInvoiceResult FunCode = "6014"
	// FunCodeQueryInvoiceType 开票类目查询（5.9）。
	FunCodeQueryInvoiceType FunCode = "6015"
	// FunCodeQueryRecharge 充值查询（5.13）。
	FunCodeQueryRecharge FunCode = "6018"
	// FunCodeQueryAvailableRecharge 查询可充值金额（5.14）。
	FunCodeQueryAvailableRecharge FunCode = "6019"
	// FunCodeApplyRecharge 申请充值（5.16）。
	FunCodeApplyRecharge FunCode = "6020"
	// FunCodeQueryRechargeResult 充值结果查询（5.17）。
	FunCodeQueryRechargeResult FunCode = "6021"
	// FunCodeBatchUpload 批量上传（5.20）。
	FunCodeBatchUpload FunCode = "6022"
	// FunCodeQueryBatchOrder 批量订单查询（5.21）。
	FunCodeQueryBatchOrder FunCode = "6023"
	// FunCodeQueryReceipt 回单查询（5.22）。
	FunCodeQueryReceipt FunCode = "6024"
	// FunCodeH5SensitiveSign H5 敏感签约（5.23）。
	FunCodeH5SensitiveSign FunCode = "6026"
	// FunCodeAllInOneRelease 一键发放（5.24）。
	FunCodeAllInOneRelease FunCode = "6029"
	// FunCodeQueryAllInOneResult 一键发放查询（5.25）。
	FunCodeQueryAllInOneResult FunCode = "6030"
	// FunCodeQueryTaskList 任务列表查询（5.26）。
	FunCodeQueryTaskList FunCode = "6031"
	// FunCodeProfessionalUnSign 自由职业者解约（5.15）。
	FunCodeProfessionalUnSign FunCode = "6036"
	// FunCodeWeChatCancelTransfer 微信取消转账（5.27）。
	FunCodeWeChatCancelTransfer FunCode = "6043"
	// FunCodeQuerySignPagination 签约分页查询（5.28）。
	FunCodeQuerySignPagination FunCode = "6044"
	// FunCodeQueryTaxReport 报税查询（5.29）。
	FunCodeQueryTaxReport FunCode = "6047"
)

// APIVersion 接口版本号，目前版本为 V1.0（V 大写）。
const APIVersion = "V1.0"

// ReqMessage 公共请求参数（4.3）。
type ReqMessage struct {
	ReqID   string  `json:"reqId,omitempty"`   // 请求序号，每次请求唯一
	FunCode FunCode `json:"funCode,omitempty"` // 接口功能码
	MerID   string  `json:"merId,omitempty"`   // 商户号
	Version string  `json:"version,omitempty"` // 接口版本号（V1.0）
	ReqData string  `json:"reqData,omitempty"` // 业务参数 JSON 经 DES-ECB 加密后的 Base64 字符串
	Remark1 string  `json:"remark1,omitempty"` // 备注字段1：身份证人像面照片（hex 字符串）
	Remark2 string  `json:"remark2,omitempty"` // 备注字段2：身份证国徽面照片（hex 字符串）
	Sign    string  `json:"sign,omitempty"`    // 对 reqData 的 RSA-SHA1 签名（Base64）
}

// RespMessage 公共返回参数（4.4）。
type RespMessage struct {
	ReqID   string  `json:"reqId,omitempty"`   // 请求序号
	FunCode FunCode `json:"funCode,omitempty"` // 接口功能码
	MerID   string  `json:"merId,omitempty"`   // 商户号
	Version string  `json:"version,omitempty"` // 接口版本号
	ResData string  `json:"resData,omitempty"` // 业务返回 JSON 经 DES-ECB 加密后的 Base64 字符串
	ResCode string  `json:"resCode,omitempty"` // 响应码，"0000" 表示成功
	ResMsg  string  `json:"resMsg,omitempty"`  // 响应信息
	Sign    string  `json:"sign,omitempty"`    // 对 resData 的 RSA-SHA1 签名（Base64）
}
