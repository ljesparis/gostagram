package gostagram

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func (c *Client) generateSignature(endpoint string, params Params) (string, error) {
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
