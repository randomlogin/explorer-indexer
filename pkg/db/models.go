// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/spacesprotocol/explorer-backend/pkg/types"
)

type CovenantAction string

const (
	CovenantActionRESERVE  CovenantAction = "RESERVE"
	CovenantActionBID      CovenantAction = "BID"
	CovenantActionTRANSFER CovenantAction = "TRANSFER"
	CovenantActionROLLOUT  CovenantAction = "ROLLOUT"
	CovenantActionREVOKE   CovenantAction = "REVOKE"
	CovenantActionREJECT   CovenantAction = "REJECT"
)

func (e *CovenantAction) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CovenantAction(s)
	case string:
		*e = CovenantAction(s)
	default:
		return fmt.Errorf("unsupported scan type for CovenantAction: %T", src)
	}
	return nil
}

type NullCovenantAction struct {
	CovenantAction CovenantAction
	Valid          bool // Valid is true if CovenantAction is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCovenantAction) Scan(value interface{}) error {
	if value == nil {
		ns.CovenantAction, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CovenantAction.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCovenantAction) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CovenantAction), nil
}

type Block struct {
	Hash           types.Bytes
	Size           int64
	StrippedSize   int64
	Weight         int32
	Height         int32
	Version        int32
	HashMerkleRoot types.Bytes
	Time           int32
	MedianTime     int32
	Nonce          int64
	Bits           types.Bytes
	Difficulty     float64
	Chainwork      types.Bytes
	Orphan         bool
}

type Rollout struct {
	Name   string
	Bid    sql.NullInt64
	Height sql.NullInt64
}

type Transaction struct {
	Txid      types.Bytes
	TxHash    types.Bytes
	Version   int32
	Size      int64
	Vsize     int64
	Weight    int64
	Locktime  int32
	Fee       int64
	BlockHash types.Bytes
	Index     sql.NullInt32
}

type TxInput struct {
	BlockHash    types.Bytes
	Txid         types.Bytes
	Index        int64
	HashPrevout  *types.Bytes
	IndexPrevout int64
	Sequence     int64
	Coinbase     *types.Bytes
	Txinwitness  []types.Bytes
}

type TxOutput struct {
	BlockHash        types.Bytes
	Txid             types.Bytes
	Index            int64
	Value            int64
	Scriptpubkey     types.Bytes
	SpenderTxid      *types.Bytes
	SpenderIndex     sql.NullInt64
	SpenderBlockHash *types.Bytes
}

type Vmetaout struct {
	BlockHash     types.Bytes
	Txid          types.Bytes
	Identifier    int64
	Priority      sql.NullInt64
	Name          sql.NullString
	Reason        sql.NullString
	Value         sql.NullInt64
	Scriptpubkey  *types.Bytes
	Action        NullCovenantAction
	BurnIncrement sql.NullInt64
	Signature     *types.Bytes
	TotalBurned   sql.NullInt64
	ClaimHeight   sql.NullInt64
	ExpireHeight  sql.NullInt64
	ScriptError   sql.NullString
}
