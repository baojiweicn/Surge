package device

type Device struct {
	config *Template
}

func (d *Device) Name() string {
	return d.config.Device.Name
}

// TODO: change to object
func (d *Device) Type() string {
	return ""
}

// TODO: change to objects
func (d *Device) Tags() []string {
	return make([]string, 0)
}

// TODO
func (d *Device) Infos() map[string]interface{} {
}

// TODO
func (d *Device) Storage() map[string]interface{} {

}
