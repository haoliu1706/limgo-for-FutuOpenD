package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/KeepAlive"
	"time"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(1004), "KeepAlive")

	var err error
	err = limgo.On("send.KeepAlive", KeepAliveSend)
	err = limgo.On("recv.KeepAlive", KeepAliveRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// KeepAliveSend handler
func KeepAliveSend(conn *limgo.Request) error {

	ftpack := &limgo.FutuPack{}

	ftpack.SetProtoID(uint32(1004))
	time := time.Now().Unix()
	reqData := &KeepAlive.Request{
		C2S: &KeepAlive.C2S{
			Time: &time,
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

// KeepAliveRecv handler
func KeepAliveRecv(data []byte) error {
	fut := &KeepAlive.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	// fmt.Println(fut)

	return nil
}
