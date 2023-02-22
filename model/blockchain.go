package model

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
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
