package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/Qot_GetRT"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(3008), "QotGetRT")

	var err error
	err = limgo.On("send.QotGetRT", QotGetRTSend)
	err = limgo.On("recv.QotGetRT", QotGetRTRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// QotGetRTSend handler
func QotGetRTSend(conn *limgo.Request, stockCode string) error {
	ftpack := &limgo.FutuPack{}
	ftpack.SetProtoID(uint32(3008))

	security := transStockCode(stockCode)

	reqData := &Qot_GetRT.Request{
		C2S: &Qot_GetRT.C2S{
			Security: security,
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

// QotGetRTRecv handler
func QotGetRTRecv(data []byte) error {
	fut := &Qot_GetRT.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(fut)

	return nil
}
