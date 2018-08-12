package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/Qot_Common"
	"limgo/futupb/Qot_GetTicker"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(3010), "QotGetTicker")

	var err error
	err = limgo.On("send.QotGetTicker", QotGetTickerSend)
	err = limgo.On("recv.QotGetTicker", QotGetTickerRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// QotGetTickerSend handler
func QotGetTickerSend(conn *limgo.Request) error {
	ftpack := &limgo.FutuPack{}
	ftpack.SetProtoID(uint32(3010))

	market := int32(11)
	code := "AAPL"
	security := &Qot_Common.Security{
		Market: &market,
		Code:   &code,
	}
	maxRetNum := int32(10)

	reqData := &Qot_GetTicker.Request{
		C2S: &Qot_GetTicker.C2S{
			Security:  security,
			MaxRetNum: &maxRetNum,
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

// QotGetTickerRecv handler
func QotGetTickerRecv(data []byte) error {
	fut := &Qot_GetTicker.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(fut)

	return nil
}
