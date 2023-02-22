package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type BlockChain struct {
	Blocks []Block
	DB     *sql.DB
}

type Block struct {
	Index     int    `json:"index"`
	Timestamp int64  `json:"timestamp"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevHash"`
}
