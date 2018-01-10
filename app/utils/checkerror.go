package utils

import (
	"fmt"
)

func CkeckError(err error,msg string)  {
	if err !=nil {
		panic(fmt.Sprintf(msg, err))
	}
}
