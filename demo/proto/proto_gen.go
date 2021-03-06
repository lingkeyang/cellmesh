// Auto generated by github.com/davyxu/cellmesh/protogen
// DO NOT EDIT!

package proto

import (
	"fmt"
	"reflect"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/json"
	_ "github.com/davyxu/cellnet/codec/binary"
)

// Make compiler import happy
var (
	_ cellnet.Peer
	_ cellnet.Codec
	_ reflect.Type
	_ fmt.Formatter
)

type ResultCode int32

const (
	ResultCode_NoError      ResultCode = 0
	ResultCode_GameNotReady ResultCode = 1
)

var (
	ResultCodeMapperValueByName = map[string]int32{
		"NoError":      0,
		"GameNotReady": 1,
	}

	ResultCodeMapperNameByValue = map[int32]string{
		0: "NoError",
		1: "GameNotReady",
	}

	ResultCodeMapperTrailingCommentByValue = map[int32]string{
		0: "",
		1: "",
	}
)

func (self ResultCode) String() string {
	return ResultCodeMapperNameByValue[int32(self)]
}

type ServerInfo struct {
	IP   string
	Port int32
}

type LoginREQ struct {
	Version  string
	Platform string
	UID      string
}

type LoginACK struct {
	Result    ResultCode
	Server    ServerInfo
	GameToken string
}

type VerifyREQ struct {
	GameToken string
}

type VerifyACK struct {
	Result ResultCode
}

type ChatREQ struct {
	Content string
}

type ChatACK struct {
	Content string
}

type ClientID struct {
	ID    int64  // 客户端在网关上的SessionID
	SvcID string // 客户端在哪个网关
}

type BindBackendACK struct {
	ID int64
}

type CloseClientACK struct {
	ID  []int64
	All bool
}

type ClientClosedACK struct {
	ID ClientID
}

type PingACK struct {
}

func (self *ServerInfo) String() string      { return fmt.Sprintf("%+v", *self) }
func (self *LoginREQ) String() string        { return fmt.Sprintf("%+v", *self) }
func (self *LoginACK) String() string        { return fmt.Sprintf("%+v", *self) }
func (self *VerifyREQ) String() string       { return fmt.Sprintf("%+v", *self) }
func (self *VerifyACK) String() string       { return fmt.Sprintf("%+v", *self) }
func (self *ChatREQ) String() string         { return fmt.Sprintf("%+v", *self) }
func (self *ChatACK) String() string         { return fmt.Sprintf("%+v", *self) }
func (self *ClientID) String() string        { return fmt.Sprintf("%+v", *self) }
func (self *BindBackendACK) String() string  { return fmt.Sprintf("%+v", *self) }
func (self *CloseClientACK) String() string  { return fmt.Sprintf("%+v", *self) }
func (self *ClientClosedACK) String() string { return fmt.Sprintf("%+v", *self) }
func (self *PingACK) String() string         { return fmt.Sprintf("%+v", *self) }

// agent
var (
	Handle_Agent_BindBackendACK = func(ev cellnet.Event) { panic("'BindBackendACK' not handled") }
	Handle_Agent_CloseClientACK = func(ev cellnet.Event) { panic("'CloseClientACK' not handled") }
	Handle_Agent_Default        func(ev cellnet.Event)
)

// agent_frontend
var (
	Handle_Agent_frontend_PingACK = func(ev cellnet.Event) { panic("'PingACK' not handled") }
	Handle_Agent_frontend_Default func(ev cellnet.Event)
)

// game
var (
	Handle_Game_ChatREQ   = func(ev cellnet.Event) { panic("'ChatREQ' not handled") }
	Handle_Game_VerifyREQ = func(ev cellnet.Event) { panic("'VerifyREQ' not handled") }
	Handle_Game_Default   func(ev cellnet.Event)
)

// login
var (
	Handle_Login_LoginREQ = func(ev cellnet.Event) { panic("'LoginREQ' not handled") }
	Handle_Login_Default  func(ev cellnet.Event)
)

func GetMessageHandler(svcName string) cellnet.EventCallback {

	switch svcName {
	case "agent":
		return func(ev cellnet.Event) {
			switch ev.Message().(type) {
			case *BindBackendACK:
				Handle_Agent_BindBackendACK(ev)
			case *CloseClientACK:
				Handle_Agent_CloseClientACK(ev)
			default:
				if Handle_Agent_Default != nil {
					Handle_Agent_Default(ev)
				}
			}
		}
	case "agent_frontend":
		return func(ev cellnet.Event) {
			switch ev.Message().(type) {
			case *PingACK:
				Handle_Agent_frontend_PingACK(ev)
			default:
				if Handle_Agent_frontend_Default != nil {
					Handle_Agent_frontend_Default(ev)
				}
			}
		}
	case "game":
		return func(ev cellnet.Event) {
			switch ev.Message().(type) {
			case *ChatREQ:
				Handle_Game_ChatREQ(ev)
			case *VerifyREQ:
				Handle_Game_VerifyREQ(ev)
			default:
				if Handle_Game_Default != nil {
					Handle_Game_Default(ev)
				}
			}
		}
	case "login":
		return func(ev cellnet.Event) {
			switch ev.Message().(type) {
			case *LoginREQ:
				Handle_Login_LoginREQ(ev)
			default:
				if Handle_Login_Default != nil {
					Handle_Login_Default(ev)
				}
			}
		}
	}

	return nil
}

func init() {

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("json"),
		Type:  reflect.TypeOf((*LoginREQ)(nil)).Elem(),
		ID:    18837,
	})

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("json"),
		Type:  reflect.TypeOf((*LoginACK)(nil)).Elem(),
		ID:    46204,
	})

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*VerifyREQ)(nil)).Elem(),
		ID:    13457,
	})

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*VerifyACK)(nil)).Elem(),
		ID:    40824,
	})

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*ChatREQ)(nil)).Elem(),
		ID:    29052,
	})

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*ChatACK)(nil)).Elem(),
		ID:    56419,
	})

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*ClientID)(nil)).Elem(),
		ID:    44352,
	})

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*BindBackendACK)(nil)).Elem(),
		ID:    5768,
	})

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*CloseClientACK)(nil)).Elem(),
		ID:    58040,
	})

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*ClientClosedACK)(nil)).Elem(),
		ID:    50844,
	})

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),
		Type:  reflect.TypeOf((*PingACK)(nil)).Elem(),
		ID:    16241,
	})

}
