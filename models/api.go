package models

// APIPartType indicates the type of the API field, and the values include request and response.
// The type of field that can be used to specify the conversion when parsing an API instance
type APIPartType int

const (
	// APIPartType indicates the request and response parts of the API
	APIPartType_ALL APIPartType = iota
	// APIPartType_Request represents the request part of the API.
	APIPartType_Request
	// APIPartType_Request represents the response part of the API.
	APIPartType_Response
)

// API represents the details of an api.
type API struct {
	Project     string   `yaml:"project"`
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Type        string   `yaml:"type"`
	Endpoint    string   `yaml:"endpoint"`
	Request     request  `yaml:"request"`
	Response    response `yaml:"response"`
}

type request struct {
	Headers map[string]string `yaml:"headers"`
	Params  map[string]field  `yaml:"params"`
}

type response struct {
	Type   string                 `yaml:"type"`
	Status int                    `yaml:"status"`
	Body   map[string]interface{} `yaml:"body"`
}

type field struct {
	Type        string      `yaml:"type"`
	Description string      `yaml:"description"`
	Checks      []fieldRule `yaml:"checks"`
}

type fieldRule struct {
	Rule   string `yaml:"rule"`
	Pass   string `yaml:"pass"`
	Reject string `yaml:"reject"`
}
