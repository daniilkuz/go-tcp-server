package main

import "net"

type Server struct {
	listenAddr string
	ln         net.Listener
	quitc      chan struct{}
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitc:      make(chan struct{}),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return nil
	}
	defer ln.Close()
	s.ln = ln
	<-s.quitc
	return nil
}

func main() {}
