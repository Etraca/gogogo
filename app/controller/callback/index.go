package callback

import (
	"log"
)

func Index() (resp interface{}) {
	return "hello world!"
}

func SayHello(hello string) (response interface{}) {
	log.Println(hello)
	return hello
}
