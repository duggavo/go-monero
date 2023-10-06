package daemon

import (
	"context"
	"encoding/json"
	"fmt"
)

// GetAlternateChains displays alternative chains seen by the node.
//
// (restricted).
func (c *Client) GetAlternateChains(
	ctx context.Context,
) (*GetAlternateChainsResult, error) {
	resp := &GetAlternateChainsResult{}

	err := c.JSONRPC(ctx, "get_alternate_chains", nil, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// RPCAccessTracking retrieves statistics that the monero daemon keeps track of
// about the use of each RPC method and endpoint.
//
// (restricted).
func (c *Client) RPCAccessTracking(
	ctx context.Context,
) (*RPCAccessTrackingResult, error) {
	resp := &RPCAccessTrackingResult{}

	err := c.JSONRPC(ctx, "rpc_access_tracking", nil, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// HardForkInfo looks up informaiton about the last hard fork.
func (c *Client) HardForkInfo(
	ctx context.Context,
) (*HardForkInfoResult, error) {
	resp := &HardForkInfoResult{}

	err := c.JSONRPC(ctx, "hard_fork_info", nil, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// GetBans retrieves the list of banned IPs.
//
// (restricted).
func (c *Client) GetBans(ctx context.Context) (*GetBansResult, error) {
	resp := &GetBansResult{}

	err := c.JSONRPC(ctx, "get_bans", nil, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

type SetBansBan struct {
	Host    string `json:"host"`
	Ban     bool   `json:"ban"`
	Seconds int64  `json:"seconds"`
}

type SetBansRequestParameters struct {
	Bans []SetBansBan `json:"bans"`
}

// SetBans bans a particular host.
//
// (restricted).
func (c *Client) SetBans(
	ctx context.Context, params SetBansRequestParameters,
) (*SetBansResult, error) {
	resp := &SetBansResult{}

	err := c.JSONRPC(ctx, "set_bans", params, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// GetVersion retrieves the version of monerod that the node uses.
//
// (restricted).
func (c *Client) GetVersion(ctx context.Context) (*GetVersionResult, error) {
	resp := &GetVersionResult{}

	err := c.JSONRPC(ctx, "get_version", nil, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// GenerateBlocksRequestParameters is the set of parameters to be passed to the
// GenerateBlocks RPC method.
type GenerateBlocksRequestParameters struct {
	// AmountOfBlocks is the number of blocks to be generated.
	AmountOfBlocks uint64 `json:"amount_of_blocks,omitempty"`

	// WalletAddress is the address of the wallet that will get the rewards
	// of the coinbase transaction for such the blocks generates.
	WalletAddress string `json:"wallet_address,omitempty"`

	// PreviousBlock TODO
	PreviousBlock string `json:"prev_block,omitempty"`

	// StartingNonce TODO
	StartingNonce uint32 `json:"starting_nonce,omitempty"`
}

// GenerateBlocks combines functionality from `GetBlockTemplate` and
// `SubmitBlock` RPC calls to allow rapid block creation.
//
// Difficulty is set permanently to 1 for regtest.
//
// (restricted).
func (c *Client) GenerateBlocks(
	ctx context.Context, params GenerateBlocksRequestParameters,
) (*GenerateBlocksResult, error) {
	resp := &GenerateBlocksResult{}

	err := c.JSONRPC(ctx, "generateblocks", params, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) GetBlockCount(
	ctx context.Context,
) (*GetBlockCountResult, error) {
	resp := &GetBlockCountResult{}

	err := c.JSONRPC(ctx, "get_block_count", nil, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) OnGetBlockHash(
	ctx context.Context, height uint64,
) (string, error) {
	resp := ""
	params := []uint64{height}

	err := c.JSONRPC(ctx, "on_get_block_hash", params, &resp)
	if err != nil {
		return "", fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) RelayTx(
	ctx context.Context, txns []string,
) (*RelayTxResult, error) {
	resp := &RelayTxResult{}
	params := map[string]interface{}{
		"txids": txns,
	}

	err := c.JSONRPC(ctx, "relay_tx", params, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// GetBlockTemplate gets a block template on which mining a new block.
func (c *Client) GetBlockTemplate(ctx context.Context, params GetBlockTemplateParams) (*GetBlockTemplateResult, error) {
	resp := &GetBlockTemplateResult{}

	err := c.JSONRPC(ctx, "get_block_template", params, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) GetConnections(
	ctx context.Context,
) (*GetConnectionsResult, error) {
	resp := &GetConnectionsResult{}

	err := c.JSONRPC(ctx, "get_connections", nil, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// GetInfo retrieves general information about the state of the node and the
// network.
func (c *Client) GetInfo(ctx context.Context) (*GetInfoResult, error) {
	resp := &GetInfoResult{}

	if err := c.JSONRPC(ctx, "get_info", nil, resp); err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) GetLastBlockHeader(
	ctx context.Context,
) (*GetLastBlockHeaderResult, error) {
	resp := &GetLastBlockHeaderResult{}

	err := c.JSONRPC(ctx, "get_last_block_header", nil, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) GetCoinbaseTxSum(
	ctx context.Context, height, count uint64,
) (*GetCoinbaseTxSumResult, error) {
	resp := &GetCoinbaseTxSumResult{}
	params := map[string]uint64{
		"height": height,
		"count":  count,
	}

	err := c.JSONRPC(ctx, "get_coinbase_tx_sum", params, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// InnerJSON parses the content of the JSON embedded in `GetBlockResult`.
func (j *GetBlockResult) InnerJSON() (*GetBlockResultJSON, error) {
	res := &GetBlockResultJSON{}

	err := json.Unmarshal([]byte(j.JSON), res)
	if err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	return res, nil
}

func (c *Client) GetBlockHeadersRange(
	ctx context.Context, params GetBlockHeadersRangeParameters,
) (*GetBlockHeadersRangeResult, error) {
	resp := &GetBlockHeadersRangeResult{}

	err := c.JSONRPC(ctx, "get_block_headers_range", params, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// GetBlockHeaderByHeight retrieves block header information for either one or
// multiple blocks.
func (c *Client) GetBlockHeaderByHeight(
	ctx context.Context, height uint64,
) (*GetBlockHeaderByHeightResult, error) {
	resp := &GetBlockHeaderByHeightResult{}
	params := map[string]interface{}{
		"height": height,
	}

	err := c.JSONRPC(ctx, "get_block_header_by_height", params, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// GetBlockHeaderByHash retrieves block header information for either one or
// multiple blocks.
func (c *Client) GetBlockHeaderByHash(
	ctx context.Context, hashes []string,
) (*GetBlockHeaderByHashResult, error) {
	resp := &GetBlockHeaderByHashResult{}
	params := map[string]interface{}{
		"hashes": hashes,
	}

	err := c.JSONRPC(ctx, "get_block_header_by_hash", params, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// GetBlock fetches full block information from a block at a particular hash OR
// height.
func (c *Client) GetBlock(
	ctx context.Context, params GetBlockRequestParameters,
) (*GetBlockResult, error) {
	resp := &GetBlockResult{}

	err := c.JSONRPC(ctx, "get_block", params, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// GetFeeEstimate gives an estimation on fees per byte.
func (c *Client) GetFeeEstimate(ctx context.Context, graceBlocks uint64) (*GetFeeEstimateResult, error) {
	resp := &GetFeeEstimateResult{}
	params := map[string]uint64{
		"grace_blocks": graceBlocks,
	}

	err := c.JSONRPC(ctx, "get_fee_estimate", params, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// SyncInfo gets synchronisation informations.
func (c *Client) SyncInfo(ctx context.Context) (*SyncInfoResult, error) {
	resp := &SyncInfoResult{}

	err := c.JSONRPC(ctx, "sync_info", nil, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// CalcPow calculates PoW hash for a block candidate.
func (c *Client) CalcPow(ctx context.Context, params CalcPowParameters) (string, error) {
	var resp string

	err := c.JSONRPC(ctx, "calc_pow", params, &resp)
	if err != nil {
		return "", fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

// SubmitBlock submits a mined block to the network.
func (c *Client) SubmitBlock(ctx context.Context, minedBlockBlob string) (*SubmitBlockResult, error) {
	resp := &SubmitBlockResult{}

	err := c.JSONRPC(ctx, "submit_block", []string{minedBlockBlob}, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}

func (c *Client) GetMinerData(ctx context.Context) (*GetMinerDataResult, error) {
	resp := &GetMinerDataResult{}

	err := c.JSONRPC(ctx, "get_miner_data", nil, resp)
	if err != nil {
		return nil, fmt.Errorf("jsonrpc: %w", err)
	}

	return resp, nil
}
