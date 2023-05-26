package wallet

import (
	"context"
	"fmt"
)

func (c *Client) GetAccounts(
	ctx context.Context, params GetAccountsRequestParameters,
) (*GetAccountsResult, error) {
	resp := &GetAccountsResult{}

	if err := c.JSONRPC(ctx, "get_accounts", params, resp); err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) GetAddress(
	ctx context.Context, params GetAddressRequestParameters,
) (*GetAddressResult, error) {
	resp := &GetAddressResult{}

	if err := c.JSONRPC(ctx, "get_address", params, resp); err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// GetAddressIndex returns the account and index of a given subaddress
func (c *Client) GetAddressIndex(ctx context.Context, addr string) (*GetAddressIndexResult, error) {
	resp := &GetAddressIndexResult{}

	if err := c.JSONRPC(ctx, "get_address_index", map[string]string{
		"address": addr,
	}, resp); err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// GetBalance gets the balance of the wallet.
func (c *Client) GetBalance(
	ctx context.Context, params GetBalanceRequestParameters,
) (*GetBalanceResult, error) {
	resp := &GetBalanceResult{}

	if err := c.JSONRPC(ctx, "get_balance", params, resp); err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) CreateAddress(
	ctx context.Context, accountIndex uint, count uint, label string,
) (*CreateAddressResult, error) {
	resp := &CreateAddressResult{}

	params := map[string]interface{}{
		"account_index": accountIndex,
		"label":         label,
		"count":         count,
	}
	if err := c.JSONRPC(ctx, "create_address", params, resp); err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) AutoRefresh(
	ctx context.Context, enable bool, period int64,
) (*AutoRefreshResult, error) {
	resp := &AutoRefreshResult{}

	params := map[string]interface{}{
		"enable": enable,
		"period": period,
	}
	if err := c.JSONRPC(ctx, "auto_refresh", params, resp); err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) Refresh(
	ctx context.Context, startHeight uint64,
) (*RefreshResult, error) {
	resp := &RefreshResult{}

	params := map[string]interface{}{
		"start_height": startHeight,
	}
	if err := c.JSONRPC(ctx, "refresh", params, resp); err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) GetHeight(ctx context.Context) (*GetHeightResult, error) {
	resp := &GetHeightResult{}

	if err := c.JSONRPC(ctx, "get_height", nil, resp); err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) Transfer(ctx context.Context, params TransferParameters) (*TransferResult, error) {
	resp := &TransferResult{}

	if err := c.JSONRPC(ctx, "transfer", params, resp); err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// Same as Transfer, but can split into more than one tx if necessary
func (c *Client) TransferSplit(ctx context.Context, params TransferParameters) (*TransferSplitResult, error) {
	resp := &TransferSplitResult{}

	if err := c.JSONRPC(ctx, "transfer_split", params, resp); err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) IncomingTransfers(ctx context.Context, params IncomingTransfersParams) (*IncomingTransfersResult, error) {
	resp := &IncomingTransfersResult{}

	if err := c.JSONRPC(ctx, "incoming_transfers", params, resp); err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) SweepAll(ctx context.Context, params SweepAllParams) (*SweepAllResult, error) {
	resp := &SweepAllResult{}

	if err := c.JSONRPC(ctx, "sweep_all", params, resp); err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil

}
