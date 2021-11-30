package encoding

import (
	"fmt"
	"math/big"
)

func IntToBigEndian(n *big.Int, size int) ([]byte, error) {
	buf := make([]byte, size)
	bs := n.Bytes()
	if len(bs) > size {
		return nil, fmt.Errorf("%s overflows max size of %d bytes", n.Text(16), size)
	}

	for i, j := size-len(bs), 0; i < size; i, j = i+1, j+1 {
		buf[i] = bs[j]
	}
	return buf, nil
}

func IntToLittleEndian(n *big.Int, size int) ([]byte, error) {
	out := make([]byte, size)
	bs := n.Bytes()
	if len(bs) > size {
		return nil, fmt.Errorf("%s overflows max size of %d bytes", n.Text(16), size)
	}

	for i, j := 0, len(bs)-1; i < len(bs); i, j = i+1, j-1 {
		out[i] = bs[j]
	}
	return out, nil
}

func IntFromBigEndian(bs []byte) *big.Int {
	return new(big.Int).SetBytes(bs)
}

func IntFromLittleEndian(bs []byte) *big.Int {
	out := make([]byte, len(bs))
	for i, j := 0, len(bs)-1; j >= 0; i, j = i+1, j-1 {
		out[i] = bs[j]
	}
	return new(big.Int).SetBytes(out)
}
