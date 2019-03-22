package languages

import (
	"amwaywave.io/adp/server/models"
	. "gopkg.in/yaml.v2"
)

type yaml struct {
}

func init() {
	register("yaml", &yaml{})
}

func (y *yaml) getExtraAttr() []LanguageExtraAttr {
	return []LanguageExtraAttr{}
}

func (y *yaml) toAPI(content []byte, api *models.API) (err error) {
	err = Unmarshal([]byte(content), api)
	return
}

func (y *yaml) fromAPI(config *ParseConfig) ([]byte, error) {
	return Marshal(config.API)
}

func (y *yaml) useSelfFromAPI() bool {
	return true
}

func (y *yaml) customFuncMap() map[string]interface{} {
	return map[string]interface{}{}
}
