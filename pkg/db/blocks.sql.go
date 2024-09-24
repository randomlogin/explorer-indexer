// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: blocks.sql

package db

import (
	"context"

	"github.com/spacesprotocol/explorer-backend/pkg/types"
)

const deleteBlocksAfterHeight = `-- name: DeleteBlocksAfterHeight :exec
DELETE FROM blocks
WHERE height > $1
`

func (q *Queries) DeleteBlocksAfterHeight(ctx context.Context, height int32) error {
	_, err := q.db.ExecContext(ctx, deleteBlocksAfterHeight, height)
	return err
}

const getBlockByHash = `-- name: GetBlockByHash :one
SELECT blocks.hash, blocks.size, blocks.stripped_size, blocks.weight, blocks.height, blocks.version, blocks.hash_merkle_root, blocks.time, blocks.median_time, blocks.nonce, blocks.bits, blocks.difficulty, blocks.chainwork, blocks.orphan, (
  SELECT COUNT(*) FROM transactions WHERE blocks.hash = transactions.block_hash
)::integer AS txs_count
FROM blocks
WHERE blocks.hash = $1
`

type GetBlockByHashRow struct {
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
	TxsCount       int32
}

func (q *Queries) GetBlockByHash(ctx context.Context, hash types.Bytes) (GetBlockByHashRow, error) {
	row := q.db.QueryRowContext(ctx, getBlockByHash, hash)
	var i GetBlockByHashRow
	err := row.Scan(
		&i.Hash,
		&i.Size,
		&i.StrippedSize,
		&i.Weight,
		&i.Height,
		&i.Version,
		&i.HashMerkleRoot,
		&i.Time,
		&i.MedianTime,
		&i.Nonce,
		&i.Bits,
		&i.Difficulty,
		&i.Chainwork,
		&i.Orphan,
		&i.TxsCount,
	)
	return i, err
}

const getBlockByHeight = `-- name: GetBlockByHeight :one
SELECT blocks.hash, blocks.size, blocks.stripped_size, blocks.weight, blocks.height, blocks.version, blocks.hash_merkle_root, blocks.time, blocks.median_time, blocks.nonce, blocks.bits, blocks.difficulty, blocks.chainwork, blocks.orphan, (
  SELECT COUNT(*) FROM transactions WHERE blocks.hash = transactions.block_hash
)::integer AS txs_count
FROM blocks
WHERE blocks.height = $1
`

type GetBlockByHeightRow struct {
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
	TxsCount       int32
}

func (q *Queries) GetBlockByHeight(ctx context.Context, height int32) (GetBlockByHeightRow, error) {
	row := q.db.QueryRowContext(ctx, getBlockByHeight, height)
	var i GetBlockByHeightRow
	err := row.Scan(
		&i.Hash,
		&i.Size,
		&i.StrippedSize,
		&i.Weight,
		&i.Height,
		&i.Version,
		&i.HashMerkleRoot,
		&i.Time,
		&i.MedianTime,
		&i.Nonce,
		&i.Bits,
		&i.Difficulty,
		&i.Chainwork,
		&i.Orphan,
		&i.TxsCount,
	)
	return i, err
}

const getBlockHashByHeight = `-- name: GetBlockHashByHeight :one
SELECT hash
FROM blocks
WHERE height = $1
`

func (q *Queries) GetBlockHashByHeight(ctx context.Context, height int32) (types.Bytes, error) {
	row := q.db.QueryRowContext(ctx, getBlockHashByHeight, height)
	var hash types.Bytes
	err := row.Scan(&hash)
	return hash, err
}

const getBlocks = `-- name: GetBlocks :many
SELECT blocks.hash, blocks.size, blocks.stripped_size, blocks.weight, blocks.height, blocks.version, blocks.hash_merkle_root, blocks.time, blocks.median_time, blocks.nonce, blocks.bits, blocks.difficulty, blocks.chainwork, blocks.orphan, (
  SELECT COUNT(*) FROM transactions WHERE blocks.hash = transactions.block_hash
)::integer AS txs_count
FROM blocks
ORDER BY height DESC
LIMIT $1 OFFSET $2
`

type GetBlocksParams struct {
	Limit  int32
	Offset int32
}

type GetBlocksRow struct {
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
	TxsCount       int32
}

func (q *Queries) GetBlocks(ctx context.Context, arg GetBlocksParams) ([]GetBlocksRow, error) {
	rows, err := q.db.QueryContext(ctx, getBlocks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetBlocksRow{}
	for rows.Next() {
		var i GetBlocksRow
		if err := rows.Scan(
			&i.Hash,
			&i.Size,
			&i.StrippedSize,
			&i.Weight,
			&i.Height,
			&i.Version,
			&i.HashMerkleRoot,
			&i.Time,
			&i.MedianTime,
			&i.Nonce,
			&i.Bits,
			&i.Difficulty,
			&i.Chainwork,
			&i.Orphan,
			&i.TxsCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBlocksMaxHeight = `-- name: GetBlocksMaxHeight :one
SELECT COALESCE(MAX(height), -1)::integer
FROM blocks
`

func (q *Queries) GetBlocksMaxHeight(ctx context.Context) (int32, error) {
	row := q.db.QueryRowContext(ctx, getBlocksMaxHeight)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}

const insertBlock = `-- name: InsertBlock :exec
INSERT INTO blocks (hash, size, stripped_size, weight, height, version, hash_merkle_root, time, median_time, nonce, bits, difficulty, chainwork)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
`

type InsertBlockParams struct {
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
}

func (q *Queries) InsertBlock(ctx context.Context, arg InsertBlockParams) error {
	_, err := q.db.ExecContext(ctx, insertBlock,
		arg.Hash,
		arg.Size,
		arg.StrippedSize,
		arg.Weight,
		arg.Height,
		arg.Version,
		arg.HashMerkleRoot,
		arg.Time,
		arg.MedianTime,
		arg.Nonce,
		arg.Bits,
		arg.Difficulty,
		arg.Chainwork,
	)
	return err
}

const setOrphanAfterHeight = `-- name: SetOrphanAfterHeight :exec
UPDATE blocks SET orphan = true WHERE height > $1
`

func (q *Queries) SetOrphanAfterHeight(ctx context.Context, height int32) error {
	_, err := q.db.ExecContext(ctx, setOrphanAfterHeight, height)
	return err
}
