package callback

import (
	"gogogo/app/model"
	"log"
)

func Count() (res interface{}) {
	count, _ := new(model.EtracaLogonCount).Get()
	count.Count++
	count.Update()
	log.Println(count)
	return count
}
