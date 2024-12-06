// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: copyfrom.go

package db

import (
	"context"
)

// iteratorForInsertBatchTxInputs implements pgx.CopyFromSource.
type iteratorForInsertBatchTxInputs struct {
	rows                 []InsertBatchTxInputsParams
	skippedFirstNextCall bool
}

func (r *iteratorForInsertBatchTxInputs) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForInsertBatchTxInputs) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].BlockHash,
		r.rows[0].Txid,
		r.rows[0].Index,
		r.rows[0].HashPrevout,
		r.rows[0].IndexPrevout,
		r.rows[0].Sequence,
		r.rows[0].Coinbase,
	}, nil
}

func (r iteratorForInsertBatchTxInputs) Err() error {
	return nil
}

func (q *Queries) InsertBatchTxInputs(ctx context.Context, arg []InsertBatchTxInputsParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"tx_inputs"}, []string{"block_hash", "txid", "index", "hash_prevout", "index_prevout", "sequence", "coinbase"}, &iteratorForInsertBatchTxInputs{rows: arg})
}

// iteratorForInsertBatchTxOutputs implements pgx.CopyFromSource.
type iteratorForInsertBatchTxOutputs struct {
	rows                 []InsertBatchTxOutputsParams
	skippedFirstNextCall bool
}

func (r *iteratorForInsertBatchTxOutputs) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForInsertBatchTxOutputs) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].BlockHash,
		r.rows[0].Txid,
		r.rows[0].Index,
		r.rows[0].Value,
		r.rows[0].Scriptpubkey,
	}, nil
}

func (r iteratorForInsertBatchTxOutputs) Err() error {
	return nil
}

func (q *Queries) InsertBatchTxOutputs(ctx context.Context, arg []InsertBatchTxOutputsParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"tx_outputs"}, []string{"block_hash", "txid", "index", "value", "scriptpubkey"}, &iteratorForInsertBatchTxOutputs{rows: arg})
}