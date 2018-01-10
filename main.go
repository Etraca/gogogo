package main

import (
	"net/http"
	"gogogo/app/controller/callback"
	"reflect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"log"
	"gogogo/app/common"
	"gogogo/app/model"
	"os"
)

func init() {
	log.Println(os.Args[1])
	common.SetConfig()
	model.SetEngine()
}

func main() {
	prefix := "/mhd/ggg"
	providerServeMux := http.NewServeMux()
	providerServeMux.HandleFunc(prefix+"/test", test)
	providerServeMux.HandleFunc(prefix+"/index", safeHandler(callback.Index))
	providerServeMux.HandleFunc(prefix+"/sayHello", safeHandler(callback.SayHello))
	providerServeMux.HandleFunc(prefix+"/login", safeHandler(callback.Logon))
	providerServeMux.HandleFunc(prefix+"/register", safeHandler(callback.Register))
	providerServeMux.HandleFunc(prefix+"/resetPwd", safeHandler(callback.ResetPwd))
	providerServeMux.HandleFunc(prefix+"/cancel", safeHandler(callback.Cancel))
	providerServeMux.HandleFunc(prefix+"/queryAll", safeHandler(callback.QueryAll))
	providerServeMux.HandleFunc(prefix+"/total", safeHandler(callback.Total))
	providerServeMux.HandleFunc(prefix+"/count", safeHandler(callback.Count))
	providerServeMux.HandleFunc(prefix+"/getRemoteUser", safeHandler(callback.GetRemoteUser))
	providerServeMux.HandleFunc(prefix+"/postRemoteUser", safeHandler(callback.PostRemoteUser))
	srv := &http.Server{
		Addr:         "localhost:9090",
		Handler:      providerServeMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("ListenAndServe error: %v", err))
	}
}

func safeHandler(handleFunc interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		callHandleFunc(handleFunc, w, r)
	}
}

func callHandleFunc(handleFunc interface{}, w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	httpStatus := http.StatusOK
	input := ""
	output := []byte{}
	code := 9999 //返回的错误码
	isJson := false

	defer func() {
		recv := recover()
		if recv != nil {
			log.Println("recv", recv)
			httpStatus = http.StatusInternalServerError
		}
		if isJson {
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
		}
		w.WriteHeader(httpStatus)
		w.Write(output)
		//请求的详细访问日志  todo
		accessLog(r, input, string(output), httpStatus, start, code)
	}()

	//
	handleFuncValue := reflect.ValueOf(handleFunc)
	if handleFuncValue.Kind() != reflect.Func {
		panic(fmt.Sprintf("handle is not func , it's : %v", handleFuncValue.Kind()))
	}

	//以下 拼装 handle 需要的两个参数
	handleFuncInput, input := parseHandleFuncInput(handleFuncValue, r)
	response := handleFuncValue.Call(handleFuncInput)
	isJson, output = parseHandleFuncOutput(response)
}

func parseHandleFuncInput(handleFuncValue reflect.Value, r *http.Request) (HandleFuncInput []reflect.Value, input string) {
	//参数 必须是  一个 struct 或者map  或者...
	handleFuncInputLen := handleFuncValue.Type().NumIn()
	handleFuncInput := make([]reflect.Value, handleFuncInputLen)
	if handleFuncInputLen == 0 {
		return handleFuncInput, input
	}
	if handleFuncInputLen > 2 {
		panic(fmt.Sprintf("handleFuncInput flied num should 1 or 2 "))
	}
	arg := handleFuncValue.Type().In(0)
	if arg.Kind() == reflect.Struct {
		var body []byte
		var err error
		if _, exist := arg.FieldByName("FromForm"); exist {
			//若字段中包括FromForm,从form 中取数据
			err = r.ParseForm()
			if err != nil {
				panic(fmt.Sprintf("r.ParseFormerror: %v", err))
			}
			param := make(map[string]string)
			for k, v := range r.Form {
				param[k] = v[0]
			}
			body, err = json.Marshal(param)
			if err != nil {
				panic(fmt.Sprintf("json.Marshal error: %v", err))
			}
		} else {
			//获取请求的body数据
			body, err = ioutil.ReadAll(r.Body)
			if err != nil {
				panic(fmt.Sprintf("ioutil.ReadAll error: %v", err))
			}
		}
		input = string(body)
		p := reflect.New(arg).Interface()
		err = json.Unmarshal([]byte(input), p)
		if err != nil {
			panic(fmt.Sprintf("json.Unmarshal error: %v, str=%s", err, input))
		}
		handleFuncInput[0] = reflect.ValueOf(p).Elem()
	} else if arg.Kind() == reflect.Map {
		err := r.ParseForm()
		if err != nil {
			panic(fmt.Sprintf("r.ParseFormerror: %v", err))
		}
		param := make(map[string]string)
		for k, v := range r.Form {
			param[k] = v[0]
		}
		handleFuncInput[0] = reflect.ValueOf(param)
	} else if arg.Kind() == reflect.String {
		err := r.ParseForm()
		if err != nil {
			panic(fmt.Sprintf("r.ParseFormerror: %v", err))
		}
		var value string
		for _, v := range r.Form {
			value = v[0]
		}
		handleFuncInput[0] = reflect.ValueOf(value)
	} else {
		//暂时 仅考虑 这种情况~ 以后可以再加
		panic("cannot handle func param ")
	}
	//如果两个参数,则第二个是 Request
	if handleFuncInputLen == 2 {
		handleFuncInput[1] = reflect.ValueOf(r)
	}
	return handleFuncInput, input
}

func parseHandleFuncOutput(response []reflect.Value) (isJson bool, output []byte) {
	//返回字段 必须是 一个 response interface{} 或者没有
	isJson = false
	lenRes := len(response)
	var err error = nil
	if lenRes == 1 {
		if response[0].Kind() == reflect.Invalid {
			return
		}
		kindRes := response[0].Elem().Kind()
		if kindRes == reflect.String {
			output = []byte(response[0].Elem().Interface().(string))
		} else {
			resp := response[0].Interface()
			output, err = json.Marshal(resp)
			isJson = true
			if err != nil {
				panic(fmt.Sprintf("json.Marshal response error: %v", err))
			}
		}

	} else {
		//暂时 仅考虑 只有一个返回值 这种情况~ 以后可以再加
	}
	return
}

func accessLog(r *http.Request, input, output string, httpStatus int, start time.Time, code int) {
	log.Println("_com_request_out", "method", r.Method, "url", r.URL.Path, "httpCode", httpStatus, "input", input, "output", string(output),
		"spent(us)", int64(time.Since(start)/time.Microsecond), "code", code)
}


func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}