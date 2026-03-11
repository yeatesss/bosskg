package bosskg

import "context"

func (c *Client) QueryMerIDBalance(ctx context.Context, req QueryMerIDBalanceReq) (*QueryMerIDBalanceResp, error) {
	var out QueryMerIDBalanceResp
	_, err := c.Do(ctx, FunCodeQueryMerIDBalance, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) BatchTransfer(ctx context.Context, req BatchTransferReq) (*BatchTransferResp, error) {
	var out BatchTransferResp
	_, err := c.Do(ctx, FunCodeBatchTransfer, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) QueryBatchTransferResult(ctx context.Context, req QueryBatchTransferResultReq) (*QueryBatchTransferResultResp, error) {
	var out QueryBatchTransferResultResp
	_, err := c.Do(ctx, FunCodeQueryBatchTransferResult, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) ProfessionalSign(ctx context.Context, req ProfessionalSignReq) (*ProfessionalSignResp, error) {
	var out ProfessionalSignResp
	_, err := c.Do(ctx, FunCodeProfessionalSign, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) QueryProfessionalSignResult(ctx context.Context, req QueryProfessionalSignResultReq) (*QueryProfessionalSignResultResp, error) {
	var out QueryProfessionalSignResultResp
	_, err := c.Do(ctx, FunCodeQueryProfessionalSignResult, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) ProfessionalUnSign(ctx context.Context, req ProfessionalUnSignReq) (*ProfessionalUnSignResp, error) {
	var out ProfessionalUnSignResp
	_, err := c.Do(ctx, FunCodeProfessionalUnSign, req, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
