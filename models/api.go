package models

// APIPartType indicates the type of the API field, and the values include request and response.
// The type of field that can be used to specify the conversion when parsing an API instance
type APIPartType int
type APIFieldType string

const (
	// APIPartType indicates the request and response parts of the API
	APIPartType_ALL APIPartType = iota
	// APIPartType_Request represents the request part of the API.
	APIPartType_Request
	// APIPartType_Request represents the response part of the API.
	APIPartType_Response

	APIFieldType_String   APIFieldType = "string"
	APIFieldType_Int      APIFieldType = "int"
	APIFieldType_Double   APIFieldType = "double"
	APIFieldType_Date     APIFieldType = "date"
	APIFieldType_DateTime APIFieldType = "datetime"
	APIFieldType_Bytes    APIFieldType = "byte[]"
)

var (
	// APIFieldTypes is all types of API fields.
	APIFieldTypes = []APIFieldType{
		APIFieldType_String,
		APIFieldType_Int,
		APIFieldType_Double,
		APIFieldType_Date,
		APIFieldType_DateTime,
		APIFieldType_Bytes,
	}
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
	Headers map[string]string   `yaml:"headers"`
	Params  map[string]APIField `yaml:"params"`
}

type response struct {
	Type   string                 `yaml:"type"`
	Status int                    `yaml:"status"`
	Body   map[string]interface{} `yaml:"body"`
}

type APIField struct {
	Type        APIFieldType `yaml:"type"`
	Description string       `yaml:"description"`
	Checks      []fieldRule  `yaml:"checks"`
}

type fieldRule struct {
	Rule   string `yaml:"rule"`
	Pass   string `yaml:"pass"`
	Reject string `yaml:"reject"`
}
