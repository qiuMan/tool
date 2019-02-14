package tool

import(
	"fmt"
	"path/filepath"
	"strings"
)
type View struct {
	Name string
	Path string
	Tpl string
}

func NewView(name string, tpl string) *View {
	name = strings.ToLower(name)
	config := NewConfig()
	path, _ := config.MvcPath(name, "v", tpl) 
	obj := &View{Name: name, Path: path, Tpl: tpl}
	return obj
}

func (this *View) ICk() bool {
	exits, _ := PathExists(this.Path)
	fmt.Println(this.Path, exits)
	return exits
}

func (this *View) IRead() []string {
	var	values []string 
	config := NewConfig()
	path, exits := config.TplPath("v", this.Tpl)
	if exits == false {
		fmt.Println("视图模板不存在")
		return values
	}
	file := NewFile()
	values = file.ReadFile(path)
	return values
}

func (this *View) IWrite(values []string) {
	db := NewDb()
	db.SetColumn(this.Name)
	if this.Tpl == "index" {
		var _tThead string
		for key, _ := range db.Columns {
			key = StrFirstToUpper(key)
			_tThead += "				<th lay-data=\"{field:'" + key + "'}\">" + key + "</th> \n"		
		}
		for key, value := range values {
			str := strings.Replace(value, "<%@#tThead%>", _tThead, -1)
			str = strings.Replace(str, "<%@#&%>", this.Name, -1)
			values[key] = str
		}
	} else if this.Tpl == "detail" {
		var item string
		tab := "		"
		for key, _ := range db.Columns {
			key = StrFirstToUpper(key)
			item += tab + "<div class=\"layui-form-item\"> \n" +
			tab + "	<label class=\"layui-form-label\">" + key + "</label> \n" +
			tab + "	<div class=\"layui-input-inline\">\n" + 
			tab + "		<input type=\"text\" name=\"" + key +"\" lay-verify=\"required\" autocomplete=\"off\" value=\"<%.data." + key + "%>\" placeholder=\"请输入\" class=\"layui-input\">\n" + 
			tab + "	<div>\n" + 
			tab + "<div>\n"

		}
		
		for key, value := range values {
			str := strings.Replace(value, "<%@#item%>", item, -1)
			str = strings.Replace(str, "<%@#&%>", this.Name, -1)
			values[key] = str
		}
	} else if (this.Tpl == "js") {
		config := NewConfig()
		path, _ := config.MvcPath(this.Name, "js")
		this.Path = path
		for key, value := range values {
			str := strings.Replace(value, "<%@#&%>", this.Name, -1)
			values[key] = str
		}
	}

	file := NewFile()
	file.WriteFile(this.Path, values)
}

func (this *View) IRm() {
	file := NewFile()
	if this.Tpl == "js" {
		config := NewConfig()
		this.Path, _ = config.MvcPath(this.Name, "js")
	}
	file.Remove(this.Path)
	exPath := filepath.Dir(this.Path)//获取文件夹
	file.Remove(exPath)
}