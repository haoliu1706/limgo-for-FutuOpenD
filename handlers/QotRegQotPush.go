package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/Qot_Common"
	"limgo/futupb/Qot_RegQotPush"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(3002), "QotRegQotPush")

	var err error
	err = limgo.On("send.QotRegQotPush", QotRegQotPushSend)
	err = limgo.On("recv.QotRegQotPush", QotRegQotPushRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// QotRegQotPushSend handler
func QotRegQotPushSend(conn *limgo.Request, stockCode string, subType string) error {
	ftpack := &limgo.FutuPack{}
	ftpack.SetProtoID(uint32(3002))

	var securityList []*Qot_Common.Security
	security := transStockCode(stockCode)
	securityList = append(securityList, security)

	var subTypeList []int32
	subTypeNum := transSubType(subType)
	subTypeList = append(subTypeList, subTypeNum)

	var regPushRehabTypeList []int32
	regPushRehabType := int32(1)
	regPushRehabTypeList = append(regPushRehabTypeList, regPushRehabType)

	isRegOrUnReg := true
	reqData := &Qot_RegQotPush.Request{
		C2S: &Qot_RegQotPush.C2S{
			SecurityList:  securityList,
			SubTypeList:   subTypeList,
			RehabTypeList: regPushRehabTypeList,
			IsRegOrUnReg:  &isRegOrUnReg,
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

// QotRegQotPushRecv handler
func QotRegQotPushRecv(data []byte) error {
	fut := &Qot_RegQotPush.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(fut)

	return nil
}
