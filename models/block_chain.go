package models

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int
	TimeStamp int64
	BPM       int
	Hash      string
	PreHash   string
}

func (this *Block) IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PreHash {
		return false
	}
	if this.calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func (this *Block) calculateHash(block Block) string {
	record := string(block.Index) + string(block.TimeStamp) + string(block.BPM) + block.PreHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (this *Block) GenerateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.TimeStamp = t.Unix()
	newBlock.BPM = BPM
	newBlock.PreHash = oldBlock.Hash
	newBlock.Hash = this.calculateHash(newBlock)
	return newBlock, nil
}
