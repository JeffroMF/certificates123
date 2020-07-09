package x509util

import "crypto/x509"

const (
	UserKey               = "User"
	SubjectKey            = "Subject"
	SANsKey               = "SANs"
	TokenKey              = "Token"
	CertificateRequestKey = "CR"
)

// TemplateData is an alias for map[string]interface{}. It represents the data
// passed to the templates.
type TemplateData map[string]interface{}

// NewTemplateData creates a new map for templates data.
func NewTemplateData() TemplateData {
	return TemplateData{}
}

// CreateTemplateData creates a new TemplateData with the given common name and SANs.
func CreateTemplateData(commonName string, sans []string) TemplateData {
	return TemplateData{
		SubjectKey: Subject{
			CommonName: commonName,
		},
		SANsKey: CreateSANs(sans),
	}
}

func (t TemplateData) Set(key string, v interface{}) {
	t[key] = v
}

func (t TemplateData) SetUserData(v Subject) {
	t[UserKey] = v
}

func (t TemplateData) SetSubject(v Subject) {
	t[SubjectKey] = v
}

func (t TemplateData) SetSANs(sans []string) {
	t[SANsKey] = CreateSANs(sans)
}

func (t TemplateData) SetToken(v interface{}) {
	t[TokenKey] = v
}

func (t TemplateData) SetCertificateRequest(cr *x509.CertificateRequest) {
	t[CertificateRequestKey] = newCertificateRequest(cr)
}

const DefaultLeafTemplate = `{
	"subject": {{ toJson .Subject }},
	"sans": {{ toJson .SANs }},
	"keyUsage": ["keyEncipherment", "digitalSignature"],
	"extKeyUsage": ["serverAuth", "clientAuth"]
}`

const DefaultIntermediateTemplate = `{
	"subject": {{ toJson .Subject }},
	"keyUsage": ["certSign", "crlSign"],
	"basicConstraints": {
		"isCA": true,
		"maxPathLen": 0
	}
}`

const DefaultRootTemplate = `{
	"subject": {{ toJson .Subject }},
	"issuer": {{ toJson .Subject }},
	"keyUsage": ["certSign", "crlSign"],
	"basicConstraints": {
		"isCA": true,
		"maxPathLen": 1
	}
}`
