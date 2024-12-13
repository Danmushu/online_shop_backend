package jwt

import (
	"encoding/json"
	"errors"
	"project/global"
	"project/utils/encrypt"
	"project/utils/types"
	"strings"
	"time"
)

const durationDefault = 24 * time.Hour

func Marshal(model any) string {
	headerMap := make(map[string]any)
	headerMap["alg"] = "HS256"
	headerMap["type"] = "JWT"

	payloadMap := make(map[string]any)
	data := types.ToMap(model)
	for k, v := range data {
		payloadMap[k] = v
	}

	// payload 过期时间 签名时间
	if payloadMap["iat"] == nil || payloadMap["exp"] == nil {
		iat := time.Now()
		exp := iat.Add(durationDefault)
		payloadMap["iat"] = iat.Unix()
		payloadMap["exp"] = exp.Unix()
	}

	header, _ := json.Marshal(headerMap)
	payload, _ := json.Marshal(payloadMap)

	headerEncode := encrypt.B64encode(string(header))
	payloadEncode := encrypt.B64encode(string(payload))

	signature := encrypt.HmacSha256(headerEncode+"."+payloadEncode, global.JWTSalt)
	jwt := headerEncode + "." + payloadEncode + "." + signature
	return jwt
}

func Unmarshal(token string, output any) error {
	count := strings.Count(token, ".")
	if count != 2 {
		return errors.New("jwt 解析错误")
	}
	li := strings.Split(token, ".")
	if encrypt.HmacSha256(li[0]+"."+li[1], global.JWTSalt) != li[2] {
		return errors.New("jwt 校验错误")
	}
	if json.Unmarshal(encrypt.B64decode(li[1]), &output) != nil {
		return errors.New("jwt 负载json解析错误")
	}
	return nil
}
