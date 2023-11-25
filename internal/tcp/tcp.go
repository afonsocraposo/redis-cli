package tcp

import (
	"net"
)

type TCPClient struct {
	conn *net.TCPConn
}

// connect to TCP server
func (c *TCPClient) Connect(address string) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}

	c.conn = conn
	return nil
}

// disconnect from TCP server
func (c *TCPClient) Disconnect() error {
	if c.conn != nil {
		err := c.conn.Close()
		if err != nil {
			return err
		}
		c.conn = nil
	}
	return nil
}

func (c *TCPClient) Write(input string) error {
	if c.conn != nil {
		_, err := c.conn.Write([]byte(input))
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *TCPClient) Read() (string, error) {
	reply := make([]byte, 1024)
	if c.conn != nil {
		_, err := c.conn.Read(reply)
		if err != nil {
			return "", err
		}
	}
    s := string(reply)
	return s, nil
}
