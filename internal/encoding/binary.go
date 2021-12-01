package encoding

import (
	"fmt"
	"math/big"
)

type Endianness int

const (
	BigEndian Endianness = iota
	LittleEndian
)

// IntToBytes converts an integer to a byte array based on its endianness.
func IntToBytes(n *big.Int, size int, endianness Endianness) ([]byte, error) {
	out := make([]byte, size)
	bs := n.Bytes()
	if len(bs) > size {
		return nil, fmt.Errorf("%s overflows max size of %d bytes", n.Text(16), size)
	}

	if endianness == BigEndian {
		for i, j := size-len(bs), 0; i < size; i, j = i+1, j+1 {
			out[i] = bs[j]
		}
		return out, nil
	}

	for i, j := 0, len(bs)-1; i < len(bs); i, j = i+1, j-1 {
		out[i] = bs[j]
	}
	return out, nil
}

// BytesToInt converts a byte array to an integer based on its endianness.
func BytesToInt(bs []byte, endianness Endianness) *big.Int {
	if endianness == BigEndian {
		return new(big.Int).SetBytes(bs)
	}

	out := make([]byte, len(bs))
	for i, j := 0, len(bs)-1; j >= 0; i, j = i+1, j-1 {
		out[i] = bs[j]
	}
	return new(big.Int).SetBytes(out)
}
