package cryptokey

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
)

// GenerateECDSAKeys generates ECDSA public and private key pair with given size for SSH.
func GenerateECDSAKeys(bitSize int, passphrase string) (pubKey string, privKey string, err error) {
	// generate private key
	var privateKey *ecdsa.PrivateKey
	if privateKey, err = ecdsa.GenerateKey(curveFromLength(bitSize), rand.Reader); err != nil {
		return
	}

	// encode public key
	var (
		bytes     []byte
		publicKey ssh.PublicKey
	)
	if publicKey, err = ssh.NewPublicKey(privateKey.Public()); err != nil {
		return
	}
	pubBytes := ssh.MarshalAuthorizedKey(publicKey)

	// encode private key
	if bytes, err = x509.MarshalECPrivateKey(privateKey); err != nil {
		return
	}
	var privBytes []byte
	privBytes, err = encodePEMBlock(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: bytes,
	}, passphrase)
	if err != nil {
		return
	}

	return string(pubBytes), string(privBytes), nil
}

func curveFromLength(l int) elliptic.Curve {
	switch l {
	case 224:
		return elliptic.P224()
	case 256:
		return elliptic.P256()
	case 384:
		return elliptic.P384()
	case 521:
		return elliptic.P521()
	}
	return elliptic.P384()
}
