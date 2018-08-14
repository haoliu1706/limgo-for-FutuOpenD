# limgo-for-FutuOpenD
云顿（天津）安防科技有限公司

golang 1.10
https://dl.google.com/go/go1.10.3.windows-amd64.msi
https://dl.google.com/go/go1.10.3.darwin-amd64.pkg

IDE:liteide X34
https://github.com/visualfc/liteide/releases/download/x34.1/liteidex34.1.windows-qt5.9.5.zip
https://github.com/visualfc/liteide/releases/download/x34.1/liteidex34.1.macosx-qt5.9.5.zip



go get github.com/golang/protobuf/protoc-gen-go

工程放入GOPATH

C:\Users\XX\go\src\


测试实列请编译成可执行程序

limgo\example\main.go


接口调用如下

	// 3010 QotGetTicker 实时L2逐笔明细 SZ SH HK US自行区分US.AAPL SH.600123 HK.00700
	limgo.Do("send.QotSub", lim, "SZ.300104", "Ticker", true)  // 3001 QotSub
	limgo.Do("send.QotRegQotPush", lim, "SZ.300104", "Ticker") // 3002 QotRegQotPush
	limgo.Do("send.QotGetTicker", lim)                         // get

	// 3004 QotGetBasicQot 实时L2买卖十档摆盘
	limgo.Do("send.QotSub", lim, "SZ.300104", "Basic", true)  // 3001 QotSub
	limgo.Do("send.QotRegQotPush", lim, "SZ.300104", "Basic") // 3002 QotRegQotPush
	limgo.Do("send.QotGetBasicQot", lim)                      // get

	// 3008 QotGetRT
	limgo.Do("send.QotSub", lim, "SZ.300104", "RT", true)  // 3001 QotSub
	limgo.Do("send.QotRegQotPush", lim, "SZ.300104", "RT") // 3002 QotRegQotPush
	limgo.Do("send.QotGetRT", lim, "US.AAPL")              // get

	// 3002 QotGetOrderBook 实时L2买卖盘
	limgo.Do("send.QotSub", lim, "SZ.300104", "OrderBook", true)    // 3001 QotSub
	limgo.Do("send.QotRegQotPush", lim, "SZ.300104", "OrderBook")   // 3002 QotRegQotPush
	limgo.Do("send.QotGetOrderBook", lim, "SZ.300104", "OrderBook") // get

	// 3003 QotGetSubInfo 查询订阅信息
	limgo.Do("send.QotGetSubInfo", lim)
