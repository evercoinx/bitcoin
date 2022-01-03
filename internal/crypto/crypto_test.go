package crypto_test

import (
	"math/big"
	"testing"

	"github.com/evercoinx/bitcoin/internal/crypto"
)

func TestSecp256k1(t *testing.T) {
	curve := crypto.Secp256k1()

	t.Run("params", func(t *testing.T) {
		p := new(big.Int).Sub(
			new(big.Int).Sub(
				new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil),
				new(big.Int).Exp(big.NewInt(2), big.NewInt(32), nil),
			),
			big.NewInt(977),
		)
		n, _ := new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
		gx, _ := new(big.Int).SetString("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 16)
		gy, _ := new(big.Int).SetString("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 16)

		test := struct {
			b, p      *big.Int
			n, gx, gy *big.Int
			bitSize   int
			name      string
		}{
			b:       big.NewInt(7),
			p:       p,
			n:       n,
			gx:      gx,
			gy:      gy,
			bitSize: 256,
			name:    "Secp256k1",
		}

		got := curve.Params()
		if got.P.Cmp(test.p) != 0 {
			t.Fatalf("p: %x != %x", got.P, test.p)
		}
		if got.N.Cmp(test.n) != 0 {
			t.Fatalf("n: %x != %x", got.N, test.n)
		}
		if got.B.Cmp(test.b) != 0 {
			t.Fatalf("b: %d != %d", got.B, test.b)
		}
		if got.Gx.Cmp(test.gx) != 0 {
			t.Fatalf("gx: %x != %x", got.Gx, test.gx)
		}
		if got.Gy.Cmp(test.gy) != 0 {
			t.Fatalf("gy: %x != %x", got.Gy, test.gy)
		}
		if got.BitSize != test.bitSize {
			t.Fatalf("bit size: %d != %d", got.BitSize, test.bitSize)
		}
		if got.Name != test.name {
			t.Fatalf("name: %s != %s", got.Name, test.name)
		}
	})

	t.Run("is on curve", func(t *testing.T) {
		tests := []struct {
			name string
			x, y string
			want bool
		}{
			{
				name: "I",
				x:    "",
				y:    "",
				want: true,
			},
			{
				name: "G",
				x:    curve.Params().Gx.Text(16),
				y:    curve.Params().Gy.Text(16),
				want: true,
			},
			{
				name: "7*G",
				x:    "5cbdf0646e5db4eaa398f365f2ea7a0e3d419b7e0330e39ce92bddedcac4f9bc",
				y:    "6aebca40ba255960a3178d6d861a54dba813d0b813fde7b5a5082628087264da",
				want: true,
			},
			{
				name: "1485*G",
				x:    "c982196a7466fbbbb0e27a940b6af926c1a74d5ad07128c82824a11b5398afda",
				y:    "7a91f9eae64438afb9ce6448a1c133db2d8fb9254e4546b6f001637d50901f55",
				want: true,
			},
			{
				name: "2^128*G",
				x:    "8f68b9d2f63b5f339239c1ad981f162ee88c5678723ea3351b7b444c9ec4c0da",
				y:    "662a9f2dba063986de1d90c2b6be215dbbea2cfe95510bfdf23cbf79501fff82",
				want: true,
			},
			{
				name: "(2^240+2^31)*G",
				x:    "9577ff57c8234558f293df502ca4f09cbc65a6572c842b39b366f21717945116",
				y:    "10b49c67fa9365ad7b90dab070be339a1daf9052373ec30ffae4f72d5e66d053",
				want: true,
			},
			{
				name: "0*G",
				x:    "0000000000000000000000000000000000000000000000000000000000000000",
				y:    "0000000000000000000000000000000000000000000000000000000000000000",
				want: false,
			},
			{
				name: "7*G-1",
				x:    "5cbdf0646e5db4eaa398f365f2ea7a0e3d419b7e0330e39ce92bddedcac4f9bb",
				y:    "6aebca40ba255960a3178d6d861a54dba813d0b813fde7b5a5082628087264db",
				want: false,
			},
			{
				name: "1485*G-1",
				x:    "c982196a7466fbbbb0e27a940b6af926c1a74d5ad07128c82824a11b5398afdb",
				y:    "7a91f9eae64438afb9ce6448a1c133db2d8fb9254e4546b6f001637d50901f54",
				want: false,
			},
			{
				name: "2^128*G-1",
				x:    "8f68b9d2f63b5f339239c1ad981f162ee88c5678723ea3351b7b444c9ec4c0db",
				y:    "662a9f2dba063986de1d90c2b6be215dbbea2cfe95510bfdf23cbf79501fff81",
				want: false,
			},
			{
				name: "(2^240+2^31)*G-1",
				x:    "9577ff57c8234558f293df502ca4f09cbc65a6572c842b39b366f21717945115",
				y:    "10b49c67fa9365ad7b90dab070be339a1daf9052373ec30ffae4f72d5e66d054",
				want: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				var x *big.Int
				if tt.x != "" {
					x, _ = new(big.Int).SetString(tt.x, 16)
				}

				var y *big.Int
				if tt.y != "" {
					y, _ = new(big.Int).SetString(tt.y, 16)
				}

				got := curve.IsOnCurve(x, y)
				if got != tt.want {
					t.Fatalf("%t != %t", got, tt.want)
				}
			})
		}
	})

	t.Run("scalar base multiplication", func(t *testing.T) {
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
	})
}
