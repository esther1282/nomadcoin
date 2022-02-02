package blockchain

import (
	"sync"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height)
	b.NewestHash = block.Hash
	b.Height = block.Height
}

func Blockchain() *blockchain { //블록체인이 초기화 되었는지 알 수 있다.
	if b == nil { //딱 한번만 실행
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}
