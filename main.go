package main

import (
	"fmt"
	"github.com/ayachain/go-aya/utils"
	"github.com/ayachain/go-ayadb/db/kv/btree"
	"time"
)

func main() {

	tdb := btree.NewBTreeKVDB()

	if utils.AFMS_IsPathExist("/kvdbtest/db") {
		utils.AFMS_RemovePath("/kvdbtest/db")
	}

	if err := tdb.CreateDB("kvdbtest","db"); err != nil {
		panic(err)
	}

	//写入数据 100w条
	fmt.Print("正在写入100w条数据:")
	stime := time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		kv := fmt.Sprintf("Key%d", i)
		tdb.Set(kv, kv)
	}
	fmt.Printf("%dms\n", (time.Now().UnixNano() - stime) / 1e6)


	//写入IPFS
	fmt.Print("正在刷写磁盘:")
	stime = time.Now().UnixNano()
	tdb.CloseDB()
	fmt.Printf("%dms\n", (time.Now().UnixNano() - stime) / 1e6)


	//读取
	fmt.Print("读取数据:")
	tdb = btree.NewBTreeKVDB()
	stime = time.Now().UnixNano()
	if err := tdb.OpenDB("kvdbtest","db"); err != nil {
		panic(err)
	}
	fmt.Printf("%dms\n", (time.Now().UnixNano() - stime) / 1e6)

}
