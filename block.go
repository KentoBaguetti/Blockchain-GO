package main

import (
	"time"
	"bytes"
	"crypto/sha256"
	"strconv"
)

type Block struct {
	TimeStamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}

func (b *Block) SetHash() {
	timeStamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timeStamp}, []byte[])
	hash := sha256.Sum256(headers)

	b.hash = hash[:]
}