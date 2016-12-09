package main

import (
	"ethos/syscall"
	"ethos/ethos"
//	ethosLog "ethos/log"
	"ethos/efmt"
	"log"
	"os"
   	"path/filepath"
	
)

func CreateData(x string, y string) MyType{
	return MyType{x,y}
}


func SearchWord(x string, y string, word string) int {
	if x==word || y==word{
		return 1
	} 
		return 0

}
 


func main () {

	me := syscall.GetUser()
	path:= "/user/"+me +"/myDir/"
	fd, statusw := ethos.OpenDirectoryPath(path)
	if statusw!= syscall.StatusOk{
		log.Fatalf ("Error %v %v", fd, statusw)
	}	
	
	
	var data MyType
	status := data.CreateDirectory("/user/said3","")
	efmt.Println("Sai",status)
	
	p1:= CreateData("Great","Human")
	p2:= CreateData("Humor","Evil")
	p3:= CreateData("Hello","World")
	p4:= CreateData("World","Fun")

	efmt.Println("Sai, Input", p1)
	efmt.Println("Sai, Input2", p2)
	efmt.Println("Sai, Input3", p3)
	efmt.Println("Sai, Input4", p4)
	
	status = p1.WriteVar("/user/said3/foo.txt")
	p2.WriteVar("/user/said3/foo1.txt")
	p3.WriteVar("/user/said3/foo2.txt")
	p4.WriteVar("/user/said3/foo3.txt")
	
	efmt.Println("Sai, status", status)
	
	
	d:= MyType{"Aa","aa"}
	d.ReadVar("/user/said3/foo.txt")
	d:= MyType{"s","aa"}
	d.ReadVar("/user/said3/foo1.txt")
	f:= MyType{"s","aa"}
	f.ReadVar("/user/said3/foo2.txt")
	g:= MyType{"s","aa"}
	g.ReadVar("/user/said3/foo3.txt")
		
	efmt.Println("Sai" , status)
	efmt.Println("Sai Read out", d)
	efmt.Println("Sai Read out2", e)
	efmt.Println("Sai Read out4", f)
	efmt.Println("Sai Read out3", g)

	if SearchWord(d.Field1, d.Field2, "Hello")==1{
	efmt.Println("Sai Found Word")
	}else{
	efmt.Println("Sai Word miss")
	}
	    searchDir := "/user/said3"

	    fileList := []string{}
	    err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	    })

	    for _, file := range fileList {
		efmt.Println(file)
	    }
//	dirFd, statusn := ethos.GetFileInformationPath("/user/said3/foo.txt")
//	efmt.Println("Sai dirf",dirFd)
//	efmt.Println("Sai statusn", statusn)

	
	}
