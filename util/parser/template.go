package parser

import (
	"fmt"
	"strings"
)

/*
example :

>> t := NewTemplate("{{name}} is {{where}}").Render(
	F("name", "bjw"),
	F("where", "flying"),
)

>> bjw is flying

*/

// Template : template is the template for parse
type Template struct {
	template string
}

// NewTemplate : create new template
func NewTemplate(template string) Template {
	return Template{
		template: template,
	}
}

// Keys : get all template keys
func (t Template) Keys() []string {
	var keys []string
	template := t.template
	for strings.Index(template, "{{") != -1 && strings.Index(template, "}}") != -1 {
		start := strings.Index(template, "{{")
		end := strings.Index(template, "}}")
		key := template[start+2 : end]
		keys = append(keys, key)
		template = template[:start] + template[end+1:]
	}
	return keys
}

// Render : render the template to string
func (t Template) Render(params []Field) string {
	template := t.template
	for _, f := range params {
		template = strings.ReplaceAll(template, fmt.Sprintf("{{%v}}", f.Key()), f.Value())
	}
	return template
}

// Field : is the field used to render
type Field struct {
	key   string
	value string
}

// F : create new field
func F(key, value string) Field {
	return Field{
		key:   key,
		value: value,
	}
}

// Key : get field key
func (f Field) Key() string {
	return f.key
}

// Key : get field value
func (f Field) Value() string {
	return f.value
}
