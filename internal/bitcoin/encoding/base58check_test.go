package encoding

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestBase58CheckEncode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		payload string
		version AddressVersion
		want    string
	}{
		{
			"public key hash",
			"5f2613791b36f667fdb8e95608b55e3df4c5f9eb",
			AddressVersionPublicKeyHash,
			"19g6oo8foQF5jfqK9gH2bLkFNwgCenRBPD",
		},
		{
			"script hash",
			"04e214163b3b927c3d2058171dd66ff6780f8708",
			AddressVersionScriptHash,
			"328qTX1KYxMohp4MjPPEDBoRomCGwrB2ag",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, err := hex.DecodeString(tt.payload)
			if err != nil {
				t.Fatal(err)
			}

			ver := AddressVersionPublicKeyHash
			if tt.version == AddressVersionScriptHash {
				ver = AddressVersionScriptHash
			}

			got := Base58CheckEncode(payload, ver)
			if tt.want != got {
				t.Fatalf("result mismatch: %s != %s", tt.want, got)
			}
		})
	}
}

func TestBase58CheckDecode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		address string
		want    string
	}{
		{
			"p2pkh address",
			"19g6oo8foQF5jfqK9gH2bLkFNwgCenRBPD",
			"5f2613791b36f667fdb8e95608b55e3df4c5f9eb",
		},
		{
			"p2sh address",
			"328qTX1KYxMohp4MjPPEDBoRomCGwrB2ag",
			"04e214163b3b927c3d2058171dd66ff6780f8708",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := hex.DecodeString(tt.want)
			if err != nil {
				t.Fatal(err)
			}

			got, err := Base58CheckDecode(tt.address)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(want, got) {
				t.Fatalf("result mismatch: %s != %x", tt.want, got)
			}
		})
	}
}
