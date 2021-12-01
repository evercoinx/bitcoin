package encoding

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestBase58CheckEncode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		data string
		want string
	}{
		{
			"case one",
			"7c076ff316692a3d7eb3c3bb0f8b1488cf72e1afcd929e29307032997a838a3d",
			"wdA2ffYs5cudrdkhFm5Ym94AuLvavacapuDBL2CAcvqYPkcvi",
		},
		{
			"case two",
			"eff69ef2b1bd93a66ed5219add4fb51e11a840f404876325a1e8ffe0529a2c",
			"Qwj1mwXNifQmo5VV2s587usAy4QRUviQsBxoe4EJXyWz4GBs",
		},
		{
			"case three",
			"c7207fee197d27c618aea621406f6bf5ef6fca38681d82b2f06fddbdce6feab6",
			"2WhRyzK3iKFveq4hvQ3VR9uau26t6qZCMhADPAVMeMR6VraBbX",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := hex.DecodeString(tt.data)
			if err != nil {
				t.Fatal(err)
			}

			got := Base58CheckEncode(data)
			if tt.want != got {
				t.Fatalf("result mismatch: %s != %s", tt.want, got)
			}
		})
	}
}

func TestBase58CheckDecode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		data string
		want string
	}{
		{
			"case one",
			"mzx5YhAH9kNHtcN481u6WkjeHjYtVeKVh2",
			"d52ad7ca9b3d096a38e752c2018e6fbc40cdf26f",
		},
		{
			"case two",
			"mnrVtF8DWjMu839VW3rBfgYaAfKk8983Xf",
			"507b27411ccf7f16f10297de6cef3f291623eddf",
		},
		{
			"case three",
			"miKegze5FQNCnGw6PKyqUbYUeBa4x2hFeM",
			"1ec51b3654c1f1d0f4929d11a1f702937eaf50c8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := hex.DecodeString(tt.want)
			if err != nil {
				t.Fatal(err)
			}

			got, err := Base58CheckDecode(tt.data)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(want, got) {
				t.Fatalf("result mismatch: %s != %x", tt.want, got)
			}
		})
	}
}
