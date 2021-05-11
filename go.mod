module github.com/go-kratos/swagger-api

go 1.15

require (
	github.com/go-kratos/kratos/v2 v2.0.0-20210511081852-42313e936823
	github.com/gogo/protobuf v1.3.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	google.golang.org/genproto v0.0.0-20210510173355-fb37daa5cd7a
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0 => github.com/longXboy/grpc-gateway/v2 v2.0.0-20210511071124-76ae0ebed08b
