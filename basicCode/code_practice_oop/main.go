/*

   @author #Kkk
   @File code_practice_oop
   @Description: NEW 一个对象的好处  (低耦合,高内聚)
   @version 0.1
   @date 2022/8/5 10:05

*/

package main

import "gin_demo/basicCode/code_practice_oop/model"

func main() {
	// 1. 结构体 实体封装;
	//user := model.UserInfo{
	//	Name:      "#Kkk",
	//	Age:       18,
	//	Height:    180.5,
	//	Schooled: "深圳",
	//	Hobby:     []string{"烧烤", "啤酒"},
	//	MoreInfo:  nil,
	//}
	//
	//println(user.Name, user.Height)

	user := model.NewModelUserInfo("#Kkk", 18, 180.5, "深圳", []string{"卷", "不卷"}, nil)

	println(user.Name, user.Height)

	product := &model.Product{}

	product.SetProductName("#Kkk")
	product.SetProductPrice(18.5)

	println(product.GetProductName(), product.GetProductPrice())

	//---------------------------------------------------------

	// 2.继承   复用,扩展,维护性; 访问流程   属性名->子结构体->夫结构体 (向上查找)

	super := &model.PaymentAges{
		AppId:       "super-AppId",
		MchId:       "super-MchId",
		Key:         "super-Key",
		CallbackUrl: "https://super-CallbackUrl",
	}
	super.Info("super")

	zfbPay := &model.ZfbPayment{
		PaymentAges: model.PaymentAges{
			AppId:       "zfbPay-AppId",
			MchId:       "zfbPay-MchId",
			Key:         "zfbPay-Key",
			CallbackUrl: "https://zfbPay-CallbackUrl",
		},
		ZfbOpenId: "zfbPay-ZfbOpenId",
	}
	// 方法重载
	zfbPay.Info("zfbPay")

	wxPay := &model.WxPayment{
		PaymentAges: model.PaymentAges{
			AppId:       "wxPay-AppId",
			MchId:       "wxPay-MchId",
			Key:         "wxPay-Key",
			CallbackUrl: "https://wxpay-CallbackUrl",
		},
		PaymentOther: model.PaymentOther{
			AppId:       "PaymentOther-AppId",
			MchId:       "PaymentOther-MchId",
			Key:         "PaymentOther-Key",
			CallbackUrl: "PaymentOther-CallbackUrl",
		},
		WxOpenId: "wxPay-WxOpenId",
	}
	println(zfbPay.AppId)
	println(wxPay.AppId)
	println(wxPay.PaymentOther.AppId)

	// 3. 多态 通过接口来实现 (继承,重写,父类引用指向子类对象)

	/**
	定义接口的所有方法的任何类型都表示隐式实现该接口。
	类型接口的变量可以保存实现该接口的任何值。
	接口的这个属性用于实现GO的多态性。
	*/

	wxPay.SayHello()

	zfbPay.SayHello()
}
