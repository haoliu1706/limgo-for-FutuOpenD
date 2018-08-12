package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/GetGlobalState"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(1002), "GetGlobalState")

	var err error
	err = limgo.On("send.GetGlobalState", GetGlobalStateSend)
	err = limgo.On("recv.GetGlobalState", GetGlobalStateRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// GetGlobalStateSend handler
func GetGlobalStateSend(conn *limgo.Request) error {

	ftpack := &limgo.FutuPack{}

	ftpack.SetProtoID(uint32(1002))
	userID := conn.InitData.LoginUserID

	reqData := &GetGlobalState.Request{
		C2S: &GetGlobalState.C2S{
			UserID: &userID,
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

// GetGlobalStateRecv handler
func GetGlobalStateRecv(data []byte) error {
	fut := &GetGlobalState.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(fut)

	return nil
}
