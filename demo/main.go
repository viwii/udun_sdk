package main

import (
	"fmt"

	sdk "github.com/viwii/udun_sdk"
)

func main() {

	udunClient := sdk.NewUdunClient("https://hk01-node.uduncloud.com",
		"xxxxxx",
		"xxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"http://192.168.2.223:8081/udun/notify")

	coinList := udunClient.ListSupportCoin(false)
	fmt.Println(coinList)
	fmt.Println(udunClient.CheckAddress("60", "xxxxxxxxxxxxxxxxxxxx"))
	fmt.Println(udunClient.Withdraw("xxxxxxxxxxxxxxxx", 1.0, "60", "60", "1", ""))
}
