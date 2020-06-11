package errors

import (
	"github.com/baojiweicn/Surge/util/parser"
)

// Error : is the struct for create new error
type Error struct {
	template parser.Template
	params   []Field
}

// Field : use the field for the Error
type Field struct {
	key   string
	value string
}

// F : create new error field
func F(key, value string) Field {
	return Field{
		key:   key,
		value: value,
	}
}

// NewError : create new error
func NewError(format string) *Error {
	template := parser.NewTemplate(format)
	return &Error{
		template: template,
	}
}

// Raise : put params to the error
func (e *Error) Raise(fields []Field) *Error {
	e.params = fields
	return e
}

// String : render error to string
func (e *Error) String() string {
	params := make([]parser.Field, 0)
	if e.params != nil {
		for _, p := range e.params {
			params = append(params, parser.F(p.key, p.value))
		}
	}
	return e.template.Render(params)
}

// Panic : panic the error
func (e *Error) Panic() {
	panic(e)
}

// Error : error is the interface as error
func (e *Error) Error() string {
	return e.String()
}
