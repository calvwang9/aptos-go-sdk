package api

import (
	"encoding/json"
	"fmt"
	"github.com/aptos-labs/aptos-go-sdk/internal/types"
)

// TransactionVariant is the type of transaction, all transactions submitted by this SDK are [TransactionVariantUserTransaction]
type TransactionVariant string

const (
	TransactionVariantPendingTransaction         TransactionVariant = "pending_transaction"
	TransactionVariantUserTransaction            TransactionVariant = "user_transaction"
	TransactionVariantGenesisTransaction         TransactionVariant = "genesis_transaction"
	TransactionVariantBlockMetadataTransaction   TransactionVariant = "block_metadata_transaction"
	TransactionVariantStateCheckpointTransaction TransactionVariant = "state_checkpoint_transaction"
	TransactionVariantValidatorTransaction       TransactionVariant = "validator_transaction"
)

// Transaction is an enum type for all possible transactions on the blockchain
type Transaction struct {
	Type  TransactionVariant
	Inner TransactionImpl
}

// Hash of the transaction for lookup on-chain
func (o *Transaction) Hash() Hash {
	return o.Inner.TxnHash()
}

// Success of the transaction.  Pending transactions, and genesis may not have a success field.
// If this is the case, it will be nil
func (o *Transaction) Success() *bool {
	return o.Inner.TxnSuccess()
}

// Version of the transaction on chain, will be nil if it is a PendingTransaction
func (o *Transaction) Version() *uint64 {
	return o.Inner.TxnVersion()
}

func (o *Transaction) UnmarshalJSON(b []byte) error {
	type inner struct {
		Type string `json:"type"`
	}
	data := &inner{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	o.Type = TransactionVariant(data.Type)
	switch o.Type {
	case TransactionVariantPendingTransaction:
		o.Inner = &PendingTransaction{}
	case TransactionVariantUserTransaction:
		o.Inner = &UserTransaction{}
	case TransactionVariantGenesisTransaction:
		o.Inner = &GenesisTransaction{}
	case TransactionVariantBlockMetadataTransaction:
		o.Inner = &BlockMetadataTransaction{}
	case TransactionVariantStateCheckpointTransaction:
		o.Inner = &StateCheckpointTransaction{}
	case TransactionVariantValidatorTransaction:
		o.Inner = &ValidatorTransaction{}
	default:
		return fmt.Errorf("unknown transaction type: %s", o.Type)
	}
	return json.Unmarshal(b, o.Inner)
}

// UserTransaction changes the transaction to a [UserTransaction]; however, it will fail if it's not one.
func (o *Transaction) UserTransaction() (*UserTransaction, error) {
	if o.Type == TransactionVariantUserTransaction {
		return o.Inner.(*UserTransaction), nil
	}
	return nil, fmt.Errorf("transaction type is not user: %s", o.Type)
}

// PendingTransaction changes the transaction to a [PendingTransaction]; however, it will fail if it's not one.
func (o *Transaction) PendingTransaction() (*PendingTransaction, error) {
	if o.Type == TransactionVariantPendingTransaction {
		return o.Inner.(*PendingTransaction), nil
	}
	return nil, fmt.Errorf("transaction type is not pending: %s", o.Type)
}

// GenesisTransaction changes the transaction to a [GenesisTransaction]; however, it will fail if it's not one.
func (o *Transaction) GenesisTransaction() (*GenesisTransaction, error) {
	if o.Type == TransactionVariantGenesisTransaction {
		return o.Inner.(*GenesisTransaction), nil
	}
	return nil, fmt.Errorf("transaction type is not genesis: %s", o.Type)
}

// BlockMetadataTransaction changes the transaction to a [BlockMetadataTransaction]; however, it will fail if it's not one.
func (o *Transaction) BlockMetadataTransaction() (*BlockMetadataTransaction, error) {
	if o.Type == TransactionVariantBlockMetadataTransaction {
		return o.Inner.(*BlockMetadataTransaction), nil
	}
	return nil, fmt.Errorf("transaction type is not block metadata: %s", o.Type)
}

// StateCheckpointTransaction changes the transaction to a [StateCheckpointTransaction]; however, it will fail if it's not one.
func (o *Transaction) StateCheckpointTransaction() (*StateCheckpointTransaction, error) {
	if o.Type == TransactionVariantStateCheckpointTransaction {
		return o.Inner.(*StateCheckpointTransaction), nil
	}
	return nil, fmt.Errorf("transaction type is not state checkpoint: %s", o.Type)
}

// ValidatorTransaction changes the transaction to a [ValidatorTransaction]; however, it will fail if it's not one.
func (o *Transaction) ValidatorTransaction() (*ValidatorTransaction, error) {
	if o.Type == TransactionVariantValidatorTransaction {
		return o.Inner.(*ValidatorTransaction), nil
	}
	return nil, fmt.Errorf("transaction type is not validator: %s", o.Type)
}

// TransactionImpl is an interface for all transactions
type TransactionImpl interface {
	// TxnSuccess tells us if the transaction is a success.  It will be nil if it doesn't apply.
	TxnSuccess() *bool

	// TxnHash gives us the hash of the transaction, this should always apply.
	TxnHash() Hash

	// TxnVersion gives us the ledger version of the transaction.
	TxnVersion() *uint64
}

// UserTransaction is a user submitted transaction as an entry function or script
type UserTransaction struct {
	Version                 uint64
	Hash                    Hash
	AccumulatorRootHash     Hash
	StateChangeHash         Hash
	EventRootHash           Hash
	GasUsed                 uint64
	Success                 bool
	VmStatus                string
	Changes                 []*WriteSetChange
	Events                  []*Event
	Sender                  *types.AccountAddress
	SequenceNumber          uint64
	MaxGasAmount            uint64
	GasUnitPrice            uint64
	ExpirationTimestampSecs uint64
	Payload                 *TransactionPayload
	Signature               *Signature
	Timestamp               uint64 // TODO: native time?
	StateCheckpointHash     Hash   // Optional
}

func (o *UserTransaction) TxnHash() Hash {
	return o.Hash
}
func (o *UserTransaction) TxnSuccess() *bool {
	return &o.Success
}
func (o *UserTransaction) TxnVersion() *uint64 {
	return &o.Version
}

func (o *UserTransaction) UnmarshalJSON(b []byte) error {
	type inner struct {
		Version                 U64                   `json:"version"`
		Hash                    Hash                  `json:"hash"`
		AccumulatorRootHash     Hash                  `json:"accumulator_root_hash"`
		StateChangeHash         Hash                  `json:"state_change_hash"`
		EventRootHash           Hash                  `json:"event_root_hash"`
		GasUsed                 U64                   `json:"gas_used"`
		Success                 bool                  `json:"success"`
		VmStatus                string                `json:"vm_status"`
		Changes                 []*WriteSetChange     `json:"changes"`
		Events                  []*Event              `json:"events"`
		Sender                  *types.AccountAddress `json:"sender"`
		SequenceNumber          U64                   `json:"sequence_number"`
		MaxGasAmount            U64                   `json:"max_gas_amount"`
		GasUnitPrice            U64                   `json:"gas_unit_price"`
		ExpirationTimestampSecs U64                   `json:"expiration_timestamp_secs"`
		Payload                 *TransactionPayload   `json:"payload"`
		Signature               *Signature            `json:"signature"`
		Timestamp               U64                   `json:"timestamp"`
		StateCheckpointHash     Hash                  `json:"state_checkpoint_hash"` // Optional
	}
	data := &inner{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	o.Version = data.Version.toUint64()
	o.Hash = data.Hash
	o.AccumulatorRootHash = data.AccumulatorRootHash
	o.StateChangeHash = data.StateChangeHash
	o.EventRootHash = data.EventRootHash
	o.GasUsed = data.GasUsed.toUint64()
	o.Success = data.Success
	o.VmStatus = data.VmStatus
	o.Changes = data.Changes
	o.Events = data.Events
	o.Sender = data.Sender
	o.SequenceNumber = data.SequenceNumber.toUint64()
	o.MaxGasAmount = data.MaxGasAmount.toUint64()
	o.GasUnitPrice = data.GasUnitPrice.toUint64()
	o.ExpirationTimestampSecs = data.ExpirationTimestampSecs.toUint64()
	o.Payload = data.Payload
	o.Signature = data.Signature
	o.Timestamp = data.Timestamp.toUint64()
	o.StateCheckpointHash = data.StateCheckpointHash
	return nil
}

type PendingTransaction struct {
	Hash                    string
	Sender                  *types.AccountAddress
	SequenceNumber          uint64
	MaxGasAmount            uint64
	GasUnitPrice            uint64
	ExpirationTimestampSecs uint64
	Payload                 *TransactionPayload
	Signature               *Signature
}

func (o *PendingTransaction) TxnHash() Hash {
	return o.Hash
}
func (o *PendingTransaction) TxnSuccess() *bool {
	return nil
}
func (o *PendingTransaction) TxnVersion() *uint64 {
	return nil
}

func (o *PendingTransaction) UnmarshalJSON(b []byte) error {
	type inner struct {
		Hash                    Hash                  `json:"hash"`
		Sender                  *types.AccountAddress `json:"sender"`
		SequenceNumber          U64                   `json:"sequence_number"`
		MaxGasAmount            U64                   `json:"max_gas_amount"`
		GasUnitPrice            U64                   `json:"gas_unit_price"`
		ExpirationTimestampSecs U64                   `json:"expiration_timestamp_secs"`
		Payload                 *TransactionPayload   `json:"payload"`
		Signature               *Signature            `json:"signature"`
	}
	data := &inner{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	o.Hash = data.Hash
	o.Sender = data.Sender
	o.SequenceNumber = data.SequenceNumber.toUint64()
	o.MaxGasAmount = data.MaxGasAmount.toUint64()
	o.GasUnitPrice = data.GasUnitPrice.toUint64()
	o.ExpirationTimestampSecs = data.ExpirationTimestampSecs.toUint64()
	o.Payload = data.Payload
	o.Signature = data.Signature
	return nil
}

type GenesisTransaction struct {
	Version             uint64
	Hash                Hash
	AccumulatorRootHash Hash
	StateChangeHash     Hash
	EventRootHash       Hash
	GasUsed             uint64
	Success             bool
	VmStatus            string
	Changes             []*WriteSetChange
	Events              []*Event
	Payload             *TransactionPayload
	StateCheckpointHash Hash // Optional
}

func (o *GenesisTransaction) TxnHash() Hash {
	return o.Hash
}
func (o *GenesisTransaction) TxnSuccess() *bool {
	return &o.Success
}
func (o *GenesisTransaction) TxnVersion() *uint64 {
	return &o.Version
}

func (o *GenesisTransaction) UnmarshalJSON(b []byte) error {
	type inner struct {
		Version             U64                 `json:"version"`
		Hash                Hash                `json:"hash"`
		AccumulatorRootHash Hash                `json:"accumulator_root_hash"`
		StateChangeHash     Hash                `json:"state_change_hash"`
		EventRootHash       Hash                `json:"event_root_hash"`
		GasUsed             U64                 `json:"gas_used"`
		Success             bool                `json:"success"`
		VmStatus            string              `json:"vm_status"`
		Changes             []*WriteSetChange   `json:"changes"`
		Events              []*Event            `json:"events"`
		Payload             *TransactionPayload `json:"payload"`
		StateCheckpointHash Hash                `json:"state_checkpoint_hash"` // Optional
	}
	data := &inner{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	o.Version = data.Version.toUint64()
	o.Hash = data.Hash
	o.AccumulatorRootHash = data.AccumulatorRootHash
	o.StateChangeHash = data.StateChangeHash
	o.EventRootHash = data.EventRootHash
	o.GasUsed = data.GasUsed.toUint64()
	o.Success = data.Success
	o.VmStatus = data.VmStatus
	o.Changes = data.Changes
	o.Events = data.Events
	o.Payload = data.Payload
	o.StateCheckpointHash = data.StateCheckpointHash
	return nil
}

type BlockMetadataTransaction struct {
	Id                       string
	Epoch                    uint64
	Round                    uint64
	PreviousBlockVotesBitvec []uint8
	Proposer                 *types.AccountAddress
	FailedProposerIndices    []uint32
	Version                  uint64
	Hash                     string
	AccumulatorRootHash      Hash
	StateChangeHash          Hash
	EventRootHash            Hash
	GasUsed                  uint64
	Success                  bool
	VmStatus                 string
	Changes                  []*WriteSetChange
	Events                   []*Event
	Timestamp                uint64
	StateCheckpointHash      Hash
}

func (o *BlockMetadataTransaction) TxnHash() Hash {
	return o.Hash
}
func (o *BlockMetadataTransaction) TxnSuccess() *bool {
	return &o.Success
}
func (o *BlockMetadataTransaction) TxnVersion() *uint64 {
	return &o.Version
}

func (o *BlockMetadataTransaction) UnmarshalJSON(b []byte) error {
	type inner struct {
		Id                       string                `json:"id"`
		Epoch                    U64                   `json:"epoch"`
		Round                    U64                   `json:"round"`
		PreviousBlockVotesBitvec []byte                `json:"previous_block_votes_bitvec"` // TODO: this had to be float64 earlier
		Proposer                 *types.AccountAddress `json:"proposer"`
		FailedProposerIndices    []uint32              `json:"failed_proposer_indices"` // TODO: verify
		Version                  U64                   `json:"version"`
		Hash                     Hash                  `json:"hash"`
		AccumulatorRootHash      Hash                  `json:"accumulator_root_hash"`
		StateChangeHash          Hash                  `json:"state_change_hash"`
		EventRootHash            Hash                  `json:"event_root_hash"`
		GasUsed                  U64                   `json:"gas_used"`
		Success                  bool                  `json:"success"`
		VmStatus                 string                `json:"vm_status"`
		Changes                  []*WriteSetChange     `json:"changes"`
		Events                   []*Event              `json:"events"`
		Timestamp                U64                   `json:"timestamp"`
		StateCheckpointHash      Hash                  `json:"state_checkpoint_hash,omitempty"` // Optional
	}
	data := &inner{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	o.Id = data.Id
	o.Epoch = data.Epoch.toUint64()
	o.Round = data.Round.toUint64()
	o.PreviousBlockVotesBitvec = data.PreviousBlockVotesBitvec
	o.Proposer = data.Proposer
	o.FailedProposerIndices = data.FailedProposerIndices
	o.Version = data.Version.toUint64()
	o.Hash = data.Hash
	o.AccumulatorRootHash = data.AccumulatorRootHash
	o.StateChangeHash = data.StateChangeHash
	o.EventRootHash = data.EventRootHash
	o.GasUsed = data.GasUsed.toUint64()
	o.Success = data.Success
	o.VmStatus = data.VmStatus
	o.Changes = data.Changes
	o.Events = data.Events
	o.Timestamp = data.Timestamp.toUint64()
	o.StateCheckpointHash = data.StateCheckpointHash
	return nil
}

type StateCheckpointTransaction struct {
	Version             uint64
	Hash                Hash
	AccumulatorRootHash Hash
	StateChangeHash     Hash
	EventRootHash       Hash
	GasUsed             uint64
	Success             bool
	VmStatus            string
	Changes             []*WriteSetChange
	Timestamp           uint64
	StateCheckpointHash Hash // This is optional
}

func (o *StateCheckpointTransaction) TxnHash() Hash {
	return o.Hash
}
func (o *StateCheckpointTransaction) TxnSuccess() *bool {
	return &o.Success
}
func (o *StateCheckpointTransaction) TxnVersion() *uint64 {
	return &o.Version
}

func (o *StateCheckpointTransaction) UnmarshalJSON(b []byte) error {
	type inner struct {
		Version             U64               `json:"version"`
		Hash                Hash              `json:"hash"`
		AccumulatorRootHash Hash              `json:"accumulator_root_hash"`
		StateChangeHash     Hash              `json:"state_change_hash"`
		EventRootHash       Hash              `json:"event_root_hash"`
		GasUsed             U64               `json:"gas_used"`
		Success             bool              `json:"success"`
		VmStatus            string            `json:"vm_status"`
		Changes             []*WriteSetChange `json:"changes"`
		Timestamp           U64               `json:"timestamp"`
		StateCheckpointHash Hash              `json:"state_checkpoint_hash"` // Optional
	}
	data := &inner{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	o.Version = data.Version.toUint64()
	o.Hash = data.Hash
	o.AccumulatorRootHash = data.AccumulatorRootHash
	o.StateChangeHash = data.StateChangeHash
	o.EventRootHash = data.EventRootHash
	o.GasUsed = data.GasUsed.toUint64()
	o.Success = data.Success
	o.VmStatus = data.VmStatus
	o.Changes = data.Changes
	o.Timestamp = data.Timestamp.toUint64()
	o.StateCheckpointHash = data.StateCheckpointHash
	return nil
}

type ValidatorTransaction struct {
	Version             uint64
	Hash                Hash
	AccumulatorRootHash Hash
	StateChangeHash     Hash
	EventRootHash       Hash
	GasUsed             uint64
	Success             bool
	VmStatus            string
	Changes             []*WriteSetChange
	Events              []*Event
	Timestamp           uint64
	StateCheckpointHash Hash // This is optional
}

func (o *ValidatorTransaction) TxnHash() Hash {
	return o.Hash
}
func (o *ValidatorTransaction) TxnSuccess() *bool {
	return &o.Success
}
func (o *ValidatorTransaction) TxnVersion() *uint64 {
	return &o.Version
}

func (o *ValidatorTransaction) UnmarshalJSON(b []byte) error {
	type inner struct {
		Version             U64               `json:"version"`
		Hash                Hash              `json:"hash"`
		AccumulatorRootHash Hash              `json:"accumulator_root_hash"`
		StateChangeHash     Hash              `json:"state_change_hash"`
		EventRootHash       Hash              `json:"event_root_hash"`
		GasUsed             U64               `json:"gas_used"`
		Success             bool              `json:"success"`
		VmStatus            string            `json:"vm_status"`
		Changes             []*WriteSetChange `json:"changes"`
		Events              []*Event          `json:"events"`
		Timestamp           U64               `json:"timestamp"`
		StateCheckpointHash Hash              `json:"state_checkpoint_hash"` // Optional
	}
	data := &inner{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	o.Version = data.Version.toUint64()
	o.Hash = data.Hash
	o.AccumulatorRootHash = data.AccumulatorRootHash
	o.StateChangeHash = data.StateChangeHash
	o.EventRootHash = data.EventRootHash
	o.GasUsed = data.GasUsed.toUint64()
	o.Success = data.Success
	o.VmStatus = data.VmStatus
	o.Changes = data.Changes
	o.Events = data.Events
	o.Timestamp = data.Timestamp.toUint64()
	o.StateCheckpointHash = data.StateCheckpointHash

	return nil
}

type SubmitTransactionResponse struct {
	Hash                    Hash
	Sender                  *types.AccountAddress
	SequenceNumber          uint64
	MaxGasAmount            uint64
	GasUnitPrice            uint64
	ExpirationTimestampSecs uint64
	Payload                 *TransactionPayload
	Signature               *Signature
}

func (o *SubmitTransactionResponse) UnmarshalJSON(b []byte) error {
	type inner struct {
		Hash                    Hash                  `json:"hash"`
		Sender                  *types.AccountAddress `json:"sender"`
		SequenceNumber          U64                   `json:"sequence_number"`
		MaxGasAmount            U64                   `json:"max_gas_amount"`
		GasUnitPrice            U64                   `json:"gas_unit_price"`
		ExpirationTimestampSecs U64                   `json:"expiration_timestamp_secs"`
		Payload                 *TransactionPayload   `json:"payload"`
		Signature               *Signature            `json:"signature"`
	}
	data := &inner{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	o.Hash = data.Hash
	o.Sender = data.Sender
	o.SequenceNumber = data.SequenceNumber.toUint64()
	o.MaxGasAmount = data.MaxGasAmount.toUint64()
	o.GasUnitPrice = data.GasUnitPrice.toUint64()
	o.ExpirationTimestampSecs = data.ExpirationTimestampSecs.toUint64()
	o.Payload = data.Payload
	o.Signature = data.Signature
	return nil
}
