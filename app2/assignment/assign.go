package main

func app1() {
	x = 1 
	*p = true
	person.name = "bob"
	count[x] = count[x] * scale
}

func tupleAsign() {
	x, y = y, x
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func fib(n int) int {
	x, y  := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	return x
}

func returnResult() {
	v, ok = m[key]
	v, ok = x.(T)
	v, ok = <-ch

	_, err = io.Copy(dst, src)
	_, ok = x.(T)
}

func assignability () {
	medals := []string{"gold", "silver", "bronze"}
	medals[0] = "gold"
}