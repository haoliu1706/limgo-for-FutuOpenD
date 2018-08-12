package limgo

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"limgo/futupb/InitConnect"
	"math/rand"
	"strconv"

	"log"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
)

// LimVersion client version
const LimVersion = 101

// Config for init Request
type Config struct {
	Host string
	Port string
}

// Request is a service struct
type Request struct {
	Conn     net.Conn
	Error    error
	SyncMode bool
	InitData struct {
		ServerVer         int32
		LoginUserID       uint64
		ConnID            uint64
		ConnAESKey        string
		KeepAliveInterval int32
	}
}

// New returns a new Request pointer
func New(conf Config) *Request {

	conn, err := net.Dial("tcp", conf.Host+":"+conf.Port)
	if err != nil {
		log.Fatalln("failed to connect opend server")
	}

	log.Println("connected server: " + conf.Host + ":" + conf.Port)

	r := &Request{
		Conn:     conn,
		Error:    err,
		SyncMode: false,
	}

	r.initRequest()

	return r
}

// Send data
func (r *Request) Send(pack *FutuPack) error {
	var packData []byte
	var err error

	// pack
	packData, err = pack.Pack()
	if err != nil {
		return err
	}

	// incr write timeout time
	r.Conn.SetWriteDeadline(time.Now().Add(3 * time.Second))

	// write
	_, err = r.Conn.Write(packData)
	if err != nil {
		return err
	}

	return err
}

// Recv data
func (r *Request) Recv() []byte {

	// scanner
	scanner := bufio.NewScanner(r.Conn)

	// set max scan token size
	scanner.Buffer([]byte{}, bufio.MaxScanTokenSize*10)

	// split
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF && data[0] == 'F' {
			if len(data) >= 44 {
				length := uint32(0)
				binary.Read(bytes.NewReader(data[12:16]), binary.LittleEndian, &length)

				if int(length)+44 <= len(data) {
					return int(length) + 44, data[:int(length)+44], nil
				}
			}
		}

		return
	})

	// scan
	for scanner.Scan() {

		// read timeout
		r.Conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		pack := new(FutuPack)
		err := pack.Unpack(scanner.Bytes())
		if err != nil {
			log.Fatalln("unpack error", err)
		}

		fmt.Println(pack.nProtoID, "==>")
		protoID := uint32(pack.nProtoID)

		if r.SyncMode {
			return pack.arrBody
		}

		// trans protoID
		handlerName, ok := TransHandlerID(protoID)
		if ok {
			recvFunc := "recv." + handlerName
			if !HasHandler(recvFunc) {
				fmt.Println("recvFunc error: ", protoID, recvFunc)
			}

			// Recv
			Do(recvFunc, pack.arrBody)

		} else {
			fmt.Println("trans handler " + strconv.Itoa(int(protoID)) + " error")
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("invalid package", err)
	}

	return nil
}

// Close connect
func (r *Request) Close() {
	r.Conn.Close()
}

// InitRequest init request
func (r *Request) initRequest() {

	rand.Seed(int64(time.Now().Nanosecond()))
	clientID := strconv.Itoa(int(rand.Int31()))
	clientVer := int32(LimVersion)

	ftpack := &FutuPack{}
	ftpack.SetProtoID(uint32(1001))

	reqData := &InitConnect.Request{
		C2S: &InitConnect.C2S{
			ClientID:  &clientID,
			ClientVer: &clientVer,
		},
	}
	pbData, err := proto.Marshal(reqData)
	if err != nil {
		log.Fatalln("marshal error: ", err)
	}

	ftpack.SetBody(pbData)
	r.Send(ftpack)

	r.SyncMode = true
	rawPack := r.Recv()

	// initConnect data
	retData := initConnectData(rawPack)

	r.InitData.ServerVer = *retData.S2C.ServerVer
	r.InitData.LoginUserID = *retData.S2C.LoginUserID
	r.InitData.ConnID = *retData.S2C.ConnID
	r.InitData.ConnAESKey = *retData.S2C.ConnAESKey
	r.InitData.KeepAliveInterval = *retData.S2C.KeepAliveInterval

	r.SyncMode = false

	return
}

// unmarshal initConnectData
func initConnectData(data []byte) *InitConnect.Response {
	fut := &InitConnect.Response{}
	err := proto.Unmarshal(data, fut)
	if err != nil {
		log.Fatal("InitConnect unmarshaling error:", err)
	}

	if fut.GetRetType() != int32(0) {
		log.Fatalln("InitConnect rettype: ", fut.GetRetType())
	}

	return fut
}

// KeepAlive keep alive
func (r *Request) KeepAlive(output bool) {
	go func() {
		// interval := r.InitData.KeepAliveInterval
		tick := time.NewTicker(3 * time.Second)

		for {
			select {
			case <-tick.C:
				if output {
					fmt.Println("tick")
				}

				// 1004 KeepAlive
				Do("send.KeepAlive", r)
			}
		}
	}()

}
