package jwt

import (
	"base64"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"reflect"
	"strings"
)

type Header struct {
	JwtType   string
	Algorithm string
}
type Payload struct {
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
	payload := Payload{
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
