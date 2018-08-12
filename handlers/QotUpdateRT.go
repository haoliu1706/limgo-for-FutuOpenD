package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/Qot_UpdateRT"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(3009), "QotUpdateRT")

	var err error
	err = limgo.On("recv.QotUpdateRT", QotUpdateRTRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// QotUpdateRTRecv handler
func QotUpdateRTRecv(data []byte) error {
	fut := &Qot_UpdateRT.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(fut)

	return nil
}
