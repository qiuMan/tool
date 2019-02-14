package tool

import(
	"fmt"
	"io"
	"os"
	"bufio"
	"strings"
	"path/filepath"
)

type File struct {
	Values []string
}

func NewFile() *File {
	obj := new(File)
	return obj
}
//写文件
func (this *File) WriteFile(path string, values []string) bool {

	exist, _ := PathExists(path)
	if exist == false {
		this.Mkdir(path)
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0600)
	defer file.Close() 

	if err != nil {
		fmt.Println(err)
		return false
	}

	for _, value := range values {
		file.WriteString(value + "\n")
	}
	fmt.Println(path)
	return true
}


//读取文件
func (this *File) ReadFile(path string) []string {
	
	exist, _ := PathExists(path)

	if exist == false {
		fmt.Println(path, "file not exist")
		return this.Values 
	}

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		fmt.Println("Failed to open the input file ", file)
	}

	br :=bufio.NewReader(file)
	values := make([]string, 0)

	for {         
		line, _, err1 := br.ReadLine() 
 
        if err1 != nil {             
			if err1 != io.EOF {                 
				err = err1             
			}             
			break         
		} 
 
		str := string(line) // 转换字符数组为字符串 
		values = append(values, str)
	}   
	return values 
}

//创建文件
func (this *File) Mkdir(path string) bool {
	exPath := filepath.Dir(path)//获取文件夹
	var exists bool
	var dir string
	dir_list := strings.Split(exPath, "\\")

	//创建文件夹
	for key, value := range dir_list {
		if key > 0 {
			dir += "/"	
		}
		dir += value 
		exists, _ = PathExists(dir)
		if exists == false {
			err := os.Mkdir(dir, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
			} else {
				fmt.Printf("mkdir success![%v]\n", dir)
			}
		}
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Printf("touch success![%v]\n", path)
	defer file.Close()

	return true
}

//删除文件

func (this *File) Remove(path string) {
	exists, _ := PathExists(path)

	if exists == true {
		fmt.Println("delete:", path)
		err := os.Remove(path)
		fmt.Println(err)
	}

}