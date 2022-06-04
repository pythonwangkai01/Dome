package e

import (
	"fmt"
)

func HandlerError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
		panic(err)
	}
}
