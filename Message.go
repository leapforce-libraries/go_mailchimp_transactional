package go_mailchimp_transactional

import (
	"cloud.google.com/go/civil"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
	"time"
)

type Message struct {
	Diag               string               `json:"diag"`
	DocumentId         string               `json:"document_id"`
	State              string               `json:"state"`
	Opens              int                  `json:"opens"`
	BgtoolsCode        interface{}          `json:"bgtools_code"`
	Reject             interface{}          `json:"reject"`
	SmtpEvents         []MessageSmtpEvent   `json:"smtp_events"`
	Email              string               `json:"email"`
	ElasticsearchIndex string               `json:"elasticsearch_index"`
	Version            string               `json:"@version"`
	Clicks             int                  `json:"clicks"`
	Ts                 int                  `json:"ts"`
	Subject            string               `json:"subject"`
	Subaccount         interface{}          `json:"subaccount"`
	Timestamp          time.Time            `json:"@timestamp"`
	LogClassification  string               `json:"log_classification"`
	Tags               []string             `json:"tags"`
	Resends            []interface{}        `json:"resends"`
	Id                 string               `json:"_id"`
	Sender             string               `json:"sender"`
	Template           interface{}          `json:"template"`
	OpensDetail        []MessageOpenDetail  `json:"opens_detail"`
	ClicksDetail       []MessageClickDetail `json:"clicks_detail"`
}

type MessageSmtpEvent struct {
	Diag          string `json:"diag"`
	Type          string `json:"type"`
	SourceIp      string `json:"source_ip"`
	DestinationIp string `json:"destination_ip"`
	Ts            int    `json:"ts"`
	Size          int    `json:"size"`
}

type MessageOpenDetail struct {
	Ua       string `json:"ua"`
	Ip       string `json:"ip"`
	Ts       int    `json:"ts"`
	Location string `json:"location"`
}

type MessageClickDetail struct {
	Ua       string `json:"ua"`
	Ip       string `json:"ip"`
	Ts       int    `json:"ts"`
	Location string `json:"location"`
	Url      string `json:"url"`
}

type SearchMessagesConfig struct {
	Key      string      `json:"key"`
	Query    *string     `json:"query,omitempty"`
	DateFrom *civil.Date `json:"date_from,omitempty"`
	DateTo   *civil.Date `json:"date_to,omitempty"`
	Tags     *[]string   `json:"tags,omitempty"`
	Senders  *[]string   `json:"senders,omitempty"`
	ApiKeys  *[]string   `json:"api_keys,omitempty"`
	Limit    *int        `json:"limit,omitempty"`
}

func (service *Service) SearchMessages(cfg *SearchMessagesConfig) (*[]Message, *errortools.Error) {
	if cfg == nil {
		return nil, errortools.ErrorMessage("SearchMessagesConfig must not be nil")
	}

	var messages []Message

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url("messages/search"),
		BodyModel:     cfg,
		ResponseModel: &messages,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &messages, nil
}
