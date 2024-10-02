// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: spaces.sql

package db

import (
	"context"
	"database/sql"

	"github.com/spacesprotocol/explorer-backend/pkg/types"
)

const getVMetaOutsByTxid = `-- name: GetVMetaOutsByTxid :many
SELECT block_hash, txid, tx_index, outpoint_txid, outpoint_index, name, burn_increment, covenant_action, claim_height, expire_height
FROM vmetaouts
WHERE txid = $1
ORDER BY tx_index
`

func (q *Queries) GetVMetaOutsByTxid(ctx context.Context, txid types.Bytes) ([]Vmetaout, error) {
	rows, err := q.db.QueryContext(ctx, getVMetaOutsByTxid, txid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Vmetaout{}
	for rows.Next() {
		var i Vmetaout
		if err := rows.Scan(
			&i.BlockHash,
			&i.Txid,
			&i.TxIndex,
			&i.OutpointTxid,
			&i.OutpointIndex,
			&i.Name,
			&i.BurnIncrement,
			&i.CovenantAction,
			&i.ClaimHeight,
			&i.ExpireHeight,
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

const getVMetaoutsByBlockAndTxid = `-- name: GetVMetaoutsByBlockAndTxid :many
SELECT block_hash, txid, tx_index, outpoint_txid, outpoint_index, name, burn_increment, covenant_action, claim_height, expire_height
FROM vmetaouts
WHERE block_hash = $1 and txid = $2
ORDER BY tx_index
`

type GetVMetaoutsByBlockAndTxidParams struct {
	BlockHash types.Bytes
	Txid      types.Bytes
}

func (q *Queries) GetVMetaoutsByBlockAndTxid(ctx context.Context, arg GetVMetaoutsByBlockAndTxidParams) ([]Vmetaout, error) {
	rows, err := q.db.QueryContext(ctx, getVMetaoutsByBlockAndTxid, arg.BlockHash, arg.Txid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Vmetaout{}
	for rows.Next() {
		var i Vmetaout
		if err := rows.Scan(
			&i.BlockHash,
			&i.Txid,
			&i.TxIndex,
			&i.OutpointTxid,
			&i.OutpointIndex,
			&i.Name,
			&i.BurnIncrement,
			&i.CovenantAction,
			&i.ClaimHeight,
			&i.ExpireHeight,
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

const insertVMetaOut = `-- name: InsertVMetaOut :exec
INSERT INTO vmetaouts (block_hash, txid, tx_index, outpoint_txid, outpoint_index, name, burn_increment, covenant_action, claim_height, expire_height)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`

type InsertVMetaOutParams struct {
	BlockHash      types.Bytes
	Txid           types.Bytes
	TxIndex        int64
	OutpointTxid   types.Bytes
	OutpointIndex  int64
	Name           string
	BurnIncrement  sql.NullInt64
	CovenantAction CovenantAction
	ClaimHeight    sql.NullInt64
	ExpireHeight   sql.NullInt64
}

func (q *Queries) InsertVMetaOut(ctx context.Context, arg InsertVMetaOutParams) error {
	_, err := q.db.ExecContext(ctx, insertVMetaOut,
		arg.BlockHash,
		arg.Txid,
		arg.TxIndex,
		arg.OutpointTxid,
		arg.OutpointIndex,
		arg.Name,
		arg.BurnIncrement,
		arg.CovenantAction,
		arg.ClaimHeight,
		arg.ExpireHeight,
	)
	return err
}
