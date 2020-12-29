package main

import (
	"BitcoinConnect/entity"
	"BitcoinConnect/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)
const RPCUSER = "qhw"
const RPCPASSWORD = "woaizcy.."
const RPCURL = "http://127.0.0.1:8332"
const RPCUCERSION = "2.0"

func main() {
	//TimeUnix := time.Now().Unix()
	//TimeUnixStr := strconv.FormatInt(TimeUnix, 10)
	//json_rpc := make(map[string]string)
	//json_rpc["id"] = TimeUnixStr
	//json_rpc["jsonrpc"] = "2.0"
	//json_rpc["method"] = "getbestblockhash"
	//client := http.Client{}
	//request, err := http.NewRequest("POST", RPCURL, nil)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//requesthead := make(map[string]string)
	//requesthead["Authorization"] = "Basic " + RPCUSER + RPCPASSWORD
	//requesthead["Content-Type"] = "application/json"
	//requesthead["Encoding"] = "UTF-8"
	//for key, value := range requesthead {
	//	request.Header.Add(key, value)
	//}
	//response, err := client.Do(request)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//defer response.Body.Close()
	//infomation, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(infomation)
	rpcReq := entity.RPCRequest{
		Id:      time.Now().Unix(),
		Method:  "getbestblockhash",
		Jsonrpc: RPCUCERSION,
	}
	reqBytes, err := json.Marshal(&rpcReq)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(reqBytes))
	body := strings.NewReader(string(reqBytes))
	//2、发送一个请求
	client := &http.Client{}
	//请求头设置
	request, err := http.NewRequest("POST", RPCURL, body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	authStr := RPCUSER + ":" + RPCPASSWORD
	authBase64 := utils.Base64ToStr(authStr)
	//fmt.Println(authBase64)
	request.Header.Add("Encoding", "UTF-8")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Basic "+authBase64)
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	code := resp.StatusCode
	fmt.Println(code)
	if code == 200 {
		fmt.Println("请求成功")
	} else {
		fmt.Println("请求失败")
	}
	respbyte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(respbyte))
}
