package utils

import "fmt"

func Check(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}
