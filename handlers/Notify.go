package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/Notify"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(1003), "Notify")

	var err error
	err = limgo.On("recv.Notify", NotifyRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// NotifyRecv handler
func NotifyRecv(data []byte) error {
	fut := &Notify.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(fut)

	return nil
}
