package common

type Result struct {
	ErrorCode int      `json:"errorCode"`
	ErrorMsg  string      `json:"errorMSg"`
	Data      interface{} `json:"data"`
}
