package languages

import (
	"amwaywave.io/adp/server/models"
	"amwaywave.io/adp/server/pkg/tag"
	"bytes"
	"errors"
	"github.com/eightpigs/i18n"
	"text/template"
)

var (
	handles             = map[string]languageHandle{}
	notFoundParserError = errors.New(i18n.Get("language.parse.error.notFoundParser").(string))
)

// ParseConfig indicates the configuration information during the parsing process.
type ParseConfig struct {
	PartType models.APIPartType
	AttrData map[string]string
	Language string
	API      *models.API
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
	getExtraAttr() []LanguageExtraAttr

	// Whether to use the self-implemented FromAPI function.
	useSelfFromAPI() bool
}

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
		// Use its own parsing logic.
		if handles[config.Language].useSelfFromAPI() {
			fromAPIData, err := handles[config.Language].fromAPI(config)
			if err != nil {
				return []byte{}, err
			}
			return fromAPIData, nil
		} else {
			return parseAPIToTmpl(config)
		}
	} else {
		return []byte{}, notFoundParserError
	}
}

// common parsing logic.
func parseAPIToTmpl(config *ParseConfig) ([]byte, error) {
	tmpl, err := template.ParseFiles("assets/templates/" + config.Language + ".tmpl")
	if err != nil {
		return []byte{}, err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, config)
	return buf.Bytes(), err
}
