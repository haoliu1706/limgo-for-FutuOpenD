package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/Qot_UpdateTicker"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(3011), "QotUpdateTicker")

	var err error
	err = limgo.On("recv.QotUpdateTicker", QotUpdateTickerRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// QotUpdateTickerRecv handler
func QotUpdateTickerRecv(data []byte) error {
	fut := &Qot_UpdateTicker.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(fut)

	return nil
}
