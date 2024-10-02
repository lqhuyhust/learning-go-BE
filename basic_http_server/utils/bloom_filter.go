package utils

import (
	"hash/fnv"
)

const BloomFilterSize = 10000

func Hash(key string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(key))
	return h.Sum64() % BloomFilterSize
}
