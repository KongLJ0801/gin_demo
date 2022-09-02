package anonymous

import "fmt"

/**
 *
 * @date 2022/7/19 10:02
 * @version 0.1
 * @author KongLingJie
 *
 */

type foo func(a, b int) int

type User struct {
	Name  string
	bay   foo
	hello func(name string) string
}

func fun1(f foo, b, c int) int {
	a := b + c
	return f(a, c)
}
func sum(a, b int) int {
	return a - b
}
func add(a, b int) int {
	return a + b
}

func AnonymousFunc() {
	ch := make(chan func(string) string, 10)
	ch <- func(string2 string) string {
		return "hello " + string2
	}

	v := <-ch
	x := v("jack")
	fmt.Println(x)

	u := User{}
	u.bay = sum

	c := u.bay(1, 2)
	fmt.Println(c)

	var s foo = sum
	o := fun1(s, 1, 2)
	fmt.Println(o)

	var f foo = add
	g := fun1(f, 1, 2)
	fmt.Println(g)
}
