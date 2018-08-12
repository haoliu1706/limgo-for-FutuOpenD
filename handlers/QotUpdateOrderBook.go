package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/Qot_UpdateOrderBook"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(3013), "QotUpdateOrderBook")

	var err error
	err = limgo.On("recv.QotUpdateOrderBook", QotUpdateOrderBookRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// QotUpdateOrderBookRecv handler
func QotUpdateOrderBookRecv(data []byte) error {
	fut := &Qot_UpdateOrderBook.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(fut)

	return nil
}
