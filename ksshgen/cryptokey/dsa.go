package cryptokey

import (
	"crypto/dsa"
	"crypto/rand"
	"encoding/asn1"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
)

// GenerateDSAKeys generates DSA public and private key pair with given size for SSH.
func GenerateDSAKeys(bitSize int, passphrase string) (pubKey string, privKey string, err error) {
	params := new(dsa.Parameters)

	// see http://golang.org/pkg/crypto/dsa/#ParameterSizes
	if err = dsa.GenerateParameters(params, rand.Reader, dsaSizeFromLength(bitSize)); err != nil {
		return
	}

	var privateKey dsa.PrivateKey
	privateKey.PublicKey.Parameters = *params

	// this generates a public & private key pair
	if err = dsa.GenerateKey(&privateKey, rand.Reader); err != nil {
		return
	}

	// generate public key
	var publicKey ssh.PublicKey
	if publicKey, err = ssh.NewPublicKey(&privateKey.PublicKey); err != nil {
		return
	}

	// encode public key
	pubBytes := ssh.MarshalAuthorizedKey(publicKey)

	// encode private key
	var (
		bytes     []byte
		privBytes []byte
	)
	if bytes, err = asn1.Marshal(privateKey); err != nil {
		return
	}
	privBytes, err = encodePEMBlock(&pem.Block{
		Type:  "DSA PRIVATE KEY",
		Bytes: bytes,
	}, passphrase)
	if err != nil {
		return
	}

	return string(pubBytes), string(privBytes), nil
}

func dsaSizeFromLength(l int) dsa.ParameterSizes {
	switch l {
	case 1024:
		return dsa.L1024N160
	case 2048:
		return dsa.L2048N224
	case 3072:
		return dsa.L3072N256
	default:
		return dsa.L2048N256
	}
}
