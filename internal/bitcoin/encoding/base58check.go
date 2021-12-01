package encoding

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

	encoded := base58Encode(dataCsum[cnt:])
	return strings.Repeat("1", cnt) + reverseStr(encoded)
}

func base58Encode(bs []byte) string {
	num := encoding.BytesToInt(bs, encoding.BigEndian)
	mod := new(big.Int)
	var out strings.Builder

	for num.Sign() == 1 {
		num.DivMod(num, num58, mod)
		sym := base58Symbols[mod.Int64()]
		out.WriteByte(sym)
	}
	return out.String()
}

func reverseStr(str string) string {
	out := []rune(str)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return string(out)
}

// Base58CheckDecode decodes a Bitcoin address into a byte array.
func Base58CheckDecode(addr string) ([]byte, error) {
	const csumSize = 4

	decoded, err := base58Decode(addr)
	if err != nil {
		return nil, err
	}
	csumStartIdx := len(decoded) - csumSize
	data := decoded[:csumStartIdx]

	expCsum := decoded[csumStartIdx:]
	actCsum := crypto.Hash256(data)[:csumSize]
	if !bytes.Equal(expCsum, actCsum) {
		return nil, errors.New("base58check: bad checksum")
	}

	return decoded[1:csumStartIdx], nil
}

func base58Decode(addr string) ([]byte, error) {
	const addrSize = 25
	num := new(big.Int)

	for _, c := range addr {
		num.Mul(num, num58)
		if idx, ok := base58SymbolToIndex[c]; ok {
			num.Add(num, big.NewInt(int64(idx)))
		}
	}

	bs, err := encoding.IntToBytes(num, addrSize, encoding.BigEndian)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
