package bosskg

// ---------------------------------------------------------------------------
// 5.1 自由职业者无感签约（6010）
// ---------------------------------------------------------------------------

// ProfessionalSignReq 自由职业者无感签约请求参数。
type ProfessionalSignReq struct {
	Name        string   `json:"name"`                 // 姓名（最长 25 位）
	CardNo      string   `json:"cardNo"`               // 银行卡号 / 支付宝账号 / 微信 OpenID（最长 25 位）
	IDCard      string   `json:"idCard"`               // 身份证号（18 位）
	Mobile      string   `json:"mobile"`               // 银行预留手机号（11 位）
	PaymentType uint64   `json:"paymentType"`          // 签约方式：0=银行卡 1=支付宝 2=微信
	ProviderID  uint64   `json:"providerId"`           // 服务商 ID
	IDCardPic1  string   `json:"idCardPic1"`           // 身份证人像面（图片 bytes 转 hex 字符串，< 1MB）
	IDCardPic2  string   `json:"idCardPic2"`           // 身份证国徽面（图片 bytes 转 hex 字符串，< 1MB）
	OtherParam  string   `json:"otherParam,omitempty"` // 透传参数（不传时可能不返回 resData/sign）
	NotifyURL   string   `json:"notifyUrl,omitempty"`  // 签约结果异步回调地址
	TagList     []string `json:"tagList,omitempty"`    // 自由职业者技能标签
}

// ProfessionalSignResp 自由职业者无感签约响应参数。
type ProfessionalSignResp struct {
	OtherParam string `json:"otherParam,omitempty"` // 透传参数
}

// ---------------------------------------------------------------------------
// 5.2 自由职业者签约查询（6011）
// ---------------------------------------------------------------------------

// QueryProfessionalSignResultReq 自由职业者签约查询请求参数。
type QueryProfessionalSignResultReq struct {
	Name       string `json:"name"`       // 姓名（最长 25 位）
	IDCard     string `json:"idCard"`     // 身份证号（18 位）
	Mobile     string `json:"mobile"`     // 银行预留手机号（11 位）
	ProviderID uint64 `json:"providerId"` // 服务商 ID
}

// QueryProfessionalSignResultResp 自由职业者签约查询响应参数。
type QueryProfessionalSignResultResp struct {
	Name            string `json:"name,omitempty"`            // 姓名
	CardNo          string `json:"cardNo,omitempty"`          // 银行卡号
	IDCard          string `json:"idCard,omitempty"`          // 身份证号
	Mobile          string `json:"mobile,omitempty"`          // 银行预留手机号
	State           uint64 `json:"state,omitempty"`           // 签约状态：0=未签约 1=已签约 2=未检索到 3=签约中 4=签约失败 5=已解约
	OtherParam      string `json:"otherParam,omitempty"`      // 透传参数
	ProviderID      uint64 `json:"providerId,omitempty"`      // 服务商 ID
	RetMsg          string `json:"retMsg,omitempty"`          // 失败原因
	FaceAuthState   string `json:"faceAuthState,omitempty"`   // 人脸认证状态：UN_AUTH / PROCESS / SUCCESS / FAILED / EXPIRED
	FaceAuthEndTime string `json:"faceAuthEndTime,omitempty"` // 人脸认证有效期限（YYYY-MM-DD）
}

// ---------------------------------------------------------------------------
// 5.15 自由职业者解约（6036）
// ---------------------------------------------------------------------------

// ProfessionalUnSignReq 自由职业者解约请求参数。
type ProfessionalUnSignReq struct {
	UserName   string  `json:"userName"`             // 姓名（最长 25 位）
	IDCardNo   string  `json:"idcardNo"`             // 身份证号（18 位）
	ProviderID *uint64 `json:"providerId,omitempty"` // 服务商 ID（可选，不传则解约该商户下所有服务商）
}

// ProfessionalUnSignResp 自由职业者解约响应参数。
type ProfessionalUnSignResp struct {
	State  string `json:"state,omitempty"`  // 解约状态：1=成功 2=失败
	RetMsg string `json:"retMsg,omitempty"` // 返回描述
}

// ---------------------------------------------------------------------------
// 5.18 人脸识别 H5（6009）
// ---------------------------------------------------------------------------

// FaceRecognitionReq 人脸识别 H5 请求参数。
type FaceRecognitionReq struct {
	Name            string `json:"name"`                      // 姓名（最长 25 位）
	IDCard          string `json:"idCard"`                    // 身份证号（18 位）
	Mobile          string `json:"mobile"`                    // 手机号（11 位）
	RedirectBtnName string `json:"redirectBtnName,omitempty"` // 跳转按钮名称
	RedirectURL     string `json:"redirectUrl,omitempty"`     // 跳转地址
	RedirectType    string `json:"redirectType,omitempty"`    // 跳转类型
	AppID           string `json:"appid,omitempty"`           // 小程序 AppID（小程序场景必填，最长 25 位）
}

// FaceRecognitionResp 人脸识别 H5 响应参数。
type FaceRecognitionResp struct {
	URL string `json:"url,omitempty"` // 人脸识别 H5 页面地址（有效期 1 天）
}

// ---------------------------------------------------------------------------
// 5.19 同步人脸识别信息（6008）
// ---------------------------------------------------------------------------

// SyncFaceAuthReq 同步人脸识别信息请求参数。
type SyncFaceAuthReq struct {
	Name        string   `json:"name"`        // 姓名（最长 25 位）
	IDCard      string   `json:"idCard"`      // 身份证号（18 位）
	Mobile      string   `json:"mobile"`      // 手机号（11 位）
	ThirdID     string   `json:"thirdId"`     // 可追溯编码（最长 50 位）
	AuthTime    string   `json:"authTime"`    // 认证时间（yyyy-MM-dd HH:mm:ss）
	URLs        []string `json:"urls"`        // 图片/视频地址列表（每张 < 2MB）
	AuthChannel string   `json:"authChannel"` // 认证渠道（01~12 渠道码，最长 2 位）
}

// SyncFaceAuthResp 同步人脸识别信息响应参数。
type SyncFaceAuthResp struct {
	FaceAuthEndTime string `json:"faceAuthEndTime,omitempty"` // 人脸认证有效期限（yyyy-MM-dd）
}

// ---------------------------------------------------------------------------
// 5.23 H5 敏感签约（6026）
// ---------------------------------------------------------------------------

// H5SensitiveSignReq H5 敏感签约请求参数。
type H5SensitiveSignReq struct {
	UserName        string `json:"userName"`                  // 姓名（最长 25 位）
	CardNo          string `json:"cardNo"`                    // 银行卡号 / 支付宝账号 / 微信 OpenID（最长 25 位）
	IDCard          string `json:"idCard"`                    // 身份证号（18 位）
	Mobile          string `json:"mobile"`                    // 银行预留手机号（11 位）
	IDCardFrontPic  string `json:"idCardFrontPic"`            // 身份证人像面（hex 字符串，< 1MB）
	IDCardBackPic   string `json:"idCardBackPic"`             // 身份证国徽面（hex 字符串，< 1MB）
	PaymentType     uint64 `json:"paymentType"`               // 签约方式：0=银行卡 1=支付宝 2=微信
	NotifyURL       string `json:"notifyUrl,omitempty"`       // 异步回调地址
	RedirectBtnName string `json:"redirectBtnName,omitempty"` // 跳转按钮名称
	RedirectURL     string `json:"redirectUrl,omitempty"`     // 跳转地址
	RedirectType    string `json:"redirectType,omitempty"`    // 跳转类型
	OtherParam      string `json:"otherParam,omitempty"`      // 透传参数（最长 200 位）
	AppID           string `json:"appid,omitempty"`           // 小程序 AppID（小程序场景必填，最长 25 位）
}

// H5SensitiveSignResp H5 敏感签约响应参数。
type H5SensitiveSignResp struct {
	ResData string `json:"resData,omitempty"` // H5 嵌入链接（含 token，有效期 90 天）
}

// ---------------------------------------------------------------------------
// 5.28 签约分页查询（6044）
// ---------------------------------------------------------------------------

// QuerySignPaginationReq 签约分页查询请求参数。
type QuerySignPaginationReq struct {
	ProviderID      uint64 `json:"providerId"`                // 服务商 ID
	CreateTimeBegin string `json:"createTimeBegin"`           // 创建时间起（yyyy-MM-dd HH:mm:ss）
	CreateTimeEnd   string `json:"createTimeEnd"`             // 创建时间止（yyyy-MM-dd HH:mm:ss）
	State           uint64 `json:"state"`                     // 签约状态：0=未签约 1=已签约 3=签约中 4=签约失败 5=已解约
	FinishTimeBegin string `json:"finishTimeBegin,omitempty"` // 完成时间起（yyyy-MM-dd HH:mm:ss，可选）
	FinishTimeEnd   string `json:"finishTimeEnd,omitempty"`   // 完成时间止（yyyy-MM-dd HH:mm:ss，可选）
	OffsetID        string `json:"offsetId,omitempty"`        // 分页偏移 ID（最长 24 位，可选）
}

// QuerySignPaginationResp 签约分页查询响应参数（数组）。
type QuerySignPaginationResp []SignPaginationItem

// SignPaginationItem 签约分页查询结果明细。
type SignPaginationItem struct {
	Name       string `json:"name,omitempty"`       // 姓名
	CardNo     string `json:"cardNo,omitempty"`     // 银行卡号
	IDCard     string `json:"idCard,omitempty"`     // 身份证号
	Mobile     string `json:"mobile,omitempty"`     // 手机号
	State      uint64 `json:"state,omitempty"`      // 签约状态
	ProviderID uint64 `json:"providerId,omitempty"` // 服务商 ID
	OffsetID   string `json:"offsetId,omitempty"`   // 分页偏移 ID（用于下一页查询）
	RetMsg     string `json:"retMsg,omitempty"`     // 失败原因
}
