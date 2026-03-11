package bosskg

type FunCode string

const (
	// 5.5 商户账户余额查询
	FunCodeQueryMerIDBalance FunCode = "6003"
	// 5.3 商户批量付款
	FunCodeBatchTransfer FunCode = "6001"
	// 5.4 商户批量付款查询
	FunCodeQueryBatchTransferResult FunCode = "6002"
	// 5.1 自由职业者无感签约
	FunCodeProfessionalSign FunCode = "6010"
	// 5.2 自由职业者签约查询
	FunCodeQueryProfessionalSignResult FunCode = "6011"
	// 5.15 自由职业者解约
	FunCodeProfessionalUnSign FunCode = "6036"
)

const APIVersion = "V1.0"

// ReqMessage 公共请求参数（4.3）。
type ReqMessage struct {
	ReqID   string  `json:"reqId,omitempty"`   // 请求序号，每次请求保持唯一，表明报文的唯一编号
	FunCode FunCode `json:"funCode,omitempty"` // 接口编码
	MerID   string  `json:"merId,omitempty"`   // 商户号
	Version string  `json:"version,omitempty"` // 接口版本号，目前版本为 V1.0（V 大写）
	ReqData string  `json:"reqData,omitempty"` // 业务参数（JSON）DES-ECB 加密后的 Base64 字符串
	Remark1 string  `json:"remark1,omitempty"` // 备注字段1（4.3）：身份证人像面照片（hex 字符串），部分场景使用
	Remark2 string  `json:"remark2,omitempty"` // 备注字段2（4.3）：身份证国徽面照片（hex 字符串），部分场景使用
	Sign    string  `json:"sign,omitempty"`    // 业务数据签名（RSA-SHA1，对 reqData 签名后 Base64）
}

// RespMessage 公共返回参数（4.4）。
type RespMessage struct {
	ReqID   string  `json:"reqId,omitempty"`   // 请求序号
	FunCode FunCode `json:"funCode,omitempty"` // 接口编码
	MerID   string  `json:"merId,omitempty"`   // 商户号
	Version string  `json:"version,omitempty"` // 接口版本号

	ResData string `json:"resData,omitempty"` // 业务返回（JSON）DES-ECB 加密后的 Base64 字符串（部分接口可能不返回）
	ResCode string `json:"resCode,omitempty"` // 响应码（详见 6.1）
	ResMsg  string `json:"resMsg,omitempty"`  // 响应信息
	Sign    string `json:"sign,omitempty"`    // 业务数据签名结果（RSA-SHA1，对 resData 签名后 Base64；resData 为空时可能为空）
}

// QueryMerIDBalanceReq 商户账户余额查询（5.5.2）。
type QueryMerIDBalanceReq struct {
	ProviderID uint64 `json:"providerId"` // 服务商ID（联系客服获取）
	// PaymentType 账户类型：0银行卡 1支付宝 2微信（可选）
	PaymentType *uint64 `json:"paymentType,omitempty"`
}

// QueryMerIDBalanceResp 商户账户余额查询返回（5.5.3）。
type QueryMerIDBalanceResp struct {
	Balance    string `json:"balance,omitempty"`    // 账户余额（单位：分）
	ProviderID uint64 `json:"providerId,omitempty"` // 服务商ID
}

// BatchTransferReq 商户批量付款（5.3.2）。
type BatchTransferReq struct {
	MerBatchID string    `json:"merBatchId"` // 商户批次号（建议 yyyymmddHHSS + 8 位随机数；商户维度唯一）
	PayItems   []PayItem `json:"payItems"`   // 付款明细
	TaskID     uint64    `json:"taskId"`     // 任务ID（商户平台获取）
	ProviderID uint64    `json:"providerId"` // 服务商ID（联系客服获取）
}

// PayItem 批量付款明细（5.3.2.1）。
type PayItem struct {
	MerOrderID  string `json:"merOrderId"`          // 商户订单号（商户维度唯一）
	Amt         uint64 `json:"amt"`                 // 付款金额（单位：分）
	PayeeName   string `json:"payeeName"`           // 收款人名称
	PayeeAcc    string `json:"payeeAcc"`            // 收款账号（银行卡/支付宝/微信openid）
	IDCard      string `json:"idCard"`              // 身份证号
	Mobile      string `json:"mobile"`              // 收款人手机号（格式校验）
	Memo        string `json:"memo,omitempty"`      // 备注（敏感词限制见 5.3.1）
	PaymentType uint64 `json:"paymentType"`         // 付款方式：0银行卡 1支付宝 2微信
	NotifyURL   string `json:"notifyUrl,omitempty"` // 异步通知地址（不填则不通知）
}

// BatchTransferResp 商户批量付款返回（5.3.3）。
type BatchTransferResp struct {
	SuccessNum    uint64      `json:"successNum,omitempty"`    // 该批次订单受理成功笔数
	FailureNum    uint64      `json:"failureNum,omitempty"`    // 该批次订单受理失败笔数
	MerBatchID    string      `json:"merBatchId,omitempty"`    // 商户批次号
	PayResultList []PayResult `json:"payResultList,omitempty"` // 付款返回数据
}

// PayResult 付款返回明细（5.3.3.1）。
type PayResult struct {
	MerOrderID  string `json:"merOrderId,omitempty"`  // 商户订单号
	OrderNo     string `json:"orderNo,omitempty"`     // 平台订单号（全局唯一，长度可达 25 位）
	Amt         uint64 `json:"amt,omitempty"`         // 付款金额（分）
	Fee         uint64 `json:"fee,omitempty"`         // 服务费（分）
	PackageInfo string `json:"packageInfo,omitempty"` // 微信零钱新模式：跳转收款页 package 信息
	MchID       string `json:"mchId,omitempty"`       // 微信零钱新模式：跳转收款页 mchId 信息
	ResCode     string `json:"resCode,omitempty"`     // 受理响应码（非交易终态）
	ResMsg      string `json:"resMsg,omitempty"`      // 受理响应信息
}

// QueryBatchTransferResultReq 商户批量付款查询（5.4.2）。
type QueryBatchTransferResultReq struct {
	MerBatchID string      `json:"merBatchId"`           // 商户批次号
	QueryItems []QueryItem `json:"queryItems,omitempty"` // 查询条件（为空返回该批次全部订单）
}

// QueryItem 付款查询项（5.4.2.1 / 5.4.3.1）。
type QueryItem struct {
	MerOrderID string `json:"merOrderId,omitempty"` // 商户订单号
	OrderNo    string `json:"orderNo,omitempty"`    // 平台订单号（长度可达 25 位）

	State uint64 `json:"state,omitempty"` // 交易状态：1付款中、3成功、4失败、6待用户确认、7已取消

	Amt        uint64 `json:"amt,omitempty"`        // 付款金额（分）
	Fee        uint64 `json:"fee,omitempty"`        // 平台管理费（分）
	UserFee    uint64 `json:"userFee,omitempty"`    // 个人服务费/个税（分）
	VATax      uint64 `json:"vaTax,omitempty"`      // 个人增值税（分）
	VAAddTax   uint64 `json:"vaAddTax,omitempty"`   // 个人增值税附加（分）
	UserDueAmt uint64 `json:"userDueAmt,omitempty"` // 个人实际到账金额（分）

	UserFeeRatio string `json:"userFeeRatio,omitempty"` // 个人服务费率/个税税率
	ResMsg       string `json:"resMsg,omitempty"`       // 响应信息
	CreateTime   string `json:"createTime,omitempty"`   // 创建时间（yyyy-MM-dd HH:mm:ss）
	EndTime      string `json:"endTime,omitempty"`      // 完成时间（yyyy-MM-dd HH:mm:ss）
}

// QueryBatchTransferResultResp 商户批量付款查询返回（5.4.3）。
type QueryBatchTransferResultResp struct {
	MerID      string      `json:"merId,omitempty"`      // 商户号
	MerBatchID string      `json:"merBatchId,omitempty"` // 商户批次号
	QueryItems []QueryItem `json:"queryItems,omitempty"` // 查询结果
}

// ProfessionalSignReq 自由职业者无感签约（5.1.2）。
type ProfessionalSignReq struct {
	Name        string   `json:"name"`                 // 姓名
	CardNo      string   `json:"cardNo"`               // 银行卡号/支付宝账号/微信openid
	IDCard      string   `json:"idCard"`               // 身份证号
	Mobile      string   `json:"mobile"`               // 银行预留手机号
	PaymentType uint64   `json:"paymentType"`          // 签约方式：0银行卡 1支付宝 2微信
	ProviderID  uint64   `json:"providerId"`           // 服务商ID（联系客服获取）
	IDCardPic1  string   `json:"idCardPic1"`           // 身份证人像面（图片 bytes 转 hex 字符串，<1M）
	IDCardPic2  string   `json:"idCardPic2"`           // 身份证国徽面（图片 bytes 转 hex 字符串，<1M）
	OtherParam  string   `json:"otherParam,omitempty"` // 透传参数（不传可能不返回 resData/sign）
	NotifyURL   string   `json:"notifyUrl,omitempty"`  // 签约结果异步回调地址
	TagList     []string `json:"tagList,omitempty"`    // 自由职业者技能标签
}

// ProfessionalSignResp 自由职业者无感签约返回（5.1.3）。
type ProfessionalSignResp struct {
	OtherParam string `json:"otherParam,omitempty"` // 透传参数
}

// QueryProfessionalSignResultReq 自由职业者签约查询（5.2.2）。
type QueryProfessionalSignResultReq struct {
	Name       string `json:"name"`       // 姓名
	IDCard     string `json:"idCard"`     // 身份证号
	Mobile     string `json:"mobile"`     // 银行预留手机号
	ProviderID uint64 `json:"providerId"` // 服务商ID（联系客服获取）
}

// QueryProfessionalSignResultResp 自由职业者签约查询返回（5.2.3）。
type QueryProfessionalSignResultResp struct {
	Name            string `json:"name,omitempty"`            // 姓名
	CardNo          string `json:"cardNo,omitempty"`          // 银行卡号
	IDCard          string `json:"idCard,omitempty"`          // 身份证号
	Mobile          string `json:"mobile,omitempty"`          // 银行预留手机号
	State           uint64 `json:"state,omitempty"`           // 签约状态：0未签约 1已签约 2未检索到 3签约中 4签约失败 5已解约
	OtherParam      string `json:"otherParam,omitempty"`      // 透传参数
	ProviderID      uint64 `json:"providerId,omitempty"`      // 服务商ID
	RetMsg          string `json:"retMsg,omitempty"`          // 失败原因
	FaceAuthState   string `json:"faceAuthState,omitempty"`   // 人脸认证状态：UN_AUTH/PROCESS/SUCCESS/FAILED/EXPIRED
	FaceAuthEndTime string `json:"faceAuthEndTime,omitempty"` // 人脸认证有效期限（YYYY-MM-DD）
}

// ProfessionalUnSignReq 自由职业者解约（5.15.1）。
type ProfessionalUnSignReq struct {
	UserName string `json:"userName"` // 姓名
	IDCardNo string `json:"idcardNo"` // 身份证号
	// ProviderID 服务商ID（可选；不传则解约该商户下所有服务商）
	ProviderID *uint64 `json:"providerId,omitempty"`
}

// ProfessionalUnSignResp 自由职业者解约返回（5.15.2）。
type ProfessionalUnSignResp struct {
	State  string `json:"state,omitempty"`  // 解约状态：1成功 2失败
	RetMsg string `json:"retMsg,omitempty"` // 返回描述
}
