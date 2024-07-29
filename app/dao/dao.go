package dao

import (
	"BlockChain/app/model"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

var Blockchain []model.Block

func calculateHash(block model.Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
func GenerateBlock(oldBlock model.Block, BPM int) (model.Block, error) {

	var newBlock model.Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}
func IsBlockValid(newBlock, oldBlock model.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
func ReplaceChain(newBlocks []model.Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
