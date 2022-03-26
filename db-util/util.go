package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {

	db, err := leveldb.OpenFile("../db", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	address1 := common.HexToAddress("0x60a71117353f3ac632f3e75fd1be532850aa5f4d")
	if err := db.Put(address1[:], address1[:], nil); err != nil {
		panic(err)
	}
	address2 := common.HexToAddress("0xha039801ccf87b2ad3a87865a557960024cec0cf")
	if err := db.Put(address2[:], address2[:], nil); err != nil {
		panic(err)
	}

	printAll(db)
}

func printAll(db *leveldb.DB) {
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Printf("%x\n", iter.Key())
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		panic(err)
	}
}
