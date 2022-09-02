/*

   @author #Kkk
   @File model
   @Description: 微信支付
   @version 0.1
   @date 2022/8/5 10:29

*/

package model

import "fmt"

type WxPayment struct {
	PaymentAges               // 匿名
	PaymentOther PaymentOther // 命名
	WxOpenId     string
}

func (w *WxPayment) SayHello() {
	fmt.Println("Hi, I am a " + w.WxOpenId)
}
