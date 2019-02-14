package tool


type IFile interface {
	ICk() bool
	IRead() []string
	IWrite(values []string) 
	IRm()
}

type Tool struct {

}

func NewTool() *Tool {
	obj := new(Tool)
	return obj
}

func (this *Tool) MkFile(str string) {
	ifiles :=[]IFile{
		NewModel(str),
		NewView(str, "index"), 
		NewView(str, "detail"),  
		NewJs(str), 
		NewController(str),
	}
	for _, ifile := range ifiles {
		ck := ifile.ICk()
		if ck == true {
			continue
		}
		values := ifile.IRead()
		ifile.IWrite(values)
	}
}
 
func (this *Tool) RmFile(str string) {
	ifiles :=[]IFile{
		NewModel(str),
		NewView(str, "index"), 
		NewView(str, "detail"),  
		NewJs(str), 
		NewController(str),
	}
	for _, ifile := range ifiles {
		ck := ifile.ICk()
		if ck == true {
			ifile.IRm()
		}
	}
}