package openapiv2

import (
	"fmt"
	"net/http"

	"github.com/go-kratos/kratos/v2/api/metadata"
	"github.com/go-kratos/kratos/v2/api/proto/kratos/api"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
	_ "github.com/go-kratos/swagger-api/openapiv2/swagger_ui" // import statik static files
	mux "github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/generator"
	"github.com/rakyll/statik/fs"
	"google.golang.org/protobuf/types/pluginpb"
)

func NewHandler() http.Handler {
	service := metadata.NewService()
	r := mux.NewRouter()
	h := http1.DefaultHandleOptions()

	r.HandleFunc("/q/services", func(w http.ResponseWriter, r *http.Request) {
		services, err := service.ListServices(r.Context())
		if err != nil {
			h.Error(w, r, err)
			return
		}

		if err := h.Encode(w, r, services); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/q/service/{name}", func(w http.ResponseWriter, r *http.Request) {
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
		if len(protoSet.File) == 0 {
			h.Error(w, r, fmt.Errorf("proto file is empty"))
			return
		}
		if protoSet.File[len(protoSet.File)-1].Name == nil {
			h.Error(w, r, fmt.Errorf("proto file's name is null"))
			return
		}
		target = *protoSet.File[len(protoSet.File)-1].Name

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
		if len(resp.File) == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write([]byte("{}"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(*resp.File[0].Content))
	}).Methods("GET")

	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	staticServer := http.FileServer(statikFS)
	sh := http.StripPrefix("/q/swagger-ui", staticServer)
	r.PathPrefix("/q/swagger-ui").Handler(sh)
	return r
}
