package utils

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Log struct {
	// Consensus fields:
	// address of the contract that generated the event
	Address common.Address `json:"address" gencodec:"required"`
	// list of topics provided by the contract.
	Topics []common.Hash `json:"topics" gencodec:"required"`
	// supplied by the contract, usually ABI-encoded
	Data []byte `json:"data" gencodec:"required"`

	// Derived fields. These fields are filled in by the node
	// but not secured by consensus.
	// block in which the transaction was included
	BlockNumber uint64 `json:"blockNumber"`
	// hash of the transaction
	TxHash common.Hash `json:"transactionHash" gencodec:"required"`
	// index of the transaction in the block
	TxIndex uint `json:"transactionIndex"`
	// hash of the block in which the transaction was included
	BlockHash common.Hash `json:"blockHash"`
	// index of the log in the block
	Index uint `json:"logIndex"`

	// The Removed field is true if this log was reverted due to a chain reorganisation.
	// You must pay attention to this field if you receive logs through a filter query.
	Removed bool `json:"removed"`
}

// MarshalJSON marshals as JSON.
func (l Log) MarshalJSON() ([]byte, error) {
	type Log struct {
		Address     common.Address `json:"address" gencodec:"required"`
		Topics      []common.Hash  `json:"topics" gencodec:"required"`
		Data        hexutil.Bytes  `json:"data" gencodec:"required"`
		BlockNumber uint64         `json:"blockNumber"`
		TxHash      common.Hash    `json:"transactionHash" gencodec:"required"`
		TxIndex     uint           `json:"transactionIndex"`
		BlockHash   common.Hash    `json:"blockHash"`
		Index       uint           `json:"logIndex"`
		Removed     bool           `json:"removed"`
	}
	var enc Log
	enc.Address = l.Address
	enc.Topics = l.Topics
	enc.Data = l.Data
	enc.BlockNumber = l.BlockNumber
	enc.TxHash = l.TxHash
	enc.TxIndex = l.TxIndex
	enc.BlockHash = l.BlockHash
	enc.Index = l.Index
	enc.Removed = l.Removed
	return json.Marshal(&enc)
}
