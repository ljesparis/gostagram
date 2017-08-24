package gostagram

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// generate a signature to send it to secure a request.
func (c Client) generateSignature(endpoint string, params Parameters) (string, error) {
	sig := endpoint
	for key, val := range params {
		sig += fmt.Sprintf("|%s=%s", key, val)
	}

	tmp := hmac.New(sha256.New, []byte(c.clientSecret))
	_, err := tmp.Write([]byte(sig))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(tmp.Sum(nil)), nil
}
