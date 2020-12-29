package main

import (
	"BitcoinConnect/utils"
	"fmt"
)

func main() {
	//1、准备一个json数据
	//hash := utils.GetBestBlockHash()
	//fmt.Println(hash)
	//blockhash,err := utils.GetBlockHash(0)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(blockhash)
	//balance,err := utils.GetBalances()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(balance)
	//balance,err:=utils.GetBalances()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(balance)
	//difficult:=utils.GetDifficult()
	//fmt.Println(difficult)
	//num:=[]interface{}{1,2,3,4,5,6,7,8,9}
	//fmt.Println(num[0])
	chaintips,err:=utils.GetChaintips()
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(chaintips[0].Hash)
}
