package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/Qot_UpdateBasicQot"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(3005), "QotUpdateBasicQot")

	var err error
	err = limgo.On("recv.QotUpdateBasicQot", QotUpdateBasicQotRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// QotUpdateBasicQotRecv handler
func QotUpdateBasicQotRecv(data []byte) error {
	fut := &Qot_UpdateBasicQot.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(fut)

	return nil
}
