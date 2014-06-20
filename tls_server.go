package rpc2

import (
	"crypto/tls"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type SecureServer struct {
	conn    net.Listener
	rpcServ *rpc.Server
}

func NewSecureServer(addr, certFile, keyFile string) (*SecureServer, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	s := &SecureServer{rpcServ: rpc.NewServer()}
	s.conn, err = tls.Listen("tcp", addr, &tls.Config{Certificates: []tls.Certificate{cert}})
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *SecureServer) RegisterName(name string, rcvr interface{}) error {
	return s.rpcServ.RegisterName(name, rcvr)
}

func (s *SecureServer) Serve() error {
	if conn, err := s.conn.Accept(); err == nil {
		go s.rpcServ.ServeCodec(jsonrpc.NewServerCodec(conn))
	} else {
		return err
	}
	return nil
}
