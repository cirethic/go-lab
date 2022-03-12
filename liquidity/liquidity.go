package main

/*
converted from:
https://github.com/atiselsts/uniswap-v3-liquidity-math/blob/master/uniswap-v3-liquidity-math.py

More info read: ：
http://atiselsts.github.io/pdfs/uniswap-v3-liquidity-math.pdf

Liquidity math adapted from
https://github.com/Uniswap/uniswap-v3-periphery/blob/main/contracts/libraries/LiquidityAmounts.sol

See the technical note "Liquidity Math in Uniswap v3" and the Uniswap v3 whitepaper
for the description of the purpose of this code.


*/

import (
	"fmt"
	"math"
)

func main() {

	// price, min, max, x
	getY(3879.16, 1982.3, 19849, 1)
	//getY(0.9973, 0.9461092501, 1.045607201, 620)

	///  https://app.uniswap.org/#/pool/668
	//getY(1999.84, 1000.3, 3999.8, 99990)

	//price, max, x, y
	getMin(3879.6, 5000, 1, 23633)

	//	getMin(2000, 3000, 2, 4000)

	// weth/usdc
	getTicks(3879.10, 300, 5000, 18, 6)

	// usdt/usdc
	getTicks(0.9973, 0.9461092501, 1.045607201, 18, 18)

	// (p , a , b , x , y)
	getBestAmounts(0.9973, 0.9461092501, 1.045607201, 460809334302698715646/1e18, 177936101400509812460/1e18)

	//getBestAmounts(3879.97, 1982.6, 19849.6, 10, 3000)

	//price, P1, min, max, x, y
	//getBalance(2000, 2500, 1333.33, 3000, 2, 4000)
	getBalance(3879.6, 3979, 1879.6, 5879.6, 1, 4020)
	getBalance(0.9917, 0.9917, 0.49, 1.8, 185.1, 132.1)

	test_1()

	test_2()

}

func test_1() {
	fmt.Println("----------------------------------------------")

	fmt.Println("test case 1")
	fmt.Println("----------------------------------------------")

	p := 3879.10
	a := 300.01
	b := 5000.01
	x := float64(1)
	y := float64(4000)
	test(x, y, p, a, b)

}
func getTicks(p float64, a float64, b float64, xDecimals float64, yDecimals float64) {

	fmt.Println("----------------------------------------------")
	fmt.Printf("tetTicks: tick, tickLower, tickUpper by given price, min, and max value \n")
	fmt.Println("----------------------------------------------")

	//calc tick  p(i) = 1.0001i

	diffDecimals := math.Pow(10, xDecimals-yDecimals)

	// log(p , 1.0001)  ==  log(p)/ log(1.0001)
	tick := math.Log(p/diffDecimals) / math.Log(1.0001)
	tickLower := math.Log(a/diffDecimals) / math.Log(1.0001)
	tickUpper := math.Log(b/diffDecimals) / math.Log(1.0001)

	fmt.Printf("tick={%.f}\n", tick)
	fmt.Printf("tickLower={%.f}\n", tickLower)
	fmt.Printf("tickUpper={%.f}\n", tickUpper)
	fmt.Printf("\n")

}

func test_2() {
	fmt.Println("----------------------------------------------")
	fmt.Println("test case 2")
	fmt.Println("----------------------------------------------")
	p := 3227.02
	a := 1626.3
	b := 4846.3
	x := float64(1)
	y := 5096.06
	test(x, y, p, a, b)

}

//
// Test a known good combination of values against the functions provided above.
//
// Some errors are expected because:
//  -- the floating point math is meant for simplicity, not accurate calculations!
//  -- ticks and tick ranges are ignored for simplicity
//  -- the test values taken from Uniswap v3 UI and are approximate
//
func test(x float64, y float64, p float64, a float64, b float64) {
	sp := math.Sqrt(p) //p**0.5
	sa := math.Sqrt(a) // a** 0.5
	sb := math.Sqrt(b) // b**0.5

	L := get_liquidity(x, y, sp, sa, sb)
	fmt.Printf("L: {:%.2f}\n", L)

	ia := calculate_a1(L, sp, sb, x, y)
	error := 100.0 * (1 - ia/a)
	fmt.Printf("a: {:%.2f} vs {:%.2f}, error {:%.6f}\n", a, ia, error)

	ia = calculate_a2(sp, sb, x, y)
	error = 100.0 * (1 - ia/a)
	fmt.Printf("a: {:%.2f} vs {:%.2f}, error {:%.6f}\n", a, ia, error)

	ib := calculate_b1(L, sp, sa, x, y)
	error = 100.0 * (1 - ib/b)
	fmt.Printf("b: {:%.2f} vs {:%.2f}, error {:%.6f}\n", b, ib, error)

	ib = calculate_b2(sp, sa, x, y)
	error = 100.0 * (1 - ib/b)
	fmt.Printf("b: {:%.2f} vs {:%.2f}, error {:%.6f}\n", b, ib, error)

	c := sb / sp
	d := sa / sp

	ic := calculate_c(p, d, x, y)
	error = 100.0 * (1 - ic/c)
	fmt.Printf("c^2: {:%.2f} vs {:%.2f}, error {:%.6f}\n", c*c, ic*ic, error)

	id := calculate_d(p, c, x, y)
	error = 100.0 * (1 - math.Pow(id, 2)/math.Pow(d, 2))
	fmt.Printf("d^2: {:%.2f} vs {:%.2f}, error {:%.6f}\n", d*d, id*id, error)

	ix := calculate_x(L, sp, sb)
	error = 100.0 * (1 - ix/x)
	fmt.Printf("x: {:%.2f} vs {:%.2f}, error {:%.6f}\n", x, ix, error)

	iy := calculate_y(L, sp, sa)
	error = 100.0 * (1 - iy/y)
	fmt.Printf("y: {:%.2f} vs {:%.2f}, error {:%.6f}\n", y, iy, error)
	fmt.Printf("")

}

//
// Example 1 from the technical note
//
func getY(p float64, a float64, b float64, x float64) {

	fmt.Println("----------------------------------------------")
	fmt.Println("Example 1: how much of USDC I need when providing 2 ETH at this price and range?")
	fmt.Println("----------------------------------------------")

	// price := float64(3879.16)
	// min := float64(300)
	// max := float64(5000)
	// x := float64(1)

	sp := math.Sqrt(p) //p * *0.5
	sa := math.Sqrt(a) //a * *0.5
	sb := math.Sqrt(b) //b * *0.5

	fmt.Println("math.Pow(p,0.5)=", math.Pow(p, 0.5), ";sp=", sp)

	L := get_liquidity_0(x, sp, sb)
	y := calculate_y(L, sp, sa)
	fmt.Printf("amount of USDC y={:%.2f}\n", y)

	// demonstrate that with the calculated y value, the given range is correct
	c := sb / sp
	d := sa / sp
	ic := calculate_c(p, d, x, y)
	fmt.Printf("ic={:%f}\n", ic)

	id := calculate_d(p, c, x, y)
	fmt.Printf("id={:%f}\n", id)

	C := math.Pow(ic, 2) //ic * *2
	fmt.Printf("C={:%f}\n", C)

	D := math.Pow(id, 2) // id * *2
	fmt.Printf("D={:%f}\n", D)

	fmt.Printf("p_a={:%.2f} ({:%.2f} of P), p_b={:%.2f} ({:%.2f} of P)\n", D*p, D*100, C*p, C*100)
	fmt.Printf("\n")
}

//
// Example 2 from the technical note
// calc min price
//
func getMin(p float64, b float64, x float64, y float64) {
	fmt.Println("----------------------------------------------")
	fmt.Println("Example 2: I have 2 ETH and 4000 USDC, range top set to 3000 USDC. What's the bottom of the range?")
	fmt.Println("----------------------------------------------")

	sp := math.Sqrt(p) // p * *0.5
	sb := math.Sqrt(b) //b * *0.5

	a := calculate_a2(sp, sb, x, y)
	fmt.Printf("lower bound of the price p_a={:%.2f}\n", a)

}

//
// Example 3 from the technical note
// calc balances
//
func getBalance(p float64, P1 float64, a float64, b float64, x float64, y float64) (float64, float64) {
	fmt.Println("----------------------------------------------")
	fmt.Printf("Example 3: Using the position created in Example 2, what are asset balances at 2500 USDC per ETH?\n")
	fmt.Println("----------------------------------------------")

	sp := math.Sqrt(p) //p * *0.5
	sa := math.Sqrt(a) //a * *0.5
	sb := math.Sqrt(b) //b * *0.5
	// calculate the initial liquidity
	L := get_liquidity(x, y, sp, sa, sb)

	//P1 := float64(2500)
	sp1 := math.Sqrt(P1) // P1 * *0.5
	x1 := calculate_x(L, sp1, sb)
	y1 := calculate_y(L, sp1, sa)
	fmt.Printf("Amount of ETH x={:%.2f} amount of USDC y={:%.2f}\n", x1, y1)

	// alternative way, directly based on the whitepaper
	delta_p := sp1 - sp
	delta_inv_p := 1/sp1 - 1/sp
	delta_x := delta_inv_p * L
	delta_y := delta_p * L
	xx1 := x + delta_x
	yy1 := y + delta_y
	fmt.Printf("alternative way, directly based on the whitepaper:")
	fmt.Printf("delta_x={:%.2f} delta_y={:%.2f}\n", delta_x, delta_y)
	fmt.Printf("Amount of ETH x={:%.2f} amount of USDC y={:%.2f}\n", xx1, yy1)

	return x1, y1

}

func getBestAmounts(p float64, a float64, b float64, x float64, y float64) (amount0 float64, amount1 float64, swapAmount float64, zeroForOne bool) {
	fmt.Println("----------------------------------------------")
	fmt.Println("Swap Amounts: getBestAmounts and swap amount")
	fmt.Println("----------------------------------------------")

	sp := math.Sqrt(p) //p * *0.5
	sa := math.Sqrt(a) //a * *0.5
	sb := math.Sqrt(b) //b * *0.5
	// calculate the initial liquidity
	L := get_liquidity(x, y, sp, sa, sb)

	P1 := p
	sp1 := math.Sqrt(P1) // P1 * *0.5
	x1 := calculate_x(L, sp1, sb)
	y1 := calculate_y(L, sp1, sa)

	fmt.Printf("x1={%.4f}\ny1={%.4f}\n", x1, y1)

	x1r := x1 / (x1 + y1/p)
	y1r := y1 / (y1 + x1*p)
	fmt.Println(x1r, y1r)

	xr := x / (x + y/p)
	yr := y / (y + x*p)
	fmt.Println(xr, yr)
	// 20/2000, 0.9908
	// 20 * 0.9908
	if x*p > y { // trade x for y
		zeroForOne = true

		r := xr - x1r

		swapAmount = math.Abs(x * r)

		amount0 = x - swapAmount

		amount1 = y + swapAmount*p

		fmt.Println("newX=", amount0)
		fmt.Println("newY=", amount1)

	} else if x*p < y { // trade y for x
		zeroForOne = false

		r := xr - x1r

		swapAmount = math.Abs(y * r)

		amount0 = x + swapAmount/p

		amount1 = y - swapAmount

		fmt.Println("newX=", amount0)
		fmt.Println("newY=", amount1)
	}

	fmt.Printf("newX={%.18f}, newY={%.6f},swapamount={%.18f},zeroForOne={%t}\n", amount0, amount1, swapAmount, zeroForOne)

	return amount0, amount1, swapAmount, zeroForOne
}

func get_liquidity_0(x float64, sa float64, sb float64) float64 {
	return x * sa * sb / (sb - sa)
}

func get_liquidity_1(y float64, sa float64, sb float64) float64 {
	return y / (sb - sa)
}

func get_liquidity(x float64, y float64, sp float64, sa float64, sb float64) float64 {
	var liquidity, liquidity0, liquidity1 float64
	if sp <= sa {
		liquidity = get_liquidity_0(x, sa, sb)
	} else if sp < sb {
		liquidity0 = get_liquidity_0(x, sp, sb)
		liquidity1 = get_liquidity_1(y, sa, sp)
		liquidity = math.Min(liquidity0, liquidity1)
	} else {
		liquidity = get_liquidity_1(y, sa, sb)
	}
	return liquidity
}

func calculate_x(L float64, sp float64, sb float64) float64 {
	return L * (sb - sp) / (sp * sb)
}
func calculate_y(L float64, sp float64, sa float64) float64 {
	return L * (sp - sa)
}

//
// Two different ways how to calculate p_a. calculate_a1() uses liquidity as an input, calculate_a2() does not.
//
func calculate_a1(L float64, sp float64, sb float64, x float64, y float64) float64 {
	// https://www.wolframalpha.com/input/?i=solve+L+%3D+y+%2F+%28sqrt%28P%29+-+a%29+for+a
	// sqrt(a) = sqrt(P) - y / L
	return (sp - y/L) * (sp - y/L)
}
func calculate_a2(sp float64, sb float64, x float64, y float64) float64 {
	// https://www.wolframalpha.com/input/?i=solve+++x+sqrt%28P%29+sqrt%28b%29+%2F+%28sqrt%28b%29++-+sqrt%28P%29%29+%3D+y+%2F+%28sqrt%28P%29+-+a%29%2C+for+a
	// sqrt(a) = (y/sqrt(b) + sqrt(P) x - y/sqrt(P))/x
	//    simplify:
	// sqrt(a) = y/(sqrt(b) x) + sqrt(P) - y/(sqrt(P) x)
	sa := y/(sb*x) + sp - y/(sp*x)
	return sa * sa
}

//
// Two different ways how to calculate p_b. calculate_b1() uses liquidity as an input, calculate_b2() does not.
//
func calculate_b1(L float64, sp float64, sa float64, x float64, y float64) float64 {
	// https://www.wolframalpha.com/input/?i=solve+L+%3D+x+sqrt%28P%29+sqrt%28b%29+%2F+%28sqrt%28b%29+-+sqrt%28P%29%29+for+b
	// sqrt(b) = (L sqrt(P)) / (L - sqrt(P) x)
	return math.Pow((L*sp)/(L-sp*x), 2)
}

func calculate_b2(sp float64, sa float64, x float64, y float64) float64 {
	// find the square root of b:
	// https://www.wolframalpha.com/input/?i=solve+++x+sqrt%28P%29+b+%2F+%28b++-+sqrt%28P%29%29+%3D+y+%2F+%28sqrt%28P%29+-+sqrt%28a%29%29%2C+for+b
	// sqrt(b) = (sqrt(P) y)/(sqrt(a) sqrt(P) x - P x + y)
	P := sp * sp
	return math.Pow((sp * y / ((sa*sp-P)*x + y)), 2)
}

//
// Calculating c and d
//
func calculate_c(p float64, d float64, x float64, y float64) float64 {
	return y / ((d-1)*p*x + y)

}

//
// Calculate x and y given liquidity and price range
//
func calculate_d(p float64, c float64, x float64, y float64) float64 {
	return 1 + y*(1-c)/(c*p*x)
}
