package openapiv2

import "github.com/go-kratos/grpc-gateway/v2/protoc-gen-openapiv2/generator"

type options struct {
	generatorOptions []generator.Option
}

type HandlerOption func(opt *options)

func WithGeneratorOptions(opts ...generator.Option) HandlerOption {
	return func(opt *options) {
		opt.generatorOptions = opts
	}
}
