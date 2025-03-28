// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: transactions.sql

package db

import (
	"context"

	"github.com/spacesprotocol/explorer-indexer/pkg/types"
)

const deleteMempoolTransactionByTxid = `-- name: DeleteMempoolTransactionByTxid :exec
DELETE FROM transactions
where txid = $1
AND block_hash = '\xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef'
`

func (q *Queries) DeleteMempoolTransactionByTxid(ctx context.Context, txid types.Bytes) error {
	_, err := q.db.Exec(ctx, deleteMempoolTransactionByTxid, txid)
	return err
}

const deleteMempoolTransactions = `-- name: DeleteMempoolTransactions :exec
DELETE FROM transactions
WHERE block_hash = '\xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef' and index <0
`

func (q *Queries) DeleteMempoolTransactions(ctx context.Context) error {
	_, err := q.db.Exec(ctx, deleteMempoolTransactions)
	return err
}

const getMempoolTransactions = `-- name: GetMempoolTransactions :many
SELECT txid, tx_hash, version, size, vsize, weight, locktime, fee, block_hash, index
FROM transactions
WHERE block_hash = '\xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef'
ORDER BY index
LIMIT $1 OFFSET $2
`

type GetMempoolTransactionsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetMempoolTransactions(ctx context.Context, arg GetMempoolTransactionsParams) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getMempoolTransactions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.Txid,
			&i.TxHash,
			&i.Version,
			&i.Size,
			&i.Vsize,
			&i.Weight,
			&i.Locktime,
			&i.Fee,
			&i.BlockHash,
			&i.Index,
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

const getMempoolTxids = `-- name: GetMempoolTxids :many
SELECT txid
FROM transactions
WHERE block_hash = '\xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef'
ORDER BY index
`

func (q *Queries) GetMempoolTxids(ctx context.Context) ([]types.Bytes, error) {
	rows, err := q.db.Query(ctx, getMempoolTxids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []types.Bytes{}
	for rows.Next() {
		var txid types.Bytes
		if err := rows.Scan(&txid); err != nil {
			return nil, err
		}
		items = append(items, txid)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTransactionsByBlockHeight = `-- name: GetTransactionsByBlockHeight :many
SELECT
  transactions.txid, transactions.tx_hash, transactions.version, transactions.size, transactions.vsize, transactions.weight, transactions.locktime, transactions.fee, transactions.block_hash, transactions.index,
  COALESCE(blocks.height, -1)::integer AS block_height_not_null
FROM
  transactions
  INNER JOIN blocks ON (transactions.block_hash = blocks.hash)
WHERE blocks.height = $1
ORDER BY transactions.index
LIMIT $2 OFFSET $3
`

type GetTransactionsByBlockHeightParams struct {
	Height int32
	Limit  int32
	Offset int32
}

type GetTransactionsByBlockHeightRow struct {
	Txid               types.Bytes
	TxHash             types.Bytes
	Version            int32
	Size               int64
	Vsize              int64
	Weight             int64
	Locktime           int32
	Fee                int64
	BlockHash          types.Bytes
	Index              int32
	BlockHeightNotNull int32
}

func (q *Queries) GetTransactionsByBlockHeight(ctx context.Context, arg GetTransactionsByBlockHeightParams) ([]GetTransactionsByBlockHeightRow, error) {
	rows, err := q.db.Query(ctx, getTransactionsByBlockHeight, arg.Height, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetTransactionsByBlockHeightRow{}
	for rows.Next() {
		var i GetTransactionsByBlockHeightRow
		if err := rows.Scan(
			&i.Txid,
			&i.TxHash,
			&i.Version,
			&i.Size,
			&i.Vsize,
			&i.Weight,
			&i.Locktime,
			&i.Fee,
			&i.BlockHash,
			&i.Index,
			&i.BlockHeightNotNull,
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

const insertMempoolTransaction = `-- name: InsertMempoolTransaction :exec
INSERT INTO transactions (
    txid, tx_hash, version, size, vsize, weight, locktime, fee, block_hash
) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9)
`

type InsertMempoolTransactionParams struct {
	Txid      types.Bytes
	TxHash    types.Bytes
	Version   int32
	Size      int64
	Vsize     int64
	Weight    int64
	Locktime  int32
	Fee       int64
	BlockHash types.Bytes
}

func (q *Queries) InsertMempoolTransaction(ctx context.Context, arg InsertMempoolTransactionParams) error {
	_, err := q.db.Exec(ctx, insertMempoolTransaction,
		arg.Txid,
		arg.TxHash,
		arg.Version,
		arg.Size,
		arg.Vsize,
		arg.Weight,
		arg.Locktime,
		arg.Fee,
		arg.BlockHash,
	)
	return err
}

const insertTransaction = `-- name: InsertTransaction :exec
INSERT INTO transactions (
    txid, tx_hash, version, size, vsize, weight, locktime, fee, block_hash, index
) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`

type InsertTransactionParams struct {
	Txid      types.Bytes
	TxHash    types.Bytes
	Version   int32
	Size      int64
	Vsize     int64
	Weight    int64
	Locktime  int32
	Fee       int64
	BlockHash types.Bytes
	Index     int32
}

func (q *Queries) InsertTransaction(ctx context.Context, arg InsertTransactionParams) error {
	_, err := q.db.Exec(ctx, insertTransaction,
		arg.Txid,
		arg.TxHash,
		arg.Version,
		arg.Size,
		arg.Vsize,
		arg.Weight,
		arg.Locktime,
		arg.Fee,
		arg.BlockHash,
		arg.Index,
	)
	return err
}
