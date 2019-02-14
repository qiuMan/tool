package tool

import (
	// "fmt"
	"strings"
)

type Config struct {
	Path string
}

func NewConfig() *Config {
	obj := new(Config)
	obj.Path = "tool/conf"
	return obj
}
func (this *Config) Cfg(name string) string  {
	var v string	
	file := NewFile()
	values := file.ReadFile(this.Path)

	for _, value := range values {
		//去除单行属性两端的空格
		s := strings.TrimSpace(string(value))

		//判断等号=在该行的位置
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		//取得等号左边的key值，判断是否为空
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}

		//取得等号右边的value值，判断是否为空
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		if key == name {
			v = value
			break
		}

	}

	return v
}
//生成文件路径
func (this *Config) MvcPath(name string, mvc ...string) (string, bool){
	path := this.Cfg("path")
	perfix := this.Cfg("filePerfix")

	switch mvc[0] {
		case "m":
			path += "models/" + name + perfix
		case "v":
			path += "views/" + name + "/" + mvc[1] + ".html"
		case "c":
			path += "controllers/" + name + perfix
		case "js":
			path += "static/js/" + name + ".js"
		default:
			path += name + perfix
	}

	exist,_ := PathExists(path)
	return path, exist
	
}

//模板路径
func (this *Config) TplPath(mvc ...string) (string, bool) {
	path := this.Cfg("tplPath")
	perfix := this.Cfg("tplPerfix")
	switch mvc[0] {
		case "m":
			path += "model" + perfix
		case "v":
			path += "view/" + mvc[1] + perfix
		case "c":
			path += "controller" + perfix
		case "js":
			path += "js" + perfix
		default:
			path += mvc[0] + perfix
	}

	exist,_ := PathExists(path)
	return path, exist
}