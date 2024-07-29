package initialize

import (
	"BlockChain/app/dao"
	"BlockChain/app/model"
)

func BuildBlockChain() {
	initBlock := model.Block{
		Index:     0,
		Timestamp: "",
		BPM:       0,
		Hash:      "",
		PrevHash:  "",
	}
	dao.Blockchain = append(dao.Blockchain, initBlock)
}
