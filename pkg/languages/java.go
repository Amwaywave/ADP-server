package languages

import (
	"amwaywave.io/adp/server/models"
	"amwaywave.io/adp/server/pkg/tag"
)

type java struct {
	tmpl string
}

type extraAttr struct {
	// Package is the package name for Java.
	Package string `attr:"required,name:包名,description:$i18n$_language.parsers.java.packageDescription_"`
	// The indentation style of the code.
	Indent int `attr:"name:缩进,description:缩进,default: 4"`
}

func init() {
	register("java", &java{})
}

func (java *java) getExtraAttr() (attrData []LanguageExtraAttr) {
	attrs := tag.GetLanguageAttrTag(extraAttr{})
	attrData = []LanguageExtraAttr{}
	for k, v := range attrs {
		attrData = append(attrData, LanguageExtraAttr{
			Name: k,
			Attr: v,
		})
	}
	return attrData
}

func (java *java) toAPI(content []byte, api *models.API) error {
	return nil
}

func (java *java) fromAPI(config *ParseConfig) ([]byte, error) {
	return []byte{}, nil
}

func (java *java) useSelfFromAPI() bool {
	return false
}
