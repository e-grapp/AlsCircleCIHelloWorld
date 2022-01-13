package cryptokey

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
)

// GenerateEd25519Keys generates Ed25519 public and private key pair with given size for SSH.
func GenerateEd25519Keys(passphrase string) (pubKey string, privKey string, err error) {
	// generate private key
	var (
		privateKey ed25519.PrivateKey
		publicKey  ed25519.PublicKey
	)
	if publicKey, privateKey, err = ed25519.GenerateKey(rand.Reader); err != nil {
		return
	}

	// encode public key
	var sshPK ssh.PublicKey
	if sshPK, err = ssh.NewPublicKey(publicKey); err != nil {
		return
	}
	pubBytes := ssh.MarshalAuthorizedKey(sshPK)

	// encode private key
	var privBytes []byte
	privBytes, err = encodePEMBlock(&pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: MarshalED25519PrivateKey(privateKey),
	}, passphrase)
	if err != nil {
		return
	}

	return string(pubBytes), string(privBytes), nil
}
