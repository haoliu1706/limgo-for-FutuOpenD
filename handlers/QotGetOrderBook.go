package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/Qot_GetOrderBook"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(3012), "QotGetOrderBook")

	var err error
	err = limgo.On("send.QotGetOrderBook", QotGetOrderBookSend)
	err = limgo.On("recv.QotGetOrderBook", QotGetOrderBookRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// QotGetOrderBookSend handler
func QotGetOrderBookSend(conn *limgo.Request, stockCode string) error {
	ftpack := &limgo.FutuPack{}
	ftpack.SetProtoID(3012)

	security := transStockCode(stockCode)
	num := int32(5)

	reqData := &Qot_GetOrderBook.Request{
		C2S: &Qot_GetOrderBook.C2S{
			Security: security,
			Num:      &num,
		},
	}
	pbData, err := proto.Marshal(reqData)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	ftpack.SetBody(pbData)
	err = conn.Send(ftpack)

	return err
}

// QotGetOrderBookRecv handler
func QotGetOrderBookRecv(data []byte) error {
	fut := &Qot_GetOrderBook.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(fut)

	return nil
}
