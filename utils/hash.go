package utils

import "hash/fnv"

func HashMultiple(value string) int {
	h := fnv.New32a()
	h.Write([]byte(value))
	return int(h.Sum32() % 8) // 对8取模得到0-7的索引
}
