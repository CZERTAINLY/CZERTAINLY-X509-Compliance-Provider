package utils

import (
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/rsa"
	"log"

	"github.com/zmap/zcrypto/x509"
)

// GetPublicKeySize returns the size of the public key in bits
func GetPublicKeySize(certificate *x509.Certificate) int {
	switch publicKey := certificate.PublicKey.(type) {
	case *rsa.PublicKey:
		return publicKey.N.BitLen()
	case *x509.AugmentedECDSA:
		return publicKey.Pub.Curve.Params().BitSize
	case *ecdsa.PublicKey:
		return publicKey.Curve.Params().BitSize
	case *dsa.PublicKey:
		return publicKey.G.BitLen()
	default:
		log.Print("unsupported private key")
		return 0
	}
}
