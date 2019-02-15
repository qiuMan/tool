package tool

import(
	"fmt"
	"strings"
)

type Model struct {
	Name string
	Path string
	// prefix string
}

func NewModel(name string) *Model {
	name = strings.ToLower(name)
	name = StrFirstToUpper(name)
	config := NewConfig()
	path, _ := config.MvcPath(name, "m")
	obj := &Model{Name: name, Path: path}
	return obj
}

func (this *Model) ICk() (bool, string) {
	exits, _ := PathExists(this.Path)
	fmt.Println(this.Path, exits)
	return exits, this.Path
}

func (this *Model) IRead() []string {
	var	values []string 
	config := NewConfig()
	path, exits := config.TplPath("m")
	
	if exits == false {
		fmt.Println("模型模板不存在")
		return values
	}
	
	file := NewFile()
	values = file.ReadFile(path)
	return values
}

func (this *Model) IWrite(values []string) {
	db := NewDb()
	db.SetColumn(this.Name)
	_struct := "type " + this.Name + " struct {\n	Id		int		`orm:\"pk;column(" +db.Pk+ ")\"`\n"

	for key, value := range db.Columns {
		if key == db.Pk {
			continue
		}

		field := StrFirstToUpper(key)
		_type := "		string"
		
		if value == "int" {
			_type = "		int"
		} 
		orm := "		`orm:\"column(" + key + ")\"`"
		_struct += "	" + field + _type + orm + "\n"			
	}
	_struct += "} "
	for key, value := range values {
		str := strings.Replace(value, "<%@#struct%>", _struct, -1)
		str = strings.Replace(str, "<%@#&%>", this.Name, -1)
		str = strings.Replace(str, "<%@#ID%>", db.Pk, -1)
		values[key] = str
	}
	file := NewFile()
	file.WriteFile(this.Path, values)
}

