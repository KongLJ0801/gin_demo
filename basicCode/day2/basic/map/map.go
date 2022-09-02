package _map

import "fmt"

/**
 *
 * @date 2022/7/14 17:26
 * @version 0.1
 * @author KongLingJie
 *
 */

func editMap() {
	obj := make(map[string]int, 5)
	//obj := make(map[string]int, 5){'name':1}

	obj["name"] = 1
	obj["age"] = 18

	fmt.Println(obj, len(obj))

	for key, value := range obj {
		value += 2
		fmt.Printf("%s %d %d\n", key, value, obj[key])
	}

	if _, err := obj["sex"]; err {
		delete(obj, "sex")
	} else {
		fmt.Println("map里不存在[sex]这个key")
	}

	fmt.Println(obj, len(obj))
}

func MapMain() {
	editMap()
}
