package sdk

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"gopkg.in/resty.v1"
)

func UdunPost(gateway string, merchantKey string, path string, body string, repObj *ResultMsg) error {
	ret, err := resty.New().NewRequest().SetBody(parseParams(merchantKey, body)).Post(gateway + path)
	if err != nil {
		return err
	}

	fmt.Println(string(ret.Body()))
	err = json.Unmarshal(ret.Body(), repObj)
	if err != nil {
		return err
	}

	return nil
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func randomString(l int) string {
	var result bytes.Buffer
	for i := 0; i < l; {
		temp := randInt(65, 90)
		result.WriteByte(byte(temp))
		i++
	}
	return result.String()
}

func sign(key string, timestamp string, nonce string, body string) string {
	raw := body + key + nonce + timestamp
	h := md5.New()
	h.Write([]byte(raw))
	return hex.EncodeToString(h.Sum(nil))
}

func parseParams(merchantKey string, body string) map[string]string {
	params := make(map[string]string)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := randomString(6)
	sign := sign(merchantKey, timestamp, nonce, body)
	params["timestamp"] = timestamp
	params["nonce"] = nonce
	params["sign"] = sign
	params["body"] = body

	return params
}

func CheckSign(key string, timestamp string, nonce string, body string, sgn string) bool {
	signStr := sign(key, timestamp, nonce, body)
	return signStr == sgn
}
