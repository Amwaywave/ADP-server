package languages

import (
	"amwaywave.io/adp/server/models"
	"errors"
)

var handles = map[string]languageHandle{}

type languageHandle interface {
	// ToAPI is the implementation of the transformation to the API.
	ToAPI([]byte, *models.API) error
	FromAPI(*models.API) ([]byte, error)
}

func register(name string, handle languageHandle) {
	handles[name] = handle
}

// To will convert the current API information to the implementation of the specified language.
func ToAPI(language string, content []byte, api *models.API) (err error) {
	if handle, ok := handles[language]; ok {
		err = handle.ToAPI(content, api)
	} else {
		err = errors.New("没有找到该语言的转换实现")
	}
	return
}

// From will convert the implementation of the specified language into API information.
func FromAPI(language string, api *models.API) ([]byte, error) {
	if handle, ok := handles[language]; ok {
		return handle.FromAPI(api)
	} else {
		return []byte{}, errors.New("没有找到该语言的转换实现")
	}
}
