package bitcoin

import (
	"bytes"
	"errors"
	"math/big"
	"strings"

	"github.com/piwtach/blockchain/internal/crypto"
	"github.com/piwtach/blockchain/internal/encoding"
)

const base58Symbols = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var (
	num58               = big.NewInt(58)
	base58SymbolToIndex = make(map[rune]int, len(base58Symbols))
)

func init() {
	for i, a := range base58Symbols {
		base58SymbolToIndex[a] = i
	}
}

// Base58CheckEncode encodes a byte array into a Bitcoin address.
func Base58CheckEncode(data []byte) string {
	csum := crypto.Hash256(data)[:4]
	dataCsum := bytes.Join([][]byte{data, csum}, nil)

	var cnt int
	for _, d := range dataCsum {
		if d != 0x00 {
			break
		}
		cnt++
	}

	num := encoding.IntFromBigEndian(dataCsum[cnt:])
	mod := new(big.Int)
	var out strings.Builder

	for num.Sign() == 1 {
		num, mod = num.DivMod(num, num58, mod)
		sym := base58Symbols[mod.Int64()]
		out.WriteByte(sym)
	}
	return strings.Repeat("1", cnt) + reverseStr(out.String())
}

// Base58CheckDecode decodes a Bitcoin address into a byte array.
func Base58CheckDecode(addr string) ([]byte, error) {
	const (
		addrSize = 25
		csumSize = 4
	)

	num := new(big.Int)
	for _, c := range addr {
		num.Mul(num, num58)
		if idx, ok := base58SymbolToIndex[c]; ok {
			num.Add(num, big.NewInt(int64(idx)))
		}
	}

	bs, err := encoding.IntToBigEndian(num, addrSize)
	if err != nil {
		return nil, err
	}

	csumStartIdx := len(bs) - csumSize
	data := bs[:csumStartIdx]

	expCsum := bs[csumStartIdx:]
	actCsum := crypto.Hash256(data)[:csumSize]
	if !bytes.Equal(expCsum, actCsum) {
		return nil, errors.New("base58check: bad checksum")
	}

	return bs[1:csumStartIdx], nil
}

func reverseStr(str string) string {
	out := []rune(str)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return string(out)
}
