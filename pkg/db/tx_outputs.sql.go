// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tx_outputs.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spacesprotocol/explorer-backend/pkg/types"
)

const getTxOutputsByBlockAndTxid = `-- name: GetTxOutputsByBlockAndTxid :many
SELECT block_hash, txid, index, value, scriptpubkey, spender_txid, spender_index, spender_block_hash
FROM tx_outputs
WHERE block_hash = $1 and txid = $2
ORDER BY index
`

type GetTxOutputsByBlockAndTxidParams struct {
	BlockHash types.Bytes
	Txid      types.Bytes
}

func (q *Queries) GetTxOutputsByBlockAndTxid(ctx context.Context, arg GetTxOutputsByBlockAndTxidParams) ([]TxOutput, error) {
	rows, err := q.db.Query(ctx, getTxOutputsByBlockAndTxid, arg.BlockHash, arg.Txid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TxOutput{}
	for rows.Next() {
		var i TxOutput
		if err := rows.Scan(
			&i.BlockHash,
			&i.Txid,
			&i.Index,
			&i.Value,
			&i.Scriptpubkey,
			&i.SpenderTxid,
			&i.SpenderIndex,
			&i.SpenderBlockHash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTxOutputsByTxid = `-- name: GetTxOutputsByTxid :many
SELECT block_hash, txid, index, value, scriptpubkey, spender_txid, spender_index, spender_block_hash
FROM tx_outputs
WHERE txid = $1
ORDER BY index
`

func (q *Queries) GetTxOutputsByTxid(ctx context.Context, txid types.Bytes) ([]TxOutput, error) {
	rows, err := q.db.Query(ctx, getTxOutputsByTxid, txid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TxOutput{}
	for rows.Next() {
		var i TxOutput
		if err := rows.Scan(
			&i.BlockHash,
			&i.Txid,
			&i.Index,
			&i.Value,
			&i.Scriptpubkey,
			&i.SpenderTxid,
			&i.SpenderIndex,
			&i.SpenderBlockHash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

type InsertBatchTxOutputsParams struct {
	BlockHash    types.Bytes
	Txid         types.Bytes
	Index        int64
	Value        int64
	Scriptpubkey types.Bytes
}

const insertTxOutput = `-- name: InsertTxOutput :exec
INSERT INTO tx_outputs (block_hash, txid, index, value, scriptPubKey)
VALUES ($1, $2, $3, $4, $5)
`

type InsertTxOutputParams struct {
	BlockHash    types.Bytes
	Txid         types.Bytes
	Index        int64
	Value        int64
	Scriptpubkey types.Bytes
}

func (q *Queries) InsertTxOutput(ctx context.Context, arg InsertTxOutputParams) error {
	_, err := q.db.Exec(ctx, insertTxOutput,
		arg.BlockHash,
		arg.Txid,
		arg.Index,
		arg.Value,
		arg.Scriptpubkey,
	)
	return err
}

const setSpender = `-- name: SetSpender :exec
UPDATE tx_outputs
SET spender_txid = $3, spender_index = $4
WHERE txid = $1 AND index = $2
`

type SetSpenderParams struct {
	Txid         types.Bytes
	Index        int64
	SpenderTxid  *types.Bytes
	SpenderIndex pgtype.Int8
}

func (q *Queries) SetSpender(ctx context.Context, arg SetSpenderParams) error {
	_, err := q.db.Exec(ctx, setSpender,
		arg.Txid,
		arg.Index,
		arg.SpenderTxid,
		arg.SpenderIndex,
	)
	return err
}
