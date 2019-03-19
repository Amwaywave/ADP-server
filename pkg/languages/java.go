package languages

import (
	"amwaywave.io/adp/server/models"
)

type java struct {
}

func init() {
	register("java", &java{})
}

func (java *java) ToAPI(content []byte, api *models.API) error {
	return nil
}

func (java *java) FromAPI(api *models.API) ([]byte, error) {
	return []byte{}, nil
}
