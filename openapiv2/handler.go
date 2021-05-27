package openapiv2

import (
	"net/http"

	"github.com/go-kratos/kratos/v2/api/metadata"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
	_ "github.com/go-kratos/swagger-api/openapiv2/swagger_ui" // import statik static files
	mux "github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
)

func NewHandler() http.Handler {
	service := New(nil)
	r := mux.NewRouter()
	h := http1.DefaultHandleOptions()

	r.HandleFunc("/q/services", func(w http.ResponseWriter, r *http.Request) {
		services, err := service.ListServices(r.Context(), &metadata.ListServicesRequest{})
		if err != nil {
			h.Error(w, r, err)
			return
		}

		if err := h.Encode(w, r, services); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/q/service/{name}", func(w http.ResponseWriter, r *http.Request) {
		var in metadata.GetServiceDescRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		if err := binding.MapProto(&in, mux.Vars(r)); err != nil {
			h.Error(w, r, err)
			return
		}

		content, err := service.GetServiceOpenAPI(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
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
