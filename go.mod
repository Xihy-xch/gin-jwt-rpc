module gin-jw-rpc

go 1.15

require (
	base64 v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.16
	jwt v0.0.0-00010101000000-000000000000
	middleWare v0.0.0-00010101000000-000000000000
	models v0.0.0-00010101000000-000000000000
	sql v0.0.0-00010101000000-000000000000
)

replace jwt => ./jwt

replace base64 => ./base64

replace cookie => ./cookie

replace models => ./models

replace sql => ./sql

replace middleWare => ./middleWare
