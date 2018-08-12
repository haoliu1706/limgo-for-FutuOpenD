package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/Qot_Common"
	"limgo/futupb/Qot_GetBasicQot"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(3004), "QotGetBasicQot")

	var err error
	err = limgo.On("send.QotGetBasicQot", QotGetBasicQotSend)
	err = limgo.On("recv.QotGetBasicQot", QotGetBasicQotRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// QotGetBasicQotSend handler
func QotGetBasicQotSend(conn *limgo.Request, stockCode string) error {
	ftpack := &limgo.FutuPack{}
	ftpack.SetProtoID(uint32(3004))

	var securityList []*Qot_Common.Security
	security := transStockCode(stockCode)
	securityList = append(securityList, security)

	reqData := &Qot_GetBasicQot.Request{
		C2S: &Qot_GetBasicQot.C2S{
			SecurityList: securityList,
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

// QotGetBasicQotRecv handler
func QotGetBasicQotRecv(data []byte) error {
	fut := &Qot_GetBasicQot.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(fut)

	return nil
}
