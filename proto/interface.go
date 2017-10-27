package proto

import (
	"github.com/newflydd/ngrok/conn"
)

type Protocol interface {
	GetName() string
	WrapConn(conn.Conn, interface{}) conn.Conn
}
