package http

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

type Crypto struct {
}

func (it *Crypto) Generate() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization:       []string{"crusj"},
		OrganizationalUnit: []string{"crusj"},
		CommonName:         "crusj",
	}
	template := x509.Certificate{
		Raw:                         nil,
		RawTBSCertificate:           nil,
		RawSubjectPublicKeyInfo:     nil,
		RawSubject:                  nil,
		RawIssuer:                   nil,
		Signature:                   nil,
		SignatureAlgorithm:          0,
		PublicKeyAlgorithm:          0,
		PublicKey:                   nil,
		Version:                     0,
		SerialNumber:                serialNumber,
		Issuer:                      pkix.Name{},
		Subject:                     subject,
		NotBefore:                   time.Now(),
		NotAfter:                    time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:                    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		Extensions:                  nil,
		ExtraExtensions:             nil,
		UnhandledCriticalExtensions: nil,
		ExtKeyUsage:                 []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		UnknownExtKeyUsage:          nil,
		BasicConstraintsValid:       false,
		IsCA:                        false,
		MaxPathLen:                  0,
		MaxPathLenZero:              false,
		SubjectKeyId:                nil,
		AuthorityKeyId:              nil,
		OCSPServer:                  nil,
		IssuingCertificateURL:       nil,
		DNSNames:                    nil,
		EmailAddresses:              nil,
		IPAddresses:                 []net.IP{net.ParseIP("127.0.0.1")},
		URIs:                        nil,
		PermittedDNSDomainsCritical: false,
		PermittedDNSDomains:         nil,
		ExcludedDNSDomains:          nil,
		PermittedIPRanges:           nil,
		ExcludedIPRanges:            nil,
		PermittedEmailAddresses:     nil,
		ExcludedEmailAddresses:      nil,
		PermittedURIDomains:         nil,
		ExcludedURIDomains:          nil,
		CRLDistributionPoints:       nil,
		PolicyIdentifiers:           nil,
	}
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, &pk)
	certOut, _ := os.Create("Cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	keyOut, _ := os.Create("Key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
