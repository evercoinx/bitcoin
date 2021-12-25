package crypto

import (
	"crypto/elliptic"
	"math/big"

	"github.com/evercoinx/kit/math"
)

var (
	one   = big.NewInt(1)
	two   = big.NewInt(2)
	three = big.NewInt(3)
)

type ellipticCurve struct {
	f       math.GaloisField
	n, a, b *big.Int
	gx, gy  *big.Int
	bitSize int
	name    string
}

func NewEllipticCurve(p, n, a, b, gx, gy *big.Int, bitSize int, name string) elliptic.Curve {
	return &ellipticCurve{
		f:       math.NewGaloisField(p),
		n:       n,
		a:       a,
		b:       b,
		gx:      gx,
		gy:      gy,
		bitSize: bitSize,
		name:    name,
	}
}

func (c *ellipticCurve) Params() *elliptic.CurveParams {
	return &elliptic.CurveParams{
		P:       c.f.Order,
		N:       c.n,
		B:       c.b,
		Gx:      c.gx,
		Gy:      c.gy,
		BitSize: c.bitSize,
		Name:    c.name,
	}
}

// IsOnCurve reports whether the given point (x,y) lies on the curve.
func (c *ellipticCurve) IsOnCurve(x, y *big.Int) bool {
	if x == nil {
		return true
	}

	lhs := c.f.Exp(y, two)
	rhs := c.f.Add(
		c.f.Exp(x, three),
		c.f.Mul(c.a, x),
		c.b,
	)
	return lhs.Cmp(rhs) == 0
}

// Add returns the sum of points (x1,y1) and (x2,y2).
func (c *ellipticCurve) Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int) {
	if x1 == nil && x2 == nil {
		return
	}

	if x1 == nil {
		x = x2
		y = y2
		return
	}

	if x2 == nil {
		x = x1
		y = y1
		return
	}

	// a line intersects the curve at three points
	if x1.Cmp(x2) != 0 {
		slope := c.f.Div(
			c.f.Sub(y2, y1),
			c.f.Sub(x2, x1),
		)
		x = c.calculateXCoord(slope, x1, x2)
		y = c.calculateYCoord(slope, x1, y1, x)
		return
	}

	// a vertical line intersects the curve at two points
	if y1.Cmp(y2) != 0 || y1.Sign() == 0 {
		return
	}

	// a tangent line intersects the curve at two points
	slope := c.f.Div(
		c.f.Add(
			c.f.Mul(three, c.f.Exp(x1, two)),
			c.a,
		),
		c.f.Mul(two, y1),
	)
	x = c.calculateXCoord(slope, x1, x1)
	y = c.calculateYCoord(slope, x1, y1, x)
	return
}

// calculateXCoord returns the x coordinate of a point of intersection.
func (c *ellipticCurve) calculateXCoord(slope, x1, x2 *big.Int) *big.Int {
	return c.f.Sub(
		c.f.Exp(slope, two),
		x1,
		x2,
	)
}

// calculateYCoord returns the y coordinate of a point of intersection.
func (c *ellipticCurve) calculateYCoord(slope, x1, y1, x *big.Int) *big.Int {
	return c.f.Sub(
		c.f.Mul(
			slope,
			c.f.Sub(x1, x),
		),
		y1,
	)
}

// Double returns 2*(x,y).
func (c *ellipticCurve) Double(x1, y1 *big.Int) (x, y *big.Int) {
	return c.Add(x1, y1, x1, y1)
}

// ScalarMult returns k*(Bx,By) where k is a number in big-endian form.
func (c *ellipticCurve) ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int) {
	if x1 == nil {
		return
	}

	kn := new(big.Int).SetBytes(k)
	kn.Mod(kn, c.n)

	for kn.Sign() != 0 {
		if new(big.Int).And(kn, one).Sign() == 1 {
			x, y = c.Add(x, y, x1, y1)
		}
		x1, y1 = c.Double(x1, y1)
		kn.Rsh(kn, 1)
	}
	return
}

// ScalarBaseMult returns k*G, where G is the base point of the group and k is
// a number in big-endian form.
func (c *ellipticCurve) ScalarBaseMult(k []byte) (x, y *big.Int) {
	return c.ScalarMult(c.gx, c.gy, k)
}
