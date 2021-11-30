package encoding

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"
)

func TestIntToBigEndian(t *testing.T) {
	tests := []struct {
		num  *big.Int
		size int
		want string
	}{
		{
			big.NewInt(0),
			32,
			"0000000000000000000000000000000000000000000000000000000000000000",
		},
		{
			big.NewInt(1),
			32,
			"0000000000000000000000000000000000000000000000000000000000000001",
		},
		{
			big.NewInt(5000),
			32,
			"0000000000000000000000000000000000000000000000000000000000001388",
		},
		{
			big.NewInt(3917405024756549),
			32,
			"000000000000000000000000000000000000000000000000000deadbeef12345",
		},
		{
			new(big.Int).Exp(big.NewInt(2018), big.NewInt(5), nil),
			32,
			"0000000000000000000000000000000000000000000000000076e54a40efb620",
		},
		{
			new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil),
			32,
			"0000000000000000000000000000000100000000000000000000000000000000",
		},
	}

	for _, tt := range tests {
		want, err := hex.DecodeString(tt.want)
		if err != nil {
			t.Fatal(err)
		}

		got, err := IntToBigEndian(tt.num, tt.size)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(got, want) {
			t.Fatalf("%x != %s", got, tt.want)
		}
	}
}

func TestIntFromBigEndian(t *testing.T) {
	tests := []struct {
		num  string
		want *big.Int
	}{
		{
			"0000000000000000000000000000000000000000000000000000000000000000",
			big.NewInt(0),
		},
		{
			"0000000000000000000000000000000000000000000000000000000000000001",
			big.NewInt(1),
		},
		{
			"0000000000000000000000000000000000000000000000000000000000001388",
			big.NewInt(5000),
		},
		{
			"000000000000000000000000000000000000000000000000000deadbeef12345",
			big.NewInt(3917405024756549),
		},
		{
			"0000000000000000000000000000000000000000000000000076e54a40efb620",
			new(big.Int).Exp(big.NewInt(2018), big.NewInt(5), nil),
		},
		{
			"0000000000000000000000000000000100000000000000000000000000000000",
			new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil),
		},
	}

	for _, tt := range tests {
		num, err := hex.DecodeString(tt.num)
		if err != nil {
			t.Fatal(err)
		}

		got := IntFromBigEndian(num)
		if got.Cmp(tt.want) != 0 {
			t.Fatalf("%s != %s", got.Text(16), tt.want.Text(16))
		}
	}
}

func TestToLittleEndian(t *testing.T) {
	tests := []struct {
		num  *big.Int
		size int
		want string
	}{
		{
			big.NewInt(0),
			32,
			"0000000000000000000000000000000000000000000000000000000000000000",
		},
		{
			big.NewInt(1),
			32,
			"0100000000000000000000000000000000000000000000000000000000000000",
		},
		{
			big.NewInt(5000),
			32,
			"8813000000000000000000000000000000000000000000000000000000000000",
		},
		{
			big.NewInt(3917405024756549),
			32,
			"4523f1eedbea0d00000000000000000000000000000000000000000000000000",
		},
		{
			new(big.Int).Exp(big.NewInt(2018), big.NewInt(5), nil),
			32,
			"20b6ef404ae57600000000000000000000000000000000000000000000000000",
		},
		{
			new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil),
			32,
			"0000000000000000000000000000000001000000000000000000000000000000",
		},
	}

	for _, tt := range tests {
		want, err := hex.DecodeString(tt.want)
		if err != nil {
			t.Fatal(err)
		}

		got, err := IntToLittleEndian(tt.num, tt.size)
		if err != nil {
			t.Fatal(err)
		}
		if bytes.Compare(got, want) != 0 {
			t.Fatalf("%x != %s", got, tt.want)
		}
	}
}

func TestFromLittleEndian(t *testing.T) {
	tests := []struct {
		num  string
		want *big.Int
	}{
		{
			"0000000000000000000000000000000000000000000000000000000000000000",
			big.NewInt(0),
		},
		{
			"0100000000000000000000000000000000000000000000000000000000000000",
			big.NewInt(1),
		},
		{
			"8813000000000000000000000000000000000000000000000000000000000000",
			big.NewInt(5000),
		},
		{
			"4523f1eedbea0d00000000000000000000000000000000000000000000000000",
			big.NewInt(3917405024756549),
		},
		{
			"20b6ef404ae57600000000000000000000000000000000000000000000000000",
			new(big.Int).Exp(big.NewInt(2018), big.NewInt(5), nil),
		},
		{
			"0000000000000000000000000000000001000000000000000000000000000000",
			new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil),
		},
	}

	for _, tt := range tests {
		num, err := hex.DecodeString(tt.num)
		if err != nil {
			t.Fatal(err)
		}

		got := IntFromLittleEndian(num)
		if got.Cmp(tt.want) != 0 {
			t.Fatalf("%s != %s", got.Text(16), tt.want.Text(16))
		}
	}
}
