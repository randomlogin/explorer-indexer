package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spacesprotocol/explorer-backend/pkg/db"
	"github.com/spacesprotocol/explorer-backend/pkg/node"

	_ "github.com/lib/pq"
	. "github.com/spacesprotocol/explorer-backend/pkg/types"
)

func main() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// bitcoinClient := node.NewClient("http://127.0.0.1:48332", "test", "test")
	spacesClient := node.NewClient("http://127.0.0.1:7224", "test", "test")
	// pg, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//
	// bc := node.BitcoinClient{client}
	//
	// err = syncBlocks(pg, &bc, &sc)
	// // err = checkW(pg, &bc)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	blockHash := "000000000000004c1d5bbc5b3f6693f25b045ba66434fd6ba46c792177f2d63e"
	sc := node.SpacesClient{spacesClient}
	block, err := sc.GetBlockData(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.MarshalIndent(block, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	log.Print(string(b))
}

// func checkW(ph *sql.DB, bc *node.BitcoinClient) error {
// 	bestBlockHash, err := bc.GetBestBlockHash(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	hashString, err := bestBlockHash.MarshalText()
// 	if err != nil {
// 		return err
// 	}
//
// 	block, err := bc.GetBlock(context.Background(), string(hashString))
// 	if err != nil {
// 		return err
// 	}
//
// 	log.Printf("best block, %+v", block)
// 	log.Printf("best next, %+v", block.NextBlockHash)
// 	log.Print(block.NextBlockHash == nil)
// 	return nil
//
// }

func syncBlocks(pg *sql.DB, bc *node.BitcoinClient, sc *node.SpacesClient) error {
	var hash *Bytes
	height, hash, err := getSyncedHead(pg, bc)
	if err != nil {
		return err
	}
	//it means we have no synced blocks
	if height == -1 {
		hash, err = bc.GetBlockHash(context.Background(), 0)
		log.Print("here")
		if err != nil {
			log.Print("aaahere")
			return err
		}
	}

	hashString, err := hash.MarshalText()
	if err != nil {
		return err
	}

	log.Print("here")
	block, err := bc.GetBlock(context.Background(), string(hashString))
	if err != nil {
		return err
	}
	nextBlockHash := block.NextBlockHash

	log.Printf("block %+v", block)
	//TODO what if the node best block changes during the sync?
	for nextBlockHash != nil {

		nextHashString, err := nextBlockHash.MarshalText()
		if err != nil {
			return err
		}
		block, err := bc.GetBlock(context.Background(), string(nextHashString))
		log.Print("trying to sync", block.Height)
		if err != nil {
			return err
		}
		err = syncBlock(pg, block)
		if err != nil {
			return err
		}
		{
			spacesBlock, err := sc.GetBlockData(context.Background(), string(nextHashString))
			if err != nil {
				return err
		}
			err = syncSpacesTransactions(pg, spacesBlock.Transactions)
			if err != nil {
				return err
			}
		}
		nextBlockHash = block.NextBlockHash
	}
	return nil

	// maxHeight, err := bc.GetBestBlockHeight(context.Background())
	// if err != nil {
	// 	return err
	// }
	// for height < maxHeight {
	// 	height += 1
	// 	block, err := bc.GetBlockByHeight(context.Background(), height)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if !bytes.Equal(hash, block.PrevBlockHash) {
	// 		break
	// 	}
	// 	err = syncBlock(pg, block)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	hash = block.Hash
	// }
	// return nil
}

// returns the height and blockhash of the last block that is identical in the db and the node
func getSyncedHead(pg *sql.DB, bc *node.BitcoinClient) (int32, *Bytes, error) {
	q := db.New(pg)
	//takes last block from the DB
	height, err := q.GetBlocksMaxHeight(context.Background())
	if err != nil {
		return -1, nil, err
	}
	//height is the height of the db block
	for height >= 0 {
		//take last block hash from the DB
		dbHash, err := q.GetBlockHashByHeight(context.Background(), height)
		if err != nil {
			return -1, nil, err
		}
		//takes the block of same height from the bitcoin node
		nodeHash, err := bc.GetBlockHash(context.Background(), int(height))
		if err != nil {
			return -1, nil, err
		}
		// nodeHash *bytes
		// dbHash Bytes
		if bytes.Equal(dbHash, *nodeHash) {
			//marking all the blocks in the DB after the sycned height as orphans
			if err := q.SetOrphanAfterHeight(context.Background(), height); err != nil { 
				return -1, nil, err
			}
			return height, &dbHash, nil
		}
		height -= 1
	}
	return -1, nil, nil
}
