package ridder

import (
	"encoding/base64"
	"encoding/xml"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type inboundXMLMessage struct {
	MessageID              string `json:"MessageId"`
	Base64EncodedXMLString string `json:"Base64EncodedXmlString"`
}

func (service *Service) SendXMLMessage(messageID string, object interface{}) *errortools.Error {
	b, err := xml.Marshal(object)
	if err != nil {
		return errortools.ErrorMessage(err)
	}
	base64encodedXMLString := base64.StdEncoding.EncodeToString(b)

	message := inboundXMLMessage{
		MessageID:              messageID,
		Base64EncodedXMLString: base64encodedXMLString,
	}

	requestConfig := go_http.RequestConfig{
		URL:       service.url("inboundxmlmessage"),
		BodyModel: message,
	}
	_, _, e := service.post(&requestConfig)
	if e != nil {
		return e
	}

	return nil
}
