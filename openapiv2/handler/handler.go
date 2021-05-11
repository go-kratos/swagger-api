package handler

import (
	"net/http"

	"github.com/go-kratos/kratos/v2/api/metadata"
	"github.com/go-kratos/kratos/v2/api/proto/kratos/api"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/generator"
	"google.golang.org/protobuf/types/pluginpb"
)

func New() http.Handler {
	service := metadata.NewService()
	r := mux.NewRouter()
	h := http1.DefaultHandleOptions()
	r.HandleFunc("/service/{name}/metadata", func(w http.ResponseWriter, r *http.Request) {
		var in api.GetServiceMetaRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		protoSet, err := service.GetServiceMeta(r.Context(), in.Name)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		var target string
		if len(protoSet.File) > 0 {
			target = *protoSet.File[len(protoSet.File)-1].Name
		}
		req := new(pluginpb.CodeGeneratorRequest)
		req.FileToGenerate = []string{target}
		var para = ""
		req.Parameter = &para
		req.ProtoFile = protoSet.File

		var g generator.Generator
		resp, err := g.Gen(req)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		if err := h.Encode(w, r, resp.File); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")
	return r
}
