package handlers

import (
	"fmt"
	"limgo"
	"limgo/futupb/Qot_Common"
	"limgo/futupb/Qot_Sub"

	"github.com/golang/protobuf/proto"
)

func init() {
	limgo.SetHandlerID(uint32(3001), "QotSub")

	var err error
	err = limgo.On("send.QotSub", QotSubSend)
	err = limgo.On("recv.QotSub", QotSubRecv)

	if err != nil {
		fmt.Println(err)
	}
}

// QotSubSend handler
func QotSubSend(conn *limgo.Request, stockCode string, subType string, isSubOrUnSub bool) error {
	ftpack := &limgo.FutuPack{}
	ftpack.SetProtoID(3001)

	var securityList []*Qot_Common.Security
	security := transStockCode(stockCode)
	securityList = append(securityList, security)

	var subTypeList []int32
	subTypeNum := transSubType(subType)
	subTypeList = append(subTypeList, subTypeNum)

	// isSubOrUnSub := true
	// isRegOrUnRegPush := true

	var regPushRehabTypeList []int32
	regPushRehabType := int32(1)
	regPushRehabTypeList = append(regPushRehabTypeList, regPushRehabType)

	reqData := &Qot_Sub.Request{
		C2S: &Qot_Sub.C2S{
			SecurityList: securityList,
			SubTypeList:  subTypeList,
			IsSubOrUnSub: &isSubOrUnSub,
			// IsRegOrUnRegPush:     &isRegOrUnRegPush,
			RegPushRehabTypeList: regPushRehabTypeList,
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

// QotSubRecv handler
func QotSubRecv(data []byte) error {
	fut := &Qot_Sub.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	fmt.Println(fut)

	return nil
}
