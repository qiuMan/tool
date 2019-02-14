package main 

import(
	"tool/tool"
	// "fmt"
	"flag"
	// "bufio"
	// "os"
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
		tool.MkFile(mkFile)
	}

	if rmFile != "" {
		tool.RmFile(rmFile)
	}
	
} 








