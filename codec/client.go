// 协议内容编码
package codec

import (
	"bytes"
	"encoding/binary"
	"github.com/cbwfree/micro-game/utils/errors"
)

// ClientHead 客户端消息头
type ClientHead struct {
	Serial  uint16
	Cmd     uint32
	DataLen uint32
}

// Client 客户端消息
type Client struct {
	bin     binary.ByteOrder // 端类型
	mixHead []uint8          // 混淆头
	mixLen  int              // 混淆长度
	headLen int              // 消息头长度
}

func (c *Client) HeadLen() int {
	return c.headLen + c.mixLen
}

// 设置混淆 (最长4位)
func (c *Client) SetMix(mix ...uint8) {
	c.mixHead = mix
	c.mixLen = len(mix)
}

// 设置大小端
func (c *Client) SetByteOrder(bin binary.ByteOrder) {
	c.bin = bin
}

// Marshal 编码消息
func (c *Client) Marshal(head *ClientHead, data []byte) (b []byte, err error) {
	var buf = new(bytes.Buffer)
	var msgLen = len(data)

	if c.mixLen > 0 {
		for i := 0; i < c.mixLen; i++ {
			err = binary.Write(buf, c.bin, c.mixHead[i])
		}
	}

	err = binary.Write(buf, c.bin, head.Serial)
	err = binary.Write(buf, c.bin, head.Cmd)
	err = binary.Write(buf, c.bin, uint32(msgLen))
	err = binary.Write(buf, c.bin, data)

	b = buf.Bytes()

	return b, err
}

// Unmarshal 解码消息
func (c *Client) Unmarshal(raw []byte) (head *ClientHead, data []byte, err error) {
	var rawLen = len(raw)
	if rawLen < c.headLen {
		return nil, nil, errors.Invalid("msg head length error")
	}

	// 校验head
	if c.mixLen > 0 {
		heads := raw[:c.mixLen]
		raw = raw[c.mixLen:]
		rawLen = rawLen - c.mixLen

		for i, head := range heads {
			if c.mixHead[i] != head {
				return nil, nil, errors.Invalid("msg head check error")
			}
		}
	}

	head = new(ClientHead)
	head.Serial = c.bin.Uint16(raw[:2])
	head.Cmd = c.bin.Uint32(raw[2:6])
	head.DataLen = c.bin.Uint32(raw[6:c.headLen])

	if rawLen > c.headLen {
		var maxLen = int(head.DataLen) + c.headLen
		if rawLen < maxLen {
			return nil, nil, errors.Invalid("msg data length error")
		}
		data = raw[c.headLen:maxLen]
	}

	return head, data, nil
}

func NewBinClient(bin binary.ByteOrder, mix ...uint8) *Client {
	c := &Client{
		bin:     bin,
		headLen: 10,
	}
	c.SetMix(mix...)
	return c
}

func NewClient(mix ...uint8) *Client {
	return NewBinClient(binary.BigEndian, mix...)
}
