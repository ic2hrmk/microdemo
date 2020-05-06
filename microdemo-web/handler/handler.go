package handler

import (
	"encoding/json"
	"net/http"

	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
)

type (
	DemoCallResponse struct {
		RequestDetails RequestDetails `json:"request_details"`
		ServiceDetails ServiceDetails `json:"service_details"`
	}

	RequestDetails struct {
		Method   string `json:"method"`
		SourceIP string `json:"source_ip"`
	}

	ServiceDetails struct {
		ServiceName string `json:"service_name"`
		Version     string `json:"version"`
	}
)

func MicrodemoCall(service web.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Infof("got a new HTTP request to service [%s]",
			service.Options().Name)

		response := DemoCallResponse{
			RequestDetails: RequestDetails{
				Method:   r.Method,
				SourceIP: r.RemoteAddr,
			},
			ServiceDetails: ServiceDetails{
				ServiceName: service.Options().Name,
				Version:     service.Options().Version,
			},
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
}
