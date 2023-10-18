package go_mailchimp_transactional

import (
	"encoding/base64"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

const (
	apiName string = "Mailchimp Transactional"
	apiUrl  string = "https://mandrillapp.com/api/1.0"
)

type Service struct {
	apiKey        string
	httpService   *go_http.Service
	errorResponse *ErrorResponse
}

type ServiceConfig struct {
	ApiKey string
}

func NewService(cfg *ServiceConfig) (*Service, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if cfg.ApiKey == "" {
		return nil, errortools.ErrorMessage("ApiKey not provided")
	}

	var service = Service{
		apiKey: cfg.ApiKey,
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	service.httpService = httpService

	return &service, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add authentication
	headers := requestConfig.NonDefaultHeaders
	if headers == nil {
		headers = &http.Header{}
	}
	headers.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("anystring:%s", service.apiKey)))))
	requestConfig.NonDefaultHeaders = headers

	// add error model
	service.errorResponse = &ErrorResponse{}
	requestConfig.ErrorModel = service.errorResponse

	request, response, e := service.httpService.HttpRequest(requestConfig)
	if e != nil {
		if service.errorResponse.Message != "" {
			e.SetMessage(service.errorResponse.Message)
		}
	}

	if e != nil {
		return request, response, e
	}

	return request, response, nil
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiUrl, path)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.apiKey
}

func (service *Service) ApiCallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) ApiReset() {
	service.httpService.ResetRequestCount()
}

func (service *Service) ErrorResponse() *ErrorResponse {
	return service.errorResponse
}
