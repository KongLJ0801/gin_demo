package herit

import "fmt"

/**
 *
 * @date 2022/7/20 9:28
 * @version 0.1
 * @author KongLingJie
 *
 */

type Plane struct {
	color string
}

func (plane Plane) Fly() {
	fmt.Println("hi tom")
}

type Humn struct {
}

func (Humn) say() {

}

type Bird struct {
	// 多个是组合
	Plane // 匿名继承
	Humn
}

func (bird Bird) Fly() { // 重写父类方法
	bird.Plane.Fly()
	fmt.Println("hi jack")
}

func HeritFunc() {
	p := Plane{}
	p.Fly()

	b := Bird{}
	fmt.Println(b.color)
	b.Fly()

	b.say()
}
