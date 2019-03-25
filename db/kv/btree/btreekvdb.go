package btree

import (
	"encoding/json"
	"github.com/ayachain/go-ayadb/db/kv"
	AUtils "github.com/ayachain/go-aya/utils"
	"github.com/ipfs/go-ipfs-api"
	"github.com/pkg/errors"
)

type BTreeKVDB struct {

	kv.KVDB
	mfspath 	string
	sh			shell.Shell

	mapping		map[string]interface{}
}


func NewBTreeKVDB() kv.KVDB {
	return &BTreeKVDB{mapping: map[string]interface{}{}}
}


func (db *BTreeKVDB) Set(key string, v interface{}) {
	db.mapping[key] = v
}

func (db *BTreeKVDB) Get(key string) interface{} {

	v, isexist := db.mapping[key]

	if !isexist {
		return nil
	} else {
		return v
	}

}

func (db *BTreeKVDB) OpenDB(basedir string, dbname string) error {

	if !AUtils.AFMS_IsPathExist("/" + basedir + "/" + dbname) {
		return errors.New("BTreeKVDB Error:MFSPath is not exist.")
	}

	db.mfspath = "/" + basedir + "/" + dbname

	bs, err := AUtils.AFMS_ReadFile(db.mfspath, 0, 0)

	if err != nil {
		return errors.New("BTreeKVDB Error:OpenDB Readfiles faild.")
	}

	if err := json.Unmarshal(bs, &db.mapping); err != nil {
		return errors.New("BTreeKVDB Error:Unmarshal read content faild.")
	}

	return nil
}

func (db *BTreeKVDB) CreateDB(basedir string, dbname string) error {

	db.mfspath = "/" + basedir + "/" + dbname

	if AUtils.AFMS_IsPathExist(db.mfspath) {
		return errors.New("BTreeKVDB Error:MFSPath are already exist.")
	}

	return AUtils.AFMS_CreateFile(db.mfspath, nil)
}

func (db *BTreeKVDB) CloseDB() bool {

	bs, err := json.Marshal(db.mapping)

	if err != nil {
		return false
	}

	if err := AUtils.AFMS_ReplaceFile(db.mfspath, bs); err != nil {
		return false
	}

	return true

}