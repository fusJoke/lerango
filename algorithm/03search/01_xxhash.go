package main

import (
	"fmt"
	"github.com/cespare/xxhash"
)

func XXHash(key []byte) uint64{
	h := xxhash.New()
	h.Write(key)
	return h.Sum64()
}

func main() {
	keys := []string{"hi", "my", "friend", "I", "love", "you", "my", "apple"}
	for _, key := range keys {
		fmt.Printf("xxhash('%s')=%d\n", key,  XXHash([]byte(key)))
	}
}
