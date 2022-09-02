package user

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

/**
 *
 * @date 2022/7/25 10:26
 * @version 0.1
 * @author KongLingJie
 *
 */

type User struct {
	Name string
	Age  int
}

func test_json() {
	u := User{"jerry", 19}

	// json.Marshal json 序列化
	if n, err := json.Marshal(u); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(n))
	}
}

//func test_easy_json() {
//	u := User{"jerry", 19}
//
//	// easyjson.Marshal json 序列化
//
//	if n, err := easyjson.Marshal(u); err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(string(n))
//	}
//}

func test_base64() {
	img, _ := os.Open("./day8/user/123564.png")
	bs := make([]byte, 1<<20)
	n, _ := img.Read(bs)
	str := base64.StdEncoding.EncodeToString(bs[:n])
	fmt.Println(str)

	bbb, _ := base64.StdEncoding.DecodeString(str)
	fileImages, _ := os.OpenFile("./day8/user/999.png", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	fileImages.Write(bbb)
	defer fileImages.Close()

}

func UserMain() {
	//test_json()

	test_base64()
}
