package main

func pointer() {
	x := 1
	p := &x
	fmt.Println(*p)

	*p = 2
	fmt.Println(x)
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}

func app() {
	var p = f()
	v := 1
	incr(&v)

	fmt.Println(incr(&v))
}

