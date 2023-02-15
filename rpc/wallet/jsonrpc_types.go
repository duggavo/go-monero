package wallet

type GetAccountsRequestParameters struct {
	Tag            string `json:"tag,omitempty"`
	StrictBalances bool   `json:"strict_balances,omitempty"`
}

type GetAccountsResult struct {
	SubaddressAccounts []struct {
		AccountIndex    uint   `json:"account_index"` // Index of the account.
		Balance         uint64 `json:"balance"`       // Balance for the account (locked + unlocked).
		BaseAddress     string `json:"base_address"`  // Main address of the account.
		Label           string `json:"label"`         // Label of the account.
		Tag             string `json:"tag"`
		UnlockedBalance uint64 `json:"unlocked_balance"` // Balance which can be spent.
	} `json:"subaddress_accounts"`

	TotalBalance         uint64 `json:"total_balance"`          // Total balance of the wallet (locked + unlocked)
	TotalUnlockedBalance uint64 `json:"total_unlocked_balance"` // Total balance which can be spent.
}

type GetAddressRequestParameters struct {
	AccountIndex   uint   `json:"account_index"`
	AddressIndices []uint `json:"address_indices"`
}

type GetAddressResult struct {
	Address   string `json:"address"`
	Addresses []struct {
		Address      string `json:"address"`
		AddressIndex uint   `json:"address_index"`
		Label        string `json:"label"`
		Used         bool   `json:"used"`
	} `json:"addresses"`
}

type GetBalanceRequestParameters struct {
	AccountIndex   uint   `json:"account_index"`
	AddressIndices []uint `json:"address_indices"`
	AllAccounts    bool   `json:"all_accounts"`
	Strict         bool   `json:"strict"`
}

type GetBalanceResult struct {
	Balance              uint64       `json:"balance"`                // Balance of the wallet (locked + unlocked).
	MultisigImportNeeded bool         `json:"multisig_import_needed"` // True if importing multisig data is needed for returning a correct balance
	PerSubaddress        []SubAddress `json:"per_subaddress"`         // Balance information for each subaddress.
	TimeToUnlock         int          `json:"time_to_unlock"`         // Time (in seconds) before balance is safe to spend.
	BlocksToUnlock       uint         `json:"blocks_to_unlock"`       // Number of blocks before balance is safe to spend.
	UnlockedBalance      int64        `json:"unlocked_balance"`       // Balance which can be spent.
}

type SubAddress struct {
	AccountIndex      uint   `json:"account_index"`       // Index of the account.
	Address           string `json:"address"`             // Textual representation of the subaddress.
	AddressIndex      uint   `json:"address_index"`       // Index of the subaddress in the account
	Balance           uint64 `json:"balance"`             // Balance for the subaddress (locked + unlocked).
	Label             string `json:"label"`               // Label of the subaddress.
	NumUnspentOutputs uint   `json:"num_unspent_outputs"` // Number of unspent outputs available for the subaddress.
	TimeToUnlock      uint   `json:"time_to_unlock"`      // Time (in seconds) before balance is safe to spend.
	BlocksToUnlock    uint   `json:"blocks_to_unlock"`    // Number of blocks before balance is safe to spend.
	UnlockedBalance   int64  `json:"unlocked_balance"`    // Balance which can be spent.
}

type CreateAddressResult struct {
	Address        string   `json:"address"`
	AddressIndex   uint     `json:"address_index"`
	AddressIndices []uint   `json:"address_indices"`
	Addresses      []string `json:"addresses"`
}

type RefreshResult struct {
	BlocksFetched uint64 `json:"blocks_fetched"`
	ReceivedMoney bool   `json:"received_money"`
}

type AutoRefreshResult struct {
}

type GetHeightResult struct {
	Height uint64 `json:"height"`
}

type Destination struct {
	Amount  uint64 `json:"amount"`
	Address string `json:"address"`
}

type TransferParameters struct {
	Destinations   []Destination `json:"destinations"`
	AccountIndex   uint          `json:"account_index,omitempty"`
	SubaddrIndices []uint        `json:"subaddr_indices,omitempty"`
	Priority       uint          `json:"priority,omitempty"`
	UnlockTime     uint          `json:"unlock_time,omitempty"`
	GetTxKey       bool          `json:"get_tx_key,omitempty"`
	DoNotRelay     bool          `json:"do_not_relay,omitempty"`
	GetTxHex       bool          `json:"get_tx_hex,omitempty"`
	GetTxMetadata  bool          `json:"get_tx_metadata,omitempty"`
}

type TransferResult struct {
	Amount        uint64 `json:"amount"`
	Fee           uint64 `json:"fee"`
	MultisigTxset string `json:"multisig_txset"`
	TxBlob        string `json:"tx_blob"`
	TxHash        string `json:"tx_hash"`
	TxKey         string `json:"tx_key"`
	TxMetadata    string `json:"tx_metadata"`
	UnsignedTxset string `json:"unsigned_txset"`
}

type SubaddrIndices struct {
	Major uint `json:"major"` // Account index for the subaddress.
	Minor uint `json:"minor"` // Index of the subaddress in the account.
}
type Transfer struct {
	Amount       uint64         `json:"amount"`    // Amount of this transfer.
	KeyImage     string         `json:"key_image"` // Key image for the incoming transfer's unspent output.
	Spent        bool           `json:"spent"`     // Indicates if this transfer has been spent.
	SubaddrIndex SubaddrIndices `json:"subaddr_index"`
	TxHash       string         `json:"tx_hash"`  // Several incoming transfers may share the same hash if they were in the same transaction.
	Frozen       bool           `json:"frozen"`   // Indicates if the output been frozen by "freeze".
	Unlocked     bool           `json:"unlocked"` // Indicates if the output is spendable.
	BlockHeight  uint64         `json:"block_height"`
	PubKey       string         `json:"pubkey"`
}

type IncomingTransfersParams struct {
	TransferType   string `json:"transfer_type"`
	AccountIndex   uint   `json:"account_index,omitempty"`
	SubaddrIndices []uint `json:"subaddr_indices,omitempty"`
}

type IncomingTransfersResult struct {
	Transfers []Transfer `json:"transfers"`
}
