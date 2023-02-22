package HydrogenChain

import "fmt"

func main() {
	db := conn.connectToDB()
	defer db.Close()

	var blocks []Block
	blockchain := BlockChain{
		Blocks: blocks,
		DB:     db,
	}

	lastBlock := getLastBlock(blockchain.DB)
	blockchain.Blocks = append(blockchain.Blocks, lastBlock)

	for {
		fmt.Print("Enter block data: ")
		var data string
		fmt.Scanln(&data)

		newBlock := createBlock(data, lastBlock.Hash, blockchain.DB)
		blockchain.Blocks = append(blockchain.Blocks, newBlock)

		fmt.Println("Block added!")
		lastBlock = newBlock
	}
}
