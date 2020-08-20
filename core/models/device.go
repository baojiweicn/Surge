package models

/*
- device.yml
	device config has three important structs:
	device itself, actions, objects
	there are three internal entities used to put device configs
	api:
		Device().Name() => string
		Device().Environment() => *Environment (interface)
		Device().Types() => []*Type (object)
		Device().Storage() => map[string]Value
		Device().Actions() => []*Action
		Device().Action(string) => *Action
		Device().Objects() => []*Action
		Device().Object(string) => *Object
		Action.Exec(map[string]Value) => Object
		Object.Name() => string
		Object.Value(string) => Value
*/

type TemplateInterface interface {
	GetDevice() DeviceInterface
	GetActions() []ActionInterface
	GetStorage()
	GetObjects()
}

type DeviceInterface interface {
	Name() string
	Type() *DeviceType
	Tags() []*DeviceTag
	InfosTemplate() InfosTemplate
}

type ActionInterface interface {
	Name() string
	Executor() string // TODO: change to executor
	Values() map[string]Value
}

type Device struct {
	config TemplateInterface
}

func (d *Device) Name() string {
	return d.config.GetDevice().Name()
}

func (d *Device) Type() *DeviceType {
	return d.config.GetDevice().Type()
}

func (d *Device) Tags() []*DeviceTag {
	return d.config.GetDevice().Tags()
}

func (d *Device) Infos() map[string]Value {
	infoTemp := d.config.GetDevice().InfosTemplate()
	actions := d.config.GetActions()
	for _, action := range actions {
		if action.Name == "infos" {

		}
	}
	// return infoTemp.Marshal(infoMethod())
}

func (d *Device) GetStorage(key string) Value {
	// return nil
}

func (d *Device) GetActions() []Action {
	// return nil
}

type InfosTemplate struct {
}

type DeviceType struct {
}

type DeviceTag struct {
}

type Action struct {
}

type Value struct {
}

type Object struct {
	name   string
	values map[string]Value
}
