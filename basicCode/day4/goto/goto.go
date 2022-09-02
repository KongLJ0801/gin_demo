package _goto

import "fmt"

/**
 *
 * @date 2022/7/18 15:41
 * @version 0.1
 * @author KongLingJie
 *
 */

func ifLabel() {
	i := 4
L1:
	i += 3
L2:
	i += 2

	if i%2 == 0 {
		fmt.Printf("AA = %d , ", i)
		if i > 50 {
			return
		}
		goto L1
	} else {
		fmt.Printf("BB = %d , ", i)
		if i > 50 {
			return
		}
		goto L2
	}

}

func GotoMain() {
	ifLabel()
}
