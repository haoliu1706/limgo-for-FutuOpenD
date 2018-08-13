package main

import (
	"fmt"
	"limgo"
	_ "limgo/handlers"
)

func main() {

	block := make(chan bool)

	lim := limgo.New(limgo.Config{Host: "safe.lim.app", Port: "11111"})

	// 1002 GetGlobalState
	limgo.Do("send.GetGlobalState", lim)

	// keepalive
	lim.KeepAlive(false)

	// recv
	go func() {
		fmt.Println("start recv data")
		lim.Recv()
	}()

	// 3010 QotGetTicker
	limgo.Do("send.QotSub", lim, "SZ.300104", "Ticker", true)  // 3001 QotSub
	limgo.Do("send.QotRegQotPush", lim, "SZ.300104", "Ticker") // 3002 QotRegQotPush
	limgo.Do("send.QotGetTicker", lim)                         // get

	// // 3004 QotGetBasicQot
	limgo.Do("send.QotSub", lim, "SZ.300104", "Basic", true)  // 3001 QotSub
	limgo.Do("send.QotRegQotPush", lim, "SZ.300104", "Basic") // 3002 QotRegQotPush
	limgo.Do("send.QotGetBasicQot", lim)                      // get

	// // 3008 QotGetRT
	limgo.Do("send.QotSub", lim, "SZ.300104", "RT", true)  // 3001 QotSub
	limgo.Do("send.QotRegQotPush", lim, "SZ.300104", "RT") // 3002 QotRegQotPush
	limgo.Do("send.QotGetRT", lim, "US.AAPL")              // get

	// // 3002 QotGetOrderBook
	limgo.Do("send.QotSub", lim, "SZ.300104", "OrderBook", true)    // 3001 QotSub
	limgo.Do("send.QotRegQotPush", lim, "SZ.300104", "OrderBook")   // 3002 QotRegQotPush
	limgo.Do("send.QotGetOrderBook", lim, "SZ.300104", "OrderBook") // get

	// 3003 QotGetSubInfo
	limgo.Do("send.QotGetSubInfo", lim)

	<-block
}
