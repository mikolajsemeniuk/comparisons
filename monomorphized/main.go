package direct

import "strings"

type GetterSetter interface {
	Get() string
	Set(string) error
}

type Data struct {
	Name string `json:"string"`
}

func (d Data) Get() string {
	return d.Name
}

func (d *Data) Set(name string) error {
	d.Name = name
	return nil
}

type Converter struct{}

func (c *Converter) ConvertData(d *Data) string {
	d.Name = strings.Repeat("b", 10000)
	return d.Name
}

func (c *Converter) ConvertInterface(d GetterSetter) string {
	d.Set(strings.Repeat("b", 10000))
	return d.Get()
}
