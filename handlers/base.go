package handlers

import (
	"fmt"
	"limgo/futupb/Qot_Common"
	"strings"
)

func transStockCode(code string) *Qot_Common.Security {
	codes := strings.Split(code, ".")

	stock := &Qot_Common.Security{
		Market: new(int32),
		Code:   new(string),
	}

	switch codes[0] {
	case "HK":
		*stock.Market = int32(1)
	case "US":
		*stock.Market = int32(11)
	case "SH":
		*stock.Market = int32(21)
	case "SZ":
		*stock.Market = int32(22)
	}

	*stock.Code = codes[1]

	return stock
}

func transSubType(t string) int32 {
	k := "SubType_" + strings.Title(t)
	if v, ok := Qot_Common.SubType_value[k]; ok {
		return v
	}
	fmt.Println("transSubType error:", t)
	return int32(0)
}
