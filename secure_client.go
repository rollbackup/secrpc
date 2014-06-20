package secrpc

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net"
)

func SecureDialWithCert(network, addr, caFile string) (net.Conn, error) {
	caPem, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	return SecureDial(network, addr, caPem)
}

func SecureDial(network, addr string, cert []byte) (net.Conn, error) {
	rootCAs := x509.NewCertPool()
	if !rootCAs.AppendCertsFromPEM(cert) {
		return nil, errors.New("unable to append cert from PEM")
	}

	conf := &tls.Config{RootCAs: rootCAs}
	conn, err := tls.Dial(network, addr, conf)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
