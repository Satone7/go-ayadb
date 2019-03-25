package kv

type KVDB interface {
	Set(key string, v interface{})
	Get(key string) interface{}
	OpenDB(basedir string, dbname string) error
	CreateDB(basedir string, dbname string) error
	CloseDB() bool
}