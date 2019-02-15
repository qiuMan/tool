package tool

import(
	"fmt"
	"strings"
)
type Js struct {
	Name string
	Path string
}

func NewJs(name string) *Js {
	name = strings.ToLower(name)
	config := NewConfig()
	path, _ := config.MvcPath(name, "js") 
	obj := &Js{Name: name, Path: path}
	return obj
}

func (this *Js) ICk() (bool, string) {
	exits, _ := PathExists(this.Path)
	fmt.Println(this.Path, exits)
	return exits, this.Path
}

func (this *Js) IRead() []string {
	var	values []string 
	config := NewConfig()
	path, exits := config.TplPath("js")
	if exits == false {
		fmt.Println("js模板不存在")
		return values
	}
	file := NewFile()
	values = file.ReadFile(path)
	return values
}

func (this *Js) IWrite(values []string) {
	db := NewDb()
	db.SetColumn(this.Name)

	config := NewConfig()
	path, _ := config.MvcPath(this.Name, "js")
	this.Path = path
	for key, value := range values {
		str := strings.Replace(value, "<%@#&%>", this.Name, -1)
		values[key] = str
	}

	file := NewFile()
	file.WriteFile(this.Path, values)
}
