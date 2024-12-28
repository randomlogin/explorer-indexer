package node

import (
	"context"
	"log"
	"sort"

	. "github.com/spacesprotocol/explorer-backend/pkg/types"
)

var mempoolChunkSize = 200

type BitcoinClient struct {
	*Client
}

func (client *BitcoinClient) GetBlockChainInfo() {
	ctx := context.Background()
	var x interface{}
	// var z []string
	err := client.Rpc(ctx, "getblockchaininfo", []interface{}{}, x)
	log.Print(err)

}

func (client *BitcoinClient) GetBlock(ctx context.Context, blockHash string) (*Block, error) {
	block := new(Block)
	err := client.Rpc(ctx, "getblock", []interface{}{blockHash, 2}, block)
	if err != nil {
		return nil, err
	}
	return block, err
}

func (client *BitcoinClient) GetBlockHash(ctx context.Context, height int) (*Bytes, error) {
	blockHash := new(Bytes)
	// hexHeight := fmt.Sprintf("%x", height)
	err := client.Rpc(ctx, "getblockhash", []interface{}{height}, blockHash)
	if err != nil {
		return nil, err
	}
	return blockHash, err
}

func (client *BitcoinClient) GetBestBlockHeight(ctx context.Context) (int32, Bytes, error) {
	blockHash, err := client.GetBestBlockHash(ctx)
	if err != nil {
		return -1, nil, err
	}
	blockH, err := blockHash.MarshalText()
	if err != nil {
		return -1, nil, err
	}
	block, err := client.GetBlock(ctx, string(blockH))
	if err != nil {
		return -1, nil, err
	}
	return block.Height, block.Hash, nil
}

func (client *BitcoinClient) GetBestBlockHash(ctx context.Context) (*Bytes, error) {
	blockHash := new(Bytes)
	// hexHeight := fmt.Sprintf("%x", height)
	err := client.Rpc(ctx, "getbestblockhash", []interface{}{}, blockHash)
	if err != nil {
		return nil, err
	}
	return blockHash, err
}

func (client *BitcoinClient) GetTransaction(ctx context.Context, txId string) (*Transaction, error) {
	tx := new(Transaction)
	err := client.Rpc(ctx, "getrawtransaction", []interface{}{txId, 2}, tx)
	if err != nil {
		return nil, err
	}
	return tx, err
}

func (client *BitcoinClient) GetMempoolTxs(ctx context.Context) ([]Transaction, error) {
	var txids []string
	var txs []Transaction
	err := client.Rpc(ctx, "getrawmempool", nil, &txids)
	if err != nil {
		return nil, err
	}
	for _, txid := range txids {
		tx, err := client.GetTransaction(context.Background(), txid)
		if err != nil {
			return nil, err
		}
		txs = append(txs, *tx)
	}
	return txs, nil
}

type MempoolTx struct {
	Time    int64    `json:"time"`
	Depends []string `json:"depends"`
}

func (client *BitcoinClient) GetMempoolTxIds(ctx context.Context) ([][]string, error) {
	response := make(map[string]MempoolTx)
	err := client.Rpc(ctx, "getrawmempool", []interface{}{true}, &response)
	if err != nil {
		return nil, err
	}

	// Build dependency graph using depends field
	dependsOn := make(map[string][]string)  // txid -> list of txs it depends on
	dependedBy := make(map[string][]string) // txid -> list of txs that depend on it

	for txid, info := range response {
		for _, dep := range info.Depends {
			dependsOn[txid] = append(dependsOn[txid], dep)
			dependedBy[dep] = append(dependedBy[dep], txid)
		}
	}

	var orderedGroups [][]string
	processed := make(map[string]bool)

	// Find independent transactions (those with no dependencies)
	var independentTxs []string
	for txid, info := range response {
		if len(info.Depends) == 0 {
			independentTxs = append(independentTxs, txid)
		}
	}

	// Sort independent transactions by time
	sort.Slice(independentTxs, func(i, j int) bool {
		return response[independentTxs[i]].Time < response[independentTxs[j]].Time
	})

	// Process dependency chains
	var processChain func(txid string, chain []string) []string
	processChain = func(txid string, chain []string) []string {
		if processed[txid] {
			return chain
		}

		processed[txid] = true
		chain = append(chain, txid)

		// Get all transactions that depend on this tx
		dependents := dependedBy[txid]

		// Sort dependents by time
		sort.Slice(dependents, func(i, j int) bool {
			return response[dependents[i]].Time < response[dependents[j]].Time
		})

		// Process each dependent
		for _, dep := range dependents {
			if !processed[dep] {
				// Check if all dependencies of this tx are processed
				allDepsProcessed := true
				for _, parentDep := range dependsOn[dep] {
					if !processed[parentDep] {
						allDepsProcessed = false
						break
					}
				}

				// Only process if all dependencies are already processed
				if allDepsProcessed {
					chain = processChain(dep, chain)
				}
			}
		}

		return chain
	}

	// Process each independent transaction and its dependency chain
	for _, indTx := range independentTxs {
		if chain := processChain(indTx, nil); len(chain) > 0 {
			orderedGroups = append(orderedGroups, chain)
		}
	}

	// Handle any remaining transactions
	var remainingTxs []string
	for txid := range response {
		if !processed[txid] {
			remainingTxs = append(remainingTxs, txid)
		}
	}

	// Sort remaining transactions by time
	sort.Slice(remainingTxs, func(i, j int) bool {
		return response[remainingTxs[i]].Time < response[remainingTxs[j]].Time
	})

	// Process remaining transactions
	for _, txid := range remainingTxs {
		if chain := processChain(txid, nil); len(chain) > 0 {
			orderedGroups = append(orderedGroups, chain)
		}
	}
	for _, x := range orderedGroups {
		log.Println(x)
	}

	return orderedGroups, nil
}
