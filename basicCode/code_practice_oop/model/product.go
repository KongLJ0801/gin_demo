/*

   @author #Kkk
   @File model
   @Description:
   @version 0.1
   @date 2022/8/5 10:15

*/

package model

type Product struct {
	productName  string
	productPrice float32
}

func (p *Product) GetProductName() string {
	return p.productName
}

func (p *Product) SetProductName(_productName string) {
	p.productName = _productName
}

func (p *Product) GetProductPrice() float32 {
	return p.productPrice
}

func (p *Product) SetProductPrice(_productPrice float32) {
	p.productPrice = _productPrice
}
