package encoding

import (
	"bytes"
	"errors"
	"math/big"
	"strings"

	"github.com/evercoinx/bitcoin/internal/hash"
	"github.com/evercoinx/kit/encoding"
)

type AddressVersion byte

const (
	AddressVersionPublicKeyHash AddressVersion = 0x00
	AddressVersionScriptHash    AddressVersion = 0x05
)

func (v AddressVersion) String() string {
	return string(v)
}

const (
	base58Symbols = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	addressSize  = 25 // in bytes; 1B is a version, 20B is a payload, 4B is a checksum
	checksumSize = 4  // in bytes
)

var (
	num58               = big.NewInt(58)
	base58SymbolToIndex = make(map[rune]int, len(base58Symbols))
)

func init() {
	for i, a := range base58Symbols {
		base58SymbolToIndex[a] = i
	}
}

// Base58CheckEncode encodes a byte slice into a Bitcoin address.
func Base58CheckEncode(payload []byte, version AddressVersion) string {
	ver := []byte{byte(version)}
	verPayload := bytes.Join([][]byte{ver, payload}, nil)

	csum := hash.Hash256(verPayload)[:checksumSize]
	bs := bytes.Join([][]byte{verPayload, csum}, nil)

	var cnt int
	for _, b := range bs {
		if b != byte(AddressVersionPublicKeyHash) {
			break
		}
		cnt++
	}

	encoded := base58Encode(bs[cnt:])
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

// Base58CheckDecode decodes a Bitcoin address into a byte slice.
func Base58CheckDecode(addr string) ([]byte, error) {
	decoded, err := base58Decode(addr)
	if err != nil {
		return nil, err
	}

	csumStartIdx := len(decoded) - checksumSize
	verPayload := decoded[:csumStartIdx]

	expCsum := decoded[csumStartIdx:]
	actCsum := hash.Hash256(verPayload)[:checksumSize]
	if !bytes.Equal(expCsum, actCsum) {
		return nil, errors.New("base58check: bad checksum")
	}

	return decoded[1:csumStartIdx], nil
}

func base58Decode(addr string) ([]byte, error) {
	num := new(big.Int)

	for _, a := range addr {
		num.Mul(num, num58)
		if idx, ok := base58SymbolToIndex[a]; ok {
			num.Add(num, big.NewInt(int64(idx)))
		}
	}

	bs, err := encoding.IntToBytes(num, addressSize, encoding.BigEndian)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
