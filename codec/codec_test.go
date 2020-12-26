package codec

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	c := NewClient(1, 2, 3, 4)

	res, err := c.Marshal(&ClientHead{0, 10001, 0}, []byte("server"))
	if err != nil {
		fmt.Printf("Server Marshal Error: %s\n", err.Error())
		return
	}

	head, data, err := c.Unmarshal(res)
	if err != nil {
		fmt.Printf("Client Unmarshal Error: %s\n", err.Error())
		return
	}

	fmt.Printf("Serial: %d, Cmd: %d, Data: %s\n", head.Serial, head.Cmd, data)
}

func TestServer(t *testing.T) {
	s := NewServer(3, 6, 7, 9)

	res, err := s.Marshal(&ServerHead{0, 10001, 0, 0}, []byte("client"))
	if err != nil {
		fmt.Printf("Client Marshal Error: %s", err.Error())
		return
	}

	head, data, err := s.Unmarshal(res)
	if err != nil {
		fmt.Printf("Server Unmarshal Error: %s\n", err.Error())
		return
	}

	fmt.Printf("Serial: %d, Cmd: %d, Code: %d, Data: %s\n", head.Serial, head.Cmd, head.Code, data)
}
