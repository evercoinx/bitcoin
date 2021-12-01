package encoding

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"
)

func TestIntToBytes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		endianness Endianness
		num        *big.Int
		size       int
		want       string
	}{
		{
			"bigendian 32-byte zero",
			BigEndian,
			big.NewInt(0),
			32,
			"0000000000000000000000000000000000000000000000000000000000000000",
		},
		{
			"littleendian 32-byte zero",
			LittleEndian,
			big.NewInt(0),
			32,
			"0000000000000000000000000000000000000000000000000000000000000000",
		},
		{
			"bigendian 32-byte one",
			BigEndian,
			big.NewInt(1),
			32,
			"0000000000000000000000000000000000000000000000000000000000000001",
		},
		{
			"littleendian 32-byte one",
			LittleEndian,
			big.NewInt(1),
			32,
			"0100000000000000000000000000000000000000000000000000000000000000",
		},
		{
			"small bigendian 32-byte number",
			BigEndian,
			big.NewInt(5000),
			32,
			"0000000000000000000000000000000000000000000000000000000000001388",
		},
		{
			"small littleendian 32-byte number",
			LittleEndian,
			big.NewInt(5000),
			32,
			"8813000000000000000000000000000000000000000000000000000000000000",
		},
		{
			"big bigendian 32-byte number",
			BigEndian,
			big.NewInt(3917405024756549),
			32,
			"000000000000000000000000000000000000000000000000000deadbeef12345",
		},
		{
			"big littleendian 32-byte number",
			LittleEndian,
			big.NewInt(3917405024756549),
			32,
			"4523f1eedbea0d00000000000000000000000000000000000000000000000000",
		},
		{
			"very big bigendian 32-byte number",
			BigEndian,
			new(big.Int).Exp(big.NewInt(2018), big.NewInt(5), nil),
			32,
			"0000000000000000000000000000000000000000000000000076e54a40efb620",
		},
		{
			"very big littleendian 32-byte number",
			LittleEndian,
			new(big.Int).Exp(big.NewInt(2018), big.NewInt(5), nil),
			32,
			"20b6ef404ae57600000000000000000000000000000000000000000000000000",
		},
		{
			"another very big bigendian 32-byte number",
			BigEndian,
			new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil),
			32,
			"0000000000000000000000000000000100000000000000000000000000000000",
		},
		{
			"another very big littleendian 32-byte number",
			LittleEndian,
			new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil),
			32,
			"0000000000000000000000000000000001000000000000000000000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := hex.DecodeString(tt.want)
			if err != nil {
				t.Fatal(err)
			}

			endianness := BigEndian
			if tt.endianness == LittleEndian {
				endianness = LittleEndian
			}

			got, err := IntToBytes(tt.num, tt.size, endianness)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(want, got) {
				t.Fatalf("result mismatch: %s != %x", tt.want, got)
			}
		})
	}
}

func TestBytesToInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		num        string
		endianness Endianness
		want       *big.Int
	}{
		{
			"bigendian 32-byte zero",
			"0000000000000000000000000000000000000000000000000000000000000000",
			BigEndian,
			big.NewInt(0),
		},
		{
			"littleendian 32-byte zero",
			"0000000000000000000000000000000000000000000000000000000000000000",
			LittleEndian,
			big.NewInt(0),
		},
		{
			"bigendian 32-byte one",
			"0000000000000000000000000000000000000000000000000000000000000001",
			BigEndian,
			big.NewInt(1),
		},
		{
			"littleendian 32-byte one",
			"0100000000000000000000000000000000000000000000000000000000000000",
			LittleEndian,
			big.NewInt(1),
		},
		{
			"small bigendian 32-byte number",
			"0000000000000000000000000000000000000000000000000000000000001388",
			BigEndian,
			big.NewInt(5000),
		},
		{
			"small littleendian 32-byte number",
			"8813000000000000000000000000000000000000000000000000000000000000",
			LittleEndian,
			big.NewInt(5000),
		},
		{
			"big bigendian 32-byte number",
			"000000000000000000000000000000000000000000000000000deadbeef12345",
			BigEndian,
			big.NewInt(3917405024756549),
		},
		{
			"big littleendian 32-byte number",
			"4523f1eedbea0d00000000000000000000000000000000000000000000000000",
			LittleEndian,
			big.NewInt(3917405024756549),
		},
		{
			"very big bigendian 32-byte number",
			"0000000000000000000000000000000000000000000000000076e54a40efb620",
			BigEndian,
			new(big.Int).Exp(big.NewInt(2018), big.NewInt(5), nil),
		},
		{
			"very big littleendian 32-byte number",
			"20b6ef404ae57600000000000000000000000000000000000000000000000000",
			LittleEndian,
			new(big.Int).Exp(big.NewInt(2018), big.NewInt(5), nil),
		},
		{
			"another very big bigendian 32-byte number",
			"0000000000000000000000000000000100000000000000000000000000000000",
			BigEndian,
			new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil),
		},
		{
			"another very big littleendian 32-byte number",
			"0000000000000000000000000000000001000000000000000000000000000000",
			LittleEndian,
			new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil),
		},
	}

	for _, tt := range tests {
		num, err := hex.DecodeString(tt.num)
		if err != nil {
			t.Fatal(err)
		}

		endianness := BigEndian
		if tt.endianness == LittleEndian {
			endianness = LittleEndian
		}

		got := BytesToInt(num, endianness)
		if tt.want.Cmp(got) != 0 {
			t.Fatalf("result match: %s != %s", tt.want.Text(16), got.Text(16))
		}
	}
}
