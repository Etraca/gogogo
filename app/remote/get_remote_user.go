package remote

import (
	"gogogo/app/model"
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"bytes"
)

func GetUser() (*model.EtracaUsers) {
	url := "http://localhost:8080/test/getUser"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	return _conver(string(body))
}

func _conver(string string) (*model.EtracaUsers) {
	user := &model.EtracaUsers{}
	json.Unmarshal([]byte(string), &user)
	return user
}

func PostUser(p string) (*model.EtracaUsers) {
	param := bytes.NewBuffer([]byte(p))
	resp, err := http.Post("http://localhost:8080/test/postUser", "application/json;charset=utf-8", param)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	return _conver(string(body))
}
