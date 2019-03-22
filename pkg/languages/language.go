package languages

import (
	"amwaywave.io/adp/server/models"
	"amwaywave.io/adp/server/pkg/tag"
	"bytes"
	"errors"
	"github.com/eightpigs/i18n"
	"strings"
	"text/template"
	"time"
)

var (
	handles             = map[string]languageHandle{}
	notFoundParserError = errors.New(i18n.Get("language.parse.error.notFoundParser").(string))
	globalFuncMap       = template.FuncMap{
		"firstToUpper": func(text string) string {
			return strings.Title(text)
		},
		"nowDate": func() string {
			return time.Now().Format("2006-01-02")
		},
	}
)

// ParseConfig indicates the configuration information during the parsing process.
type ParseConfig struct {
	PartType models.APIPartType
	AttrData map[string]string
	Language string
	API      *models.API
}

// Data used in parsing.
type parseContext struct {
	ParseConfig
	TargetName string
	Fields     map[string]models.APIField
}

// LanguageExtraAttr is the extended attribute data of the language used to customize the data in the language parsing process.
type LanguageExtraAttr struct {
	Name string
	Attr tag.LanguageAttrTag
}

type languageHandle interface {
	// ToAPI is the implementation of the transformation to the API.
	toAPI([]byte, *models.API) error
	fromAPI(config *ParseConfig) ([]byte, error)

	// Get additional attribute information for the language.
	getExtraAttr() []LanguageExtraAttr

	// Whether to use the self-implemented FromAPI function.
	useSelfFromAPI() bool
	// Get language-defined functions.
	customFuncMap() map[string]interface{}
}

// register the language to the language collection.
func register(name string, handle languageHandle) {
	handles[name] = handle
}

// GetExtraAttr will return extended attribute data for the specified language for customizing the language.
func GetExtraAttr(language string) (data []LanguageExtraAttr, err error) {
	if handle, ok := handles[language]; ok {
		data = handle.getExtraAttr()
	} else {
		err = notFoundParserError
	}
	return
}

// To will convert the current API information to the implementation of the specified language.
func ToAPI(language string, content []byte, api *models.API) (err error) {
	if handle, ok := handles[language]; ok {
		err = handle.toAPI(content, api)
	} else {
		err = notFoundParserError
	}
	return
}

// FromAPI returns a byte stream representing the parsed API.
// If the specified language implements FromAPI itself, it uses its own, otherwise it uses common parsing logic.
func FromAPI(config *ParseConfig) ([]byte, error) {
	if _, ok := handles[config.Language]; ok {
		context := getContext(config)
		// Use its own parsing logic.
		if handles[context.Language].useSelfFromAPI() {
			fromAPIData, err := handles[context.Language].fromAPI(config)
			if err != nil {
				return []byte{}, err
			}
			return fromAPIData, nil
		} else {
			return context.parseAPIToTmpl()
		}
	} else {
		return []byte{}, notFoundParserError
	}
}

func getContext(config *ParseConfig) *parseContext {
	context := parseContext{}
	context.API = config.API
	context.PartType = config.PartType
	context.Language = config.Language
	context.AttrData = config.AttrData
	context.setTargetName()
	context.setFields()
	return &context
}

// common parsing logic.
func (context *parseContext) parseAPIToTmpl() ([]byte, error) {
	tmpl, err := template.New(context.Language + ".tmpl").
		Funcs(getFuncMap(handles[context.Language].customFuncMap())).
		ParseFiles("assets/templates/" + context.Language + ".tmpl")
	if err != nil {
		return []byte{}, err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, context)
	return buf.Bytes(), err
}

func (context *parseContext) setTargetName() {
	i := 0
	if strings.Contains(context.API.Endpoint, "/") {
		i = strings.LastIndex(context.API.Endpoint, "/") + 1
	}
	context.TargetName = context.API.Endpoint[i:]
}

func (context *parseContext) setFields() {
	context.Fields = make(map[string]models.APIField, 10)
	if context.PartType == models.APIPartType_Request || context.PartType == models.APIPartType_ALL {
		if len(context.API.Request.Params) > 0 {
			for k, v := range context.API.Request.Params {
				context.Fields[k] = v
			}
		}
	}

	// TODO Parse the contents of the response body. Can generate multiple objects and then nest.
	if context.PartType == models.APIPartType_Response || context.PartType == models.APIPartType_ALL {

	}
}

func getFuncMap(funcs map[string]interface{}) template.FuncMap {
	if len(funcs) == 0 {
		return globalFuncMap
	}
	maps := template.FuncMap{}
	for k, v := range globalFuncMap {
		maps[k] = v
	}
	for k, v := range funcs {
		maps[k] = v
	}
	return maps
}
