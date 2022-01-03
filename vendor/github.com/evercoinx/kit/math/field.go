package math

import (
	"fmt"
	"math/big"
)

var (
	one = big.NewInt(1)
	two = big.NewInt(2)
)

// GaloisField describes a finite field of integers with an order which is
// a prime number.
type GaloisField struct {
	Order               *big.Int
	maxElement          *big.Int
	multInverseExponent *big.Int
}

func NewGaloisField(order *big.Int) GaloisField {
	if order.Sign() != 1 {
		panic(fmt.Sprintf("field order %v must be positive", order))
	}

	return GaloisField{
		Order:               order,
		maxElement:          new(big.Int).Sub(order, one),
		multInverseExponent: new(big.Int).Sub(order, two),
	}
}

// Add returns the sum of modular addition of two or more field elements as
// (x1+x2+...+xN) % p where p is a field order.
func (f *GaloisField) Add(elems ...*big.Int) *big.Int {
	f.validateFieldElements(elems)

	res := new(big.Int).Set(elems[0])
	for _, e := range elems[1:] {
		res.Add(res, e).Mod(res, f.Order)
	}
	return res
}

// Sub returns the difference of modular subtraction of two or more field
// elements as (x1-x2-...-xN) % p where p is a field order.
func (f *GaloisField) Sub(elems ...*big.Int) *big.Int {
	f.validateFieldElements(elems)

	res := new(big.Int).Set(elems[0])
	for _, el := range elems[1:] {
		res.Sub(res, el).Mod(res, f.Order)
	}
	return res
}

// Mul returns the result of modular multiplication of two or more field
// elements as (x1*x2*...*xN) % p where p is a field order.
func (f *GaloisField) Mul(elems ...*big.Int) *big.Int {
	f.validateFieldElements(elems)

	res := new(big.Int).Set(elems[0])
	for _, e := range elems[1:] {
		res.Mul(res, e).Mod(res, f.Order)
	}
	return res
}

// Div returns the quotient of modular division of two or more field elements as
// (x1/x2/.../xN) % p where p is a field order.
//
// Division is the inverse of multiplication so that x/y = x*y^-1. According to
// Fermat's little theorem x^(p-1) % p = 1 the multiplicative inverse of a field
// element y can be expressed as y^-1 = y^(p-2).
func (f *GaloisField) Div(elems ...*big.Int) *big.Int {
	f.validateFieldElements(elems)

	res := new(big.Int).Set(elems[0])
	for _, e := range elems[1:] {
		if e.Sign() == 0 {
			panic("division by zero is undefined")
		}
		inverseElem := new(big.Int).Exp(e, f.multInverseExponent, f.Order)
		res.Mul(res, inverseElem).Mod(res, f.Order)
	}
	return res
}

// Exp returns the result of modular exponentiation of a field element as
// x^n % p where p is a field order.
//
// By applying the modulo operation we can force out a negative exponent to be
// positive reducing it to the range within 0 and p-2. According to Fermat's
// little thereom x^(p-1) % p = 1 we can derive that x^-n = x^(p-n-1).
func (f *GaloisField) Exp(elem *big.Int, exponent *big.Int) *big.Int {
	f.validateFieldElement(elem)
	if exponent == nil {
		panic("exponent is undefined")
	}

	if elem.Sign() == 0 {
		if exponent.Sign() == 0 {
			panic("zero raised to zero is undefined")
		}
		if exponent.Sign() == -1 {
			panic("power of zero is undefined for negative exponent")
		}
		return big.NewInt(0)
	}

	res := new(big.Int).Set(elem)
	positiveExp := new(big.Int).Mod(exponent, f.maxElement)
	return res.Exp(res, positiveExp, f.Order)
}

func (f *GaloisField) validateFieldElements(elems []*big.Int) {
	if len(elems) < 2 {
		panic("not enough operands")
	}
	for _, e := range elems {
		f.validateFieldElement(e)
	}
}

func (f *GaloisField) validateFieldElement(elem *big.Int) {
	switch {
	case elem == nil:
		panic("field element is undefined")
	case elem.Sign() == -1:
		panic("field element is negative")
	case elem.Cmp(f.maxElement) == 1:
		panic("invalid field element")
	}
}
