package tag

import (
	"amwaywave.io/adp/server/pkg/expression"
	"reflect"
	"strings"
)

const tagName = "attr"

var (
	// Defines the parsing operation for each attr.
	attrParseHandleMap = map[string]func(string, *LanguageAttrTag){
		"required": func(info string, attr *LanguageAttrTag) {
			attr.Required = true
		},
		"name": func(info string, attr *LanguageAttrTag) {
			attr.Name = info
			attr.Name = expression.Parse(info)
		},
		"description": func(info string, attr *LanguageAttrTag) {
			attr.Description = info
			attr.Description = expression.Parse(attr.Description)
		},
		"default": func(info string, attr *LanguageAttrTag) {
			attr.Default = info
			attr.Default = expression.Parse(attr.Default)
		},
	}
)

// LanguageAttrTag is information used to normalize language custom attributes.
type LanguageAttrTag struct {
	Required    bool
	Name        string
	Description string
	Default     string
}

// GetLanguageAttr will return the LanguageAttr information contained in the tags.
func GetLanguageAttrTag(instance interface{}) map[string]LanguageAttrTag {
	t := reflect.TypeOf(instance)
	attrs := map[string]LanguageAttrTag{}
	if t.NumField() == 0 {
		return attrs
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tagName)
		if len(tag) > 0 {
			attr := LanguageAttrTag{}
			// get tags
			tags := strings.Split(tag, ",")
			for _, v := range tags {
				// get tag name and tag value.
				arr := strings.Split(v, ":")
				var name = arr[0]
				var val string
				if len(arr) > 1 {
					val = arr[1]
				}
				if f, ok := attrParseHandleMap[name]; ok {
					f(val, &attr)
				}
			}
			attrs[field.Name] = attr
		}
	}
	return attrs
}
