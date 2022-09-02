package _interface

import "fmt"

/**
 *
 * @date 2022/7/19 14:58
 * @version 0.1
 * @author KongLingJie
 *
 */

type Transporter interface {
	move(string, string) (int error)
	say(int)
}

type Humna interface {
	think(a int) int
}

type Car struct {
}

func (Car) move(string, string) (int, error) {
	return 1, nil
}
func (Car) say(a int) {

}

func (*Car) think(a int) (b int) {
	b = a * 2
	return
}

func InterfaceFunc() {
	var t Transporter
	car := Car{}
	car.say(1)
	t.say(1)

	var h Humna
	h = &Car{}
	x := h.think(5)
	fmt.Println(x)

	s := car.think(2)
	fmt.Println(s)

}
