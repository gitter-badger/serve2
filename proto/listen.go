package proto

import (
	"net"
	"net/http"

	"github.com/joushou/serve2/utils"
)

// ListenChecker is the provided Check function, identical to
// ProtocolHandler.Check
type ListenChecker func(header []byte) (match bool, required int)

// ListenProxy provides a net.Listener whose Accept will only return matched
// protocols.
type ListenProxy struct {
	listener *utils.ChannelListenProxy
	Checker  ListenChecker
}

func (ListenProxy) String() string {
	return "ListenProxy"
}

// Listener returns the proxy net.Listener.
func (l *ListenProxy) Listener() net.Listener {
	return l.listener
}

// Handle pushes the connection to the ListenProxy server.
func (l *ListenProxy) Handle(c net.Conn) (net.Conn, error) {
	l.Listener.Push(c)
	return nil, nil
}

// Check just calls the ListenChecker.
func (l *ListenProxy) Check(header []byte) (bool, int) {
	return l.Checker(header)
}

// NewListenProxy returns a fully initialized ListenProxy.
func NewListenProxy(checker ListenChecker, buffer int) *ListenProxy {
	listener := utils.NewChannelListenProxy(make(chan net.Conn, buffer), nil)
	return &ListenProxy{listener: listener, Checker: ListenChecker}
}
