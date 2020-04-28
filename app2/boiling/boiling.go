package main

import (
	"fmt"
	"image/gif"
	"math/rand"
)

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
}

func constDeclarations() {
	const freezingF, boilingF = 32.0, 212.0
}

func variablesDeclarations() {
	var s string
	var i, j, k int
	var b, f, m = true, 2.3, "four"

	anim := gif.GIF{LoopCount: 1}
	freq := rand.Float64() * 3.0
	t := 0.0

	x := 100
	var names []string
	var err error
}

func point() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p =2
	fmt.Println(x)
}

func incr(p *int) int {
	*p++
	return *p
}

v := 1
incr(&v)
incr(&v)


