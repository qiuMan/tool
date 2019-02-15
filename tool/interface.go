package tool


type IFile interface {
	ICk() (bool, string)
	IRead() []string
	IWrite(values []string) 
}

type Tool struct {

}

func NewTool() *Tool {
	obj := new(Tool)
	return obj
}

func (this *Tool) MkFile(ifiles []IFile, paths []string) {
	for _, ifile := range ifiles {
		values := ifile.IRead()
		ifile.IWrite(values)
	}
}
 
func (this *Tool) RmFile(ifile []IFile, paths []string) {
	file := NewFile()
	for _, path := range paths {
		file.Remove(path)
	}
}

func (this *Tool) FileList(str string) ([]IFile, []string) {
	ifiles :=[]IFile{
		NewModel(str),
		NewView(str, "index"), 
		NewView(str, "detail"),  
		NewJs(str), 
		NewController(str),
	}

	paths := make([]string, 0)

	for _, ifile := range ifiles {
		ck, path := ifile.ICk()
		if ck == true {
			paths = append(paths, path)
		}
	}

	return ifiles, paths
	
}