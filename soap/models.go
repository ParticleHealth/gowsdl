package soap

import (
	"encoding/xml"
	"strings"
)

type XMLResponseEnvelope struct {
	XMLName     xml.Name                  `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	XMLNS       string                    `xml:"xmlns:SOAP-ENV,attr,omitempty"`
	XSI         string                    `xml:"xmlns:xsi,attr,omitempty"`
	Header      interface{}               `xml:"SOAP-ENV:Header"`
	Body        SoapResponseBodyInterface `xml:"Body"`
	Attachments []MIMEMultipartAttachment `xml:"attachments,omitempty"`
}

type XMLResponseHeader struct {
	XMLName   xml.Name    `xml:"SOAP-ENV:Header"`
	Text      string      `xml:",chardata"`
	Action    interface{} `xml:"Action"`
	MessageID interface{} `xml:"MessageID"`
	To        interface{} `xml:"To"`
	RelatedTo interface{} `xml:"RelatesTo"`
}

type XMLResponseBody struct {
	Text     string `xml:",chardata"`
	Response interface{}
}

func (s *XMLResponseEnvelope) GetBody() SoapResponseBodyInterface {
	return s.Body
}

func (s *XMLResponseEnvelope) GetHeader() interface{} {
	return s.Header
}

func (s *XMLResponseEnvelope) SetBody(body SoapResponseBodyInterface) {
	s.Body = body
}

func (s *XMLResponseEnvelope) SetHeader(header interface{}) {
	s.Header = header.(*XMLResponseHeader)
}

func (s *XMLResponseEnvelope) SetXMLName(xmlName xml.Name) {
	s.XMLName = xmlName
}

func (s *XMLResponseEnvelope) GetAttachments() []MIMEMultipartAttachment {
	return s.Attachments
}

func (b *XMLResponseBody) ErrorFromFault() error {
	return nil
}

func (b *XMLResponseBody) SetContent(content interface{}) {
	b.Response = content
}

// XMLAdhocQueryBody is the Body for AdhocQueryRequest.

type XMLQueryResponse struct {
	XMLName            xml.Name `xml:"AdhocQueryResponse"`
	Text               string   `xml:",chardata"`
	Xmlns              string   `xml:"xmlns,attr"`
	Query              string   `xml:"query,attr"`
	Status             string   `xml:"status,attr"`
	RegistryObjectList struct {
		Text    string   `xml:",chardata"`
		Xmlns   string   `xml:"xmlns,attr"`
		Rim     string   `xml:"rim,attr"`
		Message []string `xml:"Message"`
	} `xml:"RegistryObjectList"`
}
type XMLQueryResponseBody struct {
	Text     string `xml:",chardata"`
	XMLName  xml.Name
	Response *XMLQueryResponse `xml:"AdhocQueryResponse"`
}

func (b *XMLQueryResponseBody) ErrorFromFault() error {
	return nil
}

func (b *XMLQueryResponseBody) SetContent(content interface{}) {
	b.Response = content.(*XMLQueryResponse)
}

func (b *XMLQueryResponseBody) HasError() bool {
	return false
}

type RetrieveDocumentSetResponse struct {
	XMLName          xml.Name `xml:"RetrieveDocumentSetResponse"`
	Text             string   `xml:",chardata"`
	Xmlns            string   `xml:"xmlns,attr"`
	Xdsb             string   `xml:"xdsb,attr"`
	RegistryResponse struct {
		Text   string `xml:",chardata"`
		Xmlns  string `xml:"xmlns,attr"`
		Rs     string `xml:"rs,attr"`
		Status string `xml:"status,attr"`
	} `xml:"RegistryResponse"`
	DocumentResponse struct {
		Text     string  `xml:",chardata"`
		MimeType string  `xml:"mimeType"`
		Document *Binary `xml:"Document"`
	} `xml:"DocumentResponse"`
}

type XMLDocRetrieveBody struct {
	Text     string `xml:",chardata"`
	XMLName  xml.Name
	Response *RetrieveDocumentSetResponse `xml:"RetrieveDocumentSetResponse"`
}

func (b *XMLDocRetrieveBody) ErrorFromFault() error {
	return nil
}

func (b *XMLDocRetrieveBody) SetContent(content interface{}) {
	b.Response = content.(*RetrieveDocumentSetResponse)
}

func (b *XMLDocRetrieveBody) HasError() bool {
	return strings.Contains(b.Response.RegistryResponse.Status, "Failure")
}
