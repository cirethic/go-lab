package main

/*

cantor  pair function  reference

 https://en.wikipedia.org/wiki/Cantor_function
 https://en.wikipedia.org/wiki/Pairing_function
 https://gist.github.com/hannesl/8031402

*/

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	a1 := cantor_pair_calculate(9, 1) // first pair 1

	a2 := cantor_pair_calculate(6, 8) // second pair

	a3 := cantor_pair_calculate(a1, a2) // third pair = first , second

	x1, x2 := cantor_pair_reverse(a3) // reverse third pair
	cantor_pair_reverse(x1)           // reverse first pair
	cantor_pair_reverse(x2)           // reverse second  pair

	return

}
func cantor_pair_calculate2(x *big.Int, y *big.Int) *big.Int {
	//result := ((x+y)*(x+y+1))/2 + y

	x.Add(x, y).Mul(x, new(big.Int).Add(x, big.NewInt(1))).Div(x, big.NewInt(2)).Add(x, y)
	fmt.Println(x)

	return x
}

/**
 * Return the source integers from a cantor pair integer.
 */
func cantor_pair_reverse2(z *big.Int) {

	negOne := big.NewFloat(-1)
	eight := big.NewFloat(8)
	one := big.NewFloat(1)
	two := big.NewFloat(2)
	three := big.NewFloat(3)
	zFloat := new(big.Float)
	zFloat.SetString(z.String())

	//t := math.floor((-1 + math.sqrt(1+8*z)) / 2)
	zFloat8 := new(big.Float).Mul(zFloat, eight)
	z81 := new(big.Float).Add(one, zFloat8)
	sqrtZ81 := new(big.Float).Sqrt(z81)
	sqrtZ81S1 := new(big.Float).Add(negOne, sqrtZ81)
	fmt.Println("z81.Sqrt(z81).Quo(z81, two)=", sqrtZ81S1)
	t := new(big.Float).Quo(sqrtZ81S1, two)

	fmt.Println("t:", t)

	t = big.NewFloat(24)

	//x := t*(t+3)/2 - float64(z)
	tx := new(big.Float)
	tx.SetString(t.String())

	t3 := new(big.Float).Add(tx, three)

	tx.Mul(tx, t3).Quo(tx, two)
	fmt.Println("t*(t+3)/2:", tx)

	tx.Sub(tx, zFloat)

	x := tx.Mul(tx, new(big.Float).Add(tx, three).Quo(tx, two))
	x.Sub(x, zFloat)

	//y := z - t*(t+1)/2
	ty := new(big.Float)
	ty.SetString(t.String())
	y := ty.Add(zFloat, ty.Mul(new(big.Float).Add(ty, one), t).Quo(ty, two))
	fmt.Println(x, y)
}

/**
Cantor pairing functions in PHP. Pass any two positive integers and get a unique integer back. Feed the unique integer back into the reverse function and get the original integers back.
https://gist.github.com/hannesl/8031402

 * Calculate a unique integer based on two integers (cantor pairing).
*/
func cantor_pair_calculate(x int64, y int64) int64 {

	fmt.Println(x, y)

	result := ((x+y)*(x+y+1))/2 + y

	fmt.Println(result)

	return result
}

/**
 * Return the source integers from a cantor pair integer.
 */
func cantor_pair_reverse(z int64) (int64, int64) {

	t := math.Floor((-1 + math.Sqrt(1+8*float64(z))) / 2)
	x := t*(t+3)/2 - float64(z)
	y := float64(z) - t*(t+1)/2

	fmt.Println(x, y)

	return int64(x), int64(y)

}
