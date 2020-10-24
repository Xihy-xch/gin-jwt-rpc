module gin-jwt-rpc

go 1.15

require (
	base64 v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	jwt v0.0.0-00010101000000-000000000000
)

replace jwt => ./jwt

replace base64 => ./base64
