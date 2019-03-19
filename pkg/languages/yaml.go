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

func (y *yaml) ToAPI(content []byte, api *models.API) (err error) {
	err = Unmarshal([]byte(content), api)
	return
}

func (y *yaml) FromAPI(api *models.API) ([]byte, error) {
	return Marshal(api)
}
