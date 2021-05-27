module github.com/go-kratos/swagger-api

go 1.15

require (
	github.com/go-kratos/kratos/v2 v2.0.0-20210527131947-b32e7d6e701f
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	github.com/rakyll/statik v0.1.7
	google.golang.org/genproto v0.0.0-20210524171403-669157292da3
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0 => github.com/longXboy/grpc-gateway/v2 v2.0.0-20210512024025-a0dff65b1b3d
