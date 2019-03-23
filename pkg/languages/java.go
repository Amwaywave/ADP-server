package languages

import (
	"amwaywave.io/adp/server/models"
	"amwaywave.io/adp/server/pkg/tag"
	"bytes"
)

type java struct {
	tmpl string
}

type extraAttr struct {
	// Package is the package name for Java.
	Package string `attr:"required,name:包名,description:$i18n$_language.parsers.java.packageDescription_"`
	// The indentation style of the code.
	Indent int    `attr:"name:缩进,description:缩进,default: 4"`
	Name   string `attr:"name:类名,description:$i18n$_language.parsers.java.className_"`
}

var (
	typeMappings = map[models.APIFieldType]string{
		models.APIFieldType_String:   "String",
		models.APIFieldType_Int:      "Integer",
		models.APIFieldType_Double:   "Double",
		models.APIFieldType_Date:     "Date",
		models.APIFieldType_DateTime: "Date",
		models.APIFieldType_Bytes:    "byte[]",
	}
	typeImports = map[models.APIFieldType]string{
		models.APIFieldType_Date:     "java.util.Date",
		models.APIFieldType_DateTime: "java.util.Date",
	}
)

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

func (java *java) customFuncMap() map[string]interface{} {
	return map[string]interface{}{
		"type": func(apiType models.APIFieldType) string {
			if v, ok := typeMappings[apiType]; ok {
				return v
			}
			return "String"
		},
		"imports": func(fields map[string]models.APIField) string {
			var buf bytes.Buffer
			var importMap = make(map[string]string, len(fields))
			for _, v := range fields {
				if i, ok := typeImports[v.Type]; ok {
					if _, ook := importMap[i]; !ook {
						buf.WriteString("import ")
						buf.WriteString(i)
						buf.WriteString(";\n")
						importMap[i] = ""
					}
				}
			}
			return buf.String()
		},
	}
}
