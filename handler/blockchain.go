package handler

import (
	"database/sql"
	"fmt"
	"time"
)

func createBlock(data string, prevHash string, db *sql.DB) Block {
	var index int
	var timestamp int64
	var hash string

	err := db.QueryRow("SELECT MAX(Index) as `Index`, MAX(Timestamp) as `Timestamp`, MAX(Hash) as `Hash` FROM Blocks").Scan(&index, &timestamp, &hash)
	if err != nil {
		panic(err.Error())
	}

	block := Block{
		Index:     index + 1,
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
		Hash:      calculateHash(fmt.Sprintf("%d%s%s%s", index+1, time.Now().Unix(), data, prevHash)),
	}

	stmt, err := db.Prepare("INSERT INTO Blocks (Index, Timestamp, Data, Hash, PrevHash) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(block.Index, block.Timestamp, block.Data, block.Hash, block.PrevHash)
	if err != nil {
		panic(err.Error())
	}

	return block
}

func getLastBlock(db *sql.DB) Block {
	var index int
	var timestamp int64
	var data string
	var hash string
	var prevHash string

	err := db.QueryRow("SELECT * FROM Blocks ORDER BY Index DESC LIMIT 1").Scan(&index, &timestamp, &data, &hash, &prevHash)
	if err != nil {
		panic(err.Error())
	}

	block := Block{
		Index:     index,
		Timestamp: timestamp,
		Data:      data,
		Hash:      hash,
		PrevHash:  prevHash,
	}

	return block
}
