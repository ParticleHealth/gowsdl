package soap

import (
	"encoding/xml"
)

// XMLRoot defines node with root attribute.
type XMLRoot struct {
	Text      string `xml:",chardata"`
	Root      string `xml:"root,attr"`
	Extension string `xml:"extension,attr,omitempty"`
}

// XMLValue defines node with value attribute.
type XMLValue struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value,attr"`
}

// XMLCode defines node with code attribute.
type XMLCode struct {
	Text       string `xml:",chardata"`
	Code       string `xml:"code,attr"`
	CodeSystem string `xml:"codeSystem,attr,omitempty"`
}

// XMLDevice defines device node.
type XMLDevice struct {
	Text           string  `xml:",chardata"`
	ClassCode      string  `xml:"classCode,attr,omitempty"`
	DeterminerCode string  `xml:"determinerCode,attr,omitempty"`
	ID             XMLRoot `xml:"id"`
}

// XMLActor defines sender/receiver node.
type XMLActor struct {
	Text     string    `xml:",chardata"`
	TypeCode string    `xml:"typeCode,attr,omitempty"`
	Device   XMLDevice `xml:"device"`
}

// XMLAddress defines address node.
type XMLAddress struct {
	Text           string `xml:",chardata"`
	MustUnderstand string `xml:"mustUnderstand,attr,omitempty"`
	Address        string `xml:"Address"`
}

// XMLIdentifier defines identifier node.
type XMLIdentifier struct {
	Text          string  `xml:",chardata"`
	Value         XMLRoot `xml:"value"`
	SemanticsText string  `xml:"semanticsText,omitempty"`
}

// XMLAction is the action in the header.
type XMLAction struct {
	Text           string `xml:",chardata"`
	XMLNS          string `xml:"xmlns,attr,omitempty"`
	MustUnderstand string `xml:"soap:mustUnderstand,attr"`
}

// XMLTo is the "To" in the header.
type XMLTo struct {
	Text           string `xml:",chardata"`
	XMLNS          string `xml:"xmlns,attr,omitempty"`
	MustUnderstand string `xml:"soap:mustUnderstand,attr"`
	D3p1           string `xml:"xmlns:d3p1,attr,omitempty"`
}

// XMLMessageID is the message ID in the header.
type XMLMessageID struct {
	Text           string `xml:",chardata"`
	XMLNS          string `xml:"xmlns,attr,omitempty"`
	MustUnderstand string `xml:"soap:mustUnderstand,attr,omitempty"`
}

// XMLSecurity is the security in the header.
type XMLSecurity struct {
	Text      string        `xml:",chardata"`
	XMLNS     string        `xml:"xmlns,attr"`
	Assertion XMLAssertion  `xml:"saml:Assertion"`
	Signature *XMLSignature `xml:"Signature,omitempty"`
}

// XMLAssertion is the assertion in the security.
type XMLAssertion struct {
	Text              string     `xml:",chardata"`
	XMLNS             string     `xml:"xmlns:saml,attr"`
	ID                string     `xml:"ID,attr"`
	IssueInstant      string     `xml:"IssueInstant,attr"`
	Version           string     `xml:"Version,attr"`
	CertificateIssuer XMLFormat  `xml:"saml:Issuer"`
	Subject           XMLSubject `xml:"saml:Subject"`
	//Conditions         XMLConditions         `xml:"Conditions"`
	AuthnStatement     XMLAuthnStatement     `xml:"saml:AuthnStatement"`
	AttributeStatement XMLAttributeStatement `xml:"saml:AttributeStatement"`
	XSI                string                `xml:"xmlns:xsi,attr,omitempty"`
}

type XMLFormat struct {
	Text   string `xml:",chardata"`
	Format string `xml:"Format,attr"`
}

// XMLSubject is the subject in the assertion.
type XMLSubject struct {
	Text string `xml:",chardata"`
	// NameID              XMLNameID              `xml:"NameID"`
	SubjectConfirmation XMLSubjectConfirmation `xml:"saml:SubjectConfirmation"`
}

// XMLNameID is the name ID in the subject.
type XMLNameID struct {
	Text   string `xml:",chardata"`
	Format string `xml:"Format,attr"`
}

// XMLSubjectConfirmation is the subject confirmation in the subject.
type XMLSubjectConfirmation struct {
	Text   string `xml:",chardata"`
	Method string `xml:"Method,attr"`
}

// XMLConditions is the conditions in the assertion.
type XMLConditions struct {
	Text         string `xml:",chardata"`
	NotBefore    string `xml:"NotBefore,attr"`
	NotOnOrAfter string `xml:"NotOnOrAfter,attr"`
}

// XMLAuthnStatement is the authentication statement in the assertion.
type XMLAuthnStatement struct {
	Text         string `xml:",chardata"`
	AuthnInstant string `xml:"AuthnInstant,attr"`
	// SubjectLocality XMLSubjectLocality `xml:"SubjectLocality"`
	AuthnContext XMLAuthnContext `xml:"saml:AuthnContext"`
}

// XMLSubjectLocality is the subject locality in the statement.
type XMLSubjectLocality struct {
	Text    string `xml:",chardata"`
	Address string `xml:"Address,attr"`
	DNSName string `xml:"DNSName,attr"`
}

// XMLAuthnContext is the authentication context in the statement.
type XMLAuthnContext struct {
	Text                    string `xml:",chardata"`
	AuthnContextClassRef    string `xml:"saml:AuthnContextClassRef"`
	AuthenticatingAuthority string `xml:"saml:AuthenticatingAuthority"`
}

// XMLAttributeStatement is the attribute statement in the assertion.
type XMLAttributeStatement struct {
	Text      string         `xml:",chardata"`
	Attribute []XMLAttribute `xml:"saml:Attribute"`
}

// XMLAttribute is the attribute in the assertion.
type XMLAttribute struct {
	Text           string            `xml:",chardata"`
	Name           string            `xml:"Name,attr"`
	NameFormat     string            `xml:"NameFormat,attr"`
	AttributeValue XMLAttributeValue `xml:"saml:AttributeValue"`
}

// XMLAttributeValue is the value in the attribute.
type XMLAttributeValue struct {
	Text string `xml:",chardata"`
	Type string `xml:"xsi:type,attr,omitempty"`
}

// XMLSignature is the signature in the security element.
type XMLSignature struct {
	XMLName        xml.Name      `xml:"Signature,omitempty"`
	Text           string        `xml:",chardata"`
	SignedInfo     XMLSignedInfo `xml:"SignedInfo"`
	SignatureValue string        `xml:"SignatureValue"`
	KeyInfo        XMLKeyInfo    `xml:"KeyInfo"`
}

// XMLSignedInfo is the signed info element in the signature.
type XMLSignedInfo struct {
	Text                   string                `xml:",chardata"`
	CanonicalizationMethod XMLAlgo               `xml:"CanonicalizationMethod"`
	SignatureMethod        XMLAlgo               `xml:"SignatureMethod"`
	Reference              XMLSignatureReference `xml:"Reference"`
}

// XMLAlgo defines an algorithm element.
type XMLAlgo struct {
	Text      string `xml:",chardata"`
	Algorithm string `xml:"Algorithm,attr"`
}

// XMLTransforms defines a transforms element.
type XMLTransforms struct {
	Text      string    `xml:",chardata"`
	Transform []XMLAlgo `xml:"Transform"`
}

// XMLSignatureReference defines a reference element in the signature.
type XMLSignatureReference struct {
	Text         string        `xml:",chardata"`
	URI          string        `xml:"URI,attr"`
	Transforms   XMLTransforms `xml:"Transforms"`
	DigestMethod XMLAlgo       `xml:"DigestMethod"`
	DigestValue  string        `xml:"DigestValue"`
}

// XMLKeyInfo is the key info element in the signature.
type XMLKeyInfo struct {
	Text     string      `xml:",chardata"`
	X509Data XMLX509Data `xml:"X509Data"`
}

// XMLX509Data is the X509 data element in the key info.
type XMLX509Data struct {
	Text            string `xml:",chardata"`
	X509Certificate string `xml:"X509Certificate"`
}

// XMLRequestEnvelope is the envelope IHE SOAP requests
type XMLRequestEnvelope struct {
	XMLName   xml.Name         `xml:"S:Envelope"`
	XMLNS     string           `xml:"xmlns:S,attr,omitempty"`
	XMLEnv    string           `xml:"xmlns:env,attr,omitempty"`
	XMLSaml   string           `xml:"xmlns:saml,attr,omitempty"`
	XMLSoap   string           `xml:"xmlns:soap,attr,omitempty"`
	XMLWSA    string           `xml:"xmlns:wsa,attr,omitempty"`
	XMLWSSE   string           `xml:"xmlns:wsse,attr,omitempty"`
	XMLWSSE11 string           `xml:"xmlns:wsse11,attr,omitempty"`
	XMLWSU    string           `xml:"xmlns:wsu,attr,omitempty"`
	XMLXS     string           `xml:"xmlns:xs,attr,omitempty"`
	XMLQuery  string           `xml:"xmlns:query,attr,omitempty"`
	Text      string           `xml:",chardata"`
	XSI       string           `xml:"xmlns:xsi,attr,omitempty"`
	XSD       string           `xml:"xmlns:xsd,attr,omitempty"`
	Header    XMLRequestHeader `xml:"env:Header"`
	Body      XMLRequestBody   `xml:"S:Body"`
}

// XMLRequestHeader is the header for the IHE SOAP requests
type XMLRequestHeader struct {
	Text      string       `xml:",chardata"`
	Action    XMLAction    `xml:"Action"`
	MessageID XMLMessageID `xml:"MessageID"`
	To        XMLTo        `xml:"To"`
	Security  XMLSecurity  `xml:"wsse:Security"`
}

// XMLRequestBody is the request body for the IHE SOAP requests
type XMLRequestBody struct {
	Text    string `xml:",chardata"`
	Request interface{}
}

type XMLResponseEnvelope struct {
	XMLName     xml.Name                  `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	XMLNS       string                    `xml:"xmlns:SOAP-ENV,attr,omitempty"`
	XSI         string                    `xml:"xmlns:xsi,attr,omitempty"`
	Header      *XMLResponseHeader        `xml:"SOAP-ENV:Header"`
	Body        SoapResponseBodyInterface `xml:"Body"`
	Attachments []MIMEMultipartAttachment `xml:"attachments,omitempty"`
}

type XMLResponseHeader struct {
	XMLName   xml.Name     `xml:"SOAP-ENV:Header"`
	Text      string       `xml:",chardata"`
	Action    XMLAction    `xml:"Action"`
	MessageID XMLMessageID `xml:"MessageID"`
	To        XMLTo        `xml:"To"`
	RelatedTo XMLTo        `xml:"RelatesTo"`
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

type XPingResponse struct {
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
type XPingReply struct {
	// XMLName xml.Name `xml:"http://example.com/service.xsd PingReply"`

	Message    string `xml:"Message,omitempty"`
	Attachment []byte `xml:"Attachment,omitempty"`
}
type XMLAdhocQueryResponseBody struct {
	Text     string `xml:",chardata"`
	XMLName  xml.Name
	Response *XPingResponse `xml:"AdhocQueryResponse"`
}

func (b *XMLAdhocQueryResponseBody) ErrorFromFault() error {
	return nil
}

func (b *XMLAdhocQueryResponseBody) SetContent(content interface{}) {
	b.Response = content.(*XPingResponse)
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

// XMLAdhocQueryRequest is the request message for Registry Stored Query [ITI-18].
type XMLAdhocQueryRequest struct {
	XMLName        xml.Name         `xml:"query:AdhocQueryRequest"`
	Text           string           `xml:",chardata"`
	Query          string           `xml:"xmlns:query,attr"`
	Rim            string           `xml:"xmlns:rim,attr"`
	Rs             string           `xml:"rs,attr"`
	ResponseOption XMLReponseOption `xml:"query:ResponseOption"`
	AdhocQuery     XMLAdhocQuery    `xml:"rim:AdhocQuery"`
}

// XMLAdhocQueryResponse is the response message for XMLAdhocQueryRequest.
type XMLAdhocQueryResponse struct {
	XMLName            xml.Name `xml:"query:AdhocQueryResponse"`
	Text               string   `xml:",chardata"`
	Status             string   `xml:"status,attr"`
	Xmlns              string   `xml:"xmlns,attr"`
	Rim                string   `xml:"rim,attr"`
	Xsi                string   `xml:"xsi,attr"`
	SchemaLocation     string   `xml:"schemaLocation,attr"`
	RegistryObjectList struct {
		Text            string               `xml:",chardata"`
		ExtrinsicObject []XMLExtrinsicObject `xml:"ExtrinsicObject"`
		ObjectRef       []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Q    string `xml:"q,attr"`
			Rim  string `xml:"rim,attr"`
		} `xml:"ObjectRef"`
	} `xml:"RegistryObjectList"`
}

// XMLReponseOption is the ResponseOption in the AdhocQueryRequest.
type XMLReponseOption struct {
	Text                  string `xml:",chardata"`
	ReturnComposedObjects string `xml:"returnComposedObjects,attr"`
	ReturnType            string `xml:"returnType,attr"`
}

// XMLAdhocQuery is the AdhocQuery in the AdhocQueryRequest.
type XMLAdhocQuery struct {
	Text string         `xml:",chardata"`
	ID   string         `xml:"id,attr"`
	Slot []XMLQuerySlot `xml:"rim:Slot"`
}

// XMLQuerySlot is a Slot in the AdhocQuery.
type XMLQuerySlot struct {
	Text      string       `xml:",chardata"`
	Name      string       `xml:"name,attr"`
	ValueList XMLValueList `xml:"rim:ValueList"`
}

// XMLValueList is a list of values in the QuerySlot.
type XMLValueList struct {
	Text  string `xml:",chardata"`
	Value string `xml:"rim:Value"`
}

// XMLExtrinsicObject is the extrinsic object in the AdhocQueryResponse.
type XMLExtrinsicObject struct {
	Text       string    `xml:",chardata"`
	ID         string    `xml:"id,attr"`
	IsOpaque   string    `xml:"isOpaque,attr"`
	MimeType   string    `xml:"mimeType,attr"`
	ObjectType string    `xml:"objectType,attr"`
	Status     string    `xml:"status,attr"`
	Q          string    `xml:"q,attr"`
	Rim        string    `xml:"rim,attr"`
	Slot       []XMLSlot `xml:"Slot"`
	Name       struct {
		Text            string `xml:",chardata"`
		LocalizedString struct {
			Text    string `xml:",chardata"`
			Charset string `xml:"charset,attr"`
			Value   string `xml:"value,attr"`
			Lang    string `xml:"lang,attr"`
		} `xml:"LocalizedString"`
	} `xml:"Name"`
	Description    string `xml:"Description"`
	Classification []struct {
		Text                 string `xml:",chardata"`
		ClassificationScheme string `xml:"classificationScheme,attr"`
		ClassifiedObject     string `xml:"classifiedObject,attr"`
		ID                   string `xml:"id,attr"`
		NodeRepresentation   string `xml:"nodeRepresentation,attr"`
		ObjectType           string `xml:"objectType,attr"`
		Slot                 struct {
			Text      string `xml:",chardata"`
			Name      string `xml:"name,attr"`
			ValueList struct {
				Text  string `xml:",chardata"`
				Value string `xml:"Value"`
			} `xml:"ValueList"`
		} `xml:"Slot"`
		Name struct {
			Text            string `xml:",chardata"`
			LocalizedString struct {
				Text    string `xml:",chardata"`
				Charset string `xml:"charset,attr"`
				Value   string `xml:"value,attr"`
				Lang    string `xml:"lang,attr"`
			} `xml:"LocalizedString"`
		} `xml:"Name"`
		Description string `xml:"Description"`
	} `xml:"Classification"`
	ExternalIdentifier []XMLExternalID `xml:"ExternalIdentifier"`
}

// XMLExternalID is the external identifier in the AdhocQueryResponse.
type XMLExternalID struct {
	Text                 string `xml:",chardata"`
	ID                   string `xml:"id,attr"`
	IdentificationScheme string `xml:"identificationScheme,attr"`
	ObjectType           string `xml:"objectType,attr"`
	RegistryObject       string `xml:"registryObject,attr"`
	Value                string `xml:"value,attr"`
	Name                 struct {
		Text            string `xml:",chardata"`
		LocalizedString struct {
			Text    string `xml:",chardata"`
			Charset string `xml:"charset,attr"`
			Value   string `xml:"value,attr"`
			Lang    string `xml:"lang,attr"`
		} `xml:"LocalizedString"`
	} `xml:"Name"`
	Description string `xml:"Description"`
}

// XMLSlot is a slot in the XMLExtrinsicObject inAdhocQueryResponse.
type XMLSlot struct {
	Text      string `xml:",chardata"`
	Name      string `xml:"name,attr"`
	ValueList struct {
		Text  string   `xml:",chardata"`
		Value []string `xml:"Value"`
	} `xml:"ValueList"`
}

type AdhocQueryResponse struct {
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
