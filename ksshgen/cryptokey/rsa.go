package cryptokey

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
)

// GenerateRSAKeys generates RSA public and private key pair with given size for SSH.
func GenerateRSAKeys(bitSize int, passphrase string) (pubKey string, privKey string, err error) {
	// generate private key
	var (
		privateKey *rsa.PrivateKey
		publicKey  ssh.PublicKey
	)
	if privateKey, err = rsa.GenerateKey(rand.Reader, rsaSizeFromLength(bitSize)); err != nil {
		return
	}

	// validate private key
	if err = privateKey.Validate(); err != nil {
		return
	}

	// encode public key
	if publicKey, err = ssh.NewPublicKey(privateKey.Public()); err != nil {
		return
	}
	pubBytes := ssh.MarshalAuthorizedKey(publicKey)

	// encode private key
	var privBytes []byte
	privBytes, err = encodePEMBlock(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}, passphrase)
	if err != nil {
		return
	}

	return string(pubBytes), string(privBytes), nil
}

func rsaSizeFromLength(l int) int {
	switch l {
	case 1024, 2048, 4096, 8192:
		return l
	default:
		return 1024
	}
}
