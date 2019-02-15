package main 

import(
	"tool/tool"
	"fmt"
	"flag"
	"bufio"
	"os"
)




var mkFile *string = flag.String("m", "", "File contains values for sorting") 
var rmFile *string = flag.String("r", "", "File contains values for sorting") 
// var outfile *string = flag.String("o", "all", "File to receive sorted values") 
// var algorithm *string = flag.String("a", "", "Sort algorithm") 

func main() {
	flag.Parse()
	mkFile := *mkFile 
	rmFile := *rmFile 
	tool := tool.NewTool()
	if mkFile != "" {
		objs, files := tool.FileList(rmFile)
		for _, file := range files {
			fmt.Println(file)
		}
		tool.MkFile(objs, files)
	} else if rmFile != "" {
		objs, files := tool.FileList(rmFile)
		for _, file := range files {
			fmt.Println(file)
		}
		fmt.Printf("确定删除文件y/n：")
		input := bufio.NewScanner(os.Stdin)//初始化一个扫表对象
    	for input.Scan() {//扫描输入内容
        	line := input.Text()//把输入内容转换为字符串
			if line == "y" || line == "n" {
				if line == "y" {
					tool.RmFile(objs, files)
				}    	
				break
			}
			
		}
	}
	
} 








