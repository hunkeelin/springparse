package klinutils

import (
	"hash/fnv"
	"strconv"
)

func Stringtoport(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.FormatUint(uint64(h.Sum32()), 10)[0:5]
}
