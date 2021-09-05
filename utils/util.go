package utils

import (
	"encoding/binary"
)

func ConvertIntToBytes(n int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(n))
	return b
}

func ConvertBytesToInt(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
