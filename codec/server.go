// 协议内容编码
package codec

import (
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
)

// ServerHead 服务器消息头
type ServerHead struct {
	Serial  uint16
	Cmd     uint32
	Code    uint32
	DataLen uint32
}

// Client 服务端消息
type Server struct {
	bin     binary.ByteOrder // 端类型
	mixHead []uint8          // 混淆头
	mixLen  int              // 混淆长度
	headLen int              // 消息头长度
}

func (c *Server) HeadLen() int {
	return c.headLen + c.mixLen
}

// 设置混淆 (最长4位)
func (c *Server) SetMix(mix ...uint8) {
	c.mixHead = mix
	c.mixLen = len(mix)
}

// 设置大小端
func (c *Server) SetByteOrder(bin binary.ByteOrder) {
	c.bin = bin
}

// Marshal 编码消息
func (c *Server) Marshal(head *ServerHead, data []byte) (b []byte, err error) {
	var buf = new(bytes.Buffer)
	var msgLen = len(data)

	if c.mixLen > 0 {
		for i := 0; i < c.mixLen; i++ {
			err = binary.Write(buf, c.bin, c.mixHead[i])
		}
	}

	err = binary.Write(buf, c.bin, head.Serial)
	err = binary.Write(buf, c.bin, head.Cmd)
	err = binary.Write(buf, c.bin, head.Code)
	err = binary.Write(buf, c.bin, uint32(msgLen))
	err = binary.Write(buf, c.bin, data)

	b = buf.Bytes()

	return b, err
}

// Unmarshal 解码消息
func (c *Server) Unmarshal(raw []byte) (head *ServerHead, data []byte, err error) {
	var rawLen = len(raw)
	if rawLen < c.headLen {
		return nil, nil, errors.New("msg head length error")
	}

	// 校验head
	if c.mixLen > 0 {
		heads := raw[:c.mixLen]
		raw = raw[c.mixLen:]
		rawLen = rawLen - c.mixLen

		for i, head := range heads {
			if c.mixHead[i] != head {
				return nil, nil, errors.New("msg head check error")
			}
		}
	}

	head = new(ServerHead)
	head.Serial = c.bin.Uint16(raw[:2])
	head.Cmd = c.bin.Uint32(raw[2:6])
	head.Code = c.bin.Uint32(raw[6:10])
	head.DataLen = c.bin.Uint32(raw[10:c.headLen])

	if rawLen > c.headLen {
		var maxLen = int(head.DataLen) + c.headLen
		if rawLen < maxLen {
			return nil, nil, errors.New("msg data length error")
		}
		data = raw[c.headLen:maxLen]
	}

	return head, data, nil
}

func NewBinServer(bin binary.ByteOrder, mix ...uint8) *Server {
	c := &Server{
		bin:     bin,
		headLen: 14,
	}
	c.SetMix(mix...)
	return c
}

func NewServer(mix ...uint8) *Server {
	return NewBinServer(binary.BigEndian, mix...)
}
