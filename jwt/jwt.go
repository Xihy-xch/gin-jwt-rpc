package jwt

import (
	"base64"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	//"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Header struct {
	JwtType   string
	Algorithm string
}
type Payload struct {
	Iss      string
	Aud      string
	Exp      string
	Iat      string
	Username string
}

type Jwt struct {
	Header  Header
	Payload Payload
	secret  string
}

func NewJwt() *Jwt {
	header := Header{
		JwtType: "Jwt",
	}
	return &Jwt{
		Header: header,
	}
}
func (jwt *Jwt) SetHeader(algorithm string) {
	jwt.Header.Algorithm = algorithm
}

func (jwt *Jwt) SetPayload(name string) {
	curTime := time.Now().Unix()
	iat := strconv.FormatInt(curTime, 10)
	exp := strconv.FormatInt(curTime+300, 10)
	//fmt.Println(strconv.FormatInt(curTime, 10))
	payload := Payload{
		Iss:      "server",
		Aud:      "client",
		Exp:      exp,
		Iat:      iat,
		Username: name,
	}
	jwt.Payload = payload
}

func (jwt *Jwt) SetSecret(secret string) {
	jwt.secret = secret
}

func (jwt *Jwt) GetHeaderToBase64() string {
	str := structToString(jwt.Header)
	// fmt.Println(str)
	return base64.EncodeToBase64(str)
}

func (jwt *Jwt) GetPayloadToBase64() string {
	str := structToString(jwt.Payload)
	// fmt.Println(str)
	return base64.EncodeToBase64(str)
}

func structToString(object interface{}) string {

	var str strings.Builder
	structValue := reflect.ValueOf(object)
	structType := reflect.TypeOf(object)

	for i := 0; i < structType.NumField(); i++ {
		str.WriteString(structType.Field(i).Name)
		str.WriteString(":")
		// fmt.Println(str.String())
		str.WriteString(structValue.Field(i).Interface().(string))
		str.WriteString(",")
		// fmt.Println(str.String())
	}
	// fmt.Println(str.String())
	return str.String()
}

func (jwt *Jwt) ComputeHmacSha256(message string) string {
	key := []byte(jwt.secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	//	fmt.Println(h.Sum(nil))
	sha := hex.EncodeToString(h.Sum(nil))
	//	fmt.Println(sha)

	//	hex.EncodeToString(h.Sum(nil))
	return sha
}

func (jwt *Jwt) GetSecrectToken(username string) string {

	jwt.SetHeader("HS256")
	jwt.SetPayload(username)
	header := jwt.GetHeaderToBase64()
	payload := jwt.GetPayloadToBase64()
	signature := jwt.ComputeHmacSha256(header + "." + payload)
	token := header + "." + payload + "." + signature
	return token
}

func ParseHeader(headerStr string) string {
	headerStr = base64.DecodeFromBase64(headerStr)
	return headerStr
}

func ParsePayload(payloadStr string) (string, bool) {
	var flag bool
	payloadStr = base64.DecodeFromBase64(payloadStr)

	attributes := strings.Split(payloadStr, ",")
	expTimeStr := strings.Split(attributes[2], ":")
	curTime := time.Now().Unix()
	if expTime, _ := strconv.ParseInt(expTimeStr[1], 10, 64); expTime > curTime {
		flag = true
	} else {
		flag = false
	}

	return payloadStr, flag
}

func (jwt *Jwt) CheckToken(token string) bool {
	jwtInfo := strings.Split(token, ".")

	_, ok := ParsePayload(jwtInfo[1])
	if !ok {
		return false
	}
	signature := jwt.ComputeHmacSha256(jwtInfo[0] + "." + jwtInfo[1])
	// fmt.Println(signature)
	if signature == jwtInfo[2] {
		return true
	} else {
		return false
	}
}
