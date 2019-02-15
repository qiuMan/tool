package tool

import(
	"fmt"
	// "path/filepath"
	"strings"
)
type Controller struct {
	Name string
	Path string
}

func NewController(name string) *Controller {
	name = strings.ToLower(name)
	name = StrFirstToUpper(name)
	cfg := NewConfig()
	path, _ := cfg.MvcPath(name, "c")
	obj := &Controller{Name: name, Path: path}
	return obj
}

func (this *Controller) ICk() (bool, string) {
	exits, _ := PathExists(this.Path)
	fmt.Println(this.Path, exits)
	return exits, this.Path
}

func (this *Controller) IRead() []string {
	var	values []string 
	config := NewConfig()
	path, exits := config.TplPath("c")
	if exits == false {
		fmt.Println("控制器模板不存在")
		return values
	}
	file := NewFile()
	values = file.ReadFile(path)
	return values
}

func (this *Controller) IWrite(values []string) {
	file := NewFile()
	
	for key, value := range values {
		str := strings.Replace(value, "<%@#&%>", this.Name, -1)
		values[key] = str
	}

	file.WriteFile(this.Path, values)
}
