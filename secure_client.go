package rpc2

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net"
)

func SecureDial(network, addr, caFile string) (net.Conn, error) {
	caPem, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	rootCAs := x509.NewCertPool()
	if !rootCAs.AppendCertsFromPEM(caPem) {
		return nil, errors.New("unable to append cert from PEM")
	}

	conf := &tls.Config{RootCAs: rootCAs}
	conn, err := tls.Dial(network, addr, conf)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
