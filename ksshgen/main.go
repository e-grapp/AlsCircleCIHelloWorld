package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"os/user"
	"strings"

	"bitbucket.org/ai69/keiki/ksshgen/cryptokey"
	"bitbucket.org/ai69/keiki/ksshgen/randomart"
	"github.com/1set/gut/yos"
	"github.com/1set/gut/ystring"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

func main() {
	// check file exists
	var (
		algoType = strings.ToLower(algoType)
		l        = log.With("algo", algoType)
	)
	if ystring.IsBlank(fileName) {
		// fileName = fmt.Sprintf("id_%s_%d", algoType, bitsNumber)
		fileName = fmt.Sprintf("id_%s", algoType)
	}
	if (yos.ExistFile(fileName) || yos.ExistFile(fileName+".pub")) && !override {
		log.Fatalw("file exists", "file", fileName)
		return
	}

	var (
		pubKey  string
		privKey string
		err     error
	)
	// generate
	switch algoType {
	case "rsa":
		if bitsNumber == 0 {
			bitsNumber = 1024
		}
		pubKey, privKey, err = cryptokey.GenerateRSAKeys(bitsNumber, passphrase)
	case "dsa":
		if bitsNumber == 0 {
			bitsNumber = 1024
		}
		pubKey, privKey, err = cryptokey.GenerateDSAKeys(bitsNumber, passphrase)
	case "ecdsa":
		if bitsNumber == 0 {
			bitsNumber = 384
		}
		pubKey, privKey, err = cryptokey.GenerateECDSAKeys(bitsNumber, passphrase)
	case "ed25519":
		bitsNumber = 256
		pubKey, privKey, err = cryptokey.GenerateEd25519Keys(passphrase)
	default:
		err = fmt.Errorf("unsupported algorithm type: %s", algoType)
	}
	if err != nil {
		l.Fatalw("generate key failed", zap.Error(err))
	}

	// comment
	if ystring.IsEmpty(title) {
		host, e1 := os.Hostname()
		usr, e2 := user.Current()
		if e1 == nil && e2 == nil {
			title = fmt.Sprintf("%s@%s", usr.Username, host)
		} else {
			l.Warnw("fail to get hostname or current username", "err_host", e1, "err_user", e2)
		}
	}
	if ystring.IsNotBlank(title) {
		pubKey = fmt.Sprintf("%s %s\n", strings.TrimSpace(pubKey), title)
	}

	// preview
	if err = previewPublicKey(fmt.Sprintf("%s %d", strings.ToUpper(algoType), bitsNumber), pubKey); err != nil {
		l.Warnw("fail to preview public key", "public_key", pubKey, zap.Error(err))
	}

	// save as files
	pubFile, privFile := fileName+".pub", fileName
	if err := writeFileString(pubFile, pubKey); err != nil {
		l.Fatalw("write public key file failed", "file", pubFile, zap.Error(err))
	}
	l.Infow("saved public key file", "file", pubFile)

	if err := writeFileString(privFile, privKey); err != nil {
		l.Fatalw("write private key file failed", "file", privFile, zap.Error(err))
	}
	l.Infow("saved private key file", "file", privFile, "has_passphrase", passphrase != "")
}

func previewPublicKey(title, pubKey string) error {
	pk, cmt, _, _, err := ssh.ParseAuthorizedKey([]byte(pubKey))
	if err != nil {
		return err
	}
	fpMD5 := ssh.FingerprintLegacyMD5(pk)
	fpSHA256 := ssh.FingerprintSHA256(pk)
	arrSHA256 := sha256.Sum256(pk.Marshal())
	rb := randomart.GenerateSubtitled(arrSHA256[:], title, "SHA256")

	log.Infow("preview generated public key", "fingerprint_md5", fpMD5, "fingerprint_sha256", fpSHA256, "algo", pk.Type(), "comment", cmt)
	fmt.Println(rb.String())
	fmt.Println(pubKey)
	return nil
}

func writeFileString(path string, content string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()
	_, err = w.WriteString(content)
	return err
}
