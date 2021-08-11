package openapiv2

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-kratos/grpc-gateway/v2/protoc-gen-openapiv2/generator"
	"github.com/go-kratos/kratos/v2/api/metadata"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
	_ "github.com/go-kratos/swagger-api/openapiv2/swagger_ui/statik" // import statik static files
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
)

func NewHandler(handlerOpts ...HandlerOption) http.Handler {
	opts := &options{
		// Compatible with default UseJSONNamesForFields is true
		generatorOptions: []generator.Option{generator.UseJSONNamesForFields(true)},
	}

	for _, o := range handlerOpts {
		o(opts)
	}

	service := New(nil, opts.generatorOptions...)
	r := mux.NewRouter()

	r.HandleFunc("/q/services", func(w http.ResponseWriter, r *http.Request) {
		services, err := service.ListServices(r.Context(), &metadata.ListServicesRequest{})
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(services)
	}).Methods("GET")

	r.HandleFunc("/q/service/{name}", func(w http.ResponseWriter, r *http.Request) {
		raws := mux.Vars(r)
		vars := make(url.Values, len(raws))
		for k, v := range raws {
			vars[k] = []string{v}
		}
		var in metadata.GetServiceDescRequest
		if err := binding.BindQuery(vars, &in); err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		content, err := service.GetServiceOpenAPI(r.Context(), &in, false)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(content))
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
