package crypto

import (
	"crypto/elliptic"
	"math/big"

	"github.com/evercoinx/kit/crypto"
)

var secp256k1 elliptic.Curve

func init() {
	p, _ := new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f", 16)
	n, _ := new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	a := big.NewInt(0)
	b := big.NewInt(7)
	gx, _ := new(big.Int).SetString("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 16)
	gy, _ := new(big.Int).SetString("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 16)
	bitSize := 256
	name := "Secp256k1"

	secp256k1 = crypto.NewEllipticCurve(p, n, a, b, gx, gy, bitSize, name)
}

func Secp256k1() elliptic.Curve {
	return secp256k1
}
