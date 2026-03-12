package bosskg

import "context"

// QueryMerIDBalance 商户账户余额查询（6003）。
func (c *Client) QueryMerIDBalance(ctx context.Context, req QueryMerIDBalanceReq) (*QueryMerIDBalanceResp, error) {
	var out QueryMerIDBalanceResp
	_, err := c.Do(ctx, FunCodeQueryMerIDBalance, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// BatchTransfer 商户批量付款（6001）。
func (c *Client) BatchTransfer(ctx context.Context, req BatchTransferReq) (*BatchTransferResp, error) {
	var out BatchTransferResp
	_, err := c.Do(ctx, FunCodeBatchTransfer, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// QueryBatchTransferResult 商户批量付款查询（6002）。
func (c *Client) QueryBatchTransferResult(ctx context.Context, req QueryBatchTransferResultReq) (*QueryBatchTransferResultResp, error) {
	var out QueryBatchTransferResultResp
	_, err := c.Do(ctx, FunCodeQueryBatchTransferResult, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// ProfessionalSign 自由职业者无感签约（6010）。
func (c *Client) ProfessionalSign(ctx context.Context, req ProfessionalSignReq) (*ProfessionalSignResp, error) {
	var out ProfessionalSignResp
	_, err := c.Do(ctx, FunCodeProfessionalSign, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// QueryProfessionalSignResult 自由职业者签约查询（6011）。
func (c *Client) QueryProfessionalSignResult(ctx context.Context, req QueryProfessionalSignResultReq) (*QueryProfessionalSignResultResp, error) {
	var out QueryProfessionalSignResultResp
	_, err := c.Do(ctx, FunCodeQueryProfessionalSignResult, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// ProfessionalUnSign 自由职业者解约（6036）。
func (c *Client) ProfessionalUnSign(ctx context.Context, req ProfessionalUnSignReq) (*ProfessionalUnSignResp, error) {
	var out ProfessionalUnSignResp
	_, err := c.Do(ctx, FunCodeProfessionalUnSign, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// DownloadReconciliation 对账文件下载（6004）。
func (c *Client) DownloadReconciliation(ctx context.Context, req DownloadReconciliationReq) (*DownloadReconciliationResp, error) {
	var out DownloadReconciliationResp
	_, err := c.Do(ctx, FunCodeDownloadReconciliation, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// QueryFreelancerBalance 自由职业者余额查询（6005）。
func (c *Client) QueryFreelancerBalance(ctx context.Context, req QueryFreelancerBalanceReq) (*QueryFreelancerBalanceResp, error) {
	var out QueryFreelancerBalanceResp
	_, err := c.Do(ctx, FunCodeQueryFreelancerBalance, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// TrialCalc 试算（6006）。
func (c *Client) TrialCalc(ctx context.Context, req TrialCalcReq) (TrialCalcResp, error) {
	var out TrialCalcResp
	_, err := c.Do(ctx, FunCodeTrialCalc, req, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryInvoiceType 开票类目查询（6015）。
func (c *Client) QueryInvoiceType(ctx context.Context, req QueryInvoiceTypeReq) (QueryInvoiceTypeResp, error) {
	var out QueryInvoiceTypeResp
	_, err := c.Do(ctx, FunCodeQueryInvoiceType, req, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryInvoiceAmount 可开票金额查询（6012）。
func (c *Client) QueryInvoiceAmount(ctx context.Context, req QueryInvoiceAmountReq) (*QueryInvoiceAmountResp, error) {
	var out QueryInvoiceAmountResp
	_, err := c.Do(ctx, FunCodeQueryInvoiceAmount, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// ApplyInvoice 申请开票（6013）。
func (c *Client) ApplyInvoice(ctx context.Context, req ApplyInvoiceReq) (*ApplyInvoiceResp, error) {
	var out ApplyInvoiceResp
	_, err := c.Do(ctx, FunCodeApplyInvoice, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// QueryInvoiceResult 开票结果查询（6014）。
func (c *Client) QueryInvoiceResult(ctx context.Context, req QueryInvoiceResultReq) (QueryInvoiceResultResp, error) {
	var out QueryInvoiceResultResp
	_, err := c.Do(ctx, FunCodeQueryInvoiceResult, req, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryRecharge 充值查询（6018）。
func (c *Client) QueryRecharge(ctx context.Context, req QueryRechargeReq) (QueryRechargeResp, error) {
	var out QueryRechargeResp
	_, err := c.Do(ctx, FunCodeQueryRecharge, req, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryAvailableRecharge 查询可充值金额（6019）。
func (c *Client) QueryAvailableRecharge(ctx context.Context, req QueryAvailableRechargeReq) (*QueryAvailableRechargeResp, error) {
	var out QueryAvailableRechargeResp
	_, err := c.Do(ctx, FunCodeQueryAvailableRecharge, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// ApplyRecharge 申请充值（6020）。
func (c *Client) ApplyRecharge(ctx context.Context, req ApplyRechargeReq) (*ApplyRechargeResp, error) {
	var out ApplyRechargeResp
	_, err := c.Do(ctx, FunCodeApplyRecharge, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// QueryRechargeResult 充值结果查询（6021）。
func (c *Client) QueryRechargeResult(ctx context.Context, req QueryRechargeResultReq) (*QueryRechargeResultResp, error) {
	var out QueryRechargeResultResp
	_, err := c.Do(ctx, FunCodeQueryRechargeResult, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// BatchUpload 批量上传（6022）。
func (c *Client) BatchUpload(ctx context.Context, req BatchUploadReq) (*BatchUploadResp, error) {
	var out BatchUploadResp
	_, err := c.Do(ctx, FunCodeBatchUpload, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// QueryBatchOrder 批量订单查询（6023）。
func (c *Client) QueryBatchOrder(ctx context.Context, req QueryBatchOrderReq) (*QueryBatchOrderResp, error) {
	var out QueryBatchOrderResp
	_, err := c.Do(ctx, FunCodeQueryBatchOrder, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// QueryReceipt 回单查询（6024）。
func (c *Client) QueryReceipt(ctx context.Context, req QueryReceiptReq) (*QueryReceiptResp, error) {
	var out QueryReceiptResp
	_, err := c.Do(ctx, FunCodeQueryReceipt, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// H5SensitiveSign H5 敏感签约（6026）。
func (c *Client) H5SensitiveSign(ctx context.Context, req H5SensitiveSignReq) (*H5SensitiveSignResp, error) {
	var out H5SensitiveSignResp
	_, err := c.Do(ctx, FunCodeH5SensitiveSign, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// AllInOneRelease 一键发放（6029）。
func (c *Client) AllInOneRelease(ctx context.Context, req AllInOneReleaseReq) (*AllInOneReleaseResp, error) {
	var out AllInOneReleaseResp
	_, err := c.Do(ctx, FunCodeAllInOneRelease, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// QueryAllInOneResult 一键发放查询（6030）。
func (c *Client) QueryAllInOneResult(ctx context.Context, req QueryAllInOneResultReq) (*QueryAllInOneResultResp, error) {
	var out QueryAllInOneResultResp
	_, err := c.Do(ctx, FunCodeQueryAllInOneResult, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// QueryTaskList 任务列表查询（6031）。
func (c *Client) QueryTaskList(ctx context.Context) (QueryTaskListResp, error) {
	var out QueryTaskListResp
	_, err := c.Do(ctx, FunCodeQueryTaskList, QueryTaskListReq{}, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeChatCancelTransfer 微信取消转账（6043）。
func (c *Client) WeChatCancelTransfer(ctx context.Context, req WeChatCancelTransferReq) (*WeChatCancelTransferResp, error) {
	var out WeChatCancelTransferResp
	_, err := c.Do(ctx, FunCodeWeChatCancelTransfer, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// QuerySignPagination 签约分页查询（6044）。
func (c *Client) QuerySignPagination(ctx context.Context, req QuerySignPaginationReq) (QuerySignPaginationResp, error) {
	var out QuerySignPaginationResp
	_, err := c.Do(ctx, FunCodeQuerySignPagination, req, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaceRecognition 人脸识别 H5（6009）。
func (c *Client) FaceRecognition(ctx context.Context, req FaceRecognitionReq) (*FaceRecognitionResp, error) {
	var out FaceRecognitionResp
	_, err := c.Do(ctx, FunCodeFaceRecognition, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// SyncFaceAuth 同步人脸识别信息（6008）。
func (c *Client) SyncFaceAuth(ctx context.Context, req SyncFaceAuthReq) (*SyncFaceAuthResp, error) {
	var out SyncFaceAuthResp
	_, err := c.Do(ctx, FunCodeSyncFaceAuth, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// QueryTaxReport 报税查询（6047）。
func (c *Client) QueryTaxReport(ctx context.Context, req QueryTaxReportReq) (*QueryTaxReportResp, error) {
	var out QueryTaxReportResp
	_, err := c.Do(ctx, FunCodeQueryTaxReport, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
