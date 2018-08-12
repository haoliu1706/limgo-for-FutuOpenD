package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/Qot_GetSubInfo"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(3003), "QotGetSubInfo")

	var err error
	err = limgo.On("send.QotGetSubInfo", QotGetSubInfoSend)
	err = limgo.On("recv.QotGetSubInfo", QotGetSubInfoRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// QotGetSubInfoSend handler
func QotGetSubInfoSend(conn *limgo.Request) error {
	ftpack := &limgo.FutuPack{}
	ftpack.SetProtoID(uint32(3003))

	isReqAllConn := true
	reqData := &Qot_GetSubInfo.Request{
		C2S: &Qot_GetSubInfo.C2S{
			IsReqAllConn: &isReqAllConn,
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

// QotGetSubInfoRecv handler
func QotGetSubInfoRecv(data []byte) error {
	fut := &Qot_GetSubInfo.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(fut)

	return nil
}
