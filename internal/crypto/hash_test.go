package crypto

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestHash160(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		data   []byte
		digest string
	}{
		{
			"empty data",
			[]byte(""),
			"b472a266d0bd89c13706a4132ccfb16f7c3b9fcb",
		},
		{
			"alphanumeric data",
			[]byte("Test123"),
			"08bb2e16081143d40acc108effa26fcad752e64e",
		},
		{
			"punctuation data",
			[]byte("Hello, world!"),
			"8d159f1c4f99d8ed858f7832310db31cb91e0745",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			digest, err := hex.DecodeString(tt.digest)
			if err != nil {
				t.Fatal(err)

			}

			res := Hash160(tt.data)
			if len(res) != 20 {
				t.Fatalf("invalid digest length: %d", len(res))
			}

			if !bytes.Equal(res, digest) {
				t.Fatalf("digests mismatch: %x != %x", digest, res)
			}
		})
	}
}

func TestHash256(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		data   []byte
		digest string
	}{
		{
			"empty data",
			[]byte(""),
			"5df6e0e2761359d30a8275058e299fcc0381534545f55cf43e41983f5d4c9456",
		},
		{
			"alphanumeric data",
			[]byte("Test123"),
			"b87660920cb5b8127fd409e87a2c093ae12444a6556752e5320a9f76c8bbf4da",
		},
		{
			"punctuation data",
			[]byte("Hello, world!"),
			"6246efc88ae4aa025e48c9c7adc723d5c97171a1fa6233623c7251ab8e57602f",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			digest, err := hex.DecodeString(tt.digest)
			if err != nil {
				t.Fatal(err)
			}

			res := Hash256(tt.data)
			if len(res) != 32 {
				t.Fatalf("invalid digest length: %d", len(res))
			}

			if !bytes.Equal(res, digest) {
				t.Fatalf("digests mismatch: %x != %x", digest, res)
			}
		})
	}
}
