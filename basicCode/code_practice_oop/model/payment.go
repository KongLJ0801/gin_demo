/*

   @author #Kkk
   @File model
   @Description:
   @version 0.1
   @date 2022/8/5 10:28

*/

package model

import "fmt"

type Person interface {
	SayHello()
}

type PaymentAges struct {
	AppId       string
	MchId       string
	Key         string
	CallbackUrl string
}

func (p *PaymentAges) Info(name string) {
	fmt.Printf("info:%s  value= %v \n", name, p)
}
