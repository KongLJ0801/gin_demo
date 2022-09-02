/*

   @author #Kkk
   @File jsonCode
   @Description:
   @version 0.1
   @date 2022/8/11 16:26

*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

type User struct {
}

func main() {
	//reflectJson()

	//structJson()

	myJson()

}

// struct 解析json code2,code1
func structJson() {
	msg := "{\"call_back_key\":\"card:power:route:50fa0cf77254497a9ade866399d12899\",\"data\":[{\"destination\":\"192.168.101.128\",\"gateway\":\"192.168.101.5\",\"netmask\":\"255.255.255.255\"},{\"destination\":\"192.168.101.128\",\"gateway\":\"192.168.101.5\",\"netmask\":\"255.255.255.255\"}],\"time_out\":\"60000 201\"}"
	//ResultMsg := &ResultData{}
	//err := json.Unmarshal([]byte(msg), ResultMsg)
	//if err != nil {
	//	fmt.Println("jsonErr 解析错误!")
	//	return
	//}
	//for i, ele := range ResultMsg.Data {
	//	fmt.Println(ele.Netmask, i)
	//	fmt.Println(ele.Gateway, i)
	//	fmt.Println(ele.Destination, i)
	//}
	//fmt.Println(ResultMsg)

	//var result ResultData
	//jsonErr := json.Unmarshal([]byte(msg), &result)
	//if jsonErr != nil {
	//	fmt.Println("jsonErr 解析错误!")
	//	return
	//}
	//fmt.Println(result)

	var MSGJson interface{}
	jsonErr := json.Unmarshal([]byte(msg), &MSGJson)
	if jsonErr != nil {
		fmt.Println("jsonErr 解析错误!")
		return
	}
	var result Result
	structErr := mapstructure.Decode(MSGJson, &result)
	if structErr != nil {
		fmt.Println("structErr 解析错误!")
		return
	}
	fmt.Println(result.Timeout)
	fmt.Println(result.CallBackKey)
	fmt.Println(result.Data)

}

// code2 (反射)
func reflectJson() {
	msg := "{\"call_back_key\":\"card:power:route:50fa0cf77254497a9ade866399d12899\",\"data\":[{\"destination\":\"192.168.101.128\",\"gateway\":\"192.168.101.5\",\"netmask\":\"255.255.255.255\"},{\"destination\":\"192.168.101.128\",\"gateway\":\"192.168.101.5\",\"netmask\":\"255.255.255.255\"}],\"time_out\":\"60000 201\"}"
	//msg := "{\"destination\":\"192.168.101.128\",\"gateway\":\"192.168.101.5\",\"netmask\":\"255.255.255.255\"}"
	result := &ResultData{}
	result, err := typeFunc[ResultData]([]byte(msg), result)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	for i, ele := range result.Data {
		val := ele.Gateway
		if val != "" {
			fmt.Println(i, ele, ele.Gateway)
		}
	}
	fmt.Printf("%+v", result)
}

//根据传入的类型通过反射获取到类型以及值
func typeFunc[E any](v any, result *E) (*E, error) {
	valueOf := reflect.ValueOf(v)
	typeOf := reflect.TypeOf(v)
	if k := typeOf.Kind(); k == reflect.Slice {
		err := json.Unmarshal(valueOf.Bytes(), result)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}

// 泛型 解析json code3
func myJson() {
	//msg := "{\"call_back_key\":\"card:power:route:50fa0cf77254497a9ade866399d12899\",\"data\":{\"destination\":\"192.168.101.128\",\"gateway\":\"192.168.101.5\",\"netmask\":\"255.255.255.255\"},\"time_out\":\"60000 201\"}"
	//var t ListResultData
	//x := structJson(t)
	//err := json.Unmarshal([]byte(msg), x)
	//if err != nil {
	//	fmt.Println("jsonErr 解析错误!")
	//	return
	//}
	//fmt.Println(x.Data.Netmask)

	//msg := "{\"call_back_key\":\"card:power:route:50fa0cf77254497a9ade866399d12899\",\"data\":[{\"destination\":\"192.168.101.128\",\"gateway\":\"192.168.101.5\",\"netmask\":\"255.255.255.255\"},{\"destination\":\"192.168.101.128\",\"gateway\":\"192.168.101.5\",\"netmask\":\"255.255.255.255\"}],\"time_out\":\"60000 201\"}"
	//var t []ListResultDataCode
	//x := structFunc(t)
	//err := json.Unmarshal([]byte(msg), x)
	//if err != nil {
	//	fmt.Println("jsonErr 解析错误!")
	//	return
	//}
	//for i, ele := range x.Data {
	//	fmt.Println(ele.Netmask, i)
	//	fmt.Println(ele.Gateway, i)
	//	fmt.Println(ele.Destination, i)
	//}
	//fmt.Printf("%+v", x)

	//msg := "{\"call_back_key\":\"card:power:route:50fa0cf77254497a9ade866399d12899\",\"data\":[{\"destination\":\"192.168.101.128\",\"gateway\":\"192.168.101.5\",\"netmask\":\"255.255.255.255\"},{\"destination\":\"192.168.101.128\",\"gateway\":\"192.168.101.5\",\"netmask\":\"255.255.255.255\"}],\"time_out\":\"60000 201\"}"
	//type anyType = []ListResultDataCode
	//anyJson := &ResultDataAny[anyType]{}
	//anyJson.StructFuncAny(msg)
	//fmt.Println(anyJson)
	//
	//msg := "{\"call_back_key\":\"card:power:route:50fa0cf77254497a9ade866399d12899\",\"data\":[{\"destination\":\"192.168.101.128\",\"gateway\":\"192.168.101.5\",\"netmask\":\"255.255.255.255\"},{\"destination\":\"192.168.101.128\",\"gateway\":\"192.168.101.5\",\"netmask\":\"255.255.255.255\"}],\"time_out\":\"60000 201\"}"
	//type anyType []ListResultDataCode
	//anyJson := &ResultDataAny5[anyType]{}
	//anyJson.StructFuncAny5(msg)
	//fmt.Println(anyJson)
	//for i, ele := range anyJson.Data {
	//	//fmt.Println(reflect.TypeOf(ele).Elem(), ele, i)
	//	fmt.Println(ele, i)
	//}

	msg := "{\"call_back_key\":\"card:power:card_execute_script:79bc1b70fdd54f11a7a61a4877806b1f\",\"data\":{\"command\":\"/data/local/tmp/getlog.sh 10\",\"script_md5\":\"141973f429aa6bd9c6b904101e43201d\",\"script_url\":\"http://192.168.13.252:18080/download/getlog.sh\"},\"time_out\":30000}"
	anyJson := &ResultDataAny5[ScriptExecutionJson]{}
	anyJson.StructFuncAny5(msg)
	fmt.Println(anyJson)
}

// code6
type ScriptExecutionJson struct {
	Command   string `json:"command"`
	ScriptMd5 string `json:"script_md5"`
	ScriptUrl string `json:"script_url"`
}

// code5

func (r *ResultDataAny5[T]) StructFuncAny5(dataJson string) error {
	err := json.Unmarshal([]byte(dataJson), r)
	if err != nil {
		fmt.Println("jsonErr 解析错误!", err)
		return err
	}
	return nil
}

type ResultDataAny5[T any] struct {
	// 消息标识
	CallBackKey string `json:"call_back_key"`
	// 消息体
	Data T `json:"data"`
	//超时时间
	Timeout int64 `json:"time_out"`
}

// code4

func (r *ResultDataAny[T]) StructFuncAny(dataJson string) error {
	err := json.Unmarshal([]byte(dataJson), r)
	if err != nil {
		fmt.Println("jsonErr 解析错误!")
		return err
	}
	return nil
}

type ResultDataAny[T MyListResultData] struct {
	// 消息标识
	CallBackKey string `json:"call_back_key"`
	//超时时间
	Timeout string `json:"time_out"`
	// 消息体
	Data T `json:"data"`
}

//code3

type MyListResultData interface {
	[]ListResultDataCode | ListResultDataCode
}

func structFunc[T MyListResultData](Data T) *ResultDataCode[T] {
	return &ResultDataCode[T]{
		CallBackKey: "",
		Timeout:     "",
		Data:        Data,
	}
}

type ResultDataCode[T MyListResultData] struct {
	// 消息标识
	CallBackKey string `json:"call_back_key"`
	//超时时间
	Timeout string `json:"time_out"`
	// 消息体
	//Data []map[string]interface{} `json:"data"`
	Data T `json:"data"`
}

type ListResultDataCode struct {
	Destination string `json:"destination"`
	Netmask     string `json:"netmask"`
	Gateway     string `json:"gateway"`
}

// code 2
type ResultData struct {
	// 消息标识
	CallBackKey string `json:"call_back_key"`
	//超时时间
	Timeout string `json:"time_out"`
	// 消息体
	//Data []map[string]interface{} `json:"data"`
	Data []ListResultData `json:"data"`
}
type ListResultData struct {
	Destination string `json:"destination"`
	Netmask     string `json:"netmask"`
	Gateway     string `json:"gateway"`
}

// code 1
type ListResult struct {
	Destination string `mapstructure:"destination"`
	Netmask     string `mapstructure:"netmask"`
	Gateway     string `mapstructure:"gateway"`
}

type Result struct {
	// 消息标识
	CallBackKey string `mapstructure:"call_back_key"`
	//超时时间
	Timeout string `mapstructure:"time_out"`
	// 消息体
	Data AutoGenerated `mapstructure:"data"`
	//Data []ListResult `mapstructure:"data"`
}

type AutoGenerated []struct {
	Destination string `json:"destination"`
	Gateway     string `json:"gateway"`
	Netmask     string `json:"netmask"`
}
