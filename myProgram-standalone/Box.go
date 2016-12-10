package main


import (
	"ethos/syscall"
	"ethos/ethos"
//	ethosLog "ethos/log"
	"ethos/efmt"
	"log"
	
)

type SliceSet map[string][]string

func (s SliceSet) Add(key, value string) {
	_, ok := s[key]
	if !ok {
		s[key] = make([]string, 0, 20)
	}
	s[key] = append(s[key], value)
}

func (s SliceSet) Peek(key string) ([]string, bool) {
	slice, ok := s[key]
	if !ok || len(slice) == 0 {
		return s[key], false
	}	
	return s[key], true
}

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
	p2.WriteVar("/user/said3/"+"foo1.txt")
	p3.WriteVar("/user/said3/"+"foo2.txt")
	p4.WriteVar("/user/said3/"+"foo3.txt")
	
	efmt.Println("Sai, status", status)
	
	
	d:= MyType{"Aa","aa"}
	d.ReadVar("/user/said3/foo.txt")
	e:= MyType{"s","aa"}
	e.ReadVar("/user/said3/foo1.txt")
	ff:= MyType{"s","aa"}
	ff.ReadVar("/user/said3/foo2.txt")
	g:= MyType{"s","aa"}
	g.ReadVar("/user/said3/foo3.txt")
		
	efmt.Println("Sai" , status)
	efmt.Println("Sai Read out", d)
	efmt.Println("Sai Read out2", e)
	efmt.Println("Sai Read out4", ff)
	efmt.Println("Sai Read out3", g)

	if SearchWord(d.Field1, d.Field2, "Hello")==1{
	efmt.Println("Sai Found Word")
	}else{
	efmt.Println("Sai Word miss")
	}
	   

	dirFd, statusn := ethos.GetFileInformationPath("/user/said3/foo.txt")
	efmt.Println("Sai dirf",dirFd)
	efmt.Println("Sai statusn", statusn)
	x := make(SliceSet)
	
	
	efmt.Println("Sai filetype", dirFd.FileType)

	typeName, statusName := ethos.TypeNameOfFile("/user/said3/foo.txt")
	efmt.Println("Sai typeName", typeName)
	efmt.Println("Sai Status", statusName)
	x.Add(d.Field1, "foo.txt")
	x.Add(d.Field2, "foo.txt")
	x.Add(e.Field1, "foo1.txt")
	x.Add(e.Field2, "foo1.txt")
	x.Add(ff.Field1, "foo2.txt")
	x.Add(ff.Field2, "foo2.txt")
	x.Add(g.Field1, "foo3.txt")
	x.Add(g.Field2, "foo4.txt")


	v, ok := x.Peek("World")
	if ok {
		efmt.Println("Sai FILE FOUND: FILE NAME",v)
	} else {
		efmt.Println(`Sai FILE NOT FOUND`)
	}
	efmt.Println("Sai",x)
	}


