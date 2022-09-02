/*

   @author #Kkk
   @File model
   @Description:支付宝
   @version 0.1
   @date 2022/8/5 10:29

*/

package model

import "fmt"

type ZfbPayment struct {
	PaymentAges
	ZfbOpenId string
}

func (z *ZfbPayment) Info(name string) {
	fmt.Printf("ZfbPayment: %s,value = %p \n", name, z)
}

func (z *ZfbPayment) SayHello() {
	fmt.Println("Hi, I am a " + z.ZfbOpenId)
}
