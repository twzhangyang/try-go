package main

func app() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}

func newInt() *int {
	return new(int)
}

func go() {
	p := new(int)
	q := new (int)
	fmt.Println(p == q)
}