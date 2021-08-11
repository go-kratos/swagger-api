package openapiv2

import (
	"context"
	"fmt"

	"github.com/go-kratos/grpc-gateway/v2/protoc-gen-openapiv2/generator"
	"github.com/go-kratos/kratos/v2/api/metadata"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/pluginpb"
)

// Service is service
type Service struct {
	ser  *metadata.Server
	opts []generator.Option
}

// New service
func New(srv *grpc.Server, opts ...generator.Option) *Service {
	return &Service{
		ser:  metadata.NewServer(srv),
		opts: opts,
	}
}

// ListServices list services
func (s *Service) ListServices(ctx context.Context, in *metadata.ListServicesRequest) (*metadata.ListServicesReply, error) {
	return s.ser.ListServices(ctx, &metadata.ListServicesRequest{})
}

// GetServiceOpenAPI get service open api
func (s *Service) GetServiceOpenAPI(ctx context.Context, in *metadata.GetServiceDescRequest, onlyRPC bool) (string, error) {
	protoSet, err := s.ser.GetServiceDesc(ctx, in)
	if err != nil {
		return "", err
	}
	files := protoSet.FileDescSet.File
	var target string
	if len(files) == 0 {
		return "", fmt.Errorf("proto file is empty")
	}
	if files[len(files)-1].Name == nil {
		return "", fmt.Errorf("proto file's name is null")
	}
	target = *files[len(files)-1].Name

	req := new(pluginpb.CodeGeneratorRequest)
	req.FileToGenerate = []string{target}
	var para = ""
	req.Parameter = &para
	req.ProtoFile = files

	g := generator.NewGenerator(s.opts...)
	resp, err := g.Gen(req, onlyRPC)
	if err != nil {
		return "", err
	}
	if len(resp.File) == 0 {
		return "{}", nil
	}
	return *resp.File[0].Content, nil
}
