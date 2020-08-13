package models

type Device struct {
	name        string
	environment string
	types       []string
	infos       []interface{}
	storage     []interface{}
}

type Info struct {
	name      string
	value     interface{}
	valueType string
}
