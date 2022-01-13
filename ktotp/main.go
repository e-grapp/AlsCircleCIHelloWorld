package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"time"

	"bitbucket.org/ai69/amoy"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
)

func main() {
	args := flag.Args()
	if len(args) <= 0 {
		log.Fatal("no arguments")
	}

	var (
		secret = strings.ToUpper(strings.Join(args, ""))
		token  string
		err    error
	)
	if detailed {
		token, err = showDetailedTOTPToken(secret)
	} else {
		token, err = getTOTPToken(secret)
	}
	if err != nil {
		log.Fatalw("failed to generate token", "secret", secret, zap.Error(err))
	}
	if newline {
		fmt.Println(token)
	} else {
		fmt.Print(token)
	}
}

func getHOTPToken(secret string, interval int64) (string, error) {
	// Add padding to the secret if it's too short
	if last := len(secret) * 5 % 8; last != 0 {
		secret += strings.Repeat("=", 8-last)
	}

	// Converts secret to base32 Encoding
	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	// Signing the value using HMAC-SHA1 Algorithm
	hash := hmac.New(sha1.New, key)
	err = binary.Write(hash, binary.BigEndian, uint64(interval))
	if err != nil {
		return "", err
	}
	h := hash.Sum(nil)

	// Get 32 bit chunk from hash starting at the offset
	offset := h[19] & 0x0f
	truncated := binary.BigEndian.Uint32(h[offset : offset+4])
	truncated &= 0x7fffffff
	code := truncated % 1000000

	return fmt.Sprintf("%06d", code), nil
}

func getTOTPToken(secret string) (string, error) {
	// The TOTP token is just a HOTP token seeded with every 30 seconds.
	interval := time.Now().Unix() / 30
	return getHOTPToken(secret, interval)
}

const timeFmt = "2006-01-02 15:04:05"

func showDetailedTOTPToken(secret string) (string, error) {
	now := time.Now()
	interval := now.Unix() / 30
	currentStart := amoy.ParseUnixTime(interval * 30)
	nextStart := amoy.ParseUnixTime((interval + 1) * 30)
	leftDur := nextStart.Sub(now)
	token, err := getHOTPToken(secret, interval)
	if err != nil {
		return "", err
	}

	OutputStderr("* Start At: %s\n", currentStart.Format(timeFmt))
	OutputStderr("*  Current: %s\n", now.Format(timeFmt))
	OutputStderr("*  Next At: %s\n", nextStart.Format(timeFmt))
	OutputStderr("* Left Sec: %.03f\n", leftDur.Seconds())
	return token, nil
}

func OutputStderr(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
}
