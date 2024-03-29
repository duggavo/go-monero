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
type GetAddressIndexResult struct {
	Index struct {
		Major uint `json:"major"` // account index
		Minor uint `json:"minor"` // subaddress index
	} `json:"index"`
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
type TransferSplitResult struct {
	TxHashList     []string `json:"tx_hash_list"`
	TxKeyList      []string `json:"tx_key_list"`
	AmountList     []uint64 `json:"amount_list"`
	FeeList        []uint64 `json:"fee_list"`
	WeightList     []uint64 `json:"weight_list"`
	TxBlobList     []string `json:"tx_blob_list"`
	TxMetadataList []string `json:"tx_metadata_list"`
	MultisigTxset  string   `json:"multisig_txset"`
	UnsignedTxset  string   `json:"unsigned_txset"`
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

type SweepAllParams struct {
	Address        string `json:"address"`                   // Destination public address.
	AccountIndex   uint   `json:"account_index"`             // Sweep transactions from this account.
	SubaddrIndices []uint `json:"subaddr_indices,omitempty"` // Sweep from this set of subaddresses in the account.
	Priority       uint8  `json:"priority,omitempty"`        // Priority for sending the sweep transfer, partially determines fee.
	Outputs        uint   `json:"outputs,omitempty"`         // Specify the number of separate outputs that will be created.
	UnlockTime     uint64 `json:"unlock_time,omitempty"`     // Number of blocks before the coins can be spent.
	PaymentId      string `json:"payment_id,omitempty"`      // The 16-bytes payment ID encoded as hex.
	GetTxKeys      bool   `json:"get_tx_keys,omitempty"`     // Return the transaction keys after sending.
	DoNotRelay     bool   `json:"do_not_relay,omitempty"`    // If true, do not relay this sweep transfer.
	GetTxHex       bool   `json:"get_tx_hex,omitempty"`      // Return the transaction as hex after sending.
	GetTxMetadata  bool   `json:"get_tx_metadata,omitempty"` // Return the transaction metadata after sending.
}

type SweepAllResult struct {
	TxHashList []string `json:"tx_hash_list"`
	TxKeyList  []string `json:"tx_key_list"`
	AmountList []uint64 `json:"amount_list"`
	FeeList    []uint64 `json:"fee_list"`
	WeightList []uint64 `json:"weight_list"`
	TxBlobList []string `json:"tx_blob_list"`
}

type RelayTxResult struct {
	TxHash string `json:"tx_hash"`
}

type CreateWalletParams struct {
	Filename string `json:"filename"`
	Password string `json:"password"`
	Language string `json:"language"`
}

type OpenWalletParams struct {
	Filename string `json:"filename"`
	Password string `json:"password"`
}

type RestoreDeterministicWalletParams struct {
	Filename        string `json:"filename"`                   // Name of the wallet.
	Password        string `json:"password"`                   // Password of the wallet.
	Seed            string `json:"seed"`                       // Mnemonic phrase of the wallet to restore.
	RestoreHeight   uint64 `json:"restore_height,omitempty"`   // Block height to restore the wallet from (default = 0).
	Language        string `json:"language,omitempty"`         // Language of the mnemonic phrase in case the old language is invalid.
	SeedOffset      string `json:"seed_offset,omitempty"`      // Offset used to derive a new seed from the given mnemonic to recover a secret wallet from the mnemonic phrase.
	AutosaveCurrent bool   `json:"autosave_current,omitempty"` // Whether to save the currently open RPC wallet before closing it (defaults to true).
}

type RestoreDeterministicWalletResult struct {
	Address       string `json:"address"`
	Info          string `json:"info"`           // Message describing the success or failure of the attempt to restore the wallet.
	Seed          string `json:"seed"`           // Mnemonic phrase of the restored wallet, which is updated if the wallet was restored from a deprecated-style mnemonic phrase.
	WasDeprecated bool   `json:"was_deprecated"` // Indicates if the restored wallet was created from a deprecated mnemonic phrase.
}

type ChangeWalletPasswordParams struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type IsMultisigResult struct {
	Multisig  bool `json:"multisig"`
	Ready     bool `json:"ready"`
	Threshold uint `json:"threshold"`
	Total     uint `json:"total"`
}

type TransferInfo struct {
	Address                         string           `json:"address"`
	Amount                          uint64           `json:"amount"`
	Amounts                         []uint64         `json:"amounts"`
	Confirmations                   uint64           `json:"confirmations"`
	DoubleSpendSeen                 bool             `json:"double_spend_seen"`
	Fee                             uint64           `json:"fee"`
	Height                          uint64           `json:"height"`
	Note                            string           `json:"note"`
	Destinations                    []Destination    `json:"destinations"`
	PaymentId                       string           `json:"payment_id"`
	SubaddrIndex                    SubaddrIndices   `json:"subaddress_index"`
	SubaddrIndices                  []SubaddrIndices `json:"subaddr_indices"`
	SuggestedConfirmationsThreshold uint64           `json:"suggested_confirmations_threshold"`
	Timestamp                       uint64           `json:"timestamp"`
	Txid                            string           `json:"txid"`
	Type                            string           `json:"type"`
	UnlockTime                      uint64           `json:"unlock_time"`
	Locked                          bool             `json:"locked"`
}

type GetTransfersParams struct {
	In      bool `json:"in,omitempty"`      // Include incoming transfers.
	Out     bool `json:"out,omitempty"`     // Include outgoing transfers.
	Pending bool `json:"pending,omitempty"` // Include pending transfers.
	Failed  bool `json:"failed,omitempty"`  // Include failed transfers.
	Pool    bool `json:"pool,omitempty"`    // Include transfers from the daemon's transaction pool.

	FilterByHeight bool   `json:"filter_by_height,omitempty"`
	MinHeight      uint64 `json:"min_height,omitempty"` // Minimum block height to scan for transfers, if filtering by height is enabled.
	MaxHeight      uint64 `json:"max_height,omitempty"` // Maximum block height to scan for transfers, if filtering by height is enabled (defaults to max block height).

	AccountIndex   uint   `json:"account_index,omitempty"`   // Index of the account to query for transfers. (defaults to 0)
	SubaddrIndices []uint `json:"subaddr_indices,omitempty"` // Defaults to empty - all indices
	AllAccounts    bool   `json:"all_accounts,omitempty"`    // Defaults to false
}

type GetTransfersResult struct {
	In      []TransferInfo
	Out     []TransferInfo
	Pending []TransferInfo
	Failed  []TransferInfo
	Pool    []TransferInfo
}
