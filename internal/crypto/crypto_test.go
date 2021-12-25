package crypto_test

import (
	"math/big"
	"testing"

	"github.com/evercoinx/bitcoin/internal/crypto"
)

func TestSecp256k1(t *testing.T) {
	curve := crypto.Secp256k1()

	tests := []struct {
		name         string
		k            *big.Int
		wantX, wantY string
	}{
		{
			name:  "N*G",
			k:     curve.Params().N,
			wantX: "",
			wantY: "",
		},
		{
			name:  "7*G",
			k:     big.NewInt(7),
			wantX: "5cbdf0646e5db4eaa398f365f2ea7a0e3d419b7e0330e39ce92bddedcac4f9bc",
			wantY: "6aebca40ba255960a3178d6d861a54dba813d0b813fde7b5a5082628087264da",
		},
		{
			name:  "1485*G",
			k:     big.NewInt(1485),
			wantX: "c982196a7466fbbbb0e27a940b6af926c1a74d5ad07128c82824a11b5398afda",
			wantY: "7a91f9eae64438afb9ce6448a1c133db2d8fb9254e4546b6f001637d50901f55",
		},
		{
			name:  "2^128*G",
			k:     new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil),
			wantX: "8f68b9d2f63b5f339239c1ad981f162ee88c5678723ea3351b7b444c9ec4c0da",
			wantY: "662a9f2dba063986de1d90c2b6be215dbbea2cfe95510bfdf23cbf79501fff82",
		},
		{
			name: "(2^240+2^31)*G",
			k: new(big.Int).Add(
				new(big.Int).Exp(big.NewInt(2), big.NewInt(240), nil),
				new(big.Int).Exp(big.NewInt(2), big.NewInt(31), nil),
			),
			wantX: "9577ff57c8234558f293df502ca4f09cbc65a6572c842b39b366f21717945116",
			wantY: "10b49c67fa9365ad7b90dab070be339a1daf9052373ec30ffae4f72d5e66d053",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wantX *big.Int
			if tt.wantX != "" {
				wantX, _ = new(big.Int).SetString(tt.wantX, 16)
			}

			var wantY *big.Int
			if tt.wantY != "" {
				wantY, _ = new(big.Int).SetString(tt.wantY, 16)
			}

			x, y := curve.ScalarBaseMult(tt.k.Bytes())
			if x.Cmp(wantX) != 0 || y.Cmp(wantY) != 0 {
				t.Fatalf("(%s,%s) != (%s,%s)", x.Text(16), y.Text(16), tt.wantX, tt.wantY)
			}
		})
	}
}
